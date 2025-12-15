package handlers_test

import (
	"context"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/your-org/go-rest-layered-template/internal/domain"
	"github.com/your-org/go-rest-layered-template/internal/handlers"
	"github.com/your-org/go-rest-layered-template/internal/services"

	"github.com/go-chi/chi/v5"
)

type repoStub struct {
	user domain.User
	err  error
}

func (r repoStub) GetByID(ctx context.Context, id int64) (domain.User, error) {
	return r.user, r.err
}

func TestUserHandler_GetByID_OK(t *testing.T) {
	svc := services.NewUserService(repoStub{
		user: domain.User{ID: 10, Email: "x@y.com", Name: "X", CreatedAt: time.Unix(1, 0).UTC()},
	})

	h := handlers.NewUserHandler(svc)

	r := chi.NewRouter()
	r.Get("/v1/users/{id}", h.GetByID)

	req := httptest.NewRequest(http.MethodGet, "/v1/users/10", nil)
	res := httptest.NewRecorder()
	r.ServeHTTP(res, req)

	if res.Code != http.StatusOK {
		t.Fatalf("expected 200, got %d: %s", res.Code, res.Body.String())
	}

	var got map[string]any
	if err := json.Unmarshal(res.Body.Bytes(), &got); err != nil {
		t.Fatalf("invalid json: %v", err)
	}
	if got["email"] != "x@y.com" {
		t.Fatalf("unexpected payload: %v", got)
	}
}

func TestUserHandler_GetByID_BadRequest(t *testing.T) {
	svc := services.NewUserService(repoStub{})
	h := handlers.NewUserHandler(svc)

	r := chi.NewRouter()
	r.Get("/v1/users/{id}", h.GetByID)

	req := httptest.NewRequest(http.MethodGet, "/v1/users/nope", nil)
	res := httptest.NewRecorder()
	r.ServeHTTP(res, req)

	if res.Code != http.StatusBadRequest {
		t.Fatalf("expected 400, got %d", res.Code)
	}
}
