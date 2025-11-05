package main

import (
	"fmt"
	"strings"
)
// ============= BÀI 8: ERROR HANDLING =============

// TODO: Validate email
func validateEmail(email string) error {
	// Check: không rỗng, có @, có .
	// Implement here
	if email == "" {
		return ValidationError{Field: "email", Message: "không được để trống"}
	}
	if !strings.Contains(email, "@") {
		return ValidationError{Field: "email", Message: "phải có @ trong địa chỉ email"}
	}
	if !strings.Contains(email, ".") {
		return ValidationError{Field: "email", Message: "phải có . trong địa chỉ email"}
	}
	return nil
}

// TODO: Custom error type
type ValidationError struct {
	Field   string
	Message string
}

func (e ValidationError) Error() string {
	// Implement here
	return fmt.Sprintf("Validation error on field '%s': %s", e.Field, e.Message)
}

// TODO: Validate user với wrapped errors
func validateUser(name, email string, age int) error {
	// Implement here
	if err := validateEmail(email); err != nil {
		return fmt.Errorf("validateUser: %w", err)
	}
	if name == "" {
		return ValidationError{Field: "name", Message: "không được để trống"}
	}
	if age < 0 {
		return ValidationError{Field: "age", Message: "phải lớn hơn 0"}
	}
	return nil
}

func exercise8() {
	fmt.Println("\n=== BÀI 8: ERROR HANDLING ===")

	emails := []string{
		"valid@example.com",
		"invalid",
		"",
		"no-at-sign.com",
	}

	for _, email := range emails {
		if err := validateEmail(email); err != nil {
			fmt.Printf("'%s': %v\n", email, err)
		} else {
			fmt.Printf("'%s': Valid ✓\n", email)
		}
	}
}

// ============= MAIN =============

func main() {
	exercise8()
}
