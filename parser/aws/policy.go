package paws

import (
	"encoding/json"
	"fmt"

	"github.com/mitchellh/mapstructure"
	"github.com/piyushsingariya/relec"
)

// Policy represents an AWS iam policy document
type Policy struct {
	Version    string      `json:"Version"`
	ID         string      `json:"ID,omitempty"`
	Statements []Statement `json:"Statement"`
}

func NewPolicy() *Policy {
	return &Policy{
		Version: "2012-10-17",
	}
}

// UnmarshalJSON decodifies input JSON info to awsPolicy type
func (p *Policy) UnmarshalJSON(policy []byte) error {
	var raw any
	var err error
	var statementList []Statement

	err = json.Unmarshal(policy, &raw)
	if err != nil {
		return err
	}

	// Parsing content of JSON element as empty interface
	switch object := raw.(type) {
	// All elelements
	case map[string]any:
		for key, value := range object {
			switch key {
			case "Version":
				p.Version = value.(string)
			case "ID":
				p.ID = value.(string)
			case "Statement":
				statementList = make([]Statement, 0)
				// Statement level - slice -> []any , single element -> map[string]interface
				switch statement := value.(type) {
				// Statement slice -> iterate over elements, parse and store into slice
				case []any:
					// statement slice
					// iterate over statements
					for _, statementValue := range statement {
						statement := Statement{}
						// Type assertion to format info
						statementMap := statementValue.(map[string]any)
						// Parse statement
						err := statement.Parse(statementMap)
						if err != nil {
							return err
						}
						// Append statement to slice
						statementList = append(statementList, statement)
					}
				// Single statement -> parse and store it into slice
				case map[string]any:
					statementMap := Statement{}
					// Parse statement
					err := statementMap.Parse(statement)
					if err != nil {
						return err
					}
					statementList = append(statementList, statementMap)
				}
				//Assign statements slice to Policy
				p.Statements = statementList
			}
		}
	}
	return err
}

func (p *Policy) ToMap() (map[string]any, error) {
	out := make(map[string]any)

	return out, relec.Unmarshal(p, &out)
}

// Statement represents body of AWS iam policy document
type Statement struct {
	StatementID  string              `json:"StatementID,omitempty"`  // Statement ID, service specific
	Effect       string              `json:"Effect"`                 // Allow or Deny
	Principal    map[string][]string `json:"Principal,omitempty"`    // principal that is allowed or denied
	NotPrincipal map[string][]string `json:"NotPrincipal,omitempty"` // exception to a list of principals
	Action       []string            `json:"Action"`                 // allowed or denied action
	NotAction    []string            `json:"NotAction,omitempty"`    // matches everything except
	Resource     []string            `json:"Resource,omitempty"`     // object or objects that the statement covers
	NotResource  []string            `json:"NotResource,omitempty"`  // matches everything except
	Condition    []string            `json:"Condition,omitempty"`    // conditions for when a policy is in effect
}

