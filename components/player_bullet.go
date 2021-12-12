package components

import (
	"Space_Invaders_Go/utilities"
	"github.com/veandco/go-sdl2/sdl"
)

type Bullet struct {
	Object *GameObject
	Speed  float64
}

const (
	bulletSize  = 32
	bulletSpeed = 8
)

func NewBullet(renderer *sdl.Renderer) (*Bullet, error) {
	tex, err := utilities.LoadTexture(renderer, "sprites/player-bullet.bmp")
	if err != nil {
		return nil, err
	}

	object := &GameObject{
		texture:  tex,
		Position: Vector{X: 0, Y: 0},
		Size:     Size{W: bulletSize, H: bulletSize},
		Active:   true}

	bullet := &Bullet{
		Object: object,
		Speed:  bulletSpeed,
	}

	return bullet, nil
}

func (bullet *Bullet) OnDraw(renderer *sdl.Renderer) error {
	size := bullet.Object.Size
	position := bullet.Object.Position

	err := renderer.Copy(
		bullet.Object.texture,
		&sdl.Rect{X: 0, Y: 0, W: size.W, H: size.H},
		&sdl.Rect{X: int32(position.X), Y: int32(position.Y), W: size.W, H: size.H})

	return err
}

func (bullet *Bullet) OnUpdate() error {
	return nil
}
