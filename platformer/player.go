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
	col        map[string]*CircCol
	box        map[string]sf.Rectf
	isGrounded bool
	canJump    bool
	vel        sf.Vector2f
}

func NewPlayer() *Player {
	p := new(Player)
	p.Sprite = sf.NewSprite(res.images["p1_stand.png"])
	p.SetPosition(sf.Vector2f{screenWidth / 2, screenHeight / 2})
	rect := p.GetGlobalBounds()
	p.SetOrigin(sf.Vector2f{rect.Width / 2, rect.Height})
	p.lives = 3
	p.crouch = false
	p.isGrounded = false
	p.col = make(map[string]*CircCol)
	p.box = make(map[string]sf.Rectf)
	//p.col["feet"] = NewCircCol(sf.Vector2f{p.GetPosition().X, p.GetPosition().Y - 10}, 10)
	p.box["feet"] = sf.Rectf{p.GetPosition().X - 25, p.GetPosition().Y - 5, 50, 5}
	//p.col["floor"] = NewCircCol(sf.Vector2f{p.GetPosition().X, p.GetPosition().Y - 10}, 11)
	p.box["floor"] = sf.Rectf{p.GetPosition().X - 25, p.GetPosition().Y - 5, 50, 7}
	return p
}

func (p *Player) Update(dt float32) {
	p.vel.X = 0

	if sf.KeyboardIsKeyPressed(sf.KeyA) && !sf.KeyboardIsKeyPressed(sf.KeyD) {
		p.Left()
	}
	if sf.KeyboardIsKeyPressed(sf.KeyD) && !sf.KeyboardIsKeyPressed(sf.KeyA) {
		p.Right()
	}
	if sf.KeyboardIsKeyPressed(sf.KeyS) && !p.crouch {
		p.Crouch()
	} else if !sf.KeyboardIsKeyPressed(sf.KeyS) && p.crouch {
		p.Uncrouch()
	}

	p.isGrounded = false
	for _, t := range tiles {
		c, _ := p.box["floor"].Intersects(t.GetGlobalBounds())
		if c {
			p.isGrounded = true
		}
		c, r := p.box["feet"].Intersects(t.GetGlobalBounds())
		if c {
			v := p.GetPosition()
			p.Move(sf.Vector2f{0, v.Y - (r.Top + r.Height)})
		}
		/*v, c := p.col["feet"].CollidesTileTop(t)
		if c {
			p.isGrounded = true
			p.Move(v)
			//fmt.Println("grounded")
		}
		_, c = p.col["floor"].CollidesTileTop(t)
		if c {
			p.isGrounded = true
		}*/
	}

	if p.isGrounded && p.vel.Y != 0 {
		p.vel.Y = 0
	} else if !p.isGrounded {
		p.vel.Y += gravity
	}

	if sf.KeyboardIsKeyPressed(sf.KeyW) {
		if p.canJump {
			p.Jump()
		}
	}
	if !sf.KeyboardIsKeyPressed(sf.KeyW) && p.isGrounded {
		p.canJump = true
	}

	//fmt.Println(p.vel.X, p.vel.Y)
	p.Move(sf.Vector2f{p.vel.X * dt, p.vel.Y * dt})

}

func (p *Player) Move(pos sf.Vector2f) {
	p.Sprite.Move(pos)
	//p.col["feet"].Vector2f = sf.Vector2f{p.GetPosition().X, p.GetPosition().Y - 10}
	v := p.GetPosition()
	p.box["feet"] = sf.Rectf{v.X - 25, v.Y - 5, 50, 5}
	p.box["floor"] = sf.Rectf{v.X - 25, v.Y - 5, 50, 7}

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

	p.vel.X = -2
}

func (p *Player) Right() {
	if p.invert {
		p.invert = false
		p.Scale(sf.Vector2f{-1, 1})
	}

	p.vel.X = 2
}

func (p *Player) Jump() {
	p.vel.Y = -5
	p.canJump = false
}
