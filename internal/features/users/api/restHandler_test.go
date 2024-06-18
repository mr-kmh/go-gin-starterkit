package api

import (
	"github.com/gin-gonic/gin"
	"github.com/mr-kmh/go-gin-starterkit/internal/features/users/entities"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"testing"
)

type MockService struct{}

func (m *MockService) GetUsers() ([]entities.User, error) {
	return make([]entities.User, 0), nil
}

func setup() (*gin.Engine, *RESTHandler) {
	router := gin.Default()
	rest := NewREST(&MockService{})
	return router, rest
}

func TestWeclome(t *testing.T) {
	router, rest := setup()
	router.GET("/welcome", rest.Welcome)

	req, _ := http.NewRequest(http.MethodGet, "/welcome", nil)

	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)

	expectedResponse := `{"message": "Welcome to Users API"}`
	assert.JSONEq(t, expectedResponse, w.Body.String())
}

func TestListUsers(t *testing.T) {
	router, rest := setup()

	router.GET("/", rest.ListUsers)

	req, _ := http.NewRequest(http.MethodGet, "/", nil)
	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)

	expectedResponse := `{"message":[]}`
	assert.JSONEq(t, expectedResponse, w.Body.String())
}
