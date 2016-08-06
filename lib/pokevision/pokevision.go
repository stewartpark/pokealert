package pokevision

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"math"
	"net/http"
)

type PokevisionPokemon struct {
	Id             int     `json:"id"`
	Data           string  `json:"data"`
	ExpirationTime int     `json:"expiration_time"`
	PokemonId      int     `json:"pokemonId"`
	Latitude       float64 `json:"latitude"`
	Longitude      float64 `json:"longitude"`
	UID            string  `json:"uid"`
	IsAlive        bool    `json:"is_alive"`
}

type PokevisionResponse struct {
	Status  string              `json:"status"`
	Pokemon []PokevisionPokemon `json:"pokemon"`
}

var seen_pokemon = make(map[string]bool)

func GetPokemonIdsWithRange(latitude, longitude, radius float64) []int {
	resp, err := http.Get(
		fmt.Sprintf("https://pokevision.com/map/data/%v/%v", latitude, longitude),
	)
	if err != nil {
		fmt.Printf("Error raised while sending a HTTP request: %v", err)
		return []int{}
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)

	if err != nil {
		fmt.Printf("Error raised while reading the buffer: %v", err)
		return []int{}
	}

	var poke_resp PokevisionResponse
	if err := json.Unmarshal(body, &poke_resp); err != nil {
		panic(err)
	}

	result := []int{}
	for _, pokemon := range poke_resp.Pokemon {
		if pokemon.IsAlive && !seen_pokemon[pokemon.UID] &&
			math.Sqrt(math.Pow(pokemon.Latitude-latitude, 2)+math.Pow(pokemon.Longitude-longitude, 2)) <= radius {
			seen_pokemon[pokemon.UID] = true
			result = append(result, pokemon.PokemonId)
		}
	}

	return result
}
