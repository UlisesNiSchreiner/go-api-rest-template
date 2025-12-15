package services_test

import (
	"context"
	"errors"
	"testing"
	"time"

	"github.com/your-org/go-rest-layered-template/internal/domain"
	"github.com/your-org/go-rest-layered-template/internal/repositories/mysqlrepo"
	"github.com/your-org/go-rest-layered-template/internal/services"
)

type mockRepo struct {
	getByID func(ctx context.Context, id int64) (domain.User, error)
}

func (m mockRepo) GetByID(ctx context.Context, id int64) (domain.User, error) {
	return m.getByID(ctx, id)
}

func TestUserService_GetByID_OK(t *testing.T) {
	expected := domain.User{ID: 7, Email: "a@b.com", Name: "Ada", CreatedAt: time.Unix(1, 0).UTC()}
	svc := services.NewUserService(mockRepo{
		getByID: func(ctx context.Context, id int64) (domain.User, error) {
			if id != 7 {
				t.Fatalf("unexpected id: %d", id)
			}
			return expected, nil
		},
	})

	got, err := svc.GetByID(context.Background(), 7)
	if err != nil {
		t.Fatalf("expected nil error, got %v", err)
	}
	if got.ID != expected.ID || got.Email != expected.Email || got.Name != expected.Name {
		t.Fatalf("unexpected user: %#v", got)
	}
}

func TestUserService_GetByID_MapsNotFound(t *testing.T) {
	svc := services.NewUserService(mockRepo{
		getByID: func(ctx context.Context, id int64) (domain.User, error) {
			return domain.User{}, mysqlrepo.ErrNotFound
		},
	})

	_, err := svc.GetByID(context.Background(), 123)
	if !errors.Is(err, services.ErrUserNotFound) {
		t.Fatalf("expected ErrUserNotFound, got %v", err)
	}
}
