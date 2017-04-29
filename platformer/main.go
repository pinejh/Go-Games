package main

import (
	"runtime"
	"time"

	sf "github.com/zyedidia/sfml/v2.3/sfml"
)

const (
	screenWidth  = 1280
	screenHeight = 720

	gravity = .25

	playerTopSpeed = 3
	playerAccel    = .35
	playerDecelG   = .35
	playerDecelA   = .95
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

	player := NewPlayer()

	view := sf.NewView()
	view.SetSize(sf.Vector2f{screenWidth, screenHeight})
	view.SetCenter(player.GetPosition())

	window.SetView(view)

	var dt float32

	for window.IsOpen() {
		start := time.Now()
		if event := window.PollEvent(); event != nil {
			switch event.Type {
			case sf.EventClosed:
				window.Close()
			}
		}

		player.Update(dt)

		view.SetCenter(player.GetPosition())
		window.SetView(view)
		window.Clear(sf.ColorWhite)
		for _, t := range level {
			window.Draw(t)
		}
		window.Draw(player)
		window.Display()

		dt = float32(time.Since(start)) / float32(time.Second) * 60
	}
}
