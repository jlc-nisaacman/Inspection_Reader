//go:build ignore

// Database initialization script
// Run this once to create all required database tables
// Usage: go run init_db.go
package main

import (
	"fmt"
	"main/parsers"
)

func main() {
	fmt.Println("=== Database Initialization ===")
	fmt.Println("Connecting to database...")

	// Connect to the PostgreSQL database
	db := parsers.ConnectDatabase()
	defer db.Close()

	fmt.Println("✓ Connected to database successfully")
	fmt.Println("\nCreating tables...")

	// Create all tables
	parsers.CreateInspectionTable(db)
	fmt.Println("✓ Inspections table created/verified")

	parsers.CreateDryTable(db)
	fmt.Println("✓ Dry systems table created/verified")

	parsers.CreatePumpTable(db)
	fmt.Println("✓ Pump systems table created/verified")

	parsers.CreateBackflowTable(db)
	fmt.Println("✓ Backflow table created/verified")

	fmt.Println("\n=== Database initialization complete! ===")
	fmt.Println("All tables are ready. You can now run your main application.")
}
