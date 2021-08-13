package pokeapi

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"pokemontrainer/drivers/databases/pokemons"
)

type Pokeapi struct {
	httpClient http.Client
}

// GetPokemonByID ...
func (pApi *Pokeapi) GetPokemonByID(ctx context.Context, pokemonID int) (pokemons.Pokemon, error) {
	fmt.Println("pokemonID", pokemonID)
	req, _ := http.NewRequest("GET", "https://pokeapi.co/api/v2/"+"pokemon/"+fmt.Sprint(pokemonID), nil)

	resp, err := pApi.httpClient.Do(req)

	if err != nil {
		return pokemons.Pokemon{}, err
	}

	defer resp.Body.Close()

	data := Response{}
	err = json.NewDecoder(resp.Body).Decode(&data)
	if err != nil {
		return pokemons.Pokemon{}, err
	}

	return data.ToDatabase(), nil
}
