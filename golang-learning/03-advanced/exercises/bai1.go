package main

import (
	"fmt"
	"reflect"
)
// ============= BÀI 1: REFLECTION BASICS =============

// TODO: Inspect struct và in ra thông tin
func inspectStruct(v interface{}) {
	// Implement here
	val := reflect.ValueOf(v)
	typ := reflect.TypeOf(v)

	if typ.Kind() != reflect.Struct {
		fmt.Println("Provided value is not a struct")
		return
	}

	fmt.Printf("Struct Type: %s\n", typ.Name())
	fmt.Printf("Number of Fields: %d\n", typ.NumField())

	for i := 0; i < typ.NumField(); i++ {
		field := typ.Field(i)
		value := val.Field(i)

		fmt.Printf("Field Name: %s\n", field.Name)
		fmt.Printf("Field Type: %s\n", field.Type)
		fmt.Printf("Field Value: %v\n", value.Interface())
		fmt.Printf("Field Tag: %s\n", field.Tag)
		fmt.Println("-----")
	}
}

// TODO: Call method qua reflection
func callMethod(v interface{}, methodName string) error {
	// Implement here
	val := reflect.ValueOf(v)
	method := val.MethodByName(methodName)
	if !method.IsValid() {
		return fmt.Errorf("method %s does not exist", methodName)
	}

	if method.Type().NumIn() != 0 || method.Type().NumOut() != 0 {
		return fmt.Errorf("method %s has parameters or return values, not supported in this example", methodName)
	}
	method.Call(nil)	
	return nil
}

// TODO: Modify field value qua reflection
func modifyField(v interface{}, fieldName string, newValue interface{}) error {
	// Implement here
	val := reflect.ValueOf(v)
	if val.Kind() != reflect.Ptr || val.Elem().Kind() != reflect.Struct {
		return fmt.Errorf("provided value is not a pointer to struct")
	}
	
	val = val.Elem()
	field := val.FieldByName(fieldName)
	if !field.IsValid() {
		return fmt.Errorf("field %s does not exist", fieldName)
	}
	if !field.CanSet() {
		return fmt.Errorf("field %s cannot be set", fieldName)
	}
	newVal := reflect.ValueOf(newValue)
	if newVal.Type() != field.Type() {
		return fmt.Errorf("provided value type %s does not match field type %s", newVal.Type(), field.Type())
	}
	field.Set(newVal)
	return nil
}

func exercise1() {
	fmt.Println("=== BÀI 1: REFLECTION BASICS ===")

	type User struct {
		Name  string `json:"name" validate:"required"`
		Email string `json:"email" validate:"email"`
		Age   int    `json:"age" validate:"min=0,max=120"`
	}

	user := User{Name: "John", Email: "john@example.com", Age: 30}

	// TODO: Test inspectStruct
	inspectStruct(user)

	// TODO: Test modifyField
	modifyField(&user, "Age", 31)
	fmt.Println("Modified user:", user)
}
func main() {
	exercise1()
}