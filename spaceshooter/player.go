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
	//p.setSpeed(0)
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

}

func (p *Player) shoot() {
	if p.canShoot {
		v := p.GetPosition()
		NewLaser(v.X, v.Y, p.GetRotation(), laserSpeed, false)
		p.canShoot = false
		go func() {
			time.Sleep(shotCooldown)
			p.canShoot = true
		}()
	}
}

func (p *Player) msg() {
	if p.canShoot {
		v := p.GetPosition()
		if msgmap[p.msgStep][0] == 0 {
			NewLaser(v.X+math.Cos((p.GetRotation())*math.Pi/180)*-30, v.Y+math.Sin((p.GetRotation())*math.Pi/180)*-30, p.GetRotation(), 8, true)
		}
		if msgmap[p.msgStep][1] == 0 {
			NewLaser(v.X+math.Cos((p.GetRotation())*math.Pi/180)*-20, v.Y+math.Sin((p.GetRotation())*math.Pi/180)*-20, p.GetRotation(), 8, true)
		}
		if msgmap[p.msgStep][2] == 0 {
			NewLaser(v.X+math.Cos((p.GetRotation())*math.Pi/180)*-10, v.Y+math.Sin((p.GetRotation())*math.Pi/180)*-10, p.GetRotation(), 8, true)
		}
		if msgmap[p.msgStep][3] == 0 {
			NewLaser(v.X, v.Y, p.GetRotation(), 8, true)
		}
		if msgmap[p.msgStep][4] == 0 {
			NewLaser(v.X+math.Cos((p.GetRotation())*math.Pi/180)*10, v.Y+math.Sin((p.GetRotation())*math.Pi/180)*10, p.GetRotation(), 8, true)
		}
		if msgmap[p.msgStep][5] == 0 {
			NewLaser(v.X+math.Cos((p.GetRotation())*math.Pi/180)*20, v.Y+math.Sin((p.GetRotation())*math.Pi/180)*20, p.GetRotation(), 8, true)
		}
		if msgmap[p.msgStep][6] == 0 {
			NewLaser(v.X+math.Cos((p.GetRotation())*math.Pi/180)*30, v.Y+math.Sin((p.GetRotation())*math.Pi/180)*30, p.GetRotation(), 8, true)
		}
		p.canShoot = false
		p.msgStep--
		if p.msgStep < 0 {
			p.msgStep = len(msgmap) - 1
		}
		go func() {
			time.Sleep(75 * time.Millisecond)
			p.canShoot = true
		}()
	}
}

func (p *Player) hack() {
	v := p.GetPosition()
	NewLaser(v.X, v.Y, p.GetRotation(), laserSpeed, false)
	NewLaser(v.X, v.Y, p.GetRotation()+10, laserSpeed, false)
	NewLaser(v.X, v.Y, p.GetRotation()-10, laserSpeed, false)
	NewLaser(v.X, v.Y, p.GetRotation()+20, laserSpeed, false)
	NewLaser(v.X, v.Y, p.GetRotation()-20, laserSpeed, false)
}
