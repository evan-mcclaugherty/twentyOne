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
	g.Play()
	g.PrintPlayers()
}

func (g *game) Play() {
	for i := 0; i < len(g.Players); i++ {
		for g.Players[i].Status == player.HIT {
			g.Players[i].Hit(g.Deck.GetOne())
			g.Players[i].CalculateScore()
			g.Players[i].DetermineStatus()
		}
	}
}

func (g *game) PrintPlayers() {
	for _, player := range g.Players {
		fmt.Println(player)
	}
}

func (g *game) CalculateScores() {
	for _, player := range g.Players {
		player.CalculateScore()
		player.DetermineStatus()
	}
}
func (g *game) Deal() {
	for _, player := range g.Players {
		cards := g.Deck.GetTwo()
		player.Hand = append(player.Hand, cards...)
		g.CalculateScores()
	}
}

func NewGame(numberOfPlayers int, deck card.Cards) game {
	players := make([]*player.Player, numberOfPlayers)
	for i := 0; i < len(players); i++ {
		players[i] = new(player.Player)
	}

	return game{
		Players: players,
		Deck:    deck,
	}
}
