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
	bar.foreground.SetOutlineThickness(1)
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
