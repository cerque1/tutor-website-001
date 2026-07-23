package service

import "context"

type UserRepository interface {
	GetAll(ctx context.Context)
}
