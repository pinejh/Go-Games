package main

import (
	"strconv"
	"time"

	sf "github.com/zyedidia/sfml/v2.3/sfml"
)

type Explosion struct {
	*sf.Sprite
	stage int
	done  bool
}

func NewExplosion(pos sf.Vector2f) *Explosion {
	e := new(Explosion)
	e.stage = 0
	e.Sprite = sf.NewSprite(res.images["expl_01_"+strconv.Itoa(e.stage)+".png"])
	e.SetPosition(pos)
	rect := e.GetGlobalBounds()
	e.SetOrigin(sf.Vector2f{rect.Width / 2, rect.Height / 2})
	e.done = false
	return e
}

func (e *Explosion) Explode() {
	go func() {
		for e.stage < 23 {
			time.Sleep(time.Millisecond * 25)
			e.stage++
			e.SetTexture(res.images["expl_01_"+strconv.Itoa(e.stage)+".png"], false)
		}
		time.Sleep(time.Millisecond * 25)
		e.done = true
	}()
}
