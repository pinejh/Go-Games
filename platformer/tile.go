package main

import (
	sf "github.com/zyedidia/sfml/v2.3/sfml"
)

type Tile struct {
	*sf.Sprite
	col bool
}

func NewTile(pos sf.Vector2f, col bool) {
	t := new(Tile)
	t.Sprite = sf.NewSprite(res.images["grass.png"])
	t.SetPosition(pos)
	rect := t.GetGlobalBounds()
	t.SetOrigin(sf.Vector2f{rect.Width / 2, rect.Height / 2})
	t.col = col
	tiles = append(tiles, t)
}
