package utilities

import (
	"github.com/veandco/go-sdl2/sdl"
)

func LoadTexture(renderer *sdl.Renderer, filename string) (*sdl.Texture, error) {
	img, err := sdl.LoadBMP(filename)
	if err != nil {
		return nil, err
	}
	defer img.Free()

	tex, err := renderer.CreateTextureFromSurface(img)
	if err != nil {
		return nil, err
	}

	return tex, nil
}
