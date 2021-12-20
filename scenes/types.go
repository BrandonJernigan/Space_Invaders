package scenes

import (
	"github.com/veandco/go-sdl2/sdl"
)

type Scene interface {
	Load(renderer *sdl.Renderer) error
	Draw(renderer *sdl.Renderer) error
	Update(renderer *sdl.Renderer) error
	Unload() error
}

type Text struct {
	Texture *sdl.Texture
	Surface *sdl.Surface
}
