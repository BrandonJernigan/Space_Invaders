package components

import "github.com/veandco/go-sdl2/sdl"

type AlienSquid struct {
	Object *GameObject
}

const (
	alienSquidSprite = "sprites/alien-squid.bmp"
	alienWidth       = 41
	alienHeight      = 33
)

func NewAlienSquid(position Vector) *AlienSquid {
	object := &GameObject{
		Size:     Size{W: alienWidth, H: alienHeight},
		Position: position,
	}
	alien := &AlienSquid{Object: object}

	return alien
}

func (alien *AlienSquid) Load(renderer *sdl.Renderer) error {
	img, err := sdl.LoadBMP(alienSquidSprite)
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

func (alien *AlienSquid) Draw(renderer *sdl.Renderer) error {
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

func (alien *AlienSquid) Update() error {
	return nil
}

func (alien *AlienSquid) Unload() error {
	err := alien.Object.Texture.Destroy()
	if err != nil {
		return err
	}
	return nil
}
