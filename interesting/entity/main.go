package main

import (
	"encoding/json"
	"fmt"
)

type Child struct {
	Name string
}

type Parent struct {
	Child
	SomeOtherField string `json:"some_other_field"`
}

func (p Parent) MarshalJSON() ([]byte, error) {
	return json.Marshal(map[string]interface{}{
		"child":            p.Child,
		"some_other_field": p.SomeOtherField,
	})
}

// func (p *Parent) UnmarshalJSON(b []byte) error {
// 	var raw map[string]interface{}
// 	if err := json.Unmarshal(b, &raw); err != nil {
// 		return err
// 	}

// 	if child, ok := raw["child"].(map[string]interface{}); ok {
// 		p.Child.Name = child["name"].(string)
// 	}

// 	if someOtherField, ok := raw["some_other_field"].(string); ok {
// 		p.SomeOtherField = someOtherField
// 	}

// 	return nil
// }

func main() {
	s := `{
	"name": "John",
	"some_other_field": "Some value"
}`
	var p Parent
	err := json.Unmarshal([]byte(s), &p)
	if err != nil {
		panic(err)
	}
	fmt.Println(p)

	b, err := json.Marshal(p)
	if err != nil {
		panic(err)
	}
	fmt.Println(string(b))
}
