package service

import (
	"github.com/MickStanciu/SC/src/domain"
	"github.com/MickStanciu/SC/src/repository"
)

type JediService interface {
	GetJediById(string) (*domain.JediMaster, *domain.ErrorModel)
}

type jediService struct {
	repository repository.JediRepository
}

func (s *jediService) GetJediById(id string) (*domain.JediMaster, *domain.ErrorModel) {
	return s.repository.GetJediById(id)
}

func NewJediService(r repository.JediRepository) JediService {
	return &jediService{
		repository: r,
	}
}
