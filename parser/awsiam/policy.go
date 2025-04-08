package awsiam

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"strings"

	"github.com/piyushsingariya/relec"
)

const (
	// See https://docs.aws.amazon.com/IAM/latest/UserGuide/reference_policies_elements_effect.html
	EffectAllow = "Allow"
	EffectDeny  = "Deny"

	// See https://docs.aws.amazon.com/IAM/latest/UserGuide/reference_policies_elements_version.html
	Version2012_10_17 = "2012-10-17"
	Version2008_10_17 = "2008-10-17"
	VersionLatest     = Version2012_10_17

	ErrorInvalidStatementSlice = "Statements is not a slice of statements"
	ErrorInvalidStatements     = "Statements must be a single Statement or a slice of Statements"
)

// Policy is a policy document.
type Policy struct {
	Id         string      `json:"Id,omitempty"`
	Statements *Statements `json:"Statement"`
	Version    string      `json:"Version"`
}

func (p *Policy) ToMap() (map[string]any, error) {
	out := make(map[string]any)

	return out, relec.Unmarshal(p, &out)
}

// Statement is a single statement in a policy document.
type Statement struct {
	Action       *StringOrSlice                        `json:"Action,omitempty"`
	Condition    map[string]map[string]*ConditionValue `json:"Condition,omitempty"`
	Effect       string                                `json:"Effect"`
	NotAction    *StringOrSlice                        `json:"NotAction,omitempty"`
	NotResource  *StringOrSlice                        `json:"NotResource,omitempty"`
	Principal    *Principal                            `json:"Principal,omitempty"`
	NotPrincipal *Principal                            `json:"NotPrincipal,omitempty"`
	Resource     *StringOrSlice                        `json:"Resource,omitempty"`
	Sid          string                                `json:"Sid,omitempty"`
}

// Statements represents Statements that can be marshaled to a single Statement or a slice of Statements.
type Statements struct {
	values   []Statement
	singular bool
}

// NewSingularStatements creates a new Statements with a single Statement.
func NewSingleStatement(statements Statement) *Statements {
	return &Statements{
		values:   []Statement{statements},
		singular: true,
	}
}

// NewStatements creates a new Statements with a slice of Statements.
func NewStatements(statements ...Statement) *Statements {
	return &Statements{
		values:   statements,
		singular: false,
	}
}

// ConditionValue is a value in a condition statement.
func (s *Statements) Add(statements ...Statement) {
	s.values = append(s.values, statements...)
	if len(s.values) > 1 {
		s.singular = false
	}
}

func (s *Statements) UnmarshalJSON(data []byte) error {
	var tmp interface{}
	err := json.Unmarshal(data, &tmp)
	if err != nil {
		return err
	}
	_, ok := tmp.([]interface{})
	if ok {
		// TODO: can we avoid strict decoding and defer to the outer
		values := []Statement{}
		decoder := json.NewDecoder(bytes.NewReader(data))
		decoder.DisallowUnknownFields()
		err = decoder.Decode(&values)
		if err != nil {
			return fmt.Errorf("%s: %v", ErrorInvalidStatementSlice, err)

		}
		s.values = values
		s.singular = false
		return nil
	}
	_, ok = tmp.(map[string]interface{})
	if ok {
		value := Statement{}
		decoder := json.NewDecoder(bytes.NewReader(data))
		decoder.DisallowUnknownFields()
		err = decoder.Decode(&value)
		if err != nil {
			return fmt.Errorf("%s: %v", ErrorInvalidStatements, err)
		}
		s.values = []Statement{value}
		s.singular = true
		return nil
	}
	return errors.New(ErrorInvalidStatements)
}

func (s *Statements) MarshalJSON() ([]byte, error) {
	buf := bytes.Buffer{}
	enc := json.NewEncoder(&buf)
	enc.SetEscapeHTML(false)

	if s.singular && len(s.values) == 1 {
		err := enc.Encode(s.values[0])
		return []byte(strings.TrimSpace(buf.String())), err
	}
	err := enc.Encode(s.values)
	return []byte(strings.TrimSpace(buf.String())), err
}

// Values returns the statement values of the Statements.
func (s *Statements) Values() []Statement {
	return s.values
}

// Singular returns true if the Statements is a single Statement.
func (s *Statements) Singular() bool {
	return s.singular
}
