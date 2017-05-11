package main

import (
	"runtime"
	"time"

	sf "github.com/zyedidia/sfml/v2.3/sfml"
)

const (
	screenWidth  = 1260
	screenHeight = 700

	gravity = .25

	playerTopSpeed = 3.75
	playerAccel    = .35
	playerDecelG   = .55
	playerDecelA   = .98
	playerJump     = 8
)

var (
	debugCollisions = true
)

var res *Resources
var levels map[string]LevelMap
var level []*Tile

func main() {
	runtime.LockOSThread()

	res = NewResources()

	levels = preloadLevels()

	window := sf.NewRenderWindow(sf.VideoMode{screenWidth, screenHeight, 32}, "Platformer", sf.StyleDefault, nil)
	window.SetVerticalSyncEnabled(true)
	window.SetFramerateLimit(60)

	loadLevel("level-1")

	p1 := NewPlayer(1, screenWidth/4, screenHeight/2)
	p2 := NewPlayer(3, screenWidth*3/4, screenHeight/2)

	/*view := sf.NewView()
	view.SetSize(sf.Vector2f{screenWidth, screenHeight})
	view.SetCenter(player.GetPosition())

	window.SetView(view)*/

	var dt float32

	for window.IsOpen() {
		start := time.Now()
		if event := window.PollEvent(); event != nil {
			switch event.Type {
			case sf.EventClosed:
				window.Close()
			}
		}

		p1.Update(dt)
		p2.Update(dt)

		/*view.SetCenter(player.GetPosition())
		window.SetView(view)*/
		window.Clear(sf.Color{209, 244, 248, 255})
		for _, t := range level {
			window.Draw(t)
		}
		window.Draw(p1)
		window.Draw(p2)
		if debugCollisions {
			for _, b := range p1.box {
				window.Draw(DrawRect(b))
			}
			for _, b := range p2.box {
				window.Draw(DrawRect(b))
			}
		}
		window.Display()

		dt = float32(time.Since(start)) / float32(time.Second) * 60
	}
}
