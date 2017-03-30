package main

import sf "github.com/zyedidia/sfml/v2.3/sfml"

type Card struct {
	*sf.RectangleShape
	show     bool
	number   int
	dispnum  string
	suit     string
	dispsuit *sf.Sprite
}

type deck []Card

func NewCard(num int, suit string) *Card {
	card := new(Card)
	card.RectangleShape = sf.NewRectangleShape(sf.Vector2f{60, 100})
	card.SetOrigin(sf.Vector2f{30, 50})
	card.SetFillColor(sf.ColorWhite)
	card.SetOutlineColor(sf.ColorWhite)
	card.SetOutlineThickness(5)
	card.show = false
	card.number = num
	if num == 1 {
		card.dispnum = "A"
	} else if num == 11 {
		card.dispnum = "J"
	} else if num == 12 {
		card.dispnum = "Q"
	} else if num == 13 {
		card.dispnum = "K"
	} else {
		card.dispnum = string(num)
	}
	card.suit = suit
	card.dispsuit = sf.NewSprite(textures[suit+".png"])
	card.dispsuit.SetOrigin(sf.Vector2f{16, 16})
	return card
}

func (c *Card) SetPos(x, y float32) {
	c.SetPosition(sf.Vector2f{x, y})
	c.dispsuit.SetPosition(sf.Vector2f{x + 9, y + 27})
}

func (c *Card) Flip() {
	c.show = !c.show
	if c.show {
		c.SetFillColor(sf.ColorWhite)
	} else {
		c.SetFillColor(sf.Color{210, 60, 60, 255})
	}

}
