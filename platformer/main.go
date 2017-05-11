package main

import (
	"runtime"
	"time"

	cm "github.com/vova616/chipmunk"
	vect "github.com/vova616/chipmunk/vect"
	sf "github.com/zyedidia/sfml/v2.3/sfml"
)

const (
	screenWidth  = 1260
	screenHeight = 700

	gravity = -900

	playerTopSpeed = 300
	playerAccel    = 15
	playerDecelG   = .55
	playerDecelA   = .98
	playerJump     = 400
)

var res *Resources
var space *cm.Space
var static []*cm.Shape
var levels map[string]LevelMap
var level []*Tile

func main() {
	runtime.LockOSThread()

	res = NewResources()

	levels = preloadLevels()

	window := sf.NewRenderWindow(sf.VideoMode{screenWidth, screenHeight, 32}, "Don't Move or Live", sf.StyleDefault, nil)
	window.SetVerticalSyncEnabled(true)
	window.SetFramerateLimit(60)

	space = cm.NewSpace()
	space.Gravity = vect.Vect{0, gravity}

	loadLevel("level-1")

	p1 := NewPlayer(1, screenWidth/4, screenHeight/2)
	space.AddBody(p1.Body)
	p2 := NewPlayer(3, screenWidth*3/4, screenHeight/2)
	space.AddBody(p2.Body)

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

		space.Step(vect.Float(1.0 / 60.0))

		p1.Update(dt)
		p2.Update(dt)

		/*view.SetCenter(player.GetPosition())
		window.SetView(view)*/
		window.Clear(sf.Color{209, 244, 248, 255})
		for _, t := range level {
			window.Draw(t.Sprite)
		}
		window.Draw(p1.Sprite)
		window.Draw(p2.Sprite)
		//rect := sf.NewRectangleShape(sf.Vector2f{p1.Body.Shapes[0].width})
		//rect.SetFillColor(sf.ColorBlack)

		//window.Draw(rect)
		window.Display()

		dt = float32(time.Since(start)) / float32(time.Second) * 60
	}
}
