package trainers_test

import (
	"context"
	"os"
	"pokemontrainer/business/pokeballs"
	trainers "pokemontrainer/business/trainers"
	trainerMock "pokemontrainer/business/trainers/mocks"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

var (
	trainerRepository      trainerMock.Repository
	trainerUseCase         trainers.UseCase
	trainerMongoRepository trainerMock.MongodbRepository
)

func setup() {
	trainerUseCase = trainers.NewTrainerUseCase(&trainerRepository, 2, &trainerMongoRepository)
}

func TestMain(m *testing.M) {
	setup()
	os.Exit(m.Run())
}

func TestRegister(t *testing.T) {
	t.Run("test 1: valid test", func(t *testing.T) {
		var domain = trainers.Domain{
			Name:     "name1",
			Username: "username1",
			Address:  "address1",
			Password: "password1",
		}
		trainerRepository.On("Register", mock.Anything, mock.AnythingOfType("string"), mock.AnythingOfType("string"),
			mock.AnythingOfType("string"), mock.AnythingOfType("string")).Return(domain, nil).Once()

		result, err := trainerUseCase.Register(context.Background(), domain.Name, domain.Address, domain.Username, domain.Password)

		assert.Nil(t, err)
		assert.Equal(t, domain.Name, result.Name)
		assert.Equal(t, domain.Address, result.Address)
		assert.Equal(t, domain.Password, result.Password)
		assert.Equal(t, domain.Username, result.Username)
	})

	t.Run("test 2: invalid input", func(t *testing.T) {
		trainerRepository.On("Register", mock.Anything, mock.AnythingOfType("string"), mock.AnythingOfType("string"),
			mock.AnythingOfType("string"), mock.AnythingOfType("string")).Return(trainers.Domain{}, trainers.ErrInvalidInput).Once()
		result, err := trainerUseCase.Register(context.Background(), "", "", "", "")

		assert.Equal(t, err, trainers.ErrInvalidInput)
		assert.Equal(t, trainers.Domain{}, result)
		// assert.Equal(t, domain.Name, result.Name)
	})
}

func TestAddGym(t *testing.T) {
	t.Run("test 1: valid test", func(t *testing.T) {
		var domain = trainers.Domain{
			ID:       1,
			Name:     "name1",
			Username: "username1",
			Address:  "address1",
		}
		trainerRepository.On("AddGym", mock.Anything, mock.AnythingOfType("int"), mock.AnythingOfType("int")).Return(domain, nil).Once()
		result, err := trainerUseCase.AddGym(context.Background(), 1, 1)

		assert.Nil(t, err)
		assert.Equal(t, domain.ID, result.ID)
	})

	t.Run("test 2: invalid id", func(t *testing.T) {
		trainerRepository.On("AddGym", mock.Anything, mock.AnythingOfType("int"), mock.AnythingOfType("int")).Return(trainers.Domain{}, trainers.ErrInvalidInput).Once()
		result, err := trainerUseCase.AddGym(context.Background(), -1, -1)

		assert.Equal(t, err, trainers.ErrInvalidInput)
		assert.Equal(t, trainers.Domain{}, result)
	})
}

func TestGetFirstBall(t *testing.T) {
	t.Run("test 1: valid test", func(t *testing.T) {
		var domain = pokeballs.Domain{
			ID:          1,
			Name:        "pokeball",
			SuccessRate: 0.51,
		}
		trainerRepository.On("GetFirstBall", mock.Anything, mock.AnythingOfType("uint")).Return(domain, nil).Once()
		result, err := trainerUseCase.GetFirstBall(context.Background(), 1)

		assert.Nil(t, err)
		assert.Equal(t, domain.ID, result.ID)
	})

	t.Run("test 2: invalid id", func(t *testing.T) {
		trainerRepository.On("GetFirstBall", mock.Anything, mock.AnythingOfType("uint")).Return(pokeballs.Domain{}, trainers.ErrInvalidInput).Once()
		result, err := trainerUseCase.GetFirstBall(context.Background(), 0)

		assert.Equal(t, err, trainers.ErrInvalidInput)
		assert.Equal(t, pokeballs.Domain{}, result)
	})
}

func TestGetTrainers(t *testing.T) {
	t.Run("test 1: valid test", func(t *testing.T) {
		var domain = []trainers.Domain{
			{
				ID:       1,
				Name:     "name1",
				Username: "username1",
				Address:  "address1",
			},
			{ID: 2,
				Name:     "name2",
				Username: "username2",
				Address:  "address2",
			},
		}
		trainerRepository.On("GetTrainers", mock.Anything).Return(domain, nil).Once()
		result, err := trainerUseCase.GetTrainers(context.Background())

		assert.Nil(t, err)
		assert.Equal(t, domain[0].ID, result[0].ID)
		assert.Equal(t, domain[1].ID, result[1].ID)
	})

	t.Run("test 1: valid test", func(t *testing.T) {

		trainerRepository.On("GetTrainers", mock.Anything).Return([]trainers.Domain{}, trainers.ErrGetTrainerFailed).Once()
		result, err := trainerUseCase.GetTrainers(context.Background())

		assert.Equal(t, []trainers.Domain{}, result)
		assert.Equal(t, trainers.ErrGetTrainerFailed, err)
	})
}

func TestLogin(t *testing.T) {
	var domain = trainers.Domain{
		ID:       1,
		Name:     "name1",
		Username: "username1",
		Address:  "address1",
		Token:    "token1",
	}
	t.Run("test 1: valid test", func(t *testing.T) {
		trainerMongoRepository.On("LoginLog", mock.Anything, mock.AnythingOfType("uint")).Return(domain, nil).Once()
		trainerRepository.On("Login", mock.Anything, mock.AnythingOfType("string"), mock.AnythingOfType("string")).Return(domain, nil).Once()
		result, err := trainerUseCase.Login(context.Background(), "username1", "password1")
		// result, err = trainerMongoRepository.LoginLog(context.Background(), 1)

		assert.NotEmpty(t, result.Token)
		assert.Equal(t, domain.ID, result.ID)
		assert.Nil(t, err)
	})

	t.Run("test 2: invalid Input", func(t *testing.T) {
		trainerMongoRepository.On("LoginLog", mock.Anything, mock.AnythingOfType("uint")).Return(domain, nil).Once()
		trainerRepository.On("Login", mock.Anything, mock.AnythingOfType("string"), mock.AnythingOfType("string")).Return(trainers.Domain{}, trainers.ErrInvalidInput).Once()
		result, err := trainerUseCase.Login(context.Background(), "", "")

		assert.Equal(t, trainers.ErrInvalidInput, err)
		assert.Equal(t, trainers.Domain{}, result)
	})
}

func TestCatchPokemon(t *testing.T) {
	var domain = trainers.Domain{
		ID:       1,
		Name:     "name1",
		Username: "username1",
		Address:  "address1",
		Token:    "token1",
	}
	t.Run("test 1: valid test", func(t *testing.T) {
		trainerRepository.On("CatchPokemon", mock.Anything, mock.AnythingOfType("int"), mock.AnythingOfType("int")).Return(domain, nil).Once()
		result, err := trainerUseCase.CatchPokemon(context.Background(), 1, 1)
		// result, err = trainerMongoRepository.LoginLog(context.Background(), 1)

		assert.Equal(t, domain.ID, result.ID)
		assert.Nil(t, err)
	})

	t.Run("test 2: invalid Input", func(t *testing.T) {
		trainerRepository.On("CatchPokemon", mock.Anything, mock.AnythingOfType("int"), mock.AnythingOfType("int")).Return(domain, nil).Once()
		result, err := trainerUseCase.CatchPokemon(context.Background(), 0, 0)

		assert.Equal(t, trainers.ErrInvalidInput, err)
		assert.Equal(t, trainers.Domain{}, result)
	})
}

func TestUpdateTrainer(t *testing.T) {
	var domain = trainers.Domain{
		ID:       1,
		Name:     "name1",
		Username: "username1",
		Address:  "address1",
		Token:    "token1",
	}
	t.Run("test 1: valid test", func(t *testing.T) {
		trainerRepository.On("UpdateTrainer",
			mock.Anything,
			mock.AnythingOfType("int"),
			mock.AnythingOfType("string"),
			mock.AnythingOfType("string"),
			mock.AnythingOfType("string"),
			mock.AnythingOfType("string")).Return(domain, nil).Once()
		result, err := trainerUseCase.UpdateTrainer(context.Background(), 1, domain.Name, domain.Address, domain.Username, "aaaaa")
		// result, err = trainerMongoRepository.LoginLog(context.Background(), 1)

		assert.Equal(t, domain.ID, result.ID)
		assert.Nil(t, err)
	})

	t.Run("test 2: invalid Input", func(t *testing.T) {
		trainerRepository.On("UpdateTrainer",
			mock.Anything,
			mock.AnythingOfType("int"),
			mock.AnythingOfType("string"),
			mock.AnythingOfType("string"),
			mock.AnythingOfType("string"),
			mock.AnythingOfType("string")).Return(domain, nil).Once()
		result, err := trainerUseCase.UpdateTrainer(context.Background(), 12, "", "", "", "")

		assert.Equal(t, trainers.ErrInvalidInput, err)
		assert.Equal(t, trainers.Domain{}, result)
	})

	t.Run("test 2: invalid ID", func(t *testing.T) {
		trainerRepository.On("UpdateTrainer",
			mock.Anything,
			mock.AnythingOfType("int"),
			mock.AnythingOfType("string"),
			mock.AnythingOfType("string"),
			mock.AnythingOfType("string"),
			mock.AnythingOfType("string")).Return(domain, nil).Once()
		result, err := trainerUseCase.UpdateTrainer(context.Background(), -12, "aaaaa", "bbbbb", "ccccc", "ddddd")

		assert.Equal(t, trainers.ErrInvalidID, err)
		assert.Equal(t, trainers.Domain{}, result)
	})
}
