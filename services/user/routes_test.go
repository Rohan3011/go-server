package user

import (
	"bytes"
	"encoding/json"
	"log"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/go-chi/chi/v5"
	"github.com/rohan3011/go-server/types"
)

func TestUserServiceHandler(t *testing.T) {

	userStore := &mockUserStore{}
	handler := NewHandler(userStore)

	t.Run("should fail if the user payload is invalid", func(t *testing.T) {
		payload := types.RegisterUserPayload{
			FirstName: "user",
			LastName:  "123",
			Email:     "test@email.com",
			Password:  "12345",
		}

		marshalled, _ := json.Marshal(payload)
		req, err := http.NewRequest(http.MethodPost, "/register", bytes.NewBuffer(marshalled))
		if err != nil {
			log.Fatal(err)
		}

		record := httptest.NewRecorder()
		router := chi.NewRouter()
		router.Post("/register", handler.handleRegister)
		router.ServeHTTP(record, req)

		if record.Code != http.StatusBadRequest {
			t.Errorf("expected status code %d, got %d", http.StatusBadRequest, record.Code)
		}
		t.Logf("Response :%v", record)
	})
}

type mockUserStore struct {
}

func (s *mockUserStore) GetUserByEmail(email string) (*types.User, error) {

	return nil, nil
}

func (s *mockUserStore) GetUserByID(id int) (*types.User, error) {
	return nil, nil
}

func (s *mockUserStore) CreateUser(user types.User) error {
	return nil
}
