package repositories

import (
	"context"

	"github.com/your-org/go-rest-layered-template/internal/domain"
)

type UserRepository interface {
	GetByID(ctx context.Context, id int64) (domain.User, error)
}
