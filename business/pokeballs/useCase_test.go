package pokeballs_test

import (
	"context"
	"os"
	pokeballs "pokemontrainer/business/pokeballs"
	pokeballMock "pokemontrainer/business/pokeballs/mocks"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

var (
	pokeballRepository pokeballMock.Repositories
	pokeballUseCase    pokeballs.UseCases
)

func setup() {
	pokeballUseCase = pokeballs.NewPokeballUseCase(&pokeballRepository, 2)
}

func TestMain(m *testing.M) {
	setup()
	os.Exit(m.Run())
}

func TestAddPokeball(t *testing.T) {
	t.Run("test 1: valid test", func(t *testing.T) {
		var domain = pokeballs.Domain{
			Name:        "greatball",
			SuccessRate: 0.50,
		}
		pokeballRepository.On("AddPokeball", mock.Anything, mock.AnythingOfType("string"), mock.AnythingOfType("float32")).Return(domain, nil).Once()
		result, err := pokeballUseCase.AddPokeball(context.Background(), "greatball", 0.50)

		assert.Nil(t, err)
		assert.Equal(t, domain.Name, result.Name)
	})

	t.Run("test 1: invalid success rate", func(t *testing.T) {
		pokeballRepository.On("AddPokeball", mock.Anything, mock.AnythingOfType("string"), mock.AnythingOfType("float32")).Return(pokeballs.Domain{}, pokeballs.ErrInvalidRate).Once()
		result, err := pokeballUseCase.AddPokeball(context.Background(), "greatball", -1)

		assert.Equal(t, err, pokeballs.ErrInvalidRate)
		assert.Equal(t, pokeballs.Domain{}, result)
		// assert.Equal(t, domain.Name, result.Name)
	})
}
