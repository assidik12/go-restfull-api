package mysql

import (
	"context"
	"database/sql"

	"github.com/assidik12/go-restfull-api/internal/domain"
)

// UserRepository mendefinisikan kontrak untuk interaksi data user.
type UserRepository interface {
	Save(ctx context.Context, user domain.User) (domain.User, error)
	FindByEmail(ctx context.Context, email string) (domain.User, error)
	FindById(ctx context.Context, id int) (domain.User, error)
}

type userRepository struct {
	db *sql.DB
}

// Perhatikan return type-nya sekarang adalah interface
func NewUserRepository(db *sql.DB) UserRepository {
	return &userRepository{db: db}
}
func (r *userRepository) Save(ctx context.Context, user domain.User) (domain.User, error) {
	query := "INSERT INTO users (name, email, password, created_at, updated_at) VALUES (?, ?, ?, ?, ?)"
	result, err := r.db.ExecContext(ctx, query, user.Name, user.Email, user.Password, user.CreatedAt, user.UpdatedAt)
	if err != nil {
		return user, err
	}

	id, err := result.LastInsertId()
	if err != nil {
		return user, err
	}

	user.ID = int(id)
	return user, nil
}

func (r *userRepository) FindByEmail(ctx context.Context, email string) (domain.User, error) {
	var user domain.User
	query := "SELECT id, name, email, password FROM users WHERE email = ?"
	err := r.db.QueryRowContext(ctx, query, email).Scan(&user.ID, &user.Name, &user.Email, &user.Password)
	if err != nil {
		return user, err // Akan mengembalikan sql.ErrNoRows jika tidak ditemukan
	}
	return user, nil
}

// FindById implements UserRepository.
func (r *userRepository) FindById(ctx context.Context, id int) (domain.User, error) {
	var user domain.User
	query := "SELECT id, name, email FROM users WHERE id = ?"
	err := r.db.QueryRowContext(ctx, query, id).Scan(&user.ID, &user.Name, &user.Email, &user.Password)
	if err != nil {
		return user, err // Akan mengembalikan sql.ErrNoRows jika tidak ditemukan
	}
	return user, nil
}
