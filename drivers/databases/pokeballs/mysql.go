package pokeballs

import (
	"context"
	"pokemontrainer/business/pokeballs"

	"gorm.io/gorm"
)

// MysqlPokeballRepository struct for NewMysqlRepository
type MysqlPokeballRepository struct {
	Conn *gorm.DB
}

// NewMysqlPokeballRepository mysqlPokeball with ORM method
func NewMysqlPokeballRepository(conn *gorm.DB) pokeballs.Repositories {
	return &MysqlPokeballRepository{
		Conn: conn,
	}
}

// AddPokeball add pokeball to database
func (repo *MysqlPokeballRepository) AddPokeball(ctx context.Context, name string, successRate float32) (pokeballs.Domain, error) {
	pokeballData := Pokeball{
		Name:        name,
		SuccessRate: successRate,
	}
	result := repo.Conn.Create(&pokeballData)

	if result.Error != nil {
		return pokeballs.Domain{}, result.Error
	}
	return ToDomain(&pokeballData), nil
}
