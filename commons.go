package relec

import (
	"fmt"
	"os"
	"reflect"
	"time"

	"github.com/goccy/go-json"
)

func Absolute[T int | int8 | int16 | int32 | int64 | float32 | float64](value T) T {
	if value < 0 {
		return -value
	}

	return value
}

func ArrayContains[T ~string | int | int8 | int16 | int32 | int64 | float32 | float64](set []T, value T) bool {
	_, found := FindInArray(set, func(elem T) bool {
		return elem == value
	})

	return found
}

func FindInArray[T any](set []T, match func(elem T) bool) (int, bool) {
	for idx, elem := range set {
		if match(elem) {
			return idx, true
		}
	}

	return -1, false
}

// Unmarshal serializes and deserializes any from into the object
// return error if occurred
func Unmarshal(from interface{}, object interface{}) error {
	reformatted := reformatInnerMaps(from)
	b, err := json.Marshal(reformatted)
	if err != nil {
		return fmt.Errorf("error marshalling object: %v", err)
	}
	err = json.Unmarshal(b, object)
	if err != nil {
		return fmt.Errorf("error unmarshalling from object: %v", err)
	}

	return nil
}

func IsInstance(val any, typ reflect.Kind) bool {
	return reflect.ValueOf(val).Kind() == typ
}

// reformatInnerMaps converts all map[interface{}]interface{} into map[string]interface{}
// because json.Marshal doesn't support map[interface{}]interface{} (supports only string keys)
// but viper produces map[interface{}]interface{} for inner maps
// return recursively converted all map[interface]interface{} to map[string]interface{}
func reformatInnerMaps(valueI interface{}) interface{} {
	switch value := valueI.(type) {
	case []interface{}:
		for i, subValue := range value {
			value[i] = reformatInnerMaps(subValue)
		}
		return value
	case map[interface{}]interface{}:
		newMap := make(map[string]interface{}, len(value))
		for k, subValue := range value {
			newMap[fmt.Sprint(k)] = reformatInnerMaps(subValue)
		}
		return newMap
	case map[string]interface{}:
		for k, subValue := range value {
			value[k] = reformatInnerMaps(subValue)
		}
		return value
	default:
		return valueI
	}
}

func CheckIfFilesExists(files ...string) error {
	for _, file := range files {
		// Check if the file or directory exists
		_, err := os.Stat(file)
		if os.IsNotExist(err) {
			return fmt.Errorf("%s does not exist: %s", file, err)
		}

		_, err = os.ReadFile(file)
		if err != nil {
			return fmt.Errorf("failed to read %s: %s", file, err)
		}
	}

	return nil
}

// func ReadFile(file string) any {
// 	content, _ := ReadFileE(file)

// 	return content
// }

func UnmarshalFile(file string, dest any) error {
	if err := CheckIfFilesExists(file); err != nil {
		return err
	}

	data, err := os.ReadFile(file)
	if err != nil {
		return fmt.Errorf("file not found : %s", err)
	}

	err = json.Unmarshal(data, dest)
	if err != nil {
		return err
	}

	return nil
}

func IsOfType(object interface{}, decidingKey string) (bool, error) {
	objectMap := make(map[string]interface{})
	if err := Unmarshal(object, &objectMap); err != nil {
		return false, err
	}

	if _, found := objectMap[decidingKey]; found {
		return true, nil
	}

	return false, nil
}

func IsSubset[T comparable](setArray, subsetArray []T) bool {
	set := make(map[T]bool)
	for _, item := range setArray {
		set[item] = true
	}

	for _, item := range subsetArray {
		if _, found := set[item]; !found {
			return false
		}
	}

	return true
}

func MaxDate(v1, v2 time.Time) time.Time {
	if v1.After(v2) {
		return v1
	}

	return v2
}


// ReplaceWithASCII replaces characters in a string with their ASCII values,
// except for alphanumeric characters, which remain unchanged.
func ReplaceWithASCII(input string) string {
	result := ""
	for _, char := range input {
		if unicode.IsLetter(char) || unicode.IsDigit(char) {
			// Keep alphanumeric characters unchanged
			result += string(char)
		} else {
			// Replace other characters with their ASCII values
			result += strconv.Itoa(int(char))
		}
	}
	return result
}