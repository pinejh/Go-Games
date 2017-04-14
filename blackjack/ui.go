package main

import sf "github.com/zyedidia/sfml/v2.3/sfml"

type button struct {
	*sf.RectangleShape
	text        *sf.Text
	onclick     func()
	activestate int
}

func NewButton(x, y, width, height float32, text string, fontsize uint, onclick func(), activestate int) *button {
	b := new(button)
	b.RectangleShape = sf.NewRectangleShape(sf.Vector2f{width, height})
	b.SetPosition(sf.Vector2f{x, y})
	b.SetOrigin(sf.Vector2f{width / 2, height / 2})
	b.SetFillColor(sf.Color{240, 103, 59, 255})
	b.text = sf.NewText(text, res.fonts["cards.ttf"], fontsize)
	b.text.SetPosition(sf.Vector2f{x, y})
	b.text.SetColor(sf.ColorWhite)
	v := b.text.GetGlobalBounds()
	b.text.SetOrigin(sf.Vector2f{v.Width / 2, v.Height/2 + 4})
	b.onclick = onclick
	b.activestate = activestate
	return b
}

func (b *button) collides(mouse sf.Vector2i) {
	v := sf.Vector2f{float32(mouse.X), float32(mouse.Y)}
	bv := b.GetGlobalBounds()
	if v.X > bv.Left && v.X < bv.Left+bv.Width && v.Y > bv.Top && v.Y < bv.Top+bv.Height {
		b.onclick()
	}
}

var title *sf.Text
var startGame *button
var exitGame *button

var drawCard *button
var stay *button

var buttons []*button

func uiInit() {
	title = sf.NewText("Blackjack", res.fonts["cards.ttf"], 48)
	title.SetPosition(sf.Vector2f{100, 100})
	startGame = NewButton(225, 250, 250, 75, "Start Game", 28, StartGame, 0)
	exitGame = NewButton(225, 350, 250, 75, "Exit Game", 28, ExitGame, 0)
	drawCard = NewButton(100, screenHeight/2-50, 150, 50, "Hit", 16, DrawCard, 1)
	stay = NewButton(100, screenHeight/2+50, 150, 50, "Stay", 16, PlayerStay, 1)
	buttons = append(buttons, startGame, exitGame, drawCard, stay)
}

func uiUpdate(window *sf.RenderWindow, focused bool, gamestate, turn int) {
	if focused && sf.IsMouseButtonPressed(sf.MouseLeft) {
		vi := sf.MouseGetPosition(window)
		if vi.X > 0 && vi.X < screenWidth && vi.Y > 0 && vi.Y < screenHeight {
			for _, b := range buttons {
				if b.activestate == gamestate {
					if gamestate == 1 {
						if turn == 0 {
							b.collides(vi)
						}
					} else {
						b.collides(vi)
					}
				}
			}
		}
	}
	if gamemode == 0 {
		window.Draw(title)
		window.Draw(startGame)
		window.Draw(startGame.text)
		window.Draw(exitGame)
		window.Draw(exitGame.text)
	} else if gamemode == 1 && turn == 0 {
		window.Draw(drawCard)
		window.Draw(drawCard.text)
		window.Draw(stay)
		window.Draw(stay.text)
	}
}
