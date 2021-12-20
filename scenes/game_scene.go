package scenes

import (
	"Space_Invaders_Go/components"
	"fmt"
	"github.com/veandco/go-sdl2/sdl"
	"github.com/veandco/go-sdl2/ttf"
	"time"
)

type GameScene struct {
	Score       int
	ScoreText   Text
	LivesText   Text
	Player      *components.Player
	Enemies     []*components.Enemy
	GameObjects []components.Renderer
	OnGameOver  func(renderer *sdl.Renderer, score int) error
}

const (
	fontBold      = "fonts/Barlow_Condensed/BarlowCondensed-SemiBold.ttf"
	smallFontSize = 24
	coolDown      = time.Millisecond * 2000
)

var speed = 1
var row = 1
var lastRowTime = time.Now().Add(-coolDown)

func NewGameScene(onGameOver func(renderer *sdl.Renderer, score int) error) *GameScene {
	return &GameScene{Score: 0, OnGameOver: onGameOver}
}

func (scene *GameScene) LoadUI(renderer *sdl.Renderer) error {
	baseFont, err := ttf.OpenFont(fontBold, smallFontSize)
	if err != nil {
		return err
	}

	scoreString := fmt.Sprintf("Score: %v", scene.Score)
	scoreTextSurface, _ := baseFont.RenderUTF8Solid(scoreString, sdl.Color{R: 255, G: 255, B: 255})
	scoreTextTexture, _ := renderer.CreateTextureFromSurface(scoreTextSurface)

	scene.ScoreText.Surface = scoreTextSurface
	scene.ScoreText.Texture = scoreTextTexture

	livesString := fmt.Sprintf("Lives: %v", 1)
	livesTextSurface, _ := baseFont.RenderUTF8Solid(livesString, sdl.Color{R: 255, G: 255, B: 255})
	livesTextTexture, _ := renderer.CreateTextureFromSurface(livesTextSurface)

	scene.LivesText.Surface = livesTextSurface
	scene.LivesText.Texture = livesTextTexture

	baseFont.Close()

	return nil
}

func (scene *GameScene) LoadPlayer(renderer *sdl.Renderer) error {
	player := components.NewPlayer(components.Vector{X: 302, Y: 800})
	err := player.Load(renderer)
	if err != nil {
		return err
	}
	scene.Player = player
	scene.GameObjects = append(scene.GameObjects, player)

	return nil
}

func (scene *GameScene) LoadAliens(renderer *sdl.Renderer, alienType string, positionY int32) error {
	for i := 0; i < 10; i++ {
		enemy := components.NewAlien(components.Vector{X: 0, Y: positionY}, alienType, int32(speed))
		positionX := i*70 + 10
		enemy.Object.Position.X = int32(positionX)

		err := enemy.Load(renderer)
		if err != nil {
			return err
		}

		scene.GameObjects = append(scene.GameObjects, enemy)
		scene.Enemies = append(scene.Enemies, enemy)
	}

	return nil
}

func (scene *GameScene) LoadEnemies(renderer *sdl.Renderer) error {
	if time.Now().After(lastRowTime.Add(coolDown)) {
		switch row {
		case 1:
			err := scene.LoadAliens(renderer, "squid", -30)
			if err != nil {
				return err
			}
			row = 2
			break
		case 2:
			err := scene.LoadAliens(renderer, "jelly", -30)
			if err != nil {
				return err
			}
			row = 3
			break
		case 3:
			err := scene.LoadAliens(renderer, "crab", -30)
			if err != nil {
				return err
			}
			row = 1
			break
		}
		lastRowTime = time.Now()
	}

	return nil
}

func (scene *GameScene) ChangeGameState(renderer *sdl.Renderer) error {
	for _, enemy := range scene.Enemies {
		enemy.Speed = 0
	}

	return nil
}

func (scene *GameScene) Load(renderer *sdl.Renderer) error {
	err := scene.LoadUI(renderer)
	if err != nil {
		return err
	}

	err = scene.LoadPlayer(renderer)
	if err != nil {
		return err
	}

	err = scene.LoadEnemies(renderer)
	if err != nil {
		return err
	}

	return nil
}

