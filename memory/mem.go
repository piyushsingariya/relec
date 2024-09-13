package memory

import (
	"context"
	"fmt"
	"reflect"
	"runtime"
	"sync"
	"sync/atomic"
	"time"
	"unsafe"

	"github.com/shirou/gopsutil/v4/mem"
)

// Global mutex and condition variables
var (
	memLockMutex  = sync.Mutex{}
	memLockCond   = sync.NewCond(&memLockMutex)
	lowMemTrigger = atomic.Bool{} // Flag to track low memory condition
)

var (
	SkipCleanupRoutine      = false
	BlockExecutionThreshold = uint64(0.1 * float64(TotalMemory()))
	TriggerCleanupThreshold = uint64(0.5 * float64(TotalMemory()))
	gcTriggered             = sync.Mutex{}
)

func SizeInString(size uint64) string {
	// Convert size to GB, MB, or KB
	var sizeInUnits float64
	var unit string

	switch {
	case size >= 1<<30:
		sizeInUnits = float64(size) / (1 << 30)
		unit = "GB"
	case size >= 1<<20:
		sizeInUnits = float64(size) / (1 << 20)
		unit = "MB"
	default:
		sizeInUnits = float64(size) / (1 << 10)
		unit = "KB"
	}

	return fmt.Sprintf("%.2f %s", sizeInUnits, unit)
}

func SizeOfString(instance any) string {
	return SizeInString(SizeOf(instance))
}

// Of returns the size of 'v' in bytes.
// If there is an error during calculation, Of returns -1.
func SizeOf(v any) uint64 {
	// Cache with every visited pointer so we don't count two pointers
	// to the same memory twice.
	cache := make(map[uintptr]bool)
	return uint64(sizeOf(reflect.Indirect(reflect.ValueOf(v)), cache))
}

// sizeOf returns the number of bytes the actual data represented by v occupies in memory.
// If there is an error, sizeOf returns -1.
func sizeOf(v reflect.Value, cache map[uintptr]bool) int {
	switch v.Kind() {

	case reflect.Array:
		sum := 0
		for i := 0; i < v.Len(); i++ {
			s := sizeOf(v.Index(i), cache)
			if s < 0 {
				return -1
			}
			sum += s
		}

		return sum + (v.Cap()-v.Len())*int(v.Type().Elem().Size())

	case reflect.Slice:
		// return 0 if this node has been visited already
		if cache[v.Pointer()] {
			return 0
		}
		cache[v.Pointer()] = true

		sum := 0
		for i := 0; i < v.Len(); i++ {
			s := sizeOf(v.Index(i), cache)
			if s < 0 {
				return -1
			}
			sum += s
		}

		sum += (v.Cap() - v.Len()) * int(v.Type().Elem().Size())

		return sum + int(v.Type().Size())

	case reflect.Struct:
		sum := 0
		for i, n := 0, v.NumField(); i < n; i++ {
			s := sizeOf(v.Field(i), cache)
			if s < 0 {
				return -1
			}
			sum += s
		}

		// Look for struct padding.
		padding := int(v.Type().Size())
		for i, n := 0, v.NumField(); i < n; i++ {
			padding -= int(v.Field(i).Type().Size())
		}

		return sum + padding

	case reflect.String:
		s := v.String()
		hdr := (*reflect.StringHeader)(unsafe.Pointer(&s))
		if cache[hdr.Data] {
			return int(v.Type().Size())
		}
		cache[hdr.Data] = true
		return len(s) + int(v.Type().Size())

	case reflect.Ptr:
		// return Ptr size if this node has been visited already (infinite recursion)
		if cache[v.Pointer()] {
			return int(v.Type().Size())
		}
		cache[v.Pointer()] = true
		if v.IsNil() {
			return int(reflect.New(v.Type()).Type().Size())
		}
		s := sizeOf(reflect.Indirect(v), cache)
		if s < 0 {
			return -1
		}
		return s + int(v.Type().Size())

	case reflect.Bool,
		reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64,
		reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64,
		reflect.Int, reflect.Uint,
		reflect.Chan,
		reflect.Uintptr,
		reflect.Float32, reflect.Float64, reflect.Complex64, reflect.Complex128,
		reflect.Func:
		return int(v.Type().Size())

	case reflect.Map:
		// return 0 if this node has been visited already (infinite recursion)
		if cache[v.Pointer()] {
			return 0
		}
		cache[v.Pointer()] = true
		sum := 0
		keys := v.MapKeys()
		for i := range keys {
			val := v.MapIndex(keys[i])
			// calculate size of key and value separately
			sv := sizeOf(val, cache)
			if sv < 0 {
				return -1
			}
			sum += sv
			sk := sizeOf(keys[i], cache)
			if sk < 0 {
				return -1
			}
			sum += sk
		}
		// Include overhead due to unused map buckets.  10.79 comes
		// from https://golang.org/src/runtime/map.go.
		return sum + int(v.Type().Size()) + int(float64(len(keys))*10.79)

	case reflect.Interface:
		return sizeOf(v.Elem(), cache) + int(v.Type().Size())

	}

	return -1
}

func runGarbageCleanup() {
	if gcTriggered.TryLock() {
		go func() {
			defer gcTriggered.Unlock()
			runtime.GC()
		}()
	}
}

func AvailableMemory() uint64 {
	stat, err := mem.VirtualMemory()
	if err != nil {
		return 0
	}

	return stat.Available
}

func TotalMemory() uint64 {
	stat, err := mem.VirtualMemory()
	if err != nil {
		return 0
	}

	return stat.Total
}

// Lock ensures only one execution of the memory checking process
// and holds the code execution where the function is called
func Lock(ctx context.Context) {
	memLockMutex.Lock()
	defer memLockMutex.Unlock()

	// Wait until memory is above the threshold
	for lowMemTrigger.Load() {
		select {
		case <-ctx.Done():
			return
		default:
			memLockCond.Wait()
		}
	}
}

func init() {
	if !SkipCleanupRoutine {
		go func() {
			for {
				availableMemory := AvailableMemory()

				memLockMutex.Lock()
				if availableMemory <= TriggerCleanupThreshold {
					runGarbageCleanup()
				}

				// If memory is below the threshold and the trigger isn't set, perform cleanup
				if availableMemory <= BlockExecutionThreshold && !lowMemTrigger.Load() {
					lowMemTrigger.Store(true)
				} else if availableMemory > BlockExecutionThreshold {
					// Reset trigger when memory is above the threshold
					lowMemTrigger.Store(false)
					memLockCond.Broadcast() // Notify all waiting goroutines
				}
				memLockMutex.Unlock()

				// Sleep for some time before checking again
				time.Sleep(time.Millisecond * 100)
			}
		}()
	}
}
