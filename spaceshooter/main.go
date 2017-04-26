package main

import (
	"runtime"
	"time"

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
var expl []*Explosion

func main() {
	runtime.LockOSThread()

	res = NewResources()

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

		for len(ast) < numAsteroids {
			GenAsteroid()
		}

		player.Update(dt)
		for _, l := range lasers {
			v := l.GetPosition()
			if v.X < -54 || v.X > screenWidth+54 || v.Y < -54 || v.Y > screenHeight+54 {
				l.dead = true
			} else {
				l.Update(dt)
			}
		}

		for _, a := range ast {
			a.Update(dt)
			for _, l := range lasers {
				if a.Collides(l.Sprite) {
					NewExplosion(a.GetPosition())
					a.dead = true
					l.dead = true
					sf.NewSound(res.sounds["sfx_explosion.wav"]).Play()
				}
			}
		}

		window.Clear(sf.ColorBlack)
		for _, e := range expl {
			window.Draw(e)
		}
		for _, l := range lasers {
			window.Draw(l)
		}
		window.Draw(player)
		for _, a := range ast {
			window.Draw(a)
		}
		window.Display()

		var tempLasers []*Laser
		for _, l := range lasers {
			if !l.dead {
				tempLasers = append(tempLasers, l)
			}
		}
		lasers = tempLasers

		var tempAst []*Asteroid
		for _, a := range ast {
			if !a.dead {
				tempAst = append(tempAst, a)
			}
		}
		ast = tempAst

		var tempExpl []*Explosion
		for _, e := range expl {
			if !e.done {
				tempExpl = append(tempExpl, e)
			}
		}
		expl = tempExpl

		dt = float32(time.Since(start)) / float32(time.Second) * 60
	}
}
