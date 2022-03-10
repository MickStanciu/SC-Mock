package service

import (
	"github.com/MickStanciu/SC/src/repository"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestJediService_GetJediById(t *testing.T) {
	jediRepo := repository.NewFakeDB()
	jediService := NewJediService(jediRepo)

	result, err := jediService.GetJediById("ahsoka tano")
	require.Nil(t, result)
	require.NotNil(t, err)
	assert.EqualValues(t, 404, err.Code)
}

func TestJediService_GetJediByIdGoodId(t *testing.T) {
	jediRepo := repository.NewFakeDB()
	jediService := NewJediService(jediRepo)

	result, err := jediService.GetJediById("kenobi-id")
	require.Nil(t, err)
	require.NotNil(t, result)
	assert.EqualValues(t, "kenobi-id", result.Id)
}
