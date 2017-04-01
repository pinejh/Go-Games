package main

import (
	"strconv"

	"github.com/pinejh/console"

	sf "github.com/zyedidia/sfml/v2.3/sfml"
)

type Card struct {
	*sf.RectangleShape
	show     bool
	number   int
	dispnum  *sf.Text
	suit     string
	dispsuit *sf.Sprite
	move     bool
	targetV  sf.Vector2f
}

type Deck []*Card

func NewCard(num int, suit string) *Card {
	card := new(Card)
	card.RectangleShape = sf.NewRectangleShape(sf.Vector2f{60, 100})
	card.SetOrigin(sf.Vector2f{30, 50})
	card.SetFillColor(sf.ColorWhite)
	card.SetOutlineColor(sf.ColorWhite)
	card.SetOutlineThickness(5)
	card.show = true
	//card.Flip()
	card.number = num
	dispnum := ""
	if num == 1 {
		dispnum = "A"
	} else if num == 11 {
		dispnum = "J"
	} else if num == 12 {
		dispnum = "Q"
	} else if num == 13 {
		dispnum = "K"
	} else {
		dispnum = strconv.Itoa(num)
	}
	card.dispnum = sf.NewText(dispnum, fonts["cards.ttf"], 24)
	card.dispnum.SetColor(sf.ColorBlack)
	card.suit = suit
	card.dispsuit = sf.NewSprite(textures[suit+".png"])
	card.dispsuit.SetOrigin(sf.Vector2f{16, 16})
	card.SetPosDeck()
	card.move = false
	card.Flip()
	return card
}

func (c *Card) SetPos(x, y float32) {
	c.SetPosition(sf.Vector2f{x, y})
	c.dispsuit.SetPosition(sf.Vector2f{x + 9, y + 27})
	c.dispnum.SetPosition(sf.Vector2f{x - 25, y - 45})
}
func (c *Card) SetPosV(v sf.Vector2f) {
	c.SetPosition(v)
	c.dispsuit.SetPosition(sf.Vector2f{v.X + 9, v.Y + 27})
	c.dispnum.SetPosition(sf.Vector2f{v.X - 25, v.Y - 45})
}
func (c *Card) SetPosDeck() {
	v := cardSlot["deck"]
	c.SetPosition(v)
	c.dispsuit.SetPosition(sf.Vector2f{v.X + 9, v.Y + 27})
	c.dispnum.SetPosition(sf.Vector2f{v.X - 25, v.Y - 45})
}
func (c *Card) SetPosM(player string, first, second int) {
	f, s := strconv.Itoa(first), strconv.Itoa(second)
	v := cardSlot[player+f+"o"+s]
	c.SetPosition(v)
	c.dispsuit.SetPosition(sf.Vector2f{v.X + 9, v.Y + 27})
	c.dispnum.SetPosition(sf.Vector2f{v.X - 25, v.Y - 45})
}
func (c *Card) MoveTo(player string, first, second int) {
	f, s := strconv.Itoa(first), strconv.Itoa(second)
	c.targetV = cardSlot[player+f+"o"+s]
	c.move = true
}

func (c *Card) Flip() {
	c.show = !c.show
	if c.show {
		c.SetFillColor(sf.ColorWhite)
	} else {
		c.SetFillColor(sf.Color{210, 60, 60, 255})
	}
}

func (c *Card) GetNum() int {
	switch c.number {
	case 11:
		return 10
	case 12:
		return 10
	case 13:
		return 10
	default:
		return c.number
	}
}

func NewDeck() *Deck {
	deck := make(Deck, 0)
	for s := 1; s <= 4; s++ {
		for n := 1; n <= 13; n++ {
			suit := ""
			switch s {
			case 1:
				suit = "spades"
			case 2:
				suit = "diamonds"
			case 3:
				suit = "clubs"
			case 4:
				suit = "hearts"
			}
			c := NewCard(n, suit)
			c.SetPosDeck()
			deck = append(deck, c)
		}
	}
	return &deck
}

func (d *Deck) Length() int {
	return len(*d)
}

func (d *Deck) pickRandom() *Card {
	i := console.RandInt(0, d.Length())
	c := (*d)[i]
	(*d) = append((*d)[:i], (*d)[i+1:]...)
	return c
}

func (d *Deck) DrawCard() *Card {
	c := (*d)[0]
	(*d) = append((*d)[1:])
	return c
}

func (d *Deck) Shuffle() {
	deck := Deck{}
	for d.Length() > 0 {
		c := d.pickRandom()
		deck = append(deck, c)
	}
	(*d) = deck
}

func (d *Deck) AddCard(deck *Deck, flip bool) int {
	c := deck.DrawCard()
	if flip {
		c.Flip()
	}
	(*d) = append((*d), c)
	return c.GetNum()
}
