package repository

import (
	"context"
	"myTube/internal/models"

	"github.com/VandiKond/vanerrors"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
)

type UsersRepo struct {
	db *pgxpool.Pool
}

func NewUserRepository(db *pgxpool.Pool) *UsersRepo {
	return &UsersRepo{db: db}
}

func (r *UsersRepo) Create(user models.User) error {
	query := `INSERT INTO users (username, email, password, created_at) VALUES ($1, $2, $3, $4) RETURNING id`
	err := r.db.QueryRow(context.Background(), query, user.Username, user.Email, user.Password, user.CreatedAt).Scan(&user.ID)
	if err != nil {
		err = vanerrors.NewWrap("unable to create user", err, vanerrors.EmptyHandler)
		return err
	}
	return nil
}

func (r *UsersRepo) GetByUsername(ctx context.Context, username string) (models.User, error) {
	query := `SELECT id, username, email, password, created_at FROM users WHERE username = $1`
	var user models.User
	err := r.db.QueryRow(ctx, query, username).Scan(&user.ID, &user.Username, &user.Email, &user.Password, &user.CreatedAt)
	if err == pgx.ErrNoRows {
		return user, vanerrors.NewName("user not found", vanerrors.EmptyHandler)
	} else if err != nil {
		err = vanerrors.NewWrap("error to get user", err, vanerrors.EmptyHandler)
		return user, err
	}
	return user, nil
}

func (r *UsersRepo) GetByCredentials(ctx context.Context, email, password string) (models.User, error) {
	query := `SELECT id, username, email, password, created_at FROM users WHERE email = $1 AND password = $2`
	var user models.User
	err := r.db.QueryRow(ctx, query, email, password).Scan(&user.ID, &user.Username, &user.Email, &user.Password, &user.CreatedAt)
	if err == pgx.ErrNoRows {
		return user, vanerrors.NewName("user not found", vanerrors.EmptyHandler)
	} else if err != nil {
		err = vanerrors.NewWrap("error to get user", err, vanerrors.EmptyHandler)
		return user, err
	}
	return user, nil
}

func (r *UsersRepo) GetByRefreshToken(ctx context.Context, refreshToken string) (models.User, error) {
	query := `SELECT id, username, email, password, created_at FROM users WHERE refresh_token = $1`
	var user models.User
	err := r.db.QueryRow(ctx, query, refreshToken).Scan(&user.ID, &user.Username, &user.Email,
		&user.Password, &user.CreatedAt)
	if err == pgx.ErrNoRows {
		return user, vanerrors.NewName("user not found", vanerrors.EmptyHandler)
	} else if err != nil {
		err = vanerrors.NewWrap("error to get user", err, vanerrors.EmptyHandler)
		return user, err
	}
	return user, nil
}

func (r *UsersRepo) Update(ctx context.Context, user models.User) error {
	query := `UPDATE users SET username = $1, email = $2, password = $3 WHERE id = $4`
	_, err := r.db.Exec(ctx, query, user.Username, user.Email, user.Password, user.ID)
	if err != nil {
		err = vanerrors.NewWrap("error to update user", err, vanerrors.EmptyHandler)
		return err
	}
	return nil
}

func (r *UsersRepo) Delete(ctx context.Context, id int) error {
	query := `DELETE FROM users WHERE id = $1`
	_, err := r.db.Exec(ctx, query, id)
	if err != nil {
		err = vanerrors.NewWrap("error to delete user", err, vanerrors.EmptyHandler)
		return err
	}
	return nil
}
