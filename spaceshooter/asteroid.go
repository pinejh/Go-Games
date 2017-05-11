package main

import (
	"strconv"

	math "github.com/chewxy/math32"
	sf "github.com/zyedidia/sfml/v2.3/sfml"
)

type Asteroid struct {
	*sf.Sprite
	width         float32
	height        float32
	rotationSpeed float32
	vel           sf.Vector2f
	radius        float32
	dead          bool
}

func NewAsteroid(pos sf.Vector2f, angle, speed, rotationSpeed float32) *Asteroid {
	a := new(Asteroid)
	a.Sprite = sf.NewSprite(res.images["asteroid"+strconv.Itoa(RandInt(1, 5))+".png"])
	a.SetPosition(pos)
	rect := a.GetGlobalBounds()
	a.width, a.height = rect.Width, rect.Height
	a.radius = (a.width + a.height) / 2 * 3 / 4
	a.SetOrigin(sf.Vector2f{a.width / 2, a.height / 2})
	angle = angle * math.Pi / 180
	a.vel = sf.Vector2f{math.Cos(angle) * speed, math.Sin(angle) * speed}
	a.rotationSpeed = rotationSpeed
	a.dead = false
	return a
}

func GenAsteroid() {
	v := sf.Vector2f{0, 0}
	if len(ast)%2 == 0 {
		v.Y = rand(0, screenHeight)
	} else {
		v.X = rand(0, screenWidth)
	}
	ast = append(ast, NewAsteroid(v, rand(-179, 180), rand(3, 6), rand(1, 5)))
}

func (a *Asteroid) Update(dt float32) {
	a.Rotate(a.rotationSpeed)
	a.Move(sf.Vector2f{a.vel.X * dt, a.vel.Y})
	Wrap(a.Sprite)
}

func (a *Asteroid) Collides(s *sf.Sprite) bool {
	apos := a.GetPosition()
	brect := s.GetGlobalBounds()
	if apos.Y-a.height/2 < brect.Top+brect.Height && apos.Y+a.height/2 > brect.Top && apos.X-a.width/2 < brect.Left+brect.Width && apos.X+a.width/2 > brect.Left {
		return true
	}
	return false
}

func (a *Asteroid) CollidesPt(pt sf.Vector2f) bool {
	return distance(a.GetPosition(), pt) < a.radius
}
