package slack

import (
	"net/http"
	"fmt"
	"strings"
	"bytes"
	"encoding/json"
)

type SlackAttachment struct {
	Fallback string `json:"fallback"`
	Color    string `json:"color"`
	Text     string `json:"text"`
	ThumbUrl string `json:"thumb_url"`
}

type SlackRequest struct {
	Attachments []SlackAttachment `json:"attachments"`
}


func GetPokemonNameById(id int) string {
	pokemon_names := []string{
		"Invalid",
		"Bulbasaur", "Ivysaur", "Venusaur", "Charmander", "Charmeleon",
		"Charizard", "Squirtle", "Wartortle", "Blastoise", "Caterpie",
		"Metapod", "Butterfree", "Weedle", "Kakuna", "Beedrill", "Pidgey", "Pidgeotto",
		"Pidgeot", "Rattata", "Raticate", "Spearow", "Fearow", "Ekans",
		"Arbok", "Pikachu", "Raichu", "Sandshrew", "Sandslash", "Nidoran", "Nidorina",
		"Nidoqueen", "Nidoran", "Nidorino", "Nidoking", "Clefairy", "Clefable", "Vulpix",
		"Ninetales", "Jigglypuff", "Wigglytuff", "Zubat", "Golbat", "Oddish", "Gloom",
		"Vileplume", "Paras", "Parasect", "Venonat", "Venomoth", "Diglett", "Dugtrio",
		"Meowth", "Persian", "Psyduck", "Golduck", "Mankey", "Primeape", "Growlithe",
		"Arcanine", "Poliwag", "Poliwhirl", "Poliwrath", "Abra", "Kadabra", "Alakazam",
		"Machop", "Machoke", "Machamp", "Bellsprout",
		"Weepinbell", "Victreebel", "Tentacool", "Tentacruel", "Geodude", "Graveler", "Golem",
		"Ponyta", "Rapidash", "Slowpoke", "Slowbro", "Magnemite", "Magneton",
		"Farfetch'd", "Doduo", "Dodrio", "Seel", "Dewgong", "Grimer", "Muk",
		"Shellder", "Cloyster", "Gastly", "Haunter", "Gengar", "Onix", "Drowzee", "Hypno",
		"Krabby", "Kingler", "Voltorb", "Electrode", "", "Exeggutor", "Cubone",
		"Marowak", "Hitmonlee", "Hitmonchan", "Lickitung", "Koffing", "Weezing",
		"Rhyhorn", "Rhydon", "Chansey", "Tangela", "Kangaskhan", "Horsea",
		"Seadra", "Goldeen", "Seaking", "", "Starmie", "Mr. mime", "Scyther",
		"Jynx", "Electabuzz", "Magmar", "Pinsir", "Tauros", "Magikarp", "Gyarados",
		"Lapras", "Ditto", "Eevee", "", "Jolteon", "Flareon", "Porygon", "Omanyte",
		"Omastar", "Kabuto", "Kabutops", "Aerodactyl", "Snorlax", "Articuno",
		"Zapdos", "Moltres", "Dratini", "Dragonair", "Dragonite", "Mewtwo", "Mew",
	}
	return pokemon_names[id]
}

func PostPokemonIds(webhook_url string, pokemonIds []int) bool {
	names := make([]string, len(pokemonIds))
	for i, v := range pokemonIds {
		names[i] = GetPokemonNameById(v)
	}

    msg := fmt.Sprintf("Pokemon have appeared nearby: %v", strings.Join(names, ", "))
	req := SlackRequest{
		Attachments: []SlackAttachment{
			SlackAttachment{
				Fallback: msg,
                Text: msg,
                Color: "#FF0000",
				ThumbUrl: fmt.Sprintf(
					"http://ugc.pokevision.com/images/pokemon/%d.png",
					pokemonIds[0],
				),
			},
		},
	}

	buf, _ := json.Marshal(req)

	http_req, _ := http.NewRequest(
		"POST",
		webhook_url,
		bytes.NewBuffer([]byte("payload=" + string(buf))),
	)
	http_req.Header.Add("Content-Type", "application/x-www-form-urlencoded")

	client := &http.Client{}
	resp, err := client.Do(http_req)

	if err != nil {
		fmt.Printf("Failed to post on Slack.\n")
		return false
	} else {
		defer resp.Body.Close()
		return true
	}
}
