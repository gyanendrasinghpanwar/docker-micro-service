package registry

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
)

func TestRegisterRoutes(t *testing.T) {
	// Set Gin to Test Mode
	gin.SetMode(gin.TestMode)

	// Create a router and register the routes
	router := gin.Default()
	RegisterRoutes(router)

	t.Run("Hello World", func(t *testing.T) {
		req, _ := http.NewRequest(http.MethodGet, "/hello", nil)
		rr := httptest.NewRecorder()
		router.ServeHTTP(rr, req)

		if status := rr.Code; status != http.StatusOK {
			t.Errorf("expected status %v but got %v", http.StatusOK, status)
		}

		expected := `{"message":"Hello, World!"}`
		if rr.Body.String() != expected {
			t.Errorf("expected %v but got %v", expected, rr.Body.String())
		}
	})

	t.Run("Greet Endpoint", func(t *testing.T) {
		jsonBody := `{"name":"John"}`
		req, _ := http.NewRequest(http.MethodPost, "/greet", bytes.NewBufferString(jsonBody))
		req.Header.Set("Content-Type", "application/json")
		rr := httptest.NewRecorder()
		router.ServeHTTP(rr, req)

		if status := rr.Code; status != http.StatusOK {
			t.Errorf("expected status %v but got %v", http.StatusOK, status)
		}

		expected := `{"message":"Hello, John"}`
		if rr.Body.String() != expected {
			t.Errorf("expected %v but got %v", expected, rr.Body.String())
		}
	})

	t.Run("Binary Search Endpoint", func(t *testing.T) {
		requestBody := RequestBody{
			Numbers: []int{1, 2, 3, 4, 5},
			Target:  3,
		}
		jsonBody, _ := json.Marshal(requestBody)
		req, _ := http.NewRequest(http.MethodPost, "/binary-search", bytes.NewBuffer(jsonBody))
		req.Header.Set("Content-Type", "application/json")
		rr := httptest.NewRecorder()
		router.ServeHTTP(rr, req)

		if status := rr.Code; status != http.StatusOK {
			t.Errorf("expected status %v but got %v", http.StatusOK, status)
		}

		expected := `{"index":2}`
		if rr.Body.String() != expected {
			t.Errorf("expected %v but got %v", expected, rr.Body.String())
		}
	})

	t.Run("Binary Search Not Found", func(t *testing.T) {
		requestBody := RequestBody{
			Numbers: []int{1, 2, 3, 4, 5},
			Target:  6,
		}
		jsonBody, _ := json.Marshal(requestBody)
		req, _ := http.NewRequest(http.MethodPost, "/binary-search", bytes.NewBuffer(jsonBody))
		req.Header.Set("Content-Type", "application/json")
		rr := httptest.NewRecorder()
		router.ServeHTTP(rr, req)

		if status := rr.Code; status != http.StatusOK {
			t.Errorf("expected status %v but got %v", http.StatusOK, status)
		}

		expected := `{"index":-1}`
		if rr.Body.String() != expected {
			t.Errorf("expected %v but got %v", expected, rr.Body.String())
		}
	})

	t.Run("Binary Search Invalid Request", func(t *testing.T) {
		invalidJSONBody := `{"numbers":[],"target":3}`
		req, _ := http.NewRequest(http.MethodPost, "/binary-search", bytes.NewBufferString(invalidJSONBody))
		req.Header.Set("Content-Type", "application/json")
		rr := httptest.NewRecorder()
		router.ServeHTTP(rr, req)

		if status := rr.Code; status != http.StatusBadRequest {
			t.Errorf("expected status %v but got %v", http.StatusBadRequest, status)
		}
		fmt.Println(rr.Body.String())
		expected := `{"error":"Numbers can not be empty array"}`
		if rr.Body.String() != expected {
			t.Errorf("expected %v but got %v", expected, rr.Body.String())
		}
	})
}
