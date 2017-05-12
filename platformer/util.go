package main

import (
	math "github.com/chewxy/math32"
	sf "github.com/zyedidia/sfml/v2.3/sfml"
)

func clamp(num, min, max float32) float32 {
	if num < min {
		return min
	}
	if num > max {
		return max
	}
	return num
}

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
	return (c.X > r.Left && c.X < r.Left+r.Width && (c.Y > r.Top && c.Y < r.Top+r.Height || c.Y < r.Top && distance(c.Vector2f, sf.Vector2f{c.X, r.Top}) < c.radius || c.Y > r.Top+r.Height && distance(c.Vector2f, sf.Vector2f{c.X, r.Top + r.Height}) < c.radius) || c.Y > r.Top && c.Y < r.Top+r.Height && (c.X < r.Left && distance(c.Vector2f, sf.Vector2f{r.Left, c.Y}) < c.radius || c.X > r.Left+r.Width && distance(c.Vector2f, sf.Vector2f{r.Left + r.Width, c.Y}) < c.radius) || c.X < r.Left && (c.Y < r.Top && distance(c.Vector2f, sf.Vector2f{r.Left, r.Top}) < c.radius || c.Y > r.Top+r.Height && distance(c.Vector2f, sf.Vector2f{r.Left, r.Top + r.Height}) < c.radius) || c.X > r.Left+r.Width && (c.Y < r.Top && distance(c.Vector2f, sf.Vector2f{r.Left + r.Width, r.Top}) < c.radius || c.Y > r.Top+r.Height && distance(c.Vector2f, sf.Vector2f{r.Left + r.Width, r.Top + r.Height}) < c.radius))
}

func (c *CircCol) CollidesPt(v sf.Vector2f) bool {
	return distance(c.Vector2f, v) < c.radius
}

func (c *CircCol) CollidesTileTop(t *Tile) (sf.Vector2f, bool) {
	r := t.GetGlobalBounds()
	if (c.X > r.Left && c.X < r.Left+r.Width) && ((c.Y > r.Top && c.Y < r.Top+r.Height) || (c.Y < r.Top && distance(c.Vector2f, sf.Vector2f{c.X, r.Top}) < c.radius)) || (c.Y < r.Top && distance(c.Vector2f, sf.Vector2f{r.Left, r.Top}) < c.radius) || (c.Y > r.Left+r.Width && distance(c.Vector2f, sf.Vector2f{r.Left + r.Width, r.Top}) < c.radius) {
		return sf.Vector2f{0, r.Top - (c.Y + c.radius)}, true
	}
	return sf.Vector2f{0, 0}, false
}

func DrawRect(rect sf.Rectf) (r *sf.RectangleShape) {
	r = sf.NewRectangleShape(sf.Vector2f{rect.Width, rect.Height})
	r.SetPosition(sf.Vector2f{rect.Left, rect.Top})
	r.SetFillColor(sf.ColorBlack)
	return
}

func NewRectangle(width, height, x, y float32) (r *sf.RectangleShape) {
	r = sf.NewRectangleShape(sf.Vector2f{width, height})
	r.SetOrigin(sf.Vector2f{width / 2, height / 2})
	r.SetPosition(sf.Vector2f{x, y})
	return
}
