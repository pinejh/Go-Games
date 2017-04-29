package main

import (
	//"fmt"

	sf "github.com/zyedidia/sfml/v2.3/sfml"
)

type Player struct {
	*sf.Sprite

	lives      int
	crouch     bool
	invert     bool
	box        map[string]sf.Rectf
	isMoving   bool
	isGrounded bool
	canJump    bool
	vel        sf.Vector2f
}

func NewPlayer() *Player {
	p := new(Player)
	p.Sprite = sf.NewSprite(res.images["p1_stand.png"])
	p.SetPosition(sf.Vector2f{0, -70})
	rect := p.GetGlobalBounds()
	p.SetOrigin(sf.Vector2f{rect.Width / 2, rect.Height})
	p.lives = 3
	p.crouch = false
	p.isGrounded = false
	p.box = make(map[string]sf.Rectf)
	p.box["feet"] = sf.Rectf{p.GetPosition().X - 20, p.GetPosition().Y - 5, 40, 5}
	return p
}

func (p *Player) Update(dt float32) {
	if sf.KeyboardIsKeyPressed(sf.KeyA) && !sf.KeyboardIsKeyPressed(sf.KeyD) {
		p.Left()
	} else if sf.KeyboardIsKeyPressed(sf.KeyD) && !sf.KeyboardIsKeyPressed(sf.KeyA) {
		p.Right()
	} else {
		p.isMoving = false
	}
	if !p.isMoving {
		if p.isGrounded {
			p.vel.X *= playerDecelG
		} else {
			p.vel.X *= playerDecelA
		}
	}

	if p.vel.X < .15 && p.vel.X > -.15 {
		p.vel.X = 0
	}
	p.vel.X = clamp(p.vel.X, -playerTopSpeed, playerTopSpeed)
	if sf.KeyboardIsKeyPressed(sf.KeyS) && !p.crouch {
		p.Crouch()
	} else if !sf.KeyboardIsKeyPressed(sf.KeyS) && p.crouch {
		p.Uncrouch()
	}

	p.isGrounded = false
	for _, t := range level {
		rect := t.GetGlobalBounds()
		c, _ := p.box["feet"].Intersects(rect)
		if c {
			p.isGrounded = true
			v := p.GetPosition()
			p.Move(sf.Vector2f{0, -v.Y + rect.Top})
		}
	}

	if p.isGrounded && p.vel.Y != 0 {
		p.vel.Y = 0
	} else if !p.isGrounded {
		p.vel.Y += gravity
	}

	if sf.KeyboardIsKeyPressed(sf.KeyW) {
		if p.canJump && p.isGrounded {
			p.Jump()
		}
	}
	if !sf.KeyboardIsKeyPressed(sf.KeyW) && p.isGrounded {
		p.canJump = true
	}

	p.Move(sf.Vector2f{p.vel.X * dt, p.vel.Y * dt})

}

func (p *Player) Move(pos sf.Vector2f) {
	p.Sprite.Move(pos)
	v := p.GetPosition()
	p.box["feet"] = sf.Rectf{v.X - 20, v.Y - 5, 40, 5}
}

func (p *Player) Crouch() {
	p.crouch = true
	p.SetTexture(res.images["p1_duck.png"], true)
	rect := p.GetGlobalBounds()
	p.SetOrigin(sf.Vector2f{rect.Width / 2, rect.Height})
}

func (p *Player) Uncrouch() {
	p.crouch = false
	p.SetTexture(res.images["p1_stand.png"], true)
	rect := p.GetGlobalBounds()
	p.SetOrigin(sf.Vector2f{rect.Width / 2, rect.Height})
}

func (p *Player) Left() {
	if !p.invert {
		p.invert = true
		p.Scale(sf.Vector2f{-1, 1})
	}

	if !p.crouch {
		p.isMoving = true
		p.vel.X -= playerAccel
	} else {
		p.isMoving = false
	}
}

func (p *Player) Right() {
	if p.invert {
		p.invert = false
		p.Scale(sf.Vector2f{-1, 1})
	}

	if !p.crouch {
		p.isMoving = true
		p.vel.X += playerAccel
	} else {
		p.isMoving = false
	}
}

func (p *Player) Jump() {
	p.vel.Y = -7.5
	p.canJump = false
}
