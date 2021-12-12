package components

import (
	"Space_Invaders_Go/utilities"
	"github.com/veandco/go-sdl2/sdl"
	"time"
)

type Player struct {
	Object       *GameObject
	Speed        float64
	ShotCoolDown time.Duration
}

const (
	playerSpeed        = 4
	playerSize         = 64
	playerShotCoolDown = time.Millisecond * 250
)

func NewPlayer(renderer *sdl.Renderer) (*Player, error) {
	tex, err := utilities.LoadTexture(renderer, "../sprites/player.bmp")
	if err != nil {
		return nil, err
	}

	object := &GameObject{
		texture:  tex,
		Position: Vector{X: 0, Y: 0},
		Size:     Size{W: playerSize, H: playerSize},
		Active:   true}

	player := &Player{
		Object:       object,
		Speed:        playerSpeed,
		ShotCoolDown: playerShotCoolDown,
	}

	return player, nil
}

func (player *Player) OnDraw(renderer *sdl.Renderer) error {
	size := player.Object.Size

	err := renderer.Copy(
		player.Object.texture,
		&sdl.Rect{X: 0, Y: 0, W: size.W, H: size.H},
		&sdl.Rect{X: 0, Y: 0, W: size.W, H: size.H})
	if err != nil {
		return err
	}
	return nil
}

func (player *Player) OnUpdate() error {
	return nil
}
