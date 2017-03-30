package main

import (
	"runtime"
	"time"

	sf "github.com/zyedidia/sfml/v2.3/sfml"
)

const (
	screenWidth  = 800
	screenHeight = 600
)

var textures map[string]*sf.Texture

func LoadTexture(filename string) {
	texture := sf.NewTexture(filename)
	textures[filename[9:]] = texture
}

func main() {
	runtime.LockOSThread()

	textures = make(map[string]*sf.Texture)

	window := sf.NewRenderWindow(sf.VideoMode{screenWidth, screenHeight, 32}, "Rectangle", sf.StyleDefault, nil)
	window.SetVerticalSyncEnabled(true)
	window.SetFramerateLimit(60)

	LoadTexture("./images/playerShip1_blue.png")

	sprite := sf.NewSprite(textures["playerShip1_blue.png"])

	size := sprite.GetGlobalBounds()
	sprite.SetOrigin(sf.Vector2f{size.Width / 2, size.Height / 2})
	sprite.SetPosition(sf.Vector2f{screenWidth / 2, screenHeight / 2})

	var dt float32
	dt = 0

	for window.IsOpen() {
		delta := time.Now()
		if event := window.PollEvent(); event != nil {
			switch event.Type {
			case sf.EventClosed:
				window.Close()
			}
		}
		sprite.Rotate(1 * dt)

		window.Clear(sf.ColorWhite)
		window.Draw(sprite)
		window.Display()
		dt = float32(time.Since(delta)) / 10000000
	}
}
