package mysqlrepo_test

import (
	"context"
	"database/sql"
	"errors"
	"regexp"
	"testing"
	"time"

	"github.com/your-org/go-rest-layered-template/internal/repositories/mysqlrepo"

	"github.com/DATA-DOG/go-sqlmock"
)

func TestUserRepository_GetByID_OK(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("sqlmock.New: %v", err)
	}
	defer func() { _ = db.Close() }()

	repo := mysqlrepo.NewUserRepository(db)

	createdAt := time.Date(2024, 1, 1, 0, 0, 0, 0, time.UTC)
	rows := sqlmock.NewRows([]string{"id", "email", "name", "created_at"}).
		AddRow(int64(1), "jane@example.com", "Jane", createdAt)

	mock.ExpectQuery(regexp.QuoteMeta(`
SELECT id, email, name, created_at
FROM users
WHERE id = ?
LIMIT 1`)).
		WithArgs(int64(1)).
		WillReturnRows(rows)

	u, err := repo.GetByID(context.Background(), 1)
	if err != nil {
		t.Fatalf("expected nil error, got %v", err)
	}
	if u.ID != 1 || u.Email != "jane@example.com" || u.Name != "Jane" || !u.CreatedAt.Equal(createdAt) {
		t.Fatalf("unexpected user: %#v", u)
	}

	if err := mock.ExpectationsWereMet(); err != nil {
		t.Fatalf("unmet expectations: %v", err)
	}
}

func TestUserRepository_GetByID_NotFound(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("sqlmock.New: %v", err)
	}
	defer func() { _ = db.Close() }()

	repo := mysqlrepo.NewUserRepository(db)

	mock.ExpectQuery(regexp.QuoteMeta(`
SELECT id, email, name, created_at
FROM users
WHERE id = ?
LIMIT 1`)).
		WithArgs(int64(999)).
		WillReturnError(sql.ErrNoRows)

	_, err = repo.GetByID(context.Background(), 999)
	if !errors.Is(err, mysqlrepo.ErrNotFound) {
		t.Fatalf("expected ErrNotFound, got %v", err)
	}

	_ = mock.ExpectationsWereMet()
}
