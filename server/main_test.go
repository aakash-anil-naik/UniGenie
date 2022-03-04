package main

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"
	"unigenie/api/models"

	"github.com/gin-gonic/gin"
)

func TestGetUniversities(t *testing.T) {
	gin.SetMode(gin.TestMode)

	r := gin.Default()
	r.GET("/getUniversities", models.GetUniversities)

	req, err := http.NewRequest(http.MethodGet, "/getUniversities", nil)

	if err != nil {
		t.Fatalf("Couldn't create a request: %v\n", err)
	}

	w := httptest.NewRecorder()

	r.ServeHTTP(w, req)
	fmt.Println(w.Body)

	if w.Code == http.StatusOK {
		t.Logf("Expected to get status %d is same ast %d\n", http.StatusOK, w.Code)
	} else {
		t.Fatalf("Expected to get status %d but instead got %d\n", http.StatusOK, w.Code)
	}
}

func TestGetUsers(t *testing.T) {
	gin.SetMode(gin.TestMode)

	r := gin.Default()
	r.GET("/getUsers", models.GetUniversities)

	req, err := http.NewRequest(http.MethodGet, "/getUsers", nil)

	if err != nil {
		t.Fatalf("Couldn't create a request: %v\n", err)
	}

	w := httptest.NewRecorder()

	r.ServeHTTP(w, req)
	fmt.Println(w.Body)

	if w.Code == http.StatusOK {
		t.Logf("Expected to get status %d is same ast %d\n", http.StatusOK, w.Code)
	} else {
		t.Fatalf("Expected to get status %d but instead got %d\n", http.StatusOK, w.Code)
	}
}
