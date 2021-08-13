package requests

import (
	"pokemontrainer/business/trainers"
)

// TrainerLogin is about what data will be accepted from FE when user login
type TrainerLogin struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

// TrainerRegister ...
type TrainerRegister struct {
	Username string `json:"username"`
	Password string `json:"password"`
	Name     string `json:"name"`
	Address  string `json:"address"`
}

// TrainerCatchPokemon ...
type TrainerCatchPokemon struct {
	TrainerID int `json:"trainer_id"`
	PokemonID int `json:"pokemon_id"`
}

// ToDomain convert request data to Domain
func (req *TrainerLogin) ToDomain() *trainers.Domain {
	return &trainers.Domain{
		Username: req.Username,
		Password: req.Password,
	}
}

// ToDomain covert register request data to domain
func (req *TrainerRegister) ToDomain() *trainers.Domain {
	return &trainers.Domain{
		Username: req.Username,
		Password: req.Password,
		Name:     req.Name,
		Address:  req.Address,
	}
}
