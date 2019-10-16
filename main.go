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

func main() {
	g := newGame(4, card.NewDeck())
	g.deal()
	g.play()
	g.printPlayers()
}

func (g *game) play() {
	for _, p := range g.Players {
		for p.Status == player.Hit {
			p.Hit(g.Deck.GetOne())
			p.CalculateScore()
			p.DetermineStatus()
		}
	}
}

func (g *game) printPlayers() {
	for _, player := range g.Players {
		fmt.Println(player)
	}
}

func (g *game) deal() {
	for _, p := range g.Players {
		cards := g.Deck.GetTwo()
		p.SetHand(cards)
		p.CalculateScore()
		p.DetermineStatus()
	}
}

func newGame(numberOfPlayers int, deck card.Cards) *game {
	players := make([]*player.Player, numberOfPlayers)

	for i := 0; i < numberOfPlayers; i++ {
		players[i] = &player.Player{}
	}

	return &game{
		Players: players,
		Deck:    deck,
	}
}

// 1. Changed a handful of functions to unexported since we don't need them exported right now. I
// didn't add comments to any of the exported funcs but generally it's good to do. If you always
// keep up with it you will have a fully documented api through godoc which is awesome!

// 2. Removed global game variable -> I generally try not to use global variables. This can get
// harder as an application gets larger and has more dependencies but there are some really cool
// patterns I've learned from it.

// 3. Inside deal() I removed the call to g.CalculateScores() and set each player's score
// when getting to them in the iteration. The call to g.CalculateScores() looks like it looped
// over all the players and set their scores on each iteration of the loop over g.Players.
// Before 10 players would have resolved to 100 loops. So O(n^2) time complexity and
// resetting the score and status for players over and over again. Now CalculateScores
// isn't used anywhere else so I deleted it.

// 4. Changed play() to a range over players like you did in deal() for consistency and added
// a SetHand method for consistency with your OOP style Go.

// 5. Changed player status constants to use iota. Right after I did that I realized you were
// using the string to print each player's status to the console so you should change them back
// but I thought otherwise it was a good opportunity to use iota!

// 6. I don't think we need the loop at the end of GetTwo()... I could be totally wrong, but
// couldn't figure out a reason to keep it.
/*
for i := 0; i < n; i++ {
	(*c)[i] = card{}
}
*/

// 7. Any other changes are based on my preferences or styleguides I like just so
// you can see the way someone else might write it -> don't treat these things
// like a PR but rather another style.
