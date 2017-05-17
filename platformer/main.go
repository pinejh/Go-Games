package main

import (
	"runtime"
	"strconv"
	"time"

	sf "github.com/zyedidia/sfml/v2.3/sfml"
)

const (
	screenWidth  = 1260
	screenHeight = 700

	gravity = .25

	gameLength = 30 //in Seconds

	playerTopSpeed = 3.75
	playerAccel    = .35
	playerDecelG   = .55
	playerDecelA   = .98
	playerJump     = 8
)

var (
	debugCollisions = false
	paused          = true
	gameStarted     = false
	gameOver        = false
	winner          = ""
	shutdown        = false
)

var res *Resources
var levels map[string]LevelMap
var level []*Tile
var score []int
var scoreBar *ScoreBar

var pauseText *sf.Text
var timer *sf.Text
var beginTime time.Time
var pausedTime time.Time

func main() {
	runtime.LockOSThread()

	res = NewResources()

	levels = preloadLevels()

	window := sf.NewRenderWindow(sf.VideoMode{screenWidth, screenHeight, 32}, "Platformer", sf.StyleDefault, nil)
	window.SetVerticalSyncEnabled(true)
	window.SetFramerateLimit(60)

	pauseText = sf.NewText("Press Enter to Begin", res.fonts["sigmar.ttf"], 72)
	pauseText.SetPosition(sf.Vector2f{screenWidth / 2, screenHeight / 2})
	r := pauseText.GetGlobalBounds()
	pauseText.SetOrigin(sf.Vector2f{r.Width / 2, r.Height / 2})

	loadLevel("level-2")

	p1 := NewPlayer(1, screenWidth/4, screenHeight-211)
	p2 := NewPlayer(3, screenWidth*3/4, screenHeight-211)

	score = make([]int, 2)
	scoreBar = newScoreBar()

	var dt float32

	for window.IsOpen() {
		start := time.Now()
		if event := window.PollEvent(); event != nil {
			switch event.Type {
			case sf.EventClosed:
				window.Close()
				break
			case sf.EventKeyReleased:
				if gameStarted && event.Key.Code == sf.KeyEscape && !gameOver {
					paused = !paused
					if paused {
						pauseTimer()
					} else {
						resumeTimer()
					}
				}
				if !gameStarted && event.Key.Code == sf.KeyReturn {
					go beginCountdown()
					gameStarted = true
				}
			}
		}

		if !paused && !gameOver {
			p1.Update(dt)
			p2.Update(dt)

			if p1.health == 0 && p2.health == 0 {
				gameOver = true
				winner = "Nobody"
			} else if p1.health == 0 {
				gameOver = true
				winner = "Player 2"
			} else if p2.health == 0 {
				gameOver = true
				winner = "Player 1"
			}

			p1Score, p2Score := 0, 0
			for _, t := range level {
				if t.GetTexture() == res.images["hillCaneGreen.png"] {
					p1Score++
				}
				if t.GetTexture() == res.images["hillCanePink.png"] {
					p2Score++
				}
			}
			score[0] = p1Score
			score[1] = p2Score

			scoreBar.Update()

			if time.Since(beginTime) > time.Duration(gameLength)*time.Second {
				gameOver = true
				if score[0] == score[1] {
					winner = "Tie"
				} else if score[0] > score[1] {
					winner = "Player 1"
				} else {
					winner = "Player 2"
				}
			} else {
				timer.SetString(strconv.Itoa(gameLength - int(time.Since(beginTime)/time.Second)))
				r := timer.GetGlobalBounds()
				timer.SetOrigin(sf.Vector2f{r.Width / 2, timer.GetOrigin().Y})
			}
		}

		/*view.SetCenter(player.GetPosition())
		window.SetView(view)*/
		window.Clear(sf.Color{209, 244, 248, 255})
		if !gameOver {
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
			window.Draw(p1.healthBar.background)
			window.Draw(p1.healthBar.bar)
			window.Draw(p1.healthBar.foreground)
			window.Draw(p2.healthBar.background)
			window.Draw(p2.healthBar.bar)
			window.Draw(p2.healthBar.foreground)
			window.Draw(scoreBar.background)
			window.Draw(scoreBar.p1)
			window.Draw(scoreBar.p2)
			window.Draw(scoreBar.foreground)
			if timer != nil && gameStarted {
				window.Draw(timer)
			}
			if paused {
				r := sf.NewRectangleShape(sf.Vector2f{screenWidth, screenHeight})
				r.SetFillColor(sf.Color{125, 125, 125, 150})
				window.Draw(r)
				window.Draw(pauseText)
			}
			window.Display()
		} else {
			winstr := ""
			if winner == "Tie" {
				winstr = "Its a Tie"
			} else if winner == "Nobody" {
				winstr = "Nobody Lived to the End"
			} else if winner == "Player 1" {
				winstr = "Player 1 is the Winner"
			} else if winner == "Player 2" {
				winstr = "Player 2 is the Winner"
			}
			win := sf.NewText(winstr, res.fonts["sigmar.ttf"], 72)
			win.SetPosition(sf.Vector2f{screenWidth / 2, screenHeight / 2})
			r := win.GetGlobalBounds()
			win.SetOrigin(sf.Vector2f{r.Width / 2, r.Height})
			window.Draw(win)
			window.Display()
			time.Sleep(time.Second * 5)
			window.Close()
		}

		dt = float32(time.Since(start)) / float32(time.Second) * 60
	}
}

func beginCountdown() {
	pauseText.SetCharacterSize(128)
	for i := 3; i >= 0; i-- {
		pauseText.SetString(strconv.Itoa(i))
		if i == 0 {
			pauseText.SetString("Start")
		}
		r := pauseText.GetGlobalBounds()
		pauseText.SetOrigin(sf.Vector2f{r.Width / 2, r.Height})
		time.Sleep(time.Second)
	}
	paused = false
	pauseText.SetString("Paused")
	r := pauseText.GetGlobalBounds()
	pauseText.SetOrigin(sf.Vector2f{r.Width / 2, pauseText.GetOrigin().Y})
	startTimer()
}

func startTimer() {
	timer = sf.NewText(strconv.Itoa(gameLength), res.fonts["sigmar.ttf"], 25)
	timer.SetPosition(sf.Vector2f{screenWidth / 2, 35})
	r := timer.GetGlobalBounds()
	timer.SetOrigin(sf.Vector2f{r.Width / 2, r.Height})
	beginTime = time.Now()
}

func pauseTimer() {
	pausedTime = time.Now()
}

func resumeTimer() {
	beginTime = beginTime.Add(time.Since(pausedTime))
}
