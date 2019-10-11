package main

import (
	"fmt"
	"github.com/evan-mcclaugherty/twentyOne/card"
	"github.com/evan-mcclaugherty/twentyOne/player"
)

type game struct {
	Players []*player.Player
	Deck    card.Cards
}

var g = game{}

func main() {
	g = NewGame(4, card.NewDeck())
	g.Deal()
	g.CalculateScores()
	g.PrintPlayers()
}

func (g *game) PrintPlayers() {
	for _, player := range g.Players {
		fmt.Println(player)
	}
}

func (g *game) CalculateScores() {
	for _, player := range g.Players {
		player.CalculateScore()
	}
}
func (g *game) Deal() {
	for _, player := range g.Players {
		cards := g.Deck.GetTwo()
		player.Hand = append(player.Hand, cards...)
		player.CalculateScore()
		player.DetermineStatus()
	}
}

func NewGame(numberOfPlayers int, deck card.Cards) game {
	players := make([]*player.Player, numberOfPlayers)
	return game{
		Players: players,
		Deck:    deck,
	}
}
