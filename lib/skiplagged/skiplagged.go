package skiplagged

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

type SkiplaggedPokemon struct {
	PokemonId      int     `json:"pokemon_id"`
	PokemonName    string  `json:"pokemon_name"`
	ExpirationTime int     `json:"expires"`
	Latitude       float64 `json:"latitude"`
	Longitude      float64 `json:"longitude"`
}

type SkiplaggedResponse struct {
	Status  string              `json:"status"`
	Pokemon []SkiplaggedPokemon `json:"pokemons"`
}

var seen_pokemon = make(map[string]bool)

func GetPokemonIdsWithRange(latitude, longitude, radius float64) []int {
	resp, err := http.Get(
		fmt.Sprintf(
			"https://skiplagged.com/api/pokemon.php?bounds=%v,%v,%v,%v",
			latitude-(radius/2.0), // Not sure it's the correct scale
			longitude-(radius/2.0),
			latitude+(radius/2.0),
			longitude+(radius/2.0),
		),
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

	var poke_resp SkiplaggedResponse
	if err := json.Unmarshal(body, &poke_resp); err != nil {
		panic(err)
	}

	result := []int{}
	for _, pokemon := range poke_resp.Pokemon {
		result = append(result, pokemon.PokemonId)
	}

	return result
}
