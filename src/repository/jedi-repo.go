package repository

import "github.com/MickStanciu/SC/src/domain"

type JediRepository interface {
	GetJediById(string) (*domain.JediMaster, *domain.ErrorModel)
}
