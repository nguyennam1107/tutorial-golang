package main

import (
	"fmt"
	"sync"
)
// ============= BÀI 6: SINGLETON PATTERN =============

// TODO: Implement Singleton
type Database struct {
	connectionString string
}

var (
	dbInstance *Database
	dbOnce     sync.Once
)

func GetDatabase() *Database {
	// Implement here
	dbOnce.Do(func() {
		dbInstance = &Database{}
	})
	return dbInstance
}

func (db *Database) Connect(connStr string) {
	db.connectionString = connStr
	fmt.Println("Connected to:", connStr)
}

func exercise6() {
	fmt.Println("\n=== BÀI 6: SINGLETON PATTERN ===")

	// TODO: Test singleton
	db1 := GetDatabase()
	db1.Connect("localhost:5432")

	db2 := GetDatabase()
	fmt.Println("Same instance?", db1 == db2)
}
func main() {
	exercise6()
}