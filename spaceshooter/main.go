package main

import (
	"runtime"
	"time"

	console "github.com/pinejh/console"
	sf "github.com/zyedidia/sfml/v2.3/sfml"
)

const (
	screenWidth  = 800
	screenHeight = 600

	shipSpeed    = 4
	rotateSpeed  = 4
	accelerate   = .25
	decelerate   = .15
	shotCooldown = 250 * time.Millisecond

	laserSpeed = 6

	numAsteroids = 6
)

var res *Resources
var lasers []*Laser
var ast []*Asteroid
var msgmap [][]int

func main() {
	runtime.LockOSThread()

	res = NewResources()

	msgmap = [][]int{[]int{0, 0, 0, 0, 0, 0, 0}, []int{0, 1, 1, 1, 1, 1, 0}, []int{0, 0, 1, 0, 0, 0, 0}, []int{0, 0, 0, 1, 0, 0, 0}, []int{0, 0, 0, 0, 1, 0, 0}, []int{0, 1, 1, 1, 1, 1, 0}, []int{0, 0, 0, 0, 0, 0, 0}, []int{0, 1, 1, 1, 1, 0, 0}, []int{0, 0, 0, 0, 0, 1, 0}, []int{0, 0, 0, 0, 0, 1, 0}, []int{0, 0, 0, 0, 0, 1, 0}, []int{0, 1, 1, 1, 1, 0, 0}, []int{0, 0, 0, 0, 0, 0, 0}, []int{0, 1, 0, 0, 0, 0, 0}, []int{0, 1, 0, 0, 0, 0, 0}, []int{0, 1, 1, 1, 1, 1, 0}, []int{0, 1, 0, 0, 0, 0, 0}, []int{0, 1, 0, 0, 0, 0, 0}, []int{0, 0, 0, 0, 0, 0, 0}, []int{0, 0, 0, 0, 0, 0, 0}, []int{0, 0, 0, 0, 0, 0, 0}}

	window := sf.NewRenderWindow(sf.VideoMode{screenWidth, screenHeight, 32}, "Spaceshooter", sf.StyleDefault, nil)
	window.SetVerticalSyncEnabled(true)
	window.SetFramerateLimit(60)

	var dt float32

	player := NewPlayer(sf.Vector2f{screenWidth / 2, screenHeight / 2})

	for len(ast) < numAsteroids {
		GenAsteroid()
	}

	for window.IsOpen() {
		start := time.Now()
		if event := window.PollEvent(); event != nil {
			switch event.Type {
			case sf.EventClosed:
				window.Close()
			}
		}

		for i := 0; i < len(ast); i++ {
			ast[i].Update(dt)
			if ast[i].dead {
				ast = append(ast[:i], ast[i+1:]...)
				i--
			}
		}
		player.Update(dt)
		for i := 0; i < len(lasers); i++ {
			v := lasers[i].GetPosition()
			if v.X < -54 || v.X > screenWidth+54 || v.Y < -54 || v.Y > screenHeight+54 {
				lasers = append(lasers[:i], lasers[i+1:]...)
				i--
			} else {
				lasers[i].Update(dt)
			}
		}

		window.Clear(sf.ColorBlack)
		for _, l := range lasers {
			window.Draw(l)
		}
		window.Draw(player)
		for _, a := range ast {
			window.Draw(a)
		}
		window.Display()
		dt = float32(time.Since(start)) / float32(time.Second) * 60
	}
}

func Wrap(g *sf.Sprite) {
	v := g.GetPosition()
	width := g.GetGlobalBounds().Width
	height := g.GetGlobalBounds().Height
	if v.X+width/2 < 0 {
		g.SetPosition(sf.Vector2f{screenWidth + width/2, v.Y})
	}
	if v.X-width/2 > screenWidth {
		g.SetPosition(sf.Vector2f{-width / 2, v.Y})
	}
	if v.Y+height/2 < 0 {
		g.SetPosition(sf.Vector2f{v.X, screenHeight + height/2})
	}
	if v.Y-height/2 > screenHeight {
		g.SetPosition(sf.Vector2f{v.X, -height / 2})
	}
}

func rand(max, min float32) float32 {
	return float32(console.RandInt(int(max), int(min)))
}
