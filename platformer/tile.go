package main

import (
	//cm "github.com/vova616/chipmunk"
	//vect "github.com/vova616/chipmunk/vect"
	sf "github.com/zyedidia/sfml/v2.3/sfml"
)

type Tile struct {
	*sf.Sprite

	col bool
}

func NewTile(pos sf.Vector2f, tilename string, col bool) *Tile {
	t := new(Tile)
	t.Sprite = sf.NewSprite(res.images[tilename])
	t.Sprite.SetPosition(pos)
	rect := t.GetGlobalBounds()
	t.SetOrigin(sf.Vector2f{rect.Width / 2, rect.Height / 2})
	t.col = col
	return t
}
