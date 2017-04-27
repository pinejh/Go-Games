package main

import (
	math "github.com/chewxy/math32"
	sf "github.com/zyedidia/sfml/v2.3/sfml"
)

func distance(a, b sf.Vector2f) float32 {
	return math.Sqrt((a.X-b.X)*(a.X-b.X) + (a.Y-b.Y)*(a.Y-b.Y))
}

type CircCol struct {
	sf.Vector2f
	radius float32
}

func NewCircCol(pos sf.Vector2f, radius float32) *CircCol {
	cc := new(CircCol)
	cc.Vector2f = pos
	cc.radius = radius
	return cc
}

func (c *CircCol) CollidesCirc(c2 *CircCol) bool {
	return distance(c.Vector2f, c2.Vector2f) < c.radius+c2.radius
}

func (c *CircCol) CollidesRect(r sf.Rectf) bool {
	if c.X > r.Left && c.X < r.Left+r.Width {
		if c.Y > r.Top && c.Y < r.Top+r.Height {
			return true
		}
		if c.Y < r.Top && distance(c.Vector2f, sf.Vector2f{c.X, r.Top}) < c.radius {
			return true
		}
		if c.Y > r.Top+r.Height && distance(c.Vector2f, sf.Vector2f{c.X, r.Top + r.Height}) < c.radius {
			return true
		}
	}
	if c.Y > r.Top && c.Y < r.Top+r.Height {
		if c.X < r.Left && distance(c.Vector2f, sf.Vector2f{r.Left, c.Y}) < c.radius {
			return true
		}
		if c.X > r.Left+r.Width && distance(c.Vector2f, sf.Vector2f{r.Left + r.Width, c.Y}) < c.radius {
			return true
		}
	}
	if c.X < r.Left {
		if c.Y < r.Top && distance(c.Vector2f, sf.Vector2f{r.Left, r.Top}) < c.radius {
			return true
		}
		if c.Y > r.Top+r.Height && distance(c.Vector2f, sf.Vector2f{r.Left, r.Top + r.Height}) < c.radius {
			return true
		}
	}
	if c.X > r.Left+r.Width {
		if c.Y < r.Top && distance(c.Vector2f, sf.Vector2f{r.Left + r.Width, r.Top}) < c.radius {
			return true
		}
		if c.Y > r.Top+r.Height && distance(c.Vector2f, sf.Vector2f{r.Left + r.Width, r.Top + r.Height}) < c.radius {
			return true
		}
	}
	return false
}
