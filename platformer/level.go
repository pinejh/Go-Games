package main

import (
	"encoding/json"
	"io/ioutil"

	sf "github.com/zyedidia/sfml/v2.3/sfml"
)

type LevelMap [][]int

type Level struct {
	Name  string  //'json:"name"'
	Level [][]int //'json:"level"'
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

func loadLevel(name string) {
	l := [][]int(levels[name])
	var tiles []*Tile
	for i, n := range l {
		for z, m := range n {
			blockName := ""
			blockCol := false
			if m == 1 {
				blockName = "grassCenter.png"
				blockCol = true
			}
			if m == 2 {
				blockName = "grassMid.png"
				blockCol = true
			}
			if m == 3 {
				blockName = "grassLeft.png"
				blockCol = true
			}
			if m == 4 {
				blockName = "grassRight.png"
				blockCol = true
			}
			if m != 0 {
				tiles = append(tiles, NewTile(sf.Vector2f{float32(z) * 70, float32(i) * 70}, blockName, blockCol))
			}
		}
	}
	level = tiles
}
