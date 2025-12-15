package mysqlrepo

import (
	"context"
	"database/sql"
	"errors"

	"github.com/your-org/go-rest-layered-template/internal/domain"
)

var ErrNotFound = errors.New("user not found")

type UserRepository struct {
	db *sql.DB
}

func NewUserRepository(db *sql.DB) *UserRepository {
	return &UserRepository{db: db}
}

func (r *UserRepository) GetByID(ctx context.Context, id int64) (domain.User, error) {
	const q = `
SELECT id, email, name, created_at
FROM users
WHERE id = ?
LIMIT 1`

	var u domain.User
	row := r.db.QueryRowContext(ctx, q, id)
	if err := row.Scan(&u.ID, &u.Email, &u.Name, &u.CreatedAt); err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return domain.User{}, ErrNotFound
		}
		return domain.User{}, err
	}
	return u, nil
}
