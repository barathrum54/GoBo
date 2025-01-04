package models

import (
	"context"
	"gobo/internal/db"
)

// Example represents an example table structure
type Example struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

// CreateExample inserts a new record into the "examples" table
func CreateExample(name string) error {
	_, err := db.Conn.Exec(context.Background(), "INSERT INTO examples (name) VALUES ($1)", name)
	return err
}

// GetExamples retrieves all records from the "examples" table
func GetExamples() ([]Example, error) {
	rows, err := db.Conn.Query(context.Background(), "SELECT id, name FROM examples")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var examples []Example
	for rows.Next() {
		var example Example
		if err := rows.Scan(&example.ID, &example.Name); err != nil {
			return nil, err
		}
		examples = append(examples, example)
	}

	return examples, nil
}
