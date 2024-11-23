package service

import (
	"myTube/internal/repository"
	"myTube/pkg/auth"
	"myTube/pkg/hash"
	"time"
)

type VideoInput struct {
	Title       string
	Description string
	Duration    int64
}

type UserSignUpInput struct {
	Username string
	Email    string
	Password string
}

type UserSignInInput struct {
	Username string
	Password string
}

type Tokens struct {
	AccessToken  string
	RefreshToken string
}

type Deps struct {
	Repos           *repository.Repositories
	Hasher          hash.PasswordHasher
	TokenManager    auth.TokenManager
	AccessTokenTTL  time.Duration
	RefreshTokenTTL time.Duration
	// Add other dependencies as needed
}

type Services struct {
	// Add service methods here
	VideoService *VideoService
	Users        *UsersService
}

func NewServices(deps *Deps) *Services {
	return &Services{
		VideoService: NewVideoService(deps.Repos.Videos),
		Users: NewUsersService(deps.Repos.Users, deps.Hasher, deps.TokenManager,
			deps.AccessTokenTTL, deps.RefreshTokenTTL),
	}
}
