package services

import (
	"context"
	"errors"

	"github.com/your-org/go-rest-layered-template/internal/domain"
	"github.com/your-org/go-rest-layered-template/internal/repositories"
	"github.com/your-org/go-rest-layered-template/internal/repositories/mysqlrepo"
)

var ErrUserNotFound = errors.New("user not found")

type UserService struct {
	repo repositories.UserRepository
}

func NewUserService(repo repositories.UserRepository) *UserService {
	return &UserService{repo: repo}
}

func (s *UserService) GetByID(ctx context.Context, id int64) (domain.User, error) {
	u, err := s.repo.GetByID(ctx, id)
	if err != nil {
		if errors.Is(err, mysqlrepo.ErrNotFound) {
			return domain.User{}, ErrUserNotFound
		}
		return domain.User{}, err
	}
	return u, nil
}
