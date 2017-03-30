package main

import (
	math "github.com/chewxy/math32"

	sf "github.com/manyminds/gosfml"
)

type ball struct {
	*sf.RectangleShape
	velX float32
	velY float32
}

func NewBall() *ball {
	b := new(ball)
	rect, _ := sf.NewRectangleShape()
	b.RectangleShape = rect
	b.velX, b.velY = 1.5, 1.5

	rect.SetSize(sf.Vector2f{10, 10})
	rect.SetFillColor(sf.ColorWhite())
	rect.SetOrigin(sf.Vector2f{5, 5})
	rect.SetPosition(sf.Vector2f{screenWidth / 2, screenHeight / 2})

	return b
}

func (b *ball) Update(dt float32, p1, p2 *player) {
	b.Move(sf.Vector2f{b.velX * dt, b.velY * dt})
	pos := b.GetPosition()
	if pos.Y+5 > screenHeight && b.velY > 0 {
		b.velY = -b.velY
	}
	if pos.Y-5 < 0 && b.velY < 0 {
		b.velY = -b.velY
	}
	if pos.X-5 > screenWidth {
		p1.Score()
		if p1.score < 7 {
			b.Reset()
		}
	}
	if pos.X+5 < 0 {
		p2.Score()
		if p2.score < 7 {
			b.Reset()
		}
	}
	if b.Collides(p1) {
		b.velX = math.Abs(b.velX)
	}
	if b.Collides(p2) {
		b.velX = -math.Abs(b.velX)
	}
	Window.Draw(b, sf.DefaultRenderStates())
}

func (b *ball) Reset() {
	b.SetPosition(sf.Vector2f{screenWidth / 2, screenHeight / 2})
}

func (b *ball) Collides(p *player) bool {
	result, _ := b.GetGlobalBounds().Intersects(p.GetGlobalBounds())
	return result
}
