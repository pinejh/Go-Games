package main

import (
	"strconv"
	"time"

	sf "github.com/zyedidia/sfml/v2.3/sfml"
)

type Explosion struct {
	*sf.Sprite
	animStage int
	animLen   int
	done      bool
}

func NewExplosion(pos sf.Vector2f) *Explosion {
	e := new(Explosion)
	e.animStage = 0
	e.animLen = 23
	e.Sprite = sf.NewSprite(res.images["expl_01_"+strconv.Itoa(e.animStage)+".png"])
	e.SetPosition(pos)
	rect := e.GetGlobalBounds()
	e.SetOrigin(sf.Vector2f{rect.Width / 2, rect.Height / 2})
	e.Scale(sf.Vector2f{2, 2})
	e.done = false
	expl = append(expl, e)
	e.Animate()
	return e
}

func (e *Explosion) Animate() {
	go func() {
		for e.animStage < e.animLen {
			time.Sleep(time.Millisecond * 20)
			e.animStage++
			e.SetTexture(res.images["expl_01_"+strconv.Itoa(e.animStage)+".png"], false)
		}
		time.Sleep(time.Millisecond * 20)
		e.done = true
	}()
}
