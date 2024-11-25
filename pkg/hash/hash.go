package hash

import (
	"crypto/sha256"
	"fmt"

	"github.com/VandiKond/vanerrors"
)

type PasswordHasher interface {
	Hash(password string) (string, error)
}

type SHA256Hasher struct {
	salt string
}

func NewSHA256Hasher(salt string) *SHA256Hasher {
	return &SHA256Hasher{salt: salt}
}

func (h *SHA256Hasher) Hash(password string) (string, error) {
	hash := sha256.New()

	if _, err := hash.Write([]byte(password)); err != nil {
		err = vanerrors.NewWrap("error to create hash", err, vanerrors.EmptyHandler)
		return "", err
	}

	return fmt.Sprintf("%x", hash.Sum([]byte(h.salt))), nil
}
