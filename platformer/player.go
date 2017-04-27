package main

import (
	sf "github.com/zyedidia/sfml/v2.3/sfml"
)

type Player struct {
	*sf.Sprite

	lives   int
	raycast map[string]sf.Vector2f
}

func NewPlayer() *Player {
	p := new(Player)
	p.Sprite = sf.NewSprite(res.images["p1_stand.png"])
	p.SetPosition(sf.Vector2f{screenWidth / 2, screenHeight / 2})
	rect := p.GetGlobalBounds()
	p.SetOrigin(sf.Vector2f{rect.Width / 2, rect.Height / 2})
	p.lives = 3
	//p.raycast["feet"]
	return p
}

func (p *Player) Update(dt float32) {
	if sf.KeyboardIsKeyPressed(sf.KeyA) && !sf.KeyboardIsKeyPressed(sf.KeyD) {

	}
	if sf.KeyboardIsKeyPressed(sf.KeyD) && !sf.KeyboardIsKeyPressed(sf.KeyA) {

	}
	if sf.KeyboardIsKeyPressed(sf.KeyW) {

	}
	if sf.KeyboardIsKeyPressed(sf.KeyS) {

	}
}
