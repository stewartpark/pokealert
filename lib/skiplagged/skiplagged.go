package skiplagged

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"time"
)

type SkiplaggedPokemon struct {
	PokemonId      int     `json:"pokemon_id"`
	PokemonName    string  `json:"pokemon_name"`
	ExpirationTime int64   `json:"expires"`
	Latitude       float64 `json:"latitude"`
	Longitude      float64 `json:"longitude"`
}

type SkiplaggedResponse struct {
	Status  string              `json:"status"`
	Pokemon []SkiplaggedPokemon `json:"pokemons"`
}

var seen_pokemon = make(map[string]bool)

func GetPokemonIdsWithRange(latitude, longitude, radius float64) ([]int, []time.Time) {
	resp, err := http.Get(
		fmt.Sprintf(
			"https://skiplagged.com/api/pokemon.php?bounds=%v,%v,%v,%v",
			latitude-(radius/2.0),
			longitude-(radius/2.0),
			latitude+(radius/2.0),
			longitude+(radius/2.0),
		),
	)
	if err != nil {
		fmt.Printf("Error raised while sending a HTTP request: %v", err)
		return []int{}, []time.Time{}
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)

	if err != nil {
		fmt.Printf("Error raised while reading the buffer: %v", err)
		return []int{}, []time.Time{}
	}

	var poke_resp SkiplaggedResponse
	if err := json.Unmarshal(body, &poke_resp); err != nil {
		panic(err)
	}

	result_ids := []int{}
	result_times := []time.Time{}
	for _, pokemon := range poke_resp.Pokemon {
		poke_UID := fmt.Sprintf(
			"%v:%v:%v:%v",
			pokemon.PokemonId,
			pokemon.ExpirationTime,
			pokemon.Latitude,
			pokemon.Longitude,
		)
		if !seen_pokemon[poke_UID] {
			seen_pokemon[poke_UID] = true
			result_ids = append(result_ids, pokemon.PokemonId)
			result_times = append(result_times, time.Unix(pokemon.ExpirationTime, 0))
		}
	}

	return result_ids, result_times
}
