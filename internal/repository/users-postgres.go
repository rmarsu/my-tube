package repository

import (
	"context"
	"fmt"
	"myTube/internal/models"

	"github.com/jackc/pgx/v5"
)

type UsersRepo struct {
	db *pgx.Conn
}

func NewUserRepository(db *pgx.Conn) *UsersRepo {
     return &UsersRepo{db: db}
}

func (r *UsersRepo) Create(user models.User) error {
	query := `INSERT INTO users (username, email, password, created_at, videos) VALUES ($1, $2, $3, $4) RETURNING id`
     err := r.db.QueryRow(context.Background(), query, user.Username, user.Email, user.Password, user.CreatedAt, user.Videos).Scan(&user.ID)
     if err!= nil {
          return err
     }
     return nil
}

func (r *UsersRepo) GetByUsername(ctx context.Context, username string) (models.User, error) {
	query := `SELECT id, username, email, password, created_at, videos FROM users WHERE username = $1`
     var user models.User
     err := r.db.QueryRow(ctx, query, username).Scan(&user.ID, &user.Username, &user.Email, &user.Password, &user.CreatedAt, &user.Videos)
     if err == pgx.ErrNoRows {
          return user, fmt.Errorf("user not found")
     } else if err!= nil {
          return user, err
     } 
     return user, nil
}

func (r *UsersRepo) GetByCredentials(ctx context.Context, email, password string) (models.User, error) {
	query := `SELECT id, username, email, password, created_at , videos FROM users WHERE email = $1 AND password = $2`
     var user models.User
     err := r.db.QueryRow(ctx, query, email, password).Scan(&user.ID, &user.Username, &user.Email, &user.Password, &user.CreatedAt, &user.Videos)
     if err == pgx.ErrNoRows {
          return user, fmt.Errorf("user not found")
     } else if err!= nil {
          return user, err
     } 
     return user, nil
}

func (r *UsersRepo) GetByRefreshToken(ctx context.Context, refreshToken string) (models.User, error) {
	query := `SELECT id, username, email, password, created_at ,videos FROM users WHERE refresh_token = $1`
     var user models.User
     err := r.db.QueryRow(ctx, query, refreshToken).Scan(&user.ID, &user.Username, &user.Email,
										&user.Password, &user.CreatedAt, &user.Videos)
     if err == pgx.ErrNoRows {
          return user, fmt.Errorf("user not found")
     } else if err!= nil {
          return user, err
     } 
     return user, nil
}

func (r *UsersRepo) Update(ctx context.Context, user models.User) error {
	query := `UPDATE users SET username = $1, email = $2, password = $3, videos = $4 WHERE id = $5`
     _, err := r.db.Exec(ctx, query, user.Username, user.Email, user.Password, user.Videos, user.ID)
     if err!= nil {
          return err
     }
     return nil
}

func (r *UsersRepo) Delete(ctx context.Context, id int) error {
	query := `DELETE FROM users WHERE id = $1`
     _, err := r.db.Exec(ctx, query, id)
     if err!= nil {
          return err
     }
     return nil
}