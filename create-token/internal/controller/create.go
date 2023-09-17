package controller

import (
	"context"
	"errors"
	"itlg_security/create-token/pkg/model"
)

var ErrNotFound = errors.New("not found")

type repo interface {
	GetUser(ctx context.Context, email string) (*model.DdbUser, error)
}

type Controller struct {
	repository repo
}

func New(repository repo) *Controller {
	return &Controller{repository}
}

func (*Controller) CreatePaseto(ctx context.Context, email string, password string) (string, error) {
	return "", nil
}
