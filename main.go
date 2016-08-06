package main

import (
	"time"
	"fmt"
	"os"
	"strconv"

	"github.com/stewartpark/pokealert/lib/slack"
	"github.com/stewartpark/pokealert/lib/skiplagged"
)


func main() {
	webhook      := os.Getenv("SLACK_WEBHOOK_URL")
	latitude, _  := strconv.ParseFloat(os.Getenv("LATITUDE"), 64)
	longitude, _ := strconv.ParseFloat(os.Getenv("LONGITUDE"), 64)
	euc_range, _ := strconv.ParseFloat(os.Getenv("RANGE"), 64)
	pull_wait, _ := strconv.ParseInt(os.Getenv("PULL_WAIT"), 10, 64)
	t            := time.NewTicker(time.Duration(pull_wait) * time.Minute)

	for {
		fmt.Printf("Loop running...\n")
		pokemons := skiplagged.GetPokemonIdsWithRange(latitude, longitude, euc_range)
		if len(pokemons) > 0 {
			if slack.PostPokemonIds(webhook, pokemons, latitude, longitude) {
				fmt.Printf("Posted to Slack!\n")
			}
		} else {
			fmt.Printf("No Pokemon found.\n")
		}
		<-t.C
	}
}
