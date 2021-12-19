package scenes

import (
	"Space_Invaders_Go/components"
	"fmt"
	"github.com/veandco/go-sdl2/sdl"
	"github.com/veandco/go-sdl2/ttf"
)

type GameScene struct {
	Score       int
	ScoreText   Text
	LivesText   Text
	GameObjects []components.Renderer
}

const (
	fontBold      = "fonts/Barlow_Condensed/BarlowCondensed-SemiBold.ttf"
	smallFontSize = 24
)

func NewGameScene() *GameScene {
	return &GameScene{Score: 0}
}

func (scene *GameScene) Load(renderer *sdl.Renderer) error {
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

	player := components.NewPlayer(components.Vector{X: 302, Y: 800})
	err = player.Load(renderer)
	if err != nil {
		return err
	}
	scene.GameObjects = append(scene.GameObjects, player)

	return nil
}

func (scene *GameScene) Draw(renderer *sdl.Renderer) error {
	for _, object := range scene.GameObjects {
		err := object.Draw(renderer)
		if err != nil {
			return err
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

func (scene *GameScene) Update() error {
	for _, object := range scene.GameObjects {
		err := object.Update()
		if err != nil {
			return err
		}
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