// Parse decodifies input JSON info into Statement type
func (statementJSON *Statement) Parse(statement map[string]any) error {
	// Definitions
	var principal, notPrincipal, action, notAction, resource, notResource, condition []string
	var err error

	/* Iterate over map elements, each key element (statementKey) is the statement element
	identifier and each value element (statementValue) the statement element value */
	for statementKey, statementValue := range statement {
		// Switch case over key type (identifying Statement elements)
		switch statementKey {
		case "StatementID":
			// Type assertion to assign
			statementJSON.StatementID = statementValue.(string)
		case "Effect":
			//Type assertion to assign
			statementJSON.Effect = statementValue.(string)
		case "Principal":
			// principal(statementValue) can be map[string][]string/string -> needs processing
			// Initialize map
			statementJSON.Principal = make(map[string][]string)
			// procesing map
			mapStatement := statementValue.(map[string]any)
			// iterate over key principal (keyPrincipal) and value principal (valuePrincipal)
			for keyPrincipal, valuePrincipal := range mapStatement {
				// valuePrincipal can be string or []string
				switch valuePrincipal := valuePrincipal.(type) {
				case string:
					// As map each element is identified with a key and has a value
					principal = make([]string, 0)
					statementJSON.Principal[keyPrincipal] = append(principal, valuePrincipal)
				case []any:
					/* If value is an interface we know we have an []string -> knowing final type
					we can use mapstructure (which uses reflect) to store as final type */
					err = mapstructure.Decode(statementValue, &statementJSON.Principal)
					if err != nil {
						return fmt.Errorf("error parsing policies : error using mapstructure parsing Policy statement principal element: %s", err)
					}
				default:
					return fmt.Errorf("error parsing policies: error using mapstructure parsing Policy statement condition element: found unknown Principal type %T %v", statementValue, statementValue)
				}
			}
		case "NotPrincipal":
			// Same case as principal
			// notprincipal has to be statementValue = map[string][]string/string -> needs processing
			// Same procedure as Principal
			// Initialize map
			statementJSON.NotPrincipal = make(map[string][]string)
			// procesing map (statementValue)
			mapStatement := statementValue.(map[string]any)
			for keyNotPrincipal, valueNotPrincipal := range mapStatement {
				// valueNotPrincipal can be string or []string
				switch vnp := valueNotPrincipal.(type) {
				case string:
					notPrincipal = make([]string, 0)
					statementJSON.NotPrincipal[keyNotPrincipal] = append(notPrincipal, vnp)
				case []any:
					err = mapstructure.Decode(statementValue, &statementJSON.NotPrincipal)
					if err != nil {
						return fmt.Errorf("error parsing policies: error using mapstructure parsing Policy statement not principal element: %s", err)
					}
				default:
					return fmt.Errorf("error parsing policies: error using mapstructure parsing Policy statement condition element: found unknown NotPrincipal type %T %v", statementValue, statementValue)
				}
			}
		case "Action":
			// We only have now string or []string, process with type assertion and mapstructure
			// Action can be string or []string
			switch statementValue := statementValue.(type) {
			case string:
				action = make([]string, 0)
				statementJSON.Action = append(action, statementValue)
			case []any:
				err = mapstructure.Decode(statementValue, &statementJSON.Action)
				if err != nil {
					return fmt.Errorf("error parsing policies: error using mapstructure parsing Policy statement action element: %s", err)
				}
			default:
				return fmt.Errorf("error parsing policies: error using mapstructure parsing Policy statement condition element: found unknown Action type %T %v", statementValue, statementValue)
			}
		case "NotAction":
			// Same as Action
			// NotAction can be string or []string
			switch statementValue := statementValue.(type) {
			case string:
				notAction = make([]string, 0)
				statementJSON.NotAction = append(notAction, statementValue)
			case []any:
				err = mapstructure.Decode(statementValue, &statementJSON.NotAction)
				if err != nil {
					return fmt.Errorf("error parsing policies: error using mapstructure parsing Policy statement not action element: %s", err)
				}
			default:
				return fmt.Errorf("error parsing policies: error using mapstructure parsing Policy statement condition element: found unknown NotAction type %T %v", statementValue, statementValue)
			}
		case "Resource":
			// Same as Action
			// Resource can be string or []string
			switch statementValue := statementValue.(type) {
			case string:
				resource = make([]string, 0)
				statementJSON.Resource = append(resource, statementValue)
			case []any:
				err = mapstructure.Decode(statementValue, &statementJSON.Resource)
				if err != nil {
					return fmt.Errorf("error parsing policies: error using mapstructure parsing Policy statement resource element: %s", err)
				}
			default:
				return fmt.Errorf("error parsing policies: error using mapstructure parsing Policy statement condition element: found unknown Resource type %T %v", statementValue, statementValue)
			}
		case "NotResource":
			// Same as Action
			// NotResource can be string or []string
			switch statementValue := statementValue.(type) {
			case string:
				notResource = make([]string, 0)
				statementJSON.NotResource = append(notResource, statementValue)
			case []any:
				err = mapstructure.Decode(statementValue, &statementJSON.NotResource)
				if err != nil {
					return fmt.Errorf("error parsing policies: error using mapstructure parsing Policy statement not resource element: %s", err)
				}
			default:
				return fmt.Errorf("error parsing policies: error using mapstructure parsing Policy statement condition element: found unknown NotResource type %T %v", statementValue, statementValue)
			}
		case "Condition":
			// Condition can be string, []string or map(lot of options)
			switch statementValue := statementValue.(type) {
			case string:
				condition = make([]string, 0)
				statementJSON.Condition = append(condition, statementValue)
			case []any:
				condition = make([]string, 0)
				for _, element := range statementValue {
					value, err := json.Marshal(element)
					if err != nil {
						return fmt.Errorf("error parsing policies: error using mapstructure parsing Policy statement condition element: %s", err)
					}
					condition = append(condition, string(value))
				}

				statementJSON.Condition = append(statementJSON.Condition, condition...)
			// If map format as raw text and store it as string
			case map[string]any:
				condition = make([]string, 0)
				value, err := json.Marshal(statementValue)
				if err != nil {
					return fmt.Errorf("error parsing policies: error using mapstructure parsing Policy statement condition element: %s", err)
				}
				statementJSON.Condition = append(condition, string(value))
			default:
				return fmt.Errorf("error parsing policies: error using mapstructure parsing Policy statement condition element: found unknown Condition type %T %v", statementValue, statementValue)
			}
		}
	}

	return nil
}
