package tests

import (
	"bytes"
	"database/sql"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
	"user-api/internal/handler"
	"user-api/internal/model"

	"github.com/gin-gonic/gin"
)

// MockUserRepository - мок-репозиторий для тестов
type MockUserRepository struct{}

func (m *MockUserRepository) Create(user *model.User) error {
	user.ID = 1
	user.CreatedAt = "2025-03-21T12:00:00Z"
	return nil
}

func (m *MockUserRepository) Get(id int) (*model.User, error) {
	if id == 1 {
		return &model.User{
			ID:        1,
			Name:      "Test User",
			Email:     "test@example.com",
			CreatedAt: "2025-03-21T12:00:00Z",
		}, nil
	}
	return nil, sql.ErrNoRows
}

func (m *MockUserRepository) Update(user *model.User) error {
	return nil
}

// setupRouter - вспомогательная функция для настройки роутера
func setupRouter() (*gin.Engine, *handler.UserHandler) {
	gin.SetMode(gin.TestMode)
	router := gin.Default()
	mockRepo := &MockUserRepository{}
	userHandler := handler.NewUserHandler(mockRepo)

	router.POST("/users", userHandler.CreateUser)
	router.GET("/users/:id", userHandler.GetUser)
	router.PUT("/users/:id", userHandler.UpdateUser)

	return router, userHandler
}

// TestCreateUser - тест создания пользователя
func TestCreateUser(t *testing.T) {
	router, _ := setupRouter()

	user := model.User{
		Name:  "Test User",
		Email: "test@example.com",
	}
	jsonValue, _ := json.Marshal(user)

	req, _ := http.NewRequest("POST", "/users", bytes.NewBuffer(jsonValue))
	req.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()

	router.ServeHTTP(w, req)

	if w.Code != http.StatusCreated {
		t.Errorf("Expected status %v, got %v", http.StatusCreated, w.Code)
	}

	var response model.User
	if err := json.Unmarshal(w.Body.Bytes(), &response); err != nil {
		t.Errorf("Failed to unmarshal response: %v", err)
	}

	if response.ID != 1 {
		t.Errorf("Expected ID 1, got %v", response.ID)
	}
	if response.Name != user.Name {
		t.Errorf("Expected Name %v, got %v", user.Name, response.Name)
	}
}

// TestGetUser - тест получения пользователя
func TestGetUser(t *testing.T) {
	router, _ := setupRouter()

	req, _ := http.NewRequest("GET", "/users/1", nil)
	w := httptest.NewRecorder()

	router.ServeHTTP(w, req)

	if w.Code != http.StatusOK {
		t.Errorf("Expected status %v, got %v", http.StatusOK, w.Code)
	}

	var response model.User
	if err := json.Unmarshal(w.Body.Bytes(), &response); err != nil {
		t.Errorf("Failed to unmarshal response: %v", err)
	}

	if response.ID != 1 {
		t.Errorf("Expected ID 1, got %v", response.ID)
	}
	if response.Name != "Test User" {
		t.Errorf("Expected Name 'Test User', got %v", response.Name)
	}
}

// TestGetUserNotFound - тест получения несуществующего пользователя
func TestGetUserNotFound(t *testing.T) {
	router, _ := setupRouter()

	req, _ := http.NewRequest("GET", "/users/999", nil)
	w := httptest.NewRecorder()

	router.ServeHTTP(w, req)

	if w.Code != http.StatusNotFound {
		t.Errorf("Expected status %v, got %v", http.StatusNotFound, w.Code)
	}
}

// TestUpdateUser - тест обновления пользователя
func TestUpdateUser(t *testing.T) {
	router, _ := setupRouter()

	user := model.User{
		Name:  "Updated User",
		Email: "updated@example.com",
	}
	jsonValue, _ := json.Marshal(user)

	req, _ := http.NewRequest("PUT", "/users/1", bytes.NewBuffer(jsonValue))
	req.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()

	router.ServeHTTP(w, req)

	if w.Code != http.StatusOK {
		t.Errorf("Expected status %v, got %v", http.StatusOK, w.Code)
	}

	var response model.User
	if err := json.Unmarshal(w.Body.Bytes(), &response); err != nil {
		t.Errorf("Failed to unmarshal response: %v", err)
	}

	if response.Name != user.Name {
		t.Errorf("Expected Name %v, got %v", user.Name, response.Name)
	}
	if response.Email != user.Email {
		t.Errorf("Expected Email %v, got %v", user.Email, response.Email)
	}
}

// TestUpdateUserInvalidID - тест обновления с некорректным ID
func TestUpdateUserInvalidID(t *testing.T) {
	router, _ := setupRouter()

	user := model.User{
		Name:  "Updated User",
		Email: "updated@example.com",
	}
	jsonValue, _ := json.Marshal(user)

	req, _ := http.NewRequest("PUT", "/users/invalid", bytes.NewBuffer(jsonValue))
	req.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()

	router.ServeHTTP(w, req)

	if w.Code != http.StatusBadRequest {
		t.Errorf("Expected status %v, got %v", http.StatusBadRequest, w.Code)
	}
}