package service

import (
	"AGZ/pkg/repository"
	"AGZ/pkg/structures"
)

type ProfileService struct {
	repo repository.Profile
}

func NewProfileService(repo repository.Profile) *ProfileService {
	return &ProfileService{repo: repo}
}
func (s *ProfileService) AddPurchase(user structures.Params) error {
	return s.repo.AddPurchase(user)
}
func (s *ProfileService) RemovePurchase(user structures.Params) error {
	return s.repo.RemovePurchase(user)
}
func (s *ProfileService) GetBasket(user structures.Params) ([]structures.Purchases, error) {
	return s.repo.GetBasket(user)
}

func (s *ProfileService) GetLinksBasket(user structures.Params) ([]structures.Links, error) {
	return s.repo.GetLinksBasket(user)
}
func (s *ProfileService) AddLink(user structures.Params) error {
	return s.repo.AddLink(user)
}
func (s *ProfileService) RemoveLink(user structures.Params) error {
	return s.repo.RemoveLink(user)
}
