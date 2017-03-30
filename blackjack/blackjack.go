package main

import (
	"runtime"

	sf "github.com/zyedidia/sfml/v2.3/sfml"
)

const (
	screenWidth  = 800
	screenHeight = 600
)

var textures map[string]*sf.Texture

func loadTexture(filename string) {
	texture := sf.NewTexture("./res/images/" + filename)
	textures[filename] = texture
}

func main() {
	runtime.LockOSThread()

	textures = make(map[string]*sf.Texture)

	loadTexture("spades.png")
	loadTexture("clubs.png")
	loadTexture("diamonds.png")
	loadTexture("hearts.png")

	window := sf.NewRenderWindow(sf.VideoMode{screenWidth, screenHeight, 32}, "Rectangle", sf.StyleDefault, nil)
	window.SetVerticalSyncEnabled(true)
	window.SetFramerateLimit(60)

	card := NewCard(3, "hearts")
	card.SetPos(screenWidth/2, screenHeight/2)

	//var dt float32
	//dt = 0

	renderCards := make([]*Card, 0)
	renderCards = append(renderCards, card)

	for window.IsOpen() {
		//delta := time.Now()
		if event := window.PollEvent(); event != nil {
			switch event.Type {
			case sf.EventKeyReleased:
				card.Flip()
			case sf.EventClosed:
				window.Close()
			}
		}

		window.Clear(sf.Color{45, 165, 75, 255})
		for _, c := range renderCards {
			window.Draw(c)
			if c.show {
				window.Draw(c.dispsuit)
				//window.Draw(c.dispnum)
			} else {
			}
		}
		window.Display()
		//dt = float32(time.Since(delta)) / 10000000
	}
}
