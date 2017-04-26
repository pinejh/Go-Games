package main

import (
	"time"

	math "github.com/chewxy/math32"
	sf "github.com/zyedidia/sfml/v2.3/sfml"
)

var msgmap = [][]int{[]int{0, 0, 0, 0, 0, 0, 0}, []int{0, 1, 1, 1, 1, 1, 0}, []int{0, 0, 1, 0, 0, 0, 0}, []int{0, 0, 0, 1, 0, 0, 0}, []int{0, 0, 0, 0, 1, 0, 0}, []int{0, 1, 1, 1, 1, 1, 0}, []int{0, 0, 0, 0, 0, 0, 0}, []int{0, 1, 1, 1, 1, 0, 0}, []int{0, 0, 0, 0, 0, 1, 0}, []int{0, 0, 0, 0, 0, 1, 0}, []int{0, 0, 0, 0, 0, 1, 0}, []int{0, 1, 1, 1, 1, 0, 0}, []int{0, 0, 0, 0, 0, 0, 0}, []int{0, 1, 0, 0, 0, 0, 0}, []int{0, 1, 0, 0, 0, 0, 0}, []int{0, 1, 1, 1, 1, 1, 0}, []int{0, 1, 0, 0, 0, 0, 0}, []int{0, 1, 0, 0, 0, 0, 0}, []int{0, 0, 0, 0, 0, 0, 0}, []int{0, 0, 0, 0, 0, 0, 0}, []int{0, 0, 0, 0, 0, 0, 0}}

func (p *Player) msg() {
	if p.canShoot {
		v := p.GetPosition()
		if msgmap[p.msgStep][0] == 0 {
			NewLaser(v.X+math.Cos((p.GetRotation())*math.Pi/180)*-30, v.Y+math.Sin((p.GetRotation())*math.Pi/180)*-30, p.GetRotation(), 8, true)
		}
		if msgmap[p.msgStep][1] == 0 {
			NewLaser(v.X+math.Cos((p.GetRotation())*math.Pi/180)*-20, v.Y+math.Sin((p.GetRotation())*math.Pi/180)*-20, p.GetRotation(), 8, true)
		}
		if msgmap[p.msgStep][2] == 0 {
			NewLaser(v.X+math.Cos((p.GetRotation())*math.Pi/180)*-10, v.Y+math.Sin((p.GetRotation())*math.Pi/180)*-10, p.GetRotation(), 8, true)
		}
		if msgmap[p.msgStep][3] == 0 {
			NewLaser(v.X, v.Y, p.GetRotation(), 8, true)
		}
		if msgmap[p.msgStep][4] == 0 {
			NewLaser(v.X+math.Cos((p.GetRotation())*math.Pi/180)*10, v.Y+math.Sin((p.GetRotation())*math.Pi/180)*10, p.GetRotation(), 8, true)
		}
		if msgmap[p.msgStep][5] == 0 {
			NewLaser(v.X+math.Cos((p.GetRotation())*math.Pi/180)*20, v.Y+math.Sin((p.GetRotation())*math.Pi/180)*20, p.GetRotation(), 8, true)
		}
		if msgmap[p.msgStep][6] == 0 {
			NewLaser(v.X+math.Cos((p.GetRotation())*math.Pi/180)*30, v.Y+math.Sin((p.GetRotation())*math.Pi/180)*30, p.GetRotation(), 8, true)
		}
		sf.NewSound(res.sounds["sfx_laser1.ogg"]).Play()
		p.canShoot = false
		p.msgStep--
		if p.msgStep < 0 {
			p.msgStep = len(msgmap) - 1
		}
		go func() {
			time.Sleep(75 * time.Millisecond)
			p.canShoot = true
		}()
	}
}

func (p *Player) hack() {
	v := p.GetPosition()
	NewLaser(v.X, v.Y, p.GetRotation(), laserSpeed, false)
	NewLaser(v.X, v.Y, p.GetRotation()+10, laserSpeed, false)
	NewLaser(v.X, v.Y, p.GetRotation()-10, laserSpeed, false)
	NewLaser(v.X, v.Y, p.GetRotation()+20, laserSpeed, false)
	NewLaser(v.X, v.Y, p.GetRotation()-20, laserSpeed, false)
	sf.NewSound(res.sounds["sfx_laser1.ogg"]).Play()
}
