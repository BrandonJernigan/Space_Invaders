package components

import "github.com/veandco/go-sdl2/sdl"

type PlayerBullet struct {
	Object *GameObject
	Speed  int
}

const (
	playerBulletSprite = "sprites/player-bullet.bmp"
	playerBulletWidth  = 3
	playerBulletHeight = 9
	playerBulletSpeed  = 6
)

func NewPlayerBullet() *PlayerBullet {
	object := &GameObject{
		Size:   Size{W: playerBulletWidth, H: playerBulletHeight},
		Active: false,
	}

	return &PlayerBullet{Object: object}
}

func (bullet *PlayerBullet) Load(renderer *sdl.Renderer) error {
	img, err := sdl.LoadBMP(playerBulletSprite)
	if err != nil {
		return err
	}

	tex, err := renderer.CreateTextureFromSurface(img)
	if err != nil {
		return err
	}

	bullet.Object.Texture = tex
	img.Free()
	return nil
}

func (bullet *PlayerBullet) SetPosition(position Vector) {
	bullet.Object.Position = position
}

func (bullet *PlayerBullet) Draw(renderer *sdl.Renderer) error {
	err := renderer.Copy(
		bullet.Object.Texture,
		&sdl.Rect{X: 0, Y: 0, W: bullet.Object.Size.W, H: bullet.Object.Size.H},
		&sdl.Rect{
			X: bullet.Object.Position.X,
			Y: bullet.Object.Position.Y,
			W: bullet.Object.Size.W,
			H: bullet.Object.Size.H})

	return err
}

func (bullet *PlayerBullet) Update() error {
	if bullet.Object.Active {
		bullet.Object.Position.Y -= playerBulletSpeed
	}

	if bullet.Object.Position.Y < 0 {
		bullet.Object.Active = false
	}

	return nil
}

func (bullet *PlayerBullet) Unload() error {
	err := bullet.Object.Texture.Destroy()
	if err != nil {
		return err
	}

	return nil
}
