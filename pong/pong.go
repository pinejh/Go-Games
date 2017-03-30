package main

import (
	"runtime"
	"time"

	sf "github.com/manyminds/gosfml"
)

const (
	screenWidth  = 400
	screenHeight = 300
	playerSpeed  = 2
)

var Window *sf.RenderWindow

var gameOver bool = false

func main() {
	runtime.LockOSThread()

	Window = sf.NewRenderWindow(sf.VideoMode{screenWidth, screenHeight, 32}, "Pong", sf.StyleDefault, sf.DefaultContextSettings())
	// Window.SetVSyncEnabled(true)
	// Window.SetFramerateLimit(60)

	p1 := NewPlayer(25)
	p2 := NewPlayer(screenWidth - 25)

	score1 := NewScore(screenWidth/2-25, 30)
	score2 := NewScore(screenWidth/2+25, 30)

	ball := NewBall()

	gameScore := []int{0, 0}

	var dt float32

	for Window.IsOpen() {
		start := time.Now()

		for event := Window.PollEvent(); event != nil; event = Window.PollEvent() {
			switch ev := event.(type) {
			case sf.EventKeyPressed:
				if ev.Code == sf.KeyW {
					p1.up = true
				}
				if ev.Code == sf.KeyS {
					p1.down = true
				}
				if ev.Code == sf.KeyUp {
					p2.up = true
				}
				if ev.Code == sf.KeyDown {
					p2.down = true
				}
				if ev.Code == sf.KeyEscape {
					Window.Close()
				}
				break
			case sf.EventKeyReleased:
				if ev.Code == sf.KeyW {
					p1.up = false
				}
				if ev.Code == sf.KeyS {
					p1.down = false
				}
				if ev.Code == sf.KeyUp {
					p2.up = false
				}
				if ev.Code == sf.KeyDown {
					p2.down = false
				}
			case sf.EventClosed:
				Window.Close()
				break
			}
		}

		Window.Clear(sf.ColorBlack())

		p1.Update(dt)
		p2.Update(dt)

		ball.Update(dt, p1, p2)

		if gameScore[0] != p1.score || gameScore[1] != p2.score {
			gameScore = []int{p1.score, p2.score}
			score1.SetNumber(gameScore[0])
			score2.SetNumber(gameScore[1])
		}
		if gameScore[0] >= 7 || gameScore[1] >= 7 {
			EndGame()
		}

		score1.Update()
		score2.Update()

		if !gameOver {
			Window.Display()
		} else {
			Window.Display()
			time.Sleep(time.Second * 5)
			Window.Close()
		}

		dt = float32(time.Since(start)) / 10000000
	}
}

func EndGame() {
	gameOver = true
}
