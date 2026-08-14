package service

import (
	"context"

	"github.com/cerque1/tutor-website-001/internal/dto"
	"github.com/cerque1/tutor-website-001/internal/security"
)

type AuthService struct{
	userRepo UserRepository
	secret []byte
}

func NewAuthService(userRepo UserRepository, secret string) *AuthService {
	return &AuthService{
		userRepo: userRepo,
		secret: []byte(secret),
	}
}

func (s *AuthService) Login(ctx context.Context, req dto.Login) (string, error) {
	user, err := s.userRepo.GetByEmail(ctx, req.Email)
	if err != nil {
		return "", err 
	}
	err = security.PasswordChecker(req.Password, user.PasswordHash)
	if err != nil {
		return "", err
	}

	token, err := security.CreateToken(user.ID, user.IsAdmin, s.secret)
	if err != nil {
		return "", err
	}

	return token, nil
}
