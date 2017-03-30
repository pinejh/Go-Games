package main

import sf "github.com/manyminds/gosfml"

type player struct {
	*sf.RectangleShape
	up    bool
	down  bool
	score int
}

func NewPlayer(x float32) *player {
	p := new(player)
	rect, _ := sf.NewRectangleShape()
	p.RectangleShape = rect
	p.up, p.down = false, false
	p.score = 0

	p.SetSize(sf.Vector2f{10, 50})
	p.SetFillColor(sf.ColorWhite())
	p.SetOrigin(sf.Vector2f{5, 25})
	p.SetPosition(sf.Vector2f{x, screenHeight / 2})
	return p
}

func (p *player) Score() {
	p.score++
}

func (p *player) Update(dt float32) {
	v := p.GetPosition()
	if p.up && !p.down {
		if v.Y-25-playerSpeed*dt >= 0 {
			p.Move(sf.Vector2f{0, -playerSpeed * dt})
		} else {
			p.SetPosition(sf.Vector2f{v.X, 25})
		}
	}
	if p.down && !p.up {
		if v.Y+25+playerSpeed*dt <= screenHeight {
			p.Move(sf.Vector2f{0, playerSpeed * dt})
		} else {
			p.SetPosition(sf.Vector2f{v.X, screenHeight - 25})
		}
	}
	Window.Draw(p, sf.DefaultRenderStates())
}
