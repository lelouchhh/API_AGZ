package service

import (
	"AGZ/pkg/repository"
	"AGZ/pkg/structures"
)

type AuthService struct {
	repo repository.Authorization
}

func NewAuthService(repo repository.Authorization) *AuthService {
	return &AuthService{repo: repo}
}
func (s *AuthService) RegisterUserEmail(user structures.Params) error {
	return s.repo.RegisterUserEmail(user)
}
func (s *AuthService) ConfirmUserEmail(user structures.Params) error {
	return s.repo.ConfirmUserEmail(user)
}
func (s *AuthService) ResetUserEmailPass(user structures.Params) error {
	return s.repo.ResetUserEmailPass(user)
}

func (s *AuthService) SignIn(user structures.Params) (structures.Tokens, error) {
	return s.repo.SignIn(user)
}
func (s *AuthService) ActionUserAccess(token structures.Tokens) error {
	return s.repo.ActionUserAccess(token)
}

func (s *AuthService) ActionUserRefresh(token structures.Tokens) (structures.Tokens, error) {
	return s.repo.ActionUserRefresh(token)
}
