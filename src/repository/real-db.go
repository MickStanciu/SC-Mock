package repository

import (
	"github.com/MickStanciu/SC/src/domain"
	"net/http"
)

type realDb struct{}

func (r *realDb) GetJediById(string) (*domain.JediMaster, *domain.ErrorModel) {
	return nil, &domain.ErrorModel{
		Code:    http.StatusNotImplemented,
		Message: "computers says no",
	}
}

func NewRealDB() JediRepository {
	return &realDb{}
}
