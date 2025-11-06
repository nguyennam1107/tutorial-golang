package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"reflect"
)
// ============= BÀI 2: JSON SERIALIZER =============

// TODO: Marshal struct to JSON
func Marshal(v interface{}) ([]byte, error) {
	// Implement here
	val := reflect.ValueOf(v)
	_ = val
	switch val.Kind() {
	case reflect.Struct:
		return json.Marshal(v)
	case reflect.Map:
		return json.Marshal(v)
	case reflect.Slice:
		return json.Marshal(v)
	default:
		return json.Marshal(v)
	}
	return nil, errors.New("not implemented")
}

// TODO: Unmarshal JSON to struct
func Unmarshal(data []byte, v interface{}) error {
	// Implement here
	val := reflect.ValueOf(v)
	if val.Kind() != reflect.Ptr || val.IsNil() {
		return errors.New("v must be a non-nil pointer")
	}
	
	return json.Unmarshal(data, v)
}

func exercise2() {
	fmt.Println("\n=== BÀI 2: JSON SERIALIZER ===")

	type Address struct {
		City    string `json:"city"`
		Country string `json:"country"`
	}

	type Person struct {
		Name    string   `json:"name"`
		Age     int      `json:"age"`
		Emails  []string `json:"emails,omitempty"`
		Address *Address `json:"address,omitempty"`
	}

	// TODO: Test Marshal
	p := Person{Name: "John", Age: 30}
	data, err := Marshal(p)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(string(data))

	// TODO: Test Unmarshal
	var p2 Person
	err = Unmarshal(data, &p2)
	fmt.Println("Unmarshaled:", p2)
}
func main() {
	exercise2()
}