// Package routes contains tests for the application's API endpoints.
// These tests validate the functionality of the registered routes, including
// CRUD operations for the Example model.
package routes

import (
	"encoding/json"
	"log"
	"net/http/httptest"
	"strings"
	"testing"

	"gobo/internal/db"
	"gobo/internal/models"

	"github.com/gofiber/fiber/v2"
	"github.com/joho/godotenv"
	"github.com/stretchr/testify/assert"
)

// setupGormTestDB initializes the test database for each test case.
// It loads environment variables, connects to the database, and runs migrations.
//
// Parameters:
// - t (*testing.T): The test context for managing test state.
func setupGormTestDB(t *testing.T) {
	log.Println("[Setup] Starting GORM test database setup...")

	// Load environment variables from the .env file.
	err := godotenv.Load("../../.env")
	if err != nil {
		t.Fatalf("[Error] Error loading .env file: %v", err)
	}

	log.Println("[Setup] Environment variables loaded successfully.")

	// Connect to the database using GORM.
	db.ConnectGORM()

	// Run database migrations for the Example model.
	err = models.AutoMigrateExamples(db.GormDB)
	if err != nil {
		t.Fatalf("[Error] Error during migrations: %v", err)
	}

	log.Println("[Setup] Test database setup completed successfully.")
}

// teardownTestDB cleans up the test database after each test.
// It drops the `examples` table to ensure a clean state for subsequent tests.
func teardownTestDB() {
	log.Println("[Teardown] Dropping test tables...")

	// Drop the `examples` table.
	db.GormDB.Exec("DROP TABLE IF EXISTS examples")

	log.Println("[Teardown] Test database cleaned up.")
}

// TestGetExamples validates the GET /examples endpoint.
// It ensures that examples can be retrieved from the database and returned in the API response.
func TestGetExamples(t *testing.T) {
	// Set up the test database.
	setupGormTestDB(t)
	defer teardownTestDB()

	// Add a test example to the database.
	testExample := models.Example{Name: "Test Example"}
	if result := db.GormDB.Create(&testExample); result.Error != nil {
		t.Fatalf("[Error] Failed to add test example: %v", result.Error)
	}

	log.Println("[Test] Added test example to the database.")

	// Create a new Fiber app instance and register routes.
	app := fiber.New()
	Register(app)

	// Perform the GET request to the /examples endpoint.
	req := httptest.NewRequest("GET", "/examples", nil)
	resp, err := app.Test(req)

	// Assert the response status code is 200 OK.
	assert.NoError(t, err)
	assert.Equal(t, 200, resp.StatusCode)

	// Parse the response body to extract examples.
	var examples []models.Example
	err = json.NewDecoder(resp.Body).Decode(&examples)
	assert.NoError(t, err)

	// Assert that the response contains at least one example.
	assert.NotEmpty(t, examples)
	assert.Equal(t, "Test Example", examples[0].Name)

	log.Println("[Test] GET /examples response validated successfully.")
}

// TestCreateExample validates the POST /examples endpoint.
// It ensures that a new example can be created and saved to the database.
func TestCreateExample(t *testing.T) {
	// Set up the test database.
	setupGormTestDB(t)
	defer teardownTestDB()

	// Create a new Fiber app instance and register routes.
	app := fiber.New()
	Register(app)

	// Define a request body for creating a new example.
	body := `{"name": "New Example"}`

	// Perform the POST request to the /examples endpoint.
	req := httptest.NewRequest("POST", "/examples", strings.NewReader(body))
	req.Header.Set("Content-Type", "application/json")
	resp, err := app.Test(req)

	// Assert the response status code is 201 Created.
	assert.NoError(t, err)
	assert.Equal(t, 201, resp.StatusCode)

	// Parse the response body to verify the result.
	var response map[string]interface{}
	err = json.NewDecoder(resp.Body).Decode(&response)
	assert.NoError(t, err)

	// Assert the success message and the presence of an ID field.
	assert.Equal(t, "Example created successfully", response["message"])
	assert.NotNil(t, response["id"])

	log.Println("[Test] POST /examples response validated successfully.")

	// Fetch the newly created example from the database.
	var example models.Example
	db.GormDB.Last(&example)
	assert.Equal(t, "New Example", example.Name)

	log.Println("[Test] Example saved successfully to the database.")
}
