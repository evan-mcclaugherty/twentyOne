package player

import (
	"fmt"
	"github.com/evan-mcclaugherty/twentyOne/card"
)

type Player struct {
	Score  int
	Hand   card.Cards
	Status string
}

func (p *Player) DetermineStatus() {
	switch {
	case p.Score == 21:
		p.Status = "win"
	case p.Score > 21:
		p.Status = "bust"
	default:
		p.Status = "stay"
	}
}

func (p *Player) CalculateScore() {
	score := 0
	numOfAces := 0
	for _, card := range p.Hand {
		if card.Value == 1 {
			numOfAces++
		}
		score += card.Value
	}

	if numOfAces > 0 {
		for i := 0; i < numOfAces; i++ {
			if score+10 <= 21 {
				score += 10
			}
		}
	}
	p.Score = score
}

func (p *Player) String() string {
	playerString := ""
	playerString += fmt.Sprintf("Player: score %v - status %v - hand %v\n", p.Score, p.Status, p.Hand)
	return playerString
}
