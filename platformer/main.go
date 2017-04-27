package main

import (
	"runtime"
	"time"

	sf "github.com/zyedidia/sfml/v2.3/sfml"
)

const (
	screenWidth  = 1280
	screenHeight = 720
)

var res *Resources

func main() {
	runtime.LockOSThread()

	res = NewResources()

	window := sf.NewRenderWindow(sf.VideoMode{screenWidth, screenHeight, 32}, "Spaceshooter", sf.StyleDefault, nil)
	window.SetVerticalSyncEnabled(true)
	window.SetFramerateLimit(60)

	player := NewPlayer()

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
		window.Draw(player)
		window.Display()

		dt = float32(time.Since(start)/time.Second) * 60
	}
}
