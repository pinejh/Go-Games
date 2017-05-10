package main

import (
	"encoding/json"
	//"fmt"
	"io/ioutil"

	cm "github.com/vova616/chipmunk"
	vect "github.com/vova616/chipmunk/vect"
	sf "github.com/zyedidia/sfml/v2.3/sfml"
)

type LevelMap [][]string

type Level struct {
	Name  string     //'json:"name"'
	Level [][]string //'json:"level"'
}

func preloadLevels() map[string]LevelMap {
	raw, _ := ioutil.ReadFile("./levels.json")

	var c []Level
	json.Unmarshal(raw, &c)
	m := make(map[string]LevelMap)
	for _, l := range c {
		m[l.Name] = LevelMap(l.Level)
	}
	return m
}

var staticBody *cm.Body

func loadLevel(name string) {
	l := [][]string(levels[name])
	var tiles []*Tile
	staticBody = cm.NewBodyStatic()
	static = []*cm.Shape{}
	for i, n := range l {
		for z, m := range n {
			blockName := ""
			blockCol := false
			if m == "1" {
				blockName = "grassCenter.png"
				blockCol = true
			}
			if m == "2" {
				blockName = "grassMid.png"
				blockCol = true
			}
			if m == "3" {
				blockName = "grass.png"
			}
			if m == "4" {
				blockName = "grassLeft.png"
				blockCol = true
			}
			if m == "5" {
				blockName = "grassRight.png"
				blockCol = true
			}
			if m == "o" {
				blockName = "castleCenter.png"
				blockCol = true
			}
			if m != "0" {
				x := screenWidth/2 + 35 + ((float32(z) - float32(len(n))/2) * 70)
				y := screenHeight + 35 - float32(len(l)-i)*70
				t := NewTile(sf.Vector2f{x, y}, blockName, blockCol)
				if t.col {
					static = append(static, cm.NewBox(vect.Vect{vect.Float(x), vect.Float(-y)}, 70, 70))
				}
				tiles = append(tiles, t)
			}
		}
	}
	for _, box := range static {
		staticBody.AddShape(box)
	}
	space.AddBody(staticBody)
	level = tiles
}
