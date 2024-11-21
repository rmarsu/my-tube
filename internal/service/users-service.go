package service

import (
	"context"
	"myTube/internal/models"
	"myTube/internal/repository"
	"myTube/pkg/auth"
	"myTube/pkg/hash"
	"strconv"
	"time"
)

type UsersService struct {
	repo repository.Users
	hasher hash.PasswordHasher	
	tokenManager auth.TokenManager

	accessTokenTTL time.Duration
	refreshTokenTTL time.Duration
}

func NewUsersService(repo repository.Users, hasher hash.PasswordHasher, tokenManager auth.TokenManager, 
	accessTTL, refreshTTL time.Duration) *UsersService {
     return &UsersService{
		repo: repo, 
		hasher: hasher,
		tokenManager: tokenManager,
		accessTokenTTL: accessTTL,
          refreshTokenTTL: refreshTTL,
	}
}

func (s *UsersService) SignUp(ctx context.Context , input UserSignUpInput) (Tokens ,error) {
	passwordHash , err := s.hasher.Hash(input.Password)
	if err != nil {
          return Tokens{} ,err
     }
	user := models.User{
		ID:        0, 
          Username: input.Username,
          Email:    input.Email,
          Password: passwordHash,
		CreatedAt: time.Now(),
		Videos:    []models.Video{},
     }
	err = s.repo.Create(user)
	if err!= nil {
          return Tokens{}, err
     }
	return s.createSession(user.ID)
}

func (s *UsersService) SignIn(ctx context.Context, input UserSignInInput) (Tokens, error) {
	passwordHash, err := s.hasher.Hash(input.Password)
	if err != nil {
          return Tokens{}, err
     }
	user, err := s.repo.GetByUsername(ctx, input.Username)
	if err!= nil {
          return Tokens{}, err
     }
	if passwordHash == input.Password {
		return s.createSession(user.ID)
     }
     return Tokens{}, auth.ErrInvalidCredentials
}

func (s *UsersService) createSession(userID int) (Tokens, error) {
	var (
		res Tokens
		err error
	)
	res.AccessToken , err = s.tokenManager.NewJWT(strconv.Itoa(userID), s.accessTokenTTL)
	if err!= nil {
          return res, err
     }
	res.RefreshToken , err = s.tokenManager.NewRefreshToken()
	if err!= nil {
          return res, err
     }
	return res, nil
}