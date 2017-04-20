package main

import (
	"strconv"

	math "github.com/chewxy/math32"
	console "github.com/pinejh/console"
	sf "github.com/zyedidia/sfml/v2.3/sfml"
)

type Asteroid struct {
	*sf.Sprite
	width         float32
	height        float32
	rotationSpeed float32
	vel           sf.Vector2f
	dead          bool
}

func NewAsteroid(pos sf.Vector2f, angle, speed, rotationSpeed float32) *Asteroid {
	a := new(Asteroid)
	a.Sprite = sf.NewSprite(res.images["asteroid"+strconv.Itoa(console.RandInt(1, 5))+".png"])
	a.SetPosition(pos)
	rect := a.GetGlobalBounds()
	a.width, a.height = rect.Width, rect.Height
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

/*
func CloneAsteroid(a *Asteroid) *Asteroid {
	ast := new(Asteroid)
	ast.Sprite = a.Sprite
	ast.SetPosition(a.GetPosition())
	ast.SetRotation(a.GetRotation())
	ast.width = a.width
	ast.height = a.height
	ast.rotationSpeed = a.rotationSpeed
	ast.vel = a.vel
	ast.dead = a.dead
	return ast
}
func CloneAsteroidAt(a *Asteroid, pos sf.Vector2f) {
	asteroid := CloneAsteroid(a)
	asteroid.SetPosition(pos)
	ast = append(ast, asteroid)
}
*/

func (a *Asteroid) Update(dt float32) {
	a.Rotate(a.rotationSpeed)
	a.Move(sf.Vector2f{a.vel.X * dt, a.vel.Y})
	Wrap(a.Sprite)
}

/*
func (a *Asteroid) Wrap() {
	v := a.GetPosition()
	width, height := a.GetGlobalBounds().Width, a.GetGlobalBounds().Height
	if v.X+width/2 < 0 {
		CloneAsteroidAt(a, sf.Vector2f{screenWidth + width/2, v.Y})
	}
	if v.X-width/2 > screenWidth {
		CloneAsteroidAt(a, sf.Vector2f{-width / 2, v.Y})
	}
	if v.Y+height/2 < 0 {
		CloneAsteroidAt(a, sf.Vector2f{v.X, screenHeight + height/2})
	}
	if v.Y-height/2 > screenHeight {
		CloneAsteroidAt(a, sf.Vector2f{v.X, -height / 2})
	}
	if v.X+width/2 < 0 || v.X-width/2 > screenWidth || v.Y+height/2 < 0 || v.Y-height/2 > screenHeight {
		a.dead = true
	}
}
*/
