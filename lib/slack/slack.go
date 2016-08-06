package slack

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"strings"
	"time"
)

type SlackAttachment struct {
	Fallback      string `json:"fallback"`
	Text          string `json:"text"`
	Color         string `json:"color"`
	AuthorName    string `json:"author_name"`
	AuthorLink    string `json:"author_link"`
	AuthorIconUrl string `json:"author_icon"`
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

func PostPokemonIds(webhook_url string, pokemonIds []int, expirationTimes []time.Time, latitude, longitude float64) bool {
	names := make([]string, len(pokemonIds))
	for i, v := range pokemonIds {
		names[i] = fmt.Sprintf("%v (:timer_clock: %v)", GetPokemonNameById(v), expirationTimes[i].Format("03:04PM"))
	}

	pokenames := strings.Join(names, ", ")
	msg := "Pokemon have appeared nearby!"
	//link := fmt.Sprintf("https://pokevision.com/#/@%v,%v", latitude, longitude)
	link := fmt.Sprintf("https://skiplagged.com/catch-that/#%v,%v,18", latitude, longitude)
	req := SlackRequest{
		Attachments: []SlackAttachment{
			SlackAttachment{
				Fallback:   msg,
				Text:       msg,
				Color:      "#DCDCDC",
				AuthorName: pokenames,
				AuthorLink: link,
				AuthorIconUrl: fmt.Sprintf(
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
		bytes.NewBuffer([]byte("payload="+string(buf))),
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
