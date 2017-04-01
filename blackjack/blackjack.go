package main

import (
	"console"
	"runtime"
	"strconv"
	"time"

	sf "github.com/zyedidia/sfml/v2.3/sfml"
)

const (
	screenWidth  = 800
	screenHeight = 600
)

var textures map[string]*sf.Texture
var fonts map[string]*sf.Font

var cardSlot map[string]sf.Vector2f

var deck *Deck
var p *Deck
var c *Deck

var turn int

func loadTexture(filename string) {
	texture := sf.NewTexture("./res/images/" + filename)
	textures[filename] = texture
}

func loadFont(filename string) {
	font := sf.NewFont("./res/fonts/" + filename)
	fonts[filename] = font
}

func main() {
	runtime.LockOSThread()

	textures = make(map[string]*sf.Texture)

	loadTexture("spades.png")
	loadTexture("clubs.png")
	loadTexture("diamonds.png")
	loadTexture("hearts.png")

	fonts = make(map[string]*sf.Font)

	loadFont("cards.ttf")

	cardSlot = make(map[string]sf.Vector2f)

	generateCardSlots()

	window := sf.NewRenderWindow(sf.VideoMode{screenWidth, screenHeight, 32}, "Blackjack", sf.StyleDefault, nil)
	window.SetVerticalSyncEnabled(true)
	window.SetFramerateLimit(60)

	var dt float32
	dt = 0

	deck := NewDeck()

	p, c := &Deck{}, &Deck{}
	pTotal, cTotal := []int{0}, []int{0}
	dispPT, dispCT := sf.NewText("", fonts["cards.ttf"], 24), sf.NewText("", fonts["cards.ttf"], 24)
	rect := dispPT.GetGlobalBounds()
	dispPT.SetOrigin(sf.Vector2f{rect.Width / 2, rect.Height / 2})
	dispPT.SetPosition(sf.Vector2f{screenWidth / 2, screenHeight - 190})
	rect2 := dispCT.GetGlobalBounds()
	dispCT.SetOrigin(sf.Vector2f{rect2.Width / 2, rect2.Height / 2})
	dispCT.SetPosition(sf.Vector2f{screenWidth / 2, 190})

	CalcPT := func() {
		str := ""
		for i, _ := range pTotal {
			str += strconv.Itoa(pTotal[i])
			if i < len(pTotal)-1 {
				str += " or "
			}
		}
		dispPT.SetString(str)
		rect := dispPT.GetGlobalBounds()
		dispPT.SetOrigin(sf.Vector2f{rect.Width / 2, rect.Height / 2})
	}
	CalcCT := func() {
		str := ""
		for i, _ := range cTotal {
			str += strconv.Itoa(cTotal[i])
			if i < len(cTotal)-1 {
				str += " or "
			}
		}
		dispCT.SetString(str)
		rect := dispCT.GetGlobalBounds()
		dispCT.SetOrigin(sf.Vector2f{rect.Width / 2, rect.Height / 2})
	}

	Repos := func() {
		for i := 1; i <= p.Length(); i++ {
			(*p)[i-1].MoveTo("p", i, p.Length())
		}
		for i := 1; i <= c.Length(); i++ {
			(*c)[i-1].MoveTo("c", i, c.Length())
		}
	}
	AddCardP := func(deck *Deck, flip bool) {
		n := p.AddCard(deck, flip)
		if n == 1 {
			l := len(pTotal)
			for i := 0; i < l; i++ {
				pTotal[i] += 1
				pTotal = append(pTotal, pTotal[i]+10)
			}
		} else {
			for i, _ := range pTotal {
				pTotal[i] += n
			}
		}
		console.Log("P: ", pTotal)
		Repos()
		CalcPT()
	}
	AddCardC := func(deck *Deck, flip bool) {
		n := c.AddCard(deck, flip)
		if n == 1 {
			l := len(cTotal)
			for i := 0; i < l; i++ {
				cTotal[i] += 1
				cTotal = append(cTotal, cTotal[i]+10)
			}
		} else {
			for i, _ := range cTotal {
				cTotal[i] += n
			}
		}
		console.Log("C: ", cTotal)
		Repos()
		CalcCT()
	}

	canNewGame := true
	NewGame := func() {
		if canNewGame {
			canNewGame = false
			p, c = &Deck{}, &Deck{}
			pTotal, cTotal = []int{0}, []int{0}
			deck = NewDeck()
			deck.Shuffle()
			AddCardP(deck, true)
			time.Sleep(time.Second)
			AddCardC(deck, false)
			time.Sleep(time.Second)
			AddCardP(deck, true)
			time.Sleep(time.Second)
			AddCardC(deck, true)
			turn = 0
			for _, v := range cTotal {
				if v == 21 {
					console.Log("COMPUTER BLACKJACK")
					(*c)[0].Flip()
				}
			}
			for _, v := range pTotal {
				if v == 21 {
					console.Log("PLAYER BLACKJACK")
				}
			}
			canNewGame = true
		}
	}

	for window.IsOpen() {
		delta := time.Now()
		if event := window.PollEvent(); event != nil {
			switch event.Type {
			case sf.EventKeyReleased:
				go NewGame()
			case sf.EventClosed:
				window.Close()
			}
		}

		window.Clear(sf.Color{45, 165, 75, 255})
		for _, card := range *deck {
			window.Draw(card)
			if card.show {
				window.Draw(card.dispnum)
				window.Draw(card.dispsuit)
			} else {
			}
		}
		for _, card := range *p {
			window.Draw(card)
			if card.move {
				v := card.GetPosition()
				if v.X == card.targetV.X && v.Y == card.targetV.Y {
					card.move = false
				} else {
					dx, dy := (card.targetV.X-v.X)*dt/10, (card.targetV.Y-v.Y)*dt/10
					if dx*-5 < 1 && dx*5 < 1 && dy*-5 < 1 && dy*5 < 1 {
						dx, dy = 0, 0
						card.SetPosV(card.targetV)
					} else {
						card.SetPosV(sf.Vector2f{v.X + dx, v.Y + dy})
					}
				}
			}
			if card.show {
				window.Draw(card.dispnum)
				window.Draw(card.dispsuit)
			} else {

			}
		}
		for _, card := range *c {
			window.Draw(card)
			if card.move {
				v := card.GetPosition()
				if v.X == card.targetV.X && v.Y == card.targetV.Y {
					card.move = false
				} else {
					dx, dy := (card.targetV.X-v.X)*dt/5, (card.targetV.Y-v.Y)*dt/5
					if dx*-5 < 1 && dx*5 < 1 && dy*-5 < 1 && dy*5 < 1 {
						dx, dy = 0, 0
						card.SetPosV(card.targetV)
					} else {
						card.SetPosV(sf.Vector2f{v.X + dx, v.Y + dy})
					}
				}
			}
			if card.show {
				window.Draw(card.dispnum)
				window.Draw(card.dispsuit)
			} else {
			}
		}
		window.Draw(dispPT)
		if turn == 1 {
			window.Draw(dispCT)
		}
		window.Display()
		dt = float32(time.Since(delta)) / 10000000
	}
}

