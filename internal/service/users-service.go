package service

import (
	"context"
	"myTube/internal/models"
	"myTube/internal/repository"
	"myTube/pkg/auth"
	"myTube/pkg/hash"
	"myTube/pkg/log"
	"strconv"
	"time"

	"github.com/VandiKond/vanerrors"
)

type UsersService struct {
	repo         repository.Users
	hasher       hash.PasswordHasher
	tokenManager auth.TokenManager

	accessTokenTTL  time.Duration
	refreshTokenTTL time.Duration
}

func NewUsersService(repo repository.Users, hasher hash.PasswordHasher, tokenManager auth.TokenManager,
	accessTTL, refreshTTL time.Duration) *UsersService {
	return &UsersService{
		repo:            repo,
		hasher:          hasher,
		tokenManager:    tokenManager,
		accessTokenTTL:  accessTTL,
		refreshTokenTTL: refreshTTL,
	}
}

func (s *UsersService) SignUp(ctx context.Context, input UserSignUpInput) (Tokens, error) {
	passwordHash, err := s.hasher.Hash(input.Password)
	if err != nil {
		err = vanerrors.NewWrap("error to hash password", err, vanerrors.EmptyHandler)
		return Tokens{}, err
	}
	user := models.User{
		Username:  input.Username,
		Email:     input.Email,
		Password:  passwordHash,
		CreatedAt: time.Now(),
	}
	err = s.repo.Create(user)
	if err != nil {
		err = vanerrors.NewWrap("error to create a user", err, vanerrors.EmptyHandler)
		return Tokens{}, err
	}

	newUser, err := s.repo.GetByUsername(ctx, input.Username)
	if err != nil {
		err = vanerrors.NewWrap("error to get by user by username", err, vanerrors.EmptyHandler)
		return Tokens{}, err
	}
	result, err := s.createSession(newUser.ID)
	if err != nil {
		err = vanerrors.NewWrap("error to create a session", err, vanerrors.EmptyHandler)
		return Tokens{}, err
	}
	return result, nil
}

func (s *UsersService) SignIn(ctx context.Context, input UserSignInInput) (Tokens, error) {
	passwordHash, err := s.hasher.Hash(input.Password)
	log.Debug(passwordHash)
	if err != nil {
		err = vanerrors.NewWrap("error to sash password", err, vanerrors.EmptyHandler)
		return Tokens{}, err
	}
	user, err := s.repo.GetByUsername(ctx, input.Username)
	if err != nil {
		err = vanerrors.NewWrap("error to get user by username", err, vanerrors.EmptyHandler)
		return Tokens{}, err
	}
	if passwordHash == user.Password {
		result, err := s.createSession(user.ID)
		if err != nil {
			err = vanerrors.NewWrap("error to create a session", err, vanerrors.EmptyHandler)
			return Tokens{}, err
		}
		return result, nil
	}
	return Tokens{}, auth.ErrInvalidCredentials
}

func (s *UsersService) createSession(userID int) (Tokens, error) {
	accessToken, err := s.tokenManager.NewJWT(strconv.Itoa(userID), s.accessTokenTTL)
	if err != nil {
		err = vanerrors.NewWrap("error to get the access token", err, vanerrors.EmptyHandler)
		return Tokens{}, err
	}
	refreshToken, err := s.tokenManager.NewRefreshToken()
	if err != nil {
		err = vanerrors.NewWrap("error to refresh token", err, vanerrors.EmptyHandler)
		return Tokens{}, err
	}
	return Tokens{
		AccessToken:  accessToken,
		RefreshToken: refreshToken,
	}, nil
}

func (s *UsersService) GetUserIdFromToken(ctx context.Context, token string) (string, error) {
	userID, err := s.tokenManager.Parse(token)
	if err != nil {
		err = vanerrors.NewWrap("error to get user id", err, vanerrors.EmptyHandler)
		return "", err
	}
	return userID, nil
}
