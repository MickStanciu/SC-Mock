package repository

import (
	"github.com/MickStanciu/SC/src/domain"
	"net/http"
)

type fakeDb struct {
}

func (r *fakeDb) GetJediById(id string) (*domain.JediMaster, *domain.ErrorModel) {
	if id == "kenobi-id" {
		return &domain.JediMaster{
			Id:   "kenobi-id",
			Name: "ObiWanKenobi",
		}, nil
	}
	return nil, &domain.ErrorModel{
		Code:    http.StatusNotFound,
		Message: "cannot find ...",
	}
}

func NewFakeDB() JediRepository {
	return &fakeDb{}
}
