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
