package main

import (
	"errors"
	"fmt"
	"strconv"
)

/*
ERROR HANDLING - XỬ LÝ LỖI

Go không có exception/try-catch
Errors là values, thường là return value cuối cùng
*/

// 1. Error là interface
// type error interface {
//     Error() string
// }

func main() {
	// 1. BASIC ERROR HANDLING
	fmt.Println("=== 1. BASIC ERROR HANDLING ===")

	// Function trả về error
	divide := func(a, b float64) (float64, error) {
		if b == 0 {
			return 0, errors.New("division by zero")
		}
		return a / b, nil
	}

	result, err := divide(10, 2)
	if err != nil {
		fmt.Println("Error:", err)
	} else {
		fmt.Printf("10 / 2 = %.2f\n", result)
	}

	result, err = divide(10, 0)
	if err != nil {
		fmt.Println("Error:", err)
	} else {
		fmt.Printf("Result: %.2f\n", result)
	}

	// 2. errors.New() vs fmt.Errorf()
	fmt.Println("\n=== 2. CREATING ERRORS ===")

	// errors.New() - simple error
	err1 := errors.New("simple error message")
	fmt.Println("Error 1:", err1)

	// fmt.Errorf() - formatted error
	userId := 123
	err2 := fmt.Errorf("user with id %d not found", userId)
	fmt.Println("Error 2:", err2)

	// 3. CUSTOM ERROR TYPES
	fmt.Println("\n=== 3. CUSTOM ERROR TYPES ===")

	type ValidationError struct {
		Field   string
		Message string
	}

	func (e ValidationError) Error() string {
		return fmt.Sprintf("validation error on field '%s': %s", e.Field, e.Message)
	}

	validateAge := func(age int) error {
		if age < 0 {
			return ValidationError{
				Field:   "age",
				Message: "must be non-negative",
			}
		}
		if age > 150 {
			return ValidationError{
				Field:   "age",
				Message: "must be less than 150",
			}
		}
		return nil
	}

	if err := validateAge(-5); err != nil {
		fmt.Println("Validation error:", err)
	}

	if err := validateAge(200); err != nil {
		fmt.Println("Validation error:", err)
	}

	if err := validateAge(25); err != nil {
		fmt.Println("Validation error:", err)
	} else {
		fmt.Println("Age is valid")
	}

	// 4. ERROR WRAPPING (Go 1.13+)
	fmt.Println("\n=== 4. ERROR WRAPPING ===")

	readConfig := func() error {
		return errors.New("file not found")
	}

	loadApp := func() error {
		err := readConfig()
		if err != nil {
			return fmt.Errorf("failed to load app: %w", err)
		}
		return nil
	}

	if err := loadApp(); err != nil {
		fmt.Println("Error:", err)

		// Unwrap error
		if unwrapped := errors.Unwrap(err); unwrapped != nil {
			fmt.Println("Unwrapped:", unwrapped)
		}
	}

	// 5. errors.Is() và errors.As()
	fmt.Println("\n=== 5. errors.Is() và errors.As() ===")

	var ErrNotFound = errors.New("not found")
	var ErrUnauthorized = errors.New("unauthorized")

	fetchUser := func(id int) error {
		if id < 0 {
			return fmt.Errorf("invalid user id: %w", ErrNotFound)
		}
		return nil
	}

	err = fetchUser(-1)
	if errors.Is(err, ErrNotFound) {
		fmt.Println("User not found!")
	}

	// errors.As() cho custom types
	type NetworkError struct {
		Code    int
		Message string
	}

	func (e NetworkError) Error() string {
		return fmt.Sprintf("network error %d: %s", e.Code, e.Message)
	}

	makeRequest := func() error {
		return NetworkError{Code: 500, Message: "internal server error"}
	}

	err = makeRequest()
	var netErr NetworkError
	if errors.As(err, &netErr) {
		fmt.Printf("Network error with code: %d\n", netErr.Code)
	}

	// 6. MULTIPLE ERROR HANDLING
	fmt.Println("\n=== 6. MULTIPLE ERROR HANDLING ===")

	processData := func(data string) error {
		if data == "" {
			return errors.New("data is empty")
		}
		if len(data) > 100 {
			return errors.New("data too long")
		}
		// Process...
		return nil
	}

	data := []string{"valid", "", "this is a very long string that exceeds the maximum allowed length for processing in this function"}

	for i, d := range data {
		if err := processData(d); err != nil {
			fmt.Printf("Data %d error: %v\n", i, err)
		} else {
			fmt.Printf("Data %d: OK\n", i)
		}
	}

	// 7. DEFER WITH ERROR HANDLING
	fmt.Println("\n=== 7. DEFER WITH ERROR ===")

	openFile := func(filename string) error {
		fmt.Printf("Opening file: %s\n", filename)
		// Simulate file operation
		return nil
	}

	closeFile := func(filename string) {
		fmt.Printf("Closing file: %s\n", filename)
	}

	readFile := func(filename string) error {
		if err := openFile(filename); err != nil {
			return fmt.Errorf("failed to open file: %w", err)
		}
		defer closeFile(filename) // Always called

		// Read operations...
		fmt.Println("Reading file...")

		return nil
	}

	if err := readFile("test.txt"); err != nil {
		fmt.Println("Error:", err)
	}

	// 8. PANIC VÀ RECOVER
	fmt.Println("\n=== 8. PANIC VÀ RECOVER ===")

	safeDivide := func(a, b int) (result int) {
		defer func() {
			if r := recover(); r != nil {
				fmt.Println("Recovered from panic:", r)
				result = 0
			}
		}()

		if b == 0 {
			panic("division by zero!")
		}
		return a / b
	}

	fmt.Println("Result:", safeDivide(10, 2))
	fmt.Println("Result:", safeDivide(10, 0))
	fmt.Println("Program continues...")

	// 9. SENTINEL ERRORS
	fmt.Println("\n=== 9. SENTINEL ERRORS ===")

	var (
		ErrInvalidInput = errors.New("invalid input")
		ErrTimeout      = errors.New("timeout")
		ErrCanceled     = errors.New("canceled")
	)

	operation := func(input string) error {
		if input == "" {
			return ErrInvalidInput
		}
		if input == "timeout" {
			return ErrTimeout
		}
		if input == "cancel" {
			return ErrCanceled
		}
		return nil
	}

	handleError := func(err error) {
		switch {
		case errors.Is(err, ErrInvalidInput):
			fmt.Println("Please provide valid input")
		case errors.Is(err, ErrTimeout):
			fmt.Println("Operation timed out, please retry")
		case errors.Is(err, ErrCanceled):
			fmt.Println("Operation was canceled")
		default:
			if err != nil {
				fmt.Println("Unknown error:", err)
			}
		}
	}

	handleError(operation(""))
	handleError(operation("timeout"))
	handleError(operation("cancel"))
	handleError(operation("valid"))

	// 10. PRACTICAL EXAMPLES
	fmt.Println("\n=== 10. PRACTICAL EXAMPLES ===")

	// User registration
	type User struct {
		Username string
		Email    string
		Age      int
	}

	validateUser := func(u User) error {
		if u.Username == "" {
			return fmt.Errorf("username is required")
		}
		if len(u.Username) < 3 {
			return fmt.Errorf("username must be at least 3 characters")
		}
		if u.Email == "" {
			return fmt.Errorf("email is required")
		}
		if u.Age < 18 {
			return fmt.Errorf("must be 18 or older")
		}
		return nil
	}

	users := []User{
		{Username: "alice", Email: "alice@example.com", Age: 25},
		{Username: "bo", Email: "bob@example.com", Age: 30},
		{Username: "charlie", Email: "", Age: 20},
		{Username: "david", Email: "david@example.com", Age: 15},
	}

	for i, user := range users {
		if err := validateUser(user); err != nil {
			fmt.Printf("User %d (%s) validation failed: %v\n", i+1, user.Username, err)
		} else {
			fmt.Printf("User %d (%s): Valid ✓\n", i+1, user.Username)
		}
	}

	// String to int conversion with error handling
	fmt.Println("\n=== String to Int Conversion ===")
	
	strings := []string{"123", "456", "abc", "789"}
	sum := 0

	for _, s := range strings {
		num, err := strconv.Atoi(s)
		if err != nil {
			fmt.Printf("Cannot convert '%s' to int: %v\n", s, err)
			continue
		}
		sum += num
	}
	fmt.Printf("Sum of valid numbers: %d\n", sum)

	// 11. ERROR BEST PRACTICES
	fmt.Println("\n=== 11. BEST PRACTICES ===")

	fmt.Println("✓ Always check errors")
	fmt.Println("✓ Don't ignore errors with _")
	fmt.Println("✓ Return errors instead of panic")
	fmt.Println("✓ Add context to errors when wrapping")
	fmt.Println("✓ Use custom error types for complex cases")
	fmt.Println("✓ Document which errors a function can return")
}
