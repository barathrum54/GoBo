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

// setupGormTestDB prepares the test database for each test case
func setupGormTestDB(t *testing.T) {
	log.Println("[Setup] Starting GORM test database setup...")

	err := godotenv.Load("../../.env")
	if err != nil {
		t.Fatalf("[Error] Error loading .env file: %v", err)
	}

	log.Println("[Setup] Environment variables loaded successfully.")

	// Connect to GORM and initialize database
	db.ConnectGORM()

	// Run database migrations
	err = models.AutoMigrateExamples(db.GormDB)
	if err != nil {
		t.Fatalf("[Error] Error during migrations: %v", err)
	}

	log.Println("[Setup] Test database setup completed successfully.")
}

// teardownTestDB cleans up the test database after each test
func teardownTestDB() {
	log.Println("[Teardown] Dropping test tables...")
	db.GormDB.Exec("DROP TABLE IF EXISTS examples")
	log.Println("[Teardown] Test database cleaned up.")
}

// Test GET /examples
func TestGetExamples(t *testing.T) {
	setupGormTestDB(t)
	defer teardownTestDB()

	// Add a test example to the database
	testExample := models.Example{Name: "Test Example"}
	if result := db.GormDB.Create(&testExample); result.Error != nil {
		t.Fatalf("[Error] Failed to add test example: %v", result.Error)
	}

	log.Println("[Test] Added test example to the database.")

	// Create a new Fiber app instance
	app := fiber.New()
	Register(app)

	// Perform the GET request
	req := httptest.NewRequest("GET", "/examples", nil)
	resp, err := app.Test(req)

	// Assert the response status code is 200 OK
	assert.NoError(t, err)
	assert.Equal(t, 200, resp.StatusCode)

	// Parse the response body to get examples
	var examples []models.Example
	err = json.NewDecoder(resp.Body).Decode(&examples)
	assert.NoError(t, err)

	// Assert we got at least one example in the response
	assert.NotEmpty(t, examples)
	assert.Equal(t, "Test Example", examples[0].Name)

	log.Println("[Test] GET /examples response validated successfully.")
}

// Test POST /examples
func TestCreateExample(t *testing.T) {
	setupGormTestDB(t)
	defer teardownTestDB()

	// Create a new Fiber app instance
	app := fiber.New()
	Register(app)

	// Create a request body for POST /examples
	body := `{"name": "New Example"}`

	// Perform the POST request
	req := httptest.NewRequest("POST", "/examples", strings.NewReader(body))
	req.Header.Set("Content-Type", "application/json")
	resp, err := app.Test(req)

	// Assert the response status code is 201 Created
	assert.NoError(t, err)
	assert.Equal(t, 201, resp.StatusCode)

	// Parse the response body
	var response map[string]interface{}
	err = json.NewDecoder(resp.Body).Decode(&response)
	assert.NoError(t, err)

	// Assert the success message and the ID field
	assert.Equal(t, "Example created successfully", response["message"])
	assert.NotNil(t, response["id"])

	log.Println("[Test] POST /examples response validated successfully.")

	// Fetch the example from the database to ensure it was created
	var example models.Example
	db.GormDB.Last(&example)
	assert.Equal(t, "New Example", example.Name)

	log.Println("[Test] Example saved successfully to the database.")
}
