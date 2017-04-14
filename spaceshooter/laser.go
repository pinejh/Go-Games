package main

import (
	math "github.com/chewxy/math32"
	sf "github.com/zyedidia/sfml/v2.3/sfml"
)

type Laser struct {
	*sf.Sprite
	angle float32
	speed float32
}

func NewLaser(x, y, angle, speed float32, special bool) {
	l := new(Laser)
	if !special {
		l.Sprite = sf.NewSprite(res.images["laserBlue.png"])
	} else {
		l.Sprite = sf.NewSprite(res.images["laserBlueShort.png"])
	}
	l.SetPosition(sf.Vector2f{x, y})
	rect := l.GetGlobalBounds()
	l.SetOrigin(sf.Vector2f{rect.Width / 2, rect.Height / 2})
	l.SetRotation(angle)
	l.angle = angle - 90
	l.speed = speed
	lasers = append(lasers, l)
}

func (l *Laser) Update() {
	l.Move(sf.Vector2f{math.Cos(l.angle*math.Pi/180) * l.speed, math.Sin(l.angle*math.Pi/180) * l.speed})
}
