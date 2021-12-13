package components

import (
	"Space_Invaders_Go/utilities"
	"github.com/veandco/go-sdl2/sdl"
)

var direction = 0

const shieldSpeed = 2

func NewShield(renderer *sdl.Renderer) (*GameObject, error) {
	tex, err := utilities.LoadTexture(renderer, "sprites/shield.bmp")
	if err != nil {
		return nil, err
	}

	shield := &GameObject{
		texture:  tex,
		Position: Vector{X: 0, Y: 500},
		Size:     Size{W: 128, H: 128},
		Active:   true}

	return shield, nil
}

func (shield *GameObject) OnDraw(renderer *sdl.Renderer) error {
	position := shield.Position
	size := shield.Size

	err := renderer.Copy(
		shield.texture,
		&sdl.Rect{X: 0, Y: 0, W: size.W, H: size.H},
		&sdl.Rect{X: int32(position.X), Y: int32(position.Y), W: size.W, H: size.H})

	return err
}

func (shield *GameObject) OnUpdate() error {
	if direction == 0 {
		if int(shield.Position.X)+shieldSpeed >= int(screenWidth)-128 {
			direction = 1
		}
		shield.Position.X += shieldSpeed
	} else {
		if int(shield.Position.X)-shieldSpeed <= 0 {
			direction = 0
		}
		shield.Position.X -= shieldSpeed
	}

	return nil
}

func (shield *GameObject) OnCollision() error {
	shield.Active = false
	return nil
}

func (shield *GameObject) CheckActive() bool {
	return shield.Active
}