func generateCardSlots() {
	cardSlot["deck"] = sf.Vector2f{screenWidth - 75, screenHeight / 2}
	genC()
	genP()
}
func genC() {
	cardSlot["c1o1"], cardSlot["c2o3"], cardSlot["c3o5"], cardSlot["c4o7"], cardSlot["c5o9"] = sf.Vector2f{screenWidth / 2, 95}, sf.Vector2f{screenWidth / 2, 95}, sf.Vector2f{screenWidth / 2, 95}, sf.Vector2f{screenWidth / 2, 95}, sf.Vector2f{screenWidth / 2, 95}
	cardSlot["c1o2"], cardSlot["c2o4"], cardSlot["c3o6"], cardSlot["c4o8"], cardSlot["c5o10"] = sf.Vector2f{screenWidth/2 - 40, 95}, sf.Vector2f{screenWidth/2 - 40, 95}, sf.Vector2f{screenWidth/2 - 40, 95}, sf.Vector2f{screenWidth/2 - 40, 95}, sf.Vector2f{screenWidth/2 - 40, 95}
	cardSlot["c2o2"], cardSlot["c3o4"], cardSlot["c4o6"], cardSlot["c5o8"], cardSlot["c6o10"] = sf.Vector2f{screenWidth/2 + 40, 95}, sf.Vector2f{screenWidth/2 + 40, 95}, sf.Vector2f{screenWidth/2 + 40, 95}, sf.Vector2f{screenWidth/2 + 40, 95}, sf.Vector2f{screenWidth/2 + 40, 95}
	cardSlot["c1o3"], cardSlot["c2o5"], cardSlot["c3o7"], cardSlot["c4o9"] = sf.Vector2f{screenWidth/2 - 80, 95}, sf.Vector2f{screenWidth/2 - 80, 95}, sf.Vector2f{screenWidth/2 - 80, 95}, sf.Vector2f{screenWidth/2 - 80, 95}
	cardSlot["c3o3"], cardSlot["c4o5"], cardSlot["c5o7"], cardSlot["c6o9"] = sf.Vector2f{screenWidth/2 + 80, 95}, sf.Vector2f{screenWidth/2 + 80, 95}, sf.Vector2f{screenWidth/2 + 80, 95}, sf.Vector2f{screenWidth/2 + 80, 95}
	cardSlot["c1o4"], cardSlot["c2o6"], cardSlot["c3o8"], cardSlot["c4o10"] = sf.Vector2f{screenWidth/2 - 120, 95}, sf.Vector2f{screenWidth/2 - 120, 95}, sf.Vector2f{screenWidth/2 - 120, 95}, sf.Vector2f{screenWidth/2 - 120, 95}
	cardSlot["c4o4"], cardSlot["c5o6"], cardSlot["c6o8"], cardSlot["c7o10"] = sf.Vector2f{screenWidth/2 + 120, 95}, sf.Vector2f{screenWidth/2 + 120, 95}, sf.Vector2f{screenWidth/2 + 120, 95}, sf.Vector2f{screenWidth/2 + 120, 95}
	cardSlot["c1o5"], cardSlot["c2o7"], cardSlot["c3o9"] = sf.Vector2f{screenWidth/2 - 160, 95}, sf.Vector2f{screenWidth/2 - 160, 95}, sf.Vector2f{screenWidth/2 - 160, 95}
	cardSlot["c5o5"], cardSlot["c6o7"], cardSlot["c7o9"] = sf.Vector2f{screenWidth/2 + 160, 95}, sf.Vector2f{screenWidth/2 + 160, 95}, sf.Vector2f{screenWidth/2 + 160, 95}
	cardSlot["c1o6"], cardSlot["c2o8"], cardSlot["c3o10"] = sf.Vector2f{screenWidth/2 - 200, 95}, sf.Vector2f{screenWidth/2 - 200, 95}, sf.Vector2f{screenWidth/2 - 200, 95}
	cardSlot["c6o6"], cardSlot["c7o8"], cardSlot["c8o10"] = sf.Vector2f{screenWidth/2 + 200, 95}, sf.Vector2f{screenWidth/2 + 200, 95}, sf.Vector2f{screenWidth/2 + 200, 95}
	cardSlot["c1o7"], cardSlot["c2o9"] = sf.Vector2f{screenWidth/2 - 240, 95}, sf.Vector2f{screenWidth/2 - 240, 95}
	cardSlot["c7o7"], cardSlot["c8o9"] = sf.Vector2f{screenWidth/2 + 240, 95}, sf.Vector2f{screenWidth/2 + 240, 95}
	cardSlot["c1o8"], cardSlot["c2o10"] = sf.Vector2f{screenWidth/2 - 280, 95}, sf.Vector2f{screenWidth/2 - 280, 95}
	cardSlot["c8o8"], cardSlot["c9o10"] = sf.Vector2f{screenWidth/2 + 280, 95}, sf.Vector2f{screenWidth/2 + 280, 95}
	cardSlot["c1o9"] = sf.Vector2f{screenWidth/2 - 320, 95}
	cardSlot["c9o9"] = sf.Vector2f{screenWidth/2 + 320, 95}
	cardSlot["c1o10"] = sf.Vector2f{screenWidth/2 - 360, 95}
	cardSlot["c10o10"] = sf.Vector2f{screenWidth/2 + 360, 95}
}
func genP() {
	cardSlot["p1o1"], cardSlot["p2o3"], cardSlot["p3o5"], cardSlot["p4o7"], cardSlot["p5o9"] = sf.Vector2f{screenWidth / 2, screenHeight - 95}, sf.Vector2f{screenWidth / 2, screenHeight - 95}, sf.Vector2f{screenWidth / 2, screenHeight - 95}, sf.Vector2f{screenWidth / 2, screenHeight - 95}, sf.Vector2f{screenWidth / 2, screenHeight - 95}
	cardSlot["p1o2"], cardSlot["p2o4"], cardSlot["p3o6"], cardSlot["p4o8"], cardSlot["p5o10"] = sf.Vector2f{screenWidth/2 - 40, screenHeight - 95}, sf.Vector2f{screenWidth/2 - 40, screenHeight - 95}, sf.Vector2f{screenWidth/2 - 40, screenHeight - 95}, sf.Vector2f{screenWidth/2 - 40, screenHeight - 95}, sf.Vector2f{screenWidth/2 - 40, screenHeight - 95}
	cardSlot["p2o2"], cardSlot["p3o4"], cardSlot["p4o6"], cardSlot["p5o8"], cardSlot["p6o10"] = sf.Vector2f{screenWidth/2 + 40, screenHeight - 95}, sf.Vector2f{screenWidth/2 + 40, screenHeight - 95}, sf.Vector2f{screenWidth/2 + 40, screenHeight - 95}, sf.Vector2f{screenWidth/2 + 40, screenHeight - 95}, sf.Vector2f{screenWidth/2 + 40, screenHeight - 95}
	cardSlot["p1o3"], cardSlot["p2o5"], cardSlot["p3o7"], cardSlot["p4o9"] = sf.Vector2f{screenWidth/2 - 80, screenHeight - 95}, sf.Vector2f{screenWidth/2 - 80, screenHeight - 95}, sf.Vector2f{screenWidth/2 - 80, screenHeight - 95}, sf.Vector2f{screenWidth/2 - 80, screenHeight - 95}
	cardSlot["p3o3"], cardSlot["p4o5"], cardSlot["p5o7"], cardSlot["p6o9"] = sf.Vector2f{screenWidth/2 + 80, screenHeight - 95}, sf.Vector2f{screenWidth/2 + 80, screenHeight - 95}, sf.Vector2f{screenWidth/2 + 80, screenHeight - 95}, sf.Vector2f{screenWidth/2 + 80, screenHeight - 95}
	cardSlot["p1o4"], cardSlot["p2o6"], cardSlot["p3o8"], cardSlot["p4o10"] = sf.Vector2f{screenWidth/2 - 120, screenHeight - 95}, sf.Vector2f{screenWidth/2 - 120, screenHeight - 95}, sf.Vector2f{screenWidth/2 - 120, screenHeight - 95}, sf.Vector2f{screenWidth/2 - 120, screenHeight - 95}
	cardSlot["p4o4"], cardSlot["p5o6"], cardSlot["p6o8"], cardSlot["p7o10"] = sf.Vector2f{screenWidth/2 + 120, screenHeight - 95}, sf.Vector2f{screenWidth/2 + 120, screenHeight - 95}, sf.Vector2f{screenWidth/2 + 120, screenHeight - 95}, sf.Vector2f{screenWidth/2 + 120, screenHeight - 95}
	cardSlot["p1o5"], cardSlot["p2o7"], cardSlot["p3o9"] = sf.Vector2f{screenWidth/2 - 160, screenHeight - 95}, sf.Vector2f{screenWidth/2 - 160, screenHeight - 95}, sf.Vector2f{screenWidth/2 - 160, screenHeight - 95}
	cardSlot["p5o5"], cardSlot["p6o7"], cardSlot["p7o9"] = sf.Vector2f{screenWidth/2 + 160, screenHeight - 95}, sf.Vector2f{screenWidth/2 + 160, screenHeight - 95}, sf.Vector2f{screenWidth/2 + 160, screenHeight - 95}
	cardSlot["p1o6"], cardSlot["p2o8"], cardSlot["p3o10"] = sf.Vector2f{screenWidth/2 - 200, screenHeight - 95}, sf.Vector2f{screenWidth/2 - 200, screenHeight - 95}, sf.Vector2f{screenWidth/2 - 200, screenHeight - 95}
	cardSlot["p6o6"], cardSlot["p7o8"], cardSlot["p8o10"] = sf.Vector2f{screenWidth/2 + 200, screenHeight - 95}, sf.Vector2f{screenWidth/2 + 200, screenHeight - 95}, sf.Vector2f{screenWidth/2 + 200, screenHeight - 95}
	cardSlot["p1o7"], cardSlot["p2o9"] = sf.Vector2f{screenWidth/2 - 240, screenHeight - 95}, sf.Vector2f{screenWidth/2 - 240, screenHeight - 95}
	cardSlot["p7o7"], cardSlot["p8o9"] = sf.Vector2f{screenWidth/2 + 240, screenHeight - 95}, sf.Vector2f{screenWidth/2 + 240, screenHeight - 95}
	cardSlot["p1o8"], cardSlot["p2o10"] = sf.Vector2f{screenWidth/2 - 280, screenHeight - 95}, sf.Vector2f{screenWidth/2 - 280, screenHeight - 95}
	cardSlot["p8o8"], cardSlot["p9o10"] = sf.Vector2f{screenWidth/2 + 280, screenHeight - 95}, sf.Vector2f{screenWidth/2 + 280, screenHeight - 95}
	cardSlot["p1o9"] = sf.Vector2f{screenWidth/2 - 320, screenHeight - 95}
	cardSlot["p9o9"] = sf.Vector2f{screenWidth/2 + 320, screenHeight - 95}
	cardSlot["p1o10"] = sf.Vector2f{screenWidth/2 - 360, screenHeight - 95}
	cardSlot["p10o10"] = sf.Vector2f{screenWidth/2 + 360, screenHeight - 95}
}
