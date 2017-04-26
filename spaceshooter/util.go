package main

import (
	console "github.com/pinejh/console"
	sf "github.com/zyedidia/sfml/v2.3/sfml"
)

func Wrap(g *sf.Sprite) {
	v := g.GetPosition()
	width := g.GetGlobalBounds().Width
	height := g.GetGlobalBounds().Height
	if v.X+width/2 < 0 {
		g.SetPosition(sf.Vector2f{screenWidth + width/2, v.Y})
	}
	if v.X-width/2 > screenWidth {
		g.SetPosition(sf.Vector2f{-width / 2, v.Y})
	}
	if v.Y+height/2 < 0 {
		g.SetPosition(sf.Vector2f{v.X, screenHeight + height/2})
	}
	if v.Y-height/2 > screenHeight {
		g.SetPosition(sf.Vector2f{v.X, -height / 2})
	}
}

func Collides(a, b *sf.Sprite) bool {
	arect, brect := a.GetGlobalBounds(), b.GetGlobalBounds()
	return arect.Top < brect.Top+brect.Height && arect.Top+arect.Height > brect.Top && arect.Left < brect.Left+brect.Width && arect.Left+arect.Width > brect.Left
}

func rand(max, min float32) float32 {
	return float32(console.RandInt(int(max), int(min)))
}
