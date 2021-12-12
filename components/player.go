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

var screenWidth int32
var screenHeight int32

func NewPlayer(renderer *sdl.Renderer) (*Player, error) {
	tex, err := utilities.LoadTexture(renderer, "sprites/player.bmp")
	if err != nil {
		return nil, err
	}

	screenWidth, screenHeight, _ = renderer.GetOutputSize()
	positionX := (screenWidth - playerSize) / 2.00
	positionY := screenHeight - playerSize

	object := &GameObject{
		texture:  tex,
		Position: Vector{X: float64(positionX), Y: float64(positionY)},
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
	position := player.Object.Position

	err := renderer.Copy(
		player.Object.texture,
		&sdl.Rect{X: 0, Y: 0, W: size.W, H: size.H},
		&sdl.Rect{X: int32(position.X), Y: int32(position.Y), W: size.W, H: size.H})

	return err
}

func (player *Player) OnUpdate() error {
	keys := sdl.GetKeyboardState()

	if keys[sdl.SCANCODE_A] == 1 {
		if player.Object.Position.X-playerSpeed > 0 {
			player.Object.Position.X -= player.Speed
		}
	} else if keys[sdl.SCANCODE_D] == 1 {
		if player.Object.Position.X+playerSpeed <= float64(screenWidth)-playerSize {
			player.Object.Position.X += player.Speed
		}
	}

	return nil
}
