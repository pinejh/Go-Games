package main

import (
	"runtime"
	"time"

	sf "github.com/zyedidia/sfml/v2.3/sfml"
)

const (
	screenWidth  = 1280
	screenHeight = 720

	gravity = .15
)

var res *Resources
var tiles []*Tile

func main() {
	runtime.LockOSThread()

	res = NewResources()

	window := sf.NewRenderWindow(sf.VideoMode{screenWidth, screenHeight, 32}, "Platformer", sf.StyleDefault, nil)
	window.SetVerticalSyncEnabled(true)
	window.SetFramerateLimit(60)

	player := NewPlayer()

	NewTile(sf.Vector2f{screenWidth / 2, screenHeight - 100}, true)

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

		window.Clear(sf.ColorWhite)
		for _, t := range tiles {
			window.Draw(t)
		}
		window.Draw(player)
		window.Display()

		dt = float32(time.Since(start)) / float32(time.Second) * 60
	}
}
