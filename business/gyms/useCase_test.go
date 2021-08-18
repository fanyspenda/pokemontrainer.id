package gyms_test

import (
	"context"
	"os"
	gyms "pokemontrainer/business/gyms"
	gymMock "pokemontrainer/business/gyms/mocks"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

var (
	gymRepository gymMock.Repositories
	gymUseCase    gyms.UseCases
)

func setup() {
	gymUseCase = gyms.NewGymUseCases(&gymRepository, 2)
}

func TestMain(m *testing.M) {
	setup()
	os.Exit(m.Run())
}

func TestAddGym(t *testing.T) {
	t.Run("test 1: valid test", func(t *testing.T) {
		var domain = gyms.Domain{
			Name:    "gym1",
			Address: "surabaya",
		}
		gymRepository.On("AddGym", mock.Anything, mock.AnythingOfType("string"), mock.AnythingOfType("string")).Return(domain, nil).Once()
		result, err := gymUseCase.AddGym(context.Background(), domain.Name, domain.Address)

		assert.Nil(t, err)
		assert.Equal(t, domain.Name, result.Name)
		assert.Equal(t, domain.Address, result.Address)
	})

	t.Run("test 2: empty input", func(t *testing.T) {
		gymRepository.On("AddGym", mock.Anything, mock.AnythingOfType("string"), mock.AnythingOfType("string")).Return(gyms.Domain{}, gyms.ErrInvalidInput).Once()
		result, err := gymUseCase.AddGym(context.Background(), "", "")

		assert.Equal(t, err, gyms.ErrInvalidInput)
		assert.Equal(t, gyms.Domain{}, result)
	})
}

func TestUpdateGym(t *testing.T) {
	t.Run("test 1: valid test", func(t *testing.T) {
		var domain = gyms.Domain{
			ID:      1,
			Name:    "gym1",
			Address: "surabaya",
		}
		gymRepository.On("UpdateGym",
			mock.Anything, mock.AnythingOfType("int"), mock.AnythingOfType("string"), mock.AnythingOfType("string")).Return(domain, nil).Once()
		result, err := gymUseCase.UpdateGym(context.Background(), int(domain.ID), domain.Name, domain.Address)

		assert.Nil(t, err)
		assert.Equal(t, domain.ID, result.ID)
		assert.Equal(t, domain.Name, result.Name)
		assert.Equal(t, domain.Address, result.Address)
	})

	t.Run("test 2: empty input", func(t *testing.T) {
		gymRepository.On("UpdateGym",
			mock.Anything, mock.AnythingOfType("int"), mock.AnythingOfType("string"), mock.AnythingOfType("string")).Return(gyms.Domain{}, gyms.ErrInvalidInput).Once()
		result, err := gymUseCase.UpdateGym(context.Background(), 2, "", "")

		assert.Equal(t, err, gyms.ErrInvalidInput)
		assert.Equal(t, gyms.Domain{}, result)
	})

	t.Run("test 3: invalid ID", func(t *testing.T) {
		gymRepository.On("UpdateGym",
			mock.Anything, mock.AnythingOfType("int"), mock.AnythingOfType("string"), mock.AnythingOfType("string")).Return(gyms.Domain{}, gyms.ErrInvalidInput).Once()
		result, err := gymUseCase.UpdateGym(context.Background(), -1, "aaa", "bbbb")

		assert.Equal(t, err, gyms.ErrInvalidID)
		assert.Equal(t, gyms.Domain{}, result)
	})
}

func TestGetGyms(t *testing.T) {
	t.Run("test 1: valid test", func(t *testing.T) {
		var domain = []gyms.Domain{
			{
				ID:      1,
				Name:    "gym1",
				Address: "surabaya",
			},
			{
				ID:      2,
				Name:    "gym2",
				Address: "malang",
			},
		}
		gymRepository.On("GetGyms", mock.Anything).Return(domain, nil).Once()
		result, err := gymUseCase.GetGyms(context.Background())

		assert.Nil(t, err)
		assert.Equal(t, domain[0].ID, result[0].ID)
		assert.Equal(t, domain[0].Name, result[0].Name)
		assert.Equal(t, domain[0].Address, result[0].Address)
		assert.Equal(t, domain[1].ID, result[1].ID)
		assert.Equal(t, domain[1].Name, result[1].Name)
		assert.Equal(t, domain[1].Address, result[1].Address)
	})
}
