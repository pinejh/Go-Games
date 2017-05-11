package main

import (
	r "math/rand"
	"time"

	math "github.com/chewxy/math32"
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

func RandInt(min, max int) int {
	r.Seed(time.Now().UTC().UnixNano())
	return r.Intn(max-min) + min
}

func rand(max, min float32) float32 {
	return float32(RandInt(int(max), int(min)))
}

func distance(a, b sf.Vector2f) float32 {
	return math.Sqrt((a.X-b.X)*(a.X-b.X) + (a.Y-b.Y)*(a.Y-b.Y))
}
