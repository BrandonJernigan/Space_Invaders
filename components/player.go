package components

import (
	"github.com/veandco/go-sdl2/sdl"
	"time"
)

type Player struct {
	Object       *GameObject
	Speed        float64
	ShotCoolDown time.Duration
}

const (
	playerSprite       = "sprites/player.bmp"
	playerWidth        = 96
	playerHeight       = 60
	playerSpeed        = 1
	playerShotCoolDown = time.Millisecond * 250
)

var playerBullets []*PlayerBullet
var lastShotTime time.Time

func NewPlayer(position Vector) *Player {
	object := &GameObject{
		Size:     Size{W: playerWidth, H: playerHeight},
		Position: position,
	}

	player := &Player{
		Object: object,
		Speed:  playerSpeed}

	return player
}

func (player *Player) LoadPlayerBullets(renderer *sdl.Renderer) error {
	for i := 0; i < 30; i++ {
		bullet := NewPlayerBullet()
		err := bullet.Load(renderer)
		if err != nil {
			return err
		}
		playerBullets = append(playerBullets, bullet)
		lastShotTime = time.Now()
	}
	return nil
}

func (player *Player) Load(renderer *sdl.Renderer) error {
	img, err := sdl.LoadBMP(playerSprite)
	if err != nil {
		return err
	}

	tex, err := renderer.CreateTextureFromSurface(img)
	if err != nil {
		return err
	}

	player.Object.Texture = tex
	img.Free()

	err = player.LoadPlayerBullets(renderer)
	if err != nil {
		return err
	}
	return nil
}

func (player *Player) Draw(renderer *sdl.Renderer) error {
	for _, bullet := range playerBullets {
		if bullet.Object.Active {
			err := bullet.Draw(renderer)
			if err != nil {
				return err
			}
		}
	}
	err := renderer.Copy(
		player.Object.Texture,
		&sdl.Rect{X: 0, Y: 0, W: player.Object.Size.W, H: player.Object.Size.H},
		&sdl.Rect{
			X: player.Object.Position.X,
			Y: player.Object.Position.Y,
			W: player.Object.Size.W,
			H: player.Object.Size.H})

	return err
}

func (player *Player) Update() error {
	for _, bullet := range playerBullets {
		err := bullet.Update()
		if err != nil {
			return err
		}
	}

	keys := sdl.GetKeyboardState()
	if keys[sdl.SCANCODE_A] == 1 && (player.Object.Position.X-playerSpeed) >= 0 {
		player.Object.Position.X -= playerSpeed
	} else if keys[sdl.SCANCODE_D] == 1 && (player.Object.Position.X+playerSpeed <= 604) {
		player.Object.Position.X += playerSpeed
	}

	if keys[sdl.SCANCODE_SPACE] == 1 {
		if time.Now().After(lastShotTime.Add(playerShotCoolDown)) {
			player.Shoot()
		}
	}

	return nil
}

func (player *Player) Unload() error {
	err := player.Object.Texture.Destroy()
	if err != nil {
		return err
	}
	return nil
}

func (player *Player) Shoot() {
	for _, bullet := range playerBullets {
		if !bullet.Object.Active {
			positionX := player.Object.Position.X + (player.Object.Size.W / 2.0)
			positionY := player.Object.Position.Y

			bullet.SetPosition(Vector{X: positionX, Y: positionY})
			bullet.Object.Active = true
			lastShotTime = time.Now()

			return
		}
	}
}
