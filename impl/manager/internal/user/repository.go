package user

import (
	"context"

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
		r.Logger.Errorf("error when creating users detail: %v, err: %s", user, err)
		return user, err
	}

	return user, nil
}

func (r *Repository) GetByID() (User, error) {
	user := User{}
	return user, nil
}

func (r *Repository) GetByEmail() (User, error) {
	user := User{}
	return user, nil
}
