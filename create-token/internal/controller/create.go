package controller

import (
	"aidanwoods.dev/go-paseto"
	"context"
	"errors"
	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
	"itlg_security/create-token/internal/repository"
	"itlg_security/create-token/pkg/model"
	"time"
)

var (
	ErrNotFound = errors.New("not found")
	ErrInvalidCreds = errors.New("invalid credentials")
)

var key = paseto.NewV4SymmetricKey()

type repo interface {
	GetUser(ctx context.Context, email string) (*model.DdbUser, error)
}

type Controller struct {
	repository repo
}

func New(repository repo) *Controller {
	return &Controller{repository}
}

func (c *Controller) CreateToken(ctx context.Context, email string, password string) (string, error) {
	user, err := c.repository.GetUser(ctx, email)
	if err != nil && errors.Is(err, repository.ErrNotFound) {
		return "", ErrNotFound
	} else if err != nil {
		return "", err
	}
	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))
	if err != nil {
		return "", ErrInvalidCreds
	}
	now := time.Now()
	token := paseto.NewToken()
	token.SetIssuer("itlg-security")
	token.SetSubject(user.ID)
	token.SetAudience("itlg-fe")
	token.SetNotBefore(now)
	token.SetIssuedAt(now)
	token.SetExpiration(now.Add(5 * time.Minute))
	token.SetJti(uuid.New().String())
	encrypted := token.V4Encrypt(key, nil)
	return encrypted, nil
}
