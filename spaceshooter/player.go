package main

import (
	"time"

	math "github.com/chewxy/math32"
	sf "github.com/zyedidia/sfml/v2.3/sfml"
)

type Player struct {
	*sf.Sprite
	speed    float32
	vel      sf.Vector2f
	keys     [5]sf.KeyCode
	canShoot bool
	msgStep  int
}

func NewPlayer(pos sf.Vector2f) *Player {
	p := new(Player)
	p.Sprite = sf.NewSprite(res.images["playerShip1_blue.png"])
	p.SetPosition(pos)
	v := p.GetGlobalBounds()
	p.SetOrigin(sf.Vector2f{v.Width / 2, v.Height / 2})

	p.keys = [5]sf.KeyCode{sf.KeyW, sf.KeyS, sf.KeyA, sf.KeyD, sf.KeySpace}
	p.canShoot = true

	p.msgStep = 0

	return p
}

func (p *Player) setSpeed(s float32) {
	p.speed = s
}

func (p *Player) Update(dt float32) {
	if sf.KeyboardIsKeyPressed(p.keys[0]) {
		if p.speed < shipSpeed*dt {
			p.speed += accelerate
		} else {
			p.setSpeed(shipSpeed * dt)
		}
	} else {
		if p.speed > 0 {
			p.speed -= decelerate
		} else {
			p.setSpeed(0)
		}
	}
	if sf.KeyboardIsKeyPressed(p.keys[2]) && !sf.KeyboardIsKeyPressed(p.keys[3]) {
		p.Rotate(-rotateSpeed * dt)
	}
	if sf.KeyboardIsKeyPressed(p.keys[3]) && !sf.KeyboardIsKeyPressed(p.keys[2]) {
		p.Rotate(rotateSpeed * dt)
	}
	if sf.KeyboardIsKeyPressed(p.keys[4]) {
		p.shoot()
	}
	if sf.KeyboardIsKeyPressed(sf.KeyN) && sf.KeyboardIsKeyPressed(sf.KeyU) && sf.KeyboardIsKeyPressed(sf.KeyT) {
		p.msg()
	}
	if sf.KeyboardIsKeyPressed(sf.KeyB) && sf.KeyboardIsKeyPressed(sf.KeyU) && sf.KeyboardIsKeyPressed(sf.KeyN) {
		p.hack()
	}
	p.vel.X = math.Cos((p.GetRotation()-90)*math.Pi/180) * p.speed
	p.vel.Y = math.Sin((p.GetRotation()-90)*math.Pi/180) * p.speed

	p.Move(p.vel)

	Wrap(p.Sprite)
}

func (p *Player) shoot() {
	if p.canShoot {
		v := p.GetPosition()
		NewLaser(v.X, v.Y, p.GetRotation(), laserSpeed, false)
		sf.NewSound(res.sounds["sfx_laser1.ogg"]).Play()
		p.canShoot = false
		go func() {
			time.Sleep(shotCooldown)
			p.canShoot = true
		}()
	}
}
