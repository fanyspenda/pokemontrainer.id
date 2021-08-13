package pokeapi

import (
	"pokemontrainer/drivers/databases/pokemons"
)

// Response data format from pokeapi
type Response struct {
	ID             int
	Name           string
	BaseExperience int
	Order          int
	Weight         int
}

// ToDatabase convert data from API to Database
func (res *Response) ToDatabase() pokemons.Pokemon {
	return pokemons.Pokemon{
		ID:   res.ID,
		Name: res.Name,
	}
}
