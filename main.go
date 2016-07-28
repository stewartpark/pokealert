package main

import (
	"time"
	"fmt"
	"os"
	"strconv"

	"github.com/stewartpark/pokealert/lib/slack"
	"github.com/stewartpark/pokealert/lib/pokevision"
)


func main() {
	t            := time.NewTicker(time.Minute)
	webhook      := os.Getenv("SLACK_WEBHOOK_URL")
	latitude, _  := strconv.ParseFloat(os.Getenv("LATITUDE"), 64)
	longitude, _ := strconv.ParseFloat(os.Getenv("LONGITUDE"), 64)
	euc_range, _ := strconv.ParseFloat(os.Getenv("RANGE"), 64)

	for {
		fmt.Printf("Loop running...\n")
		pokemons := pokevision.GetPokemonIdsWithRange(latitude, longitude, euc_range)
		if len(pokemons) > 0 {
			if slack.PostPokemonIds(webhook, pokemons) {
				fmt.Printf("Posted to Slack!\n")
			}
		} else {
			fmt.Printf("No Pokemon found.\n")
		}
		<-t.C
	}
}
