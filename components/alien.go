package components

import (
	"github.com/veandco/go-sdl2/sdl"
)

type Enemy struct {
	Object *GameObject
	Speed  int32
	Type   string
}

const (
	alienSquidSprite = "sprites/alien-squid.bmp"
	alienJellySprite = "sprites/alien-jelly.bmp"
	alienCrabSprite  = "sprites/alien-crab.bmp"
	alienSquidWidth  = 41
	alienSquidHeight = 33
	alienJellyWidth  = 49
	alienJellyHeight = 31
	alienCrabWidth   = 52
	alienCrabHeight  = 27
)

func NewAlien(position Vector, alienType string, speed int32) *Enemy {
	var width int
	var height int

	switch alienType {
	case "squid":
		width = alienSquidWidth
		height = alienSquidHeight
		break
	case "jelly":
		width = alienJellyWidth
		height = alienJellyHeight
		break
	case "crab":
		width = alienCrabWidth
		height = alienCrabHeight
	}

	object := &GameObject{
		Size:     Size{W: int32(width), H: int32(height)},
		Position: position,
		Active:   true,
	}
	alien := &Enemy{Object: object, Type: alienType, Speed: speed}

	return alien
}

func (alien *Enemy) Load(renderer *sdl.Renderer) error {
	var sprite string
	switch alien.Type {
	case "squid":
		sprite = alienSquidSprite
		break
	case "jelly":
		sprite = alienJellySprite
		break
	case "crab":
		sprite = alienCrabSprite
		break
	}

	img, err := sdl.LoadBMP(sprite)
	if err != nil {
		return err
	}

	tex, err := renderer.CreateTextureFromSurface(img)
	if err != nil {
		return err
	}

	alien.Object.Texture = tex
	img.Free()
	return nil
}

func (alien *Enemy) Draw(renderer *sdl.Renderer) error {
	err := renderer.Copy(
		alien.Object.Texture,
		&sdl.Rect{X: 0, Y: 0, W: alien.Object.Size.W, H: alien.Object.Size.H},
		&sdl.Rect{
			X: alien.Object.Position.X,
			Y: alien.Object.Position.Y,
			W: alien.Object.Size.W,
			H: alien.Object.Size.H})

	return err
}

func (alien *Enemy) Update() error {
	alien.Object.Position.Y += alien.Speed
	return nil
}

func (alien *Enemy) Unload() error {
	err := alien.Object.Texture.Destroy()
	if err != nil {
		return err
	}
	return nil
}

func (alien *Enemy) CheckActive() bool {
	return alien.Object.Active
}

func (alien *Enemy) OnCollision() {
	alien.Object.Active = false
}
