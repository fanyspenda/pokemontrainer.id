package trainers

import (
	"context"
	pokeballBusiness "pokemontrainer/business/pokeballs"
	"pokemontrainer/business/trainers"
	pokeballDatabase "pokemontrainer/drivers/databases/pokeballs"
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

// GetFirstBall add 1 ball to join table trainer-pokeball
func (repo *MysqlTrainerRepository) GetFirstBall(ctx context.Context, trainerID uint) (pokeballBusiness.Domain, error) {

	var lowestPokeball = pokeballDatabase.Pokeball{}
	result := repo.Conn.Order("success_rate asc").First(&lowestPokeball)
	if result.Error != nil {
		return pokeballBusiness.Domain{}, result.Error
	}
	var joinTableData = TrainerPokeballs{
		TrainerID:  trainerID,
		PokeballID: lowestPokeball.ID,
		Quantity:   100,
	}

	result = repo.Conn.Create(&joinTableData)
	if result.Error != nil {
		return pokeballBusiness.Domain{}, result.Error
	}
	return pokeballDatabase.ToDomain(&lowestPokeball), nil
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

	// check if pokemon is exist
	res, err := pokeapiStruct.GetPokemonByID(ctx, pokemonID)
	if err != nil {
		return trainers.Domain{}, err
	}

	pokemonData := &pokemons.Pokemon{ID: pokemonID, Name: res.Name}
	trainerPokemon := &TrainerPokemon{
		PokemonID: pokemonID,
		TrainerID: ID,
	}

	createErr := repo.Conn.Create(&pokemonData)
	if createErr.Error != nil {
		return trainers.Domain{}, createErr.Error
	}
	createErr = repo.Conn.Create(&trainerPokemon)
	if createErr != nil {
		return trainers.Domain{}, createErr.Error
	}
	return trainers.Domain{}, nil
}

// AddGym register trainer to gym
func (repo *MysqlTrainerRepository) AddGym(ctx context.Context, ID, gymID int) (trainers.Domain, error) {
	trainerGym := TrainerGym{
		GymID:     gymID,
		TrainerID: ID,
	}
	result := repo.Conn.Create(&trainerGym)
	if result.Error != nil {
		return trainers.Domain{}, result.Error
	}
	res := Trainer{ID: uint(ID)}
	return res.toDomain(), nil
}

// Login as Trainer
func (repo *MysqlTrainerRepository) Login(ctx context.Context, username, password string) (trainers.Domain, error) {
	var trainerLogin = Trainer{}
	result := repo.Conn.Where("username = ? AND password = ?", username, password).First(&trainerLogin)
	if result.Error != nil {
		return trainers.Domain{}, result.Error
	}
	return trainerLogin.toDomain(), nil
}

// UpdateTrainer update data trainer
func (repo *MysqlTrainerRepository) UpdateTrainer(ctx context.Context, trainerID int, name, address, username, password string) (trainers.Domain, error) {
	var trainerCollection = Trainer{}
	result := repo.Conn.Find(&trainerCollection).Where("id = ?", trainerID).Updates(&Trainer{
		Name:     name,
		Address:  address,
		Username: username,
		Password: password,
	})

	if result.Error != nil {
		return trainers.Domain{}, result.Error
	}

	return trainerCollection.toDomain(), nil
}
