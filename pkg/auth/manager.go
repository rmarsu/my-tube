package auth

import (
	"fmt"
	"math/rand"
	"time"

	"github.com/VandiKond/vanerrors"
	"github.com/golang-jwt/jwt/v5"
)

type MyCustomClaims struct {
	Id string `json:"id"`
	jwt.RegisteredClaims
}

var (
	ErrInvalidCredentials = vanerrors.NewName("invalid credentials", vanerrors.EmptyHandler)
)

type TokenManager interface {
	NewJWT(userId string, ttl time.Duration) (string, error)
	Parse(accessToken string) (string, error)
	NewRefreshToken() (string, error)
}

type Manager struct {
	signingKey string
}

func NewManager(signingKey string) (*Manager, error) {
	if signingKey == "" {
		return nil, vanerrors.NewName("empty signing key", vanerrors.EmptyHandler)
	}

	return &Manager{signingKey: signingKey}, nil
}

func (m *Manager) NewJWT(userId string, ttl time.Duration) (string, error) {
	claims := MyCustomClaims{
		userId,
		jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(ttl)),
			Subject:   userId,
			IssuedAt:  jwt.NewNumericDate(time.Now()),
			Issuer:    "myTube",
			Audience:  []string{"*"},
			NotBefore: jwt.NewNumericDate(time.Now()),
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	result, err := token.SignedString([]byte(m.signingKey))
	if err != nil {
		err = vanerrors.NewWrap("error to get singed string", err, vanerrors.EmptyHandler)
		return result, err
	}
	return result, nil
}

func (m *Manager) Parse(accessToken string) (string, error) {
	token, err := jwt.Parse(accessToken, func(token *jwt.Token) (i interface{}, err error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, vanerrors.NewBasic("unexpected signing method", fmt.Sprint(token.Header["alg"]), vanerrors.EmptyHandler)
		}

		return []byte(m.signingKey), nil
	})
	if err != nil {
		return "", err
	}

	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok {
		return "", vanerrors.NewName("error get user claims from token", vanerrors.EmptyHandler)
	}

	return claims["sub"].(string), nil
}

func (m *Manager) NewRefreshToken() (string, error) {
	b := make([]byte, 32)

	s := rand.NewSource(time.Now().Unix())
	r := rand.New(s)

	if _, err := r.Read(b); err != nil {
		err = vanerrors.NewWrap("error to read new token", err, vanerrors.EmptyHandler)
		return "", err
	}

	return fmt.Sprintf("%x", b), nil
}

func (m *Manager) ExtractUserIdFromToken(accessToken string) (string, error) {
	userId, err := m.Parse(accessToken)
	if err != nil {
		err = vanerrors.NewWrap("error to parse the access token", err, vanerrors.EmptyHandler)
		return "", err
	}

	return userId, nil
}
