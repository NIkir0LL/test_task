package repository

import (
    "database/sql"
    "user-api/internal/model"
)

type UserRepository struct {
    db *sql.DB
}

func NewUserRepository(db *sql.DB) *UserRepository {
    return &UserRepository{db: db}
}

func (r *UserRepository) Create(user *model.User) error {
    return r.db.QueryRow(
        "INSERT INTO users (name, email) VALUES ($1, $2) RETURNING id, created_at",
        user.Name, user.Email,
    ).Scan(&user.ID, &user.CreatedAt)
}

func (r *UserRepository) Get(id int) (*model.User, error) {
    user := &model.User{}
    err := r.db.QueryRow(
        "SELECT id, name, email, created_at FROM users WHERE id = $1",
        id,
    ).Scan(&user.ID, &user.Name, &user.Email, &user.CreatedAt)
    if err != nil {
        return nil, err
    }
    return user, nil
}

func (r *UserRepository) Update(user *model.User) error {
    _, err := r.db.Exec(
        "UPDATE users SET name = $1, email = $2 WHERE id = $3",
        user.Name, user.Email, user.ID,
    )
    return err
}