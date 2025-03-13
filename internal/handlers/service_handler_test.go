package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"services-api/internal/database"
	"services-api/internal/models"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

func setupTestRouter() *gin.Engine {
	gin.SetMode(gin.TestMode)
	router := gin.Default()
	router.GET("/services", ListServices)
	router.GET("/services/:id", GetService)
	return router
}

func TestIntegration(t *testing.T) {
	// Setup test database
	database.SetupTestDB()
	defer func() {
		if err := database.ClearDB(); err != nil {
			t.Errorf("Failed to clear test database: %v", err)
		}
	}()

	// Setup router
	router := setupTestRouter()

	// First get the list of services to get a valid ID
	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/services", nil)
	router.ServeHTTP(w, req)

	var listResponse models.ServiceResponse
	err := json.Unmarshal(w.Body.Bytes(), &listResponse)
	assert.NoError(t, err)
	assert.NotEmpty(t, listResponse.Services, "Should have at least one service")

	validID := listResponse.Services[0].ID

	t.Run("List Services", func(t *testing.T) {
		w := httptest.NewRecorder()
		req, _ := http.NewRequest("GET", "/services", nil)
		router.ServeHTTP(w, req)

		// Check status code
		assert.Equal(t, http.StatusOK, w.Code)

		// Parse response
		var response models.ServiceResponse
		err := json.Unmarshal(w.Body.Bytes(), &response)
		assert.NoError(t, err)

		// Verify we got our seeded services
		assert.Equal(t, 5, len(response.Services)) // We seeded exactly 5 services
		assert.Equal(t, int64(5), response.Total)  // Total should match our seed count
	})

	t.Run("Get Service", func(t *testing.T) {
		w := httptest.NewRecorder()
		req, _ := http.NewRequest("GET", fmt.Sprintf("/services/%d", validID), nil)
		router.ServeHTTP(w, req)

		// Check status code
		assert.Equal(t, http.StatusOK, w.Code)

		// Parse response
		var service models.Service
		err := json.Unmarshal(w.Body.Bytes(), &service)
		assert.NoError(t, err)

		// Verify we got the correct service
		assert.Equal(t, validID, service.ID)
		assert.NotEmpty(t, service.Name)
		assert.NotEmpty(t, service.Description)
	})

	t.Run("Get Non-existent Service", func(t *testing.T) {
		w := httptest.NewRecorder()
		req, _ := http.NewRequest("GET", "/services/999", nil)
		router.ServeHTTP(w, req)

		// Check status code
		assert.Equal(t, http.StatusNotFound, w.Code)
	})
}
