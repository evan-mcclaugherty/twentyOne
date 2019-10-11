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

const (
	HIT  string = "hit"
	STAY string = "stay"
	WIN  string = "win"
	BUST string = "bust"
)

func (p *Player) DetermineStatus() {
	switch {
	case p.Score == 21:
		p.Status = WIN
	case p.Score > 21:
		p.Status = BUST
	case p.Score <= 16:
		p.Status = HIT
	default:
		p.Status = STAY
	}
}

func (p *Player) Hit(oneCard card.Cards) {
	p.Hand = append(p.Hand, oneCard...)
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
	playerString += fmt.Sprintf("Player: score %v - status %v - hand %v", p.Score, p.Status, p.Hand)
	return playerString
}
