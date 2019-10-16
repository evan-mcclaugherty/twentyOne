package card

import (
	"math/rand"
	"time"
)

type card struct {
	Suite string
	Value int
}

type Cards []card

var suites = []string{"hearts", "diamonds", "spades", "clubs"}
var values = map[string]int{
	"ace":   1,
	"two":   2,
	"three": 3,
	"four":  4,
	"five":  5,
	"six":   6,
	"seven": 7,
	"eight": 8,
	"nine":  9,
	"ten":   10,
	"jack":  10,
	"queen": 10,
	"king":  10,
}

type NotEnoughCards struct {
	msg string
}

func (nec *NotEnoughCards) Error() string {
	return nec.msg
}

func (c *Cards) GetOne() []card {
	return c.withdraw(1)
}

func (c *Cards) GetTwo() []card {
	return c.withdraw(2)
}

func (c *Cards) withdraw(n int) []card {
	if len(*c) < n {
		newDeck := NewDeck()
		newDeck.Shuffle()
		*c = append(*c, newDeck...)
	}

	pulled := make([]card, n)
	copy(pulled, (*c)[:n])
	*c = (*c)[n:]

	return pulled
}

func (c *Cards) Shuffle() {
	rand.Seed(time.Now().UnixNano())
	rand.Shuffle(len(*c), func(i, j int) {
		(*c)[i], (*c)[j] = (*c)[j], (*c)[i]
	})
}

func NewDeck() Cards {
	c := Cards{}

	for _, value := range values {
		for _, suite := range suites {
			c = append(c, card{
				Suite: suite,
				Value: value,
			})
		}
	}

	c.Shuffle()

	return c
}
