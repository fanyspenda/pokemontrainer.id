package trainers

import (
	"context"
	"pokemontrainer/business/trainers"
	"pokemontrainer/drivers/databases/pokemons"
	"pokemontrainer/drivers/thirdParties/pokeapi"

	"gorm.io/gorm"
)

//MysqlTrainerRepository ...
type MysqlTrainerRepository struct {
	Conn *gorm.DB
}

// NewMysqlTrainerRepository menghubungkan business ke database
func NewMysqlTrainerRepository(conn *gorm.DB) trainers.Repository {
	return &MysqlTrainerRepository{
		Conn: conn,
	}
}

// Register as Trainer
func (repo *MysqlTrainerRepository) Register(ctx context.Context, name, address, username, password string) (trainers.Domain, error) {
	var trainerRegister = Trainer{
		Name:     name,
		Address:  address,
		Username: username,
		Password: password,
	}
	result := repo.Conn.Create(&trainerRegister)
	if result.Error != nil {
		return trainers.Domain{}, result.Error
	}
	return trainerRegister.toDomain(), nil
}

// GetTrainers return all trainer data
func (repo *MysqlTrainerRepository) GetTrainers(ctx context.Context) ([]trainers.Domain, error) {
	var trainersCollection = []Trainer{}

	result := repo.Conn.Find(&trainersCollection)

	if result.Error != nil {
		return nil, result.Error
	}

	return ToSliceDomain(trainersCollection), nil
}

// CatchPokemon Add pokemon to relation table
func (repo *MysqlTrainerRepository) CatchPokemon(ctx context.Context, ID, pokemonID int) (trainers.Domain, error) {
	pokeapiStruct := pokeapi.Pokeapi{}

	// check if pokemon id exist
	res, err := pokeapiStruct.GetPokemonByID(ctx, pokemonID)
	if err != nil {
		return trainers.Domain{}, err
	}

	pokemonData := &pokemons.Pokemon{ID: pokemonID, Name: res.Name}
	trainerPokemon := &TrainerPokemon{
		PokemonID: pokemonID,
		TrainerID: ID,
	}

	repo.Conn.Create(&pokemonData)
	repo.Conn.Create(&trainerPokemon)
	return trainers.Domain{}, nil
}

// AddGym register trainer to gym
// func AddGym(ctx context.Context, ID, gymID int) (Domain, error) {
// 	result, err := useCase.Repository.AddGym(ctx, ID, gymID)
// 	if err != nil {
// 		return Domain{}, err
// 	}
// 	return result, nil
// }

// Login as Trainer
func (repo *MysqlTrainerRepository) Login(ctx context.Context, username, password string) (trainers.Domain, error) {
	var trainerLogin = Trainer{}
	result := repo.Conn.Where("username = ? AND password = ?", username, password).First(&trainerLogin)
	if result.Error != nil {
		return trainers.Domain{}, result.Error
	}
	return trainerLogin.toDomain(), nil
}

// CatchPokemon catch pokemon as Trainer
// func CatchPokemon(ctx context.Context, ID, pokemonID int) (Domain, error) {
// 	result, err := useCase.CatchPokemon(ctx, ID, pokemonID)
// 	if err != nil {
// 		return Domain{}, err
// 	}
// 	return result, nil
// }
