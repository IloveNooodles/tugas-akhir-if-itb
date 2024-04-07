package user

import (
	"context"

	"github.com/google/uuid"
	"github.com/jmoiron/sqlx"
	"github.com/sirupsen/logrus"
)

type Repository struct {
	Logger *logrus.Logger
	DB     *sqlx.DB
}

func NewRepository(db *sqlx.DB, logger *logrus.Logger) *Repository {
	return &Repository{
		DB:     db,
		Logger: logger,
	}
}

func (r *Repository) Create(ctx context.Context, u User) (User, error) {
	user := User{}
	q := `INSERT INTO users (name, email, password) VALUES ($1, $2, $3) RETURNING *`
	err := r.DB.GetContext(ctx, &user, q, u.Name, u.Email, u.Password)

	if err != nil {
		r.Logger.Errorf("error when creating users %v, err: %s", user, err)
		return user, err
	}

	return user, nil
}

func (r *Repository) GetByID(ctx context.Context, id uuid.UUID) (User, error) {
	user := User{}
	q := `SELECT * FROM users WHERE id = $1`
	err := r.DB.GetContext(ctx, &user, q, id)

	if err != nil {
		r.Logger.Errorf("error when get users with id: %s, err: %s", id, err)
		return user, err
	}

	return user, nil
}

func (r *Repository) GetByEmail(ctx context.Context, email string) (User, error) {
	user := User{}
	q := `SELECT * FROM users WHERE email = $1`
	err := r.DB.GetContext(ctx, &user, q, email)

	if err != nil {
		r.Logger.Errorf("error when get users with id: %s, err: %s", email, err)
		return user, err
	}

	return user, nil
}
