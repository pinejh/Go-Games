package main

import (
	"fmt"
	"io/ioutil"
	"path/filepath"

	sf "github.com/zyedidia/sfml/v2.3/sfml"
)

type Resources struct {
	images map[string]*sf.Texture
	sounds map[string]*sf.SoundBuffer
	fonts  map[string]*sf.Font
}

func NewResources() *Resources {
	r := new(Resources)

	r.images = make(map[string]*sf.Texture)
	r.sounds = make(map[string]*sf.SoundBuffer)
	r.fonts = make(map[string]*sf.Font)

	r.LoadImages("./assets/images")
	r.LoadSounds("./assets/sounds")
	r.LoadFonts("./assets/fonts")

	return r
}

func (r *Resources) LoadImages(dir string) {
	files, err := ioutil.ReadDir(dir)

	if err != nil {
		fmt.Println(err)
		return
	}

	for _, f := range files {
		if f.IsDir() {
			r.LoadImages(dir + "/" + f.Name())
		} else if filepath.Ext(f.Name()) == ".png" {
			texture := sf.NewTexture(dir + "/" + f.Name())
			r.images[f.Name()] = texture
		}
	}
}

func (r *Resources) LoadSounds(dir string) {
	files, err := ioutil.ReadDir(dir)

	if err != nil {
		fmt.Println(err)
		return
	}

	for _, f := range files {
		if f.IsDir() {
			r.LoadImages(dir + "/" + f.Name())
		} else if filepath.Ext(f.Name()) == ".ogg" || filepath.Ext(f.Name()) == ".wav" {
			soundBuffer := sf.NewSoundBuffer(dir + "/" + f.Name())
			r.sounds[f.Name()] = soundBuffer
		}
	}
}

func (r *Resources) LoadFonts(dir string) {
	files, err := ioutil.ReadDir(dir)

	if err != nil {
		fmt.Println(err)
		return
	}

	for _, f := range files {
		if f.IsDir() {
			r.LoadImages(dir + "/" + f.Name())
		} else if filepath.Ext(f.Name()) == ".ttf" {
			font := sf.NewFont(dir + "/" + f.Name())
			r.fonts[f.Name()] = font
		}
	}
}

func ParsePlayerSpritesheet() {
	pTextures = make(map[string]sf.Recti)
	pTextures["p1_front"] = sf.Recti{0, 196, 66, 92}
	pTextures["p1_stand"] = sf.Recti{67, 196, 66, 92}
	pTextures["p1_jump"] = sf.Recti{439, 94, 66, 92}
	pTextures["p1_crouch"] = sf.Recti{366, 98, 66, 71}
	pTextures["p1_walk0"] = sf.Recti{3, 2, 66, 92}
	pTextures["p1_walk1"] = sf.Recti{73 + 3, 2, 66, 92}
	pTextures["p1_walk2"] = sf.Recti{146 + 3, 2, 66, 92}
	pTextures["p1_walk3"] = sf.Recti{3, 98 + 2, 66, 92}
	pTextures["p1_walk4"] = sf.Recti{73 + 3, 98 + 2, 66, 92}
	pTextures["p1_walk5"] = sf.Recti{146 + 3, 98 + 2, 66, 92}
	pTextures["p1_walk6"] = sf.Recti{219 + 3, 2, 66, 92}
	pTextures["p1_walk7"] = sf.Recti{292 + 3, 2, 66, 92}
	pTextures["p1_walk8"] = sf.Recti{219 + 3, 98 + 2, 66, 92}
	pTextures["p1_walk9"] = sf.Recti{365 + 3, 2, 66, 92}
	pTextures["p1_walk10"] = sf.Recti{292 + 3, 98 + 2, 66, 92}
}
