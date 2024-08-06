package service

import (
	"makves/internal/entity"
)

type Repository interface {
	GetByIds(ids []int) ([]entity.Item, error)
}

type Service struct {
	repo Repository
}

func (s *Service) GetItems(ids []int) ([]entity.Item, error) {
	return s.repo.GetByIds(ids)
}

func NewService(repo Repository) *Service {
	return &Service{repo: repo}
}