func (scene *GameScene) Draw(renderer *sdl.Renderer) error {
	for _, object := range scene.GameObjects {
		if object.CheckActive() {
			err := object.Draw(renderer)
			if err != nil {
				return err
			}
		}
	}

	err := renderer.Copy(
		scene.ScoreText.Texture,
		nil,
		&sdl.Rect{
			X: 20,
			Y: 860,
			W: scene.ScoreText.Surface.W,
			H: scene.ScoreText.Surface.H})
	if err != nil {
		return err
	}

	err = renderer.Copy(
		scene.LivesText.Texture,
		nil,
		&sdl.Rect{
			X: 620,
			Y: 860,
			W: scene.LivesText.Surface.W,
			H: scene.LivesText.Surface.H})
	if err != nil {
		return err
	}

	return nil
}

func (scene *GameScene) CheckPositions(positionOne, positionTwo components.Vector, offset components.Size) bool {
	if positionOne.Y <= positionTwo.Y+offset.H &&
		positionOne.Y >= positionTwo.Y {

		if positionOne.X <= positionTwo.X+offset.W &&
			positionOne.X >= positionTwo.X {
			return true
		}
	}
	return false
}

func (scene *GameScene) LoadExplosion(renderer *sdl.Renderer, sprite string,
	size components.Size, position components.Vector) error {

	img, err := sdl.LoadBMP(sprite)
	if err != nil {
		return err
	}

	tex, err := renderer.CreateTextureFromSurface(img)
	if err != nil {
		return err
	}

	err = renderer.Copy(
		tex,
		&sdl.Rect{X: 0, Y: 0, W: size.W, H: size.H},
		&sdl.Rect{X: position.X, Y: position.Y, W: size.W, H: size.H})
	if err != nil {
		return err
	}

	return nil
}

func (scene *GameScene) CheckBulletCollisions(renderer *sdl.Renderer) error {
	for _, bullet := range scene.Player.Bullets {
		bulletPosition := bullet.Object.Position

		for _, enemy := range scene.Enemies {
			enemyPosition := enemy.Object.Position

			if bullet.Object.Active && enemy.Object.Active {
				if scene.CheckPositions(bulletPosition, enemyPosition, enemy.Object.Size) {
					scene.Score += 50
					err := scene.LoadUI(renderer)
					if err != nil {
						return err
					}

					err = scene.LoadExplosion(
						renderer,
						"sprites/alien-explosion.bmp",
						components.Size{W: 67, H: 41},
						enemyPosition)
					if err != nil {
						return err
					}

					bullet.OnCollision()
					enemy.OnCollision()
				}
			}
		}
	}

	return nil
}

func (scene *GameScene) CheckPlayerCollisions(renderer *sdl.Renderer) error {
	player := scene.Player
	playerPosition := player.Object.Position
	playerSize := player.Object.Size

	for _, enemy := range scene.Enemies {
		enemyPosition := enemy.Object.Position
		if scene.CheckPositions(enemyPosition, playerPosition, playerSize) {
			err := scene.LoadExplosion(
				renderer,
				"sprites/alien-explosion.bmp",
				components.Size{W: 67, H: 41},
				enemyPosition)
			if err != nil {
				return err
			}

			err = scene.LoadExplosion(
				renderer,
				"sprites/player-explosion.bmp",
				components.Size{W: 117, H: 63},
				enemyPosition)
			if err != nil {
				return err
			}

			player.OnCollision()
			enemy.OnCollision()

			err = scene.OnGameOver(renderer, scene.Score)
			if err != nil {
				return err
			}
		}
	}

	return nil
}

func (scene *GameScene) Update(renderer *sdl.Renderer) error {
	err := scene.CheckBulletCollisions(renderer)
	if err != nil {
		return err
	}

	err = scene.CheckPlayerCollisions(renderer)
	if err != nil {
		return err
	}

	err = scene.Player.Update()
	if err != nil {
		return err
	}

	for _, object := range scene.GameObjects {
		if object.CheckActive() {
			err = object.Update()
			if err != nil {
				return err
			}
		}
	}

	err = scene.LoadEnemies(renderer)
	if err != nil {
		return err
	}
	return nil
}

func (scene *GameScene) Unload() error {
	scene.ScoreText.Surface.Free()
	err := scene.ScoreText.Texture.Destroy()
	if err != nil {
		return err
	}

	return nil
}
