package main

import (
	sf "github.com/zyedidia/sfml/v2.3/sfml"
)

type HealthBar struct {
	width, height float32
	x, y          float32

	background *sf.RectangleShape
	bar        *sf.RectangleShape
	foreground *sf.RectangleShape
}

func newHealthBar(width, height, x, y float32) (bar *HealthBar) {
	bar = new(HealthBar)
	bar.width, bar.height = width, height
	bar.x, bar.y = x, y
	bar.background = NewRectangle(bar.width, bar.height, bar.x, bar.y)
	bar.background.SetFillColor(sf.Color{125, 125, 125, 255})
	bar.foreground = NewRectangle(bar.width, bar.height, bar.x, bar.y)
	bar.foreground.SetFillColor(sf.Color{0, 0, 0, 0})
	bar.foreground.SetOutlineColor(sf.ColorWhite)
	bar.foreground.SetOutlineThickness(10)
	bar.bar = sf.NewRectangleShape(sf.Vector2f{bar.width, bar.height})
	bar.bar.SetOrigin(sf.Vector2f{0, bar.height / 2})
	bar.bar.SetPosition(sf.Vector2f{bar.x - bar.width/2, bar.y})
	bar.bar.SetFillColor(sf.Color{75, 255, 75, 255})
	return
}

func newHealthBarStd(x, y float32) *HealthBar {
	return newHealthBar(200, 25, x, y)
}

func (b *HealthBar) SetHealth(health float32) {
	health *= 2
	b.bar.SetSize(sf.Vector2f{health, b.height})
	if health > b.width/2 {
		b.bar.SetFillColor(sf.Color{75, 255, 75, 255})
	} else if health > b.width/4 {
		b.bar.SetFillColor(sf.Color{255, 255, 75, 255})
	} else {
		b.bar.SetFillColor(sf.Color{255, 75, 75, 255})
	}
}

type ScoreBar struct {
	background *sf.RectangleShape
	p1         *sf.RectangleShape
	p2         *sf.RectangleShape
	foreground *sf.RectangleShape
}

func newScoreBar() (bar *ScoreBar) {
	bar = new(ScoreBar)
	bar.background = sf.NewRectangleShape(sf.Vector2f{250, 35})
	bar.background.SetPosition(sf.Vector2f{screenWidth / 2, 35})
	bar.background.SetOrigin(sf.Vector2f{125, 17.5})
	bar.background.SetFillColor(sf.Color{125, 125, 125, 255})
	bar.foreground = sf.NewRectangleShape(sf.Vector2f{250, 35})
	bar.foreground.SetPosition(sf.Vector2f{screenWidth / 2, 35})
	bar.foreground.SetOrigin(sf.Vector2f{125, 17.5})
	bar.foreground.SetFillColor(sf.Color{0, 0, 0, 0})
	bar.foreground.SetOutlineColor(sf.Color{125, 125, 125, 255})
	bar.foreground.SetOutlineThickness(10)
	bar.p1 = sf.NewRectangleShape(sf.Vector2f{125, 35})
	bar.p1.SetPosition(sf.Vector2f{screenWidth/2 - 125, 35})
	bar.p1.SetOrigin(sf.Vector2f{0, 17.5})
	bar.p1.SetFillColor(sf.Color{139, 207, 186, 255})
	bar.p2 = sf.NewRectangleShape(sf.Vector2f{125, 35})
	bar.p2.SetPosition(sf.Vector2f{screenWidth/2 + 125, 35})
	rect := bar.p2.GetGlobalBounds()
	bar.p2.SetOrigin(sf.Vector2f{rect.Width, 17.5})
	bar.p2.SetFillColor(sf.Color{243, 155, 183, 255})
	return
}

func (bar *ScoreBar) Update() {
	p1, p2 := float32(score[0]), float32(score[1])
	total := p1 + p2
	bar.p1.SetSize(sf.Vector2f{p1 / total * 250, 35})
	bar.p1.SetOrigin(sf.Vector2f{0, 17.5})
	bar.p2.SetSize(sf.Vector2f{p2 / total * 250, 35})
	rect := bar.p2.GetGlobalBounds()
	bar.p2.SetOrigin(sf.Vector2f{rect.Width, 17.5})
}

type Timer struct {
	time     float32
	currTime float32

	background *sf.CircleShape
	timer      *sf.CircleShape
	foreground *sf.CircleShape
}
