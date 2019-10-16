package player

import (
	"fmt"

	"github.com/evan-mcclaugherty/twentyOne/card"
)

type Player struct {
	Score  int
	Hand   card.Cards
	Status int
}

const (
	Hit = iota + 1
	Stay
	Win
	Bust
)

func (p *Player) DetermineStatus() {
	switch {
	case p.Score == 21:
		p.Status = Win
	case p.Score > 21:
		p.Status = Bust
	case p.Score <= 16:
		p.Status = Hit
	default:
		p.Status = Stay
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

func (p *Player) SetHand(cards card.Cards) {
	p.Hand = cards
}

func (p *Player) String() string {
	return fmt.Sprintf("Player: score %v - status %v - hand %v", p.Score, p.Status, p.Hand)
}
