package main

import (
	"fmt"
)	
// ============= BÀI 7: FACTORY PATTERN =============

// TODO: Factory Pattern
type Logger interface {
	Log(message string)
}

type ConsoleLogger struct{}

func (cl ConsoleLogger) Log(message string) {
	fmt.Println("[CONSOLE]", message)
}

type FileLogger struct {
	filename string
}

func (fl FileLogger) Log(message string) {
	fmt.Printf("[FILE:%s] %s\n", fl.filename, message)
}

// TODO: Logger factory
func NewLogger(loggerType string, params ...string) Logger {
	// Implement here
	switch loggerType {
	case "console":
		return ConsoleLogger{}
	case "file":
		if len(params) > 0 {
			return FileLogger{filename: params[0]}
		}
	}
	return nil
}

func exercise7() {
	fmt.Println("\n=== BÀI 7: FACTORY PATTERN ===")

	// TODO: Test factory
	logger1 := NewLogger("console")
	logger1.Log("Hello from console")

	logger2 := NewLogger("file", "app.log")
	logger2.Log("Hello from file")
}
func main() {
	exercise7()
}
