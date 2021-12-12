package components

import (
	"Space_Invaders_Go/utilities"
	"github.com/veandco/go-sdl2/sdl"
)

type Enemy struct {
	Object *GameObject
	Speed  float64
}

var enemies []*Enemy

const (
	enemySpeed = 4
	enemySize  = 64
)

func NewEnemy(renderer *sdl.Renderer, enemyType string, position Vector) (*Enemy, error) {
	filepath := "sprites/enemy-" + enemyType + ".bmp"
	tex, err := utilities.LoadTexture(renderer, filepath)
	if err != nil {
		return nil, err
	}

	object := &GameObject{
		texture:  tex,
		Position: Vector{X: position.X, Y: position.Y},
		Size:     Size{W: enemySize, H: enemySize},
		Active:   true}

	enemy := &Enemy{
		Object: object,
		Speed:  enemySpeed,
	}

	enemies = append(enemies, enemy)
	return enemy, nil
}

func (enemy *Enemy) OnDraw(renderer *sdl.Renderer) error {
	size := enemy.Object.Size
	position := enemy.Object.Position

	err := renderer.Copy(
		enemy.Object.texture,
		&sdl.Rect{X: 0, Y: 0, W: size.W, H: size.H},
		&sdl.Rect{X: int32(position.X), Y: int32(position.Y), W: size.W, H: size.H})

	return err
}

func (enemy *Enemy) OnUpdate() error {
	return nil
}

func (enemy *Enemy) CheckActive() bool {
	return enemy.Object.Active
}
