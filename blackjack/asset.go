package main

import (
	"io/ioutil"

	sf "github.com/zyedidia/sfml/v2.3/sfml"
)

type Resources struct {
	images map[string]*sf.Texture
	fonts  map[string]*sf.Font
	sounds map[string]*sf.SoundBuffer
}

func NewResources(dir string) *Resources {
	r := &Resources{make(map[string]*sf.Texture), make(map[string]*sf.Font), make(map[string]*sf.SoundBuffer)}
	r.LoadImages(dir + "/images")
	r.LoadFonts(dir + "/fonts")
	r.LoadSounds(dir + "/sounds")
	return r
}

func (r *Resources) LoadImages(dir string) {
	files, _ := ioutil.ReadDir(dir)
	for _, f := range files {
		if f.IsDir() {
			r.LoadImages(dir + "/" + f.Name())
		} else {
			r.images[f.Name()] = sf.NewTexture(dir + "/" + f.Name())
		}
	}
}
func (r *Resources) LoadFonts(dir string) {
	files, _ := ioutil.ReadDir(dir)
	for _, f := range files {
		if f.IsDir() {
			r.LoadFonts(dir + "/" + f.Name())
		} else {
			r.fonts[f.Name()] = sf.NewFont(dir + "/" + f.Name())
		}
	}
}
func (r *Resources) LoadSounds(dir string) {
	files, _ := ioutil.ReadDir(dir)
	for _, f := range files {
		if f.IsDir() {
			r.LoadSounds(dir + "/" + f.Name())
		} else {
			r.sounds[f.Name()] = sf.NewSoundBuffer(dir + "/" + f.Name())
		}
	}
}
