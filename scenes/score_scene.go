package scenes

import (
	"fmt"
	"github.com/veandco/go-sdl2/sdl"
	"github.com/veandco/go-sdl2/ttf"
)

type ScoreScene struct {
	Score      int
	ScoreText  Text
	PlayButton Text
}

func NewScoreScene(score int) *ScoreScene {
	return &ScoreScene{Score: score}
}

func (scene *ScoreScene) Load(renderer *sdl.Renderer) error {
	baseFont, err := ttf.OpenFont(fontBase, baseFontSize)
	if err != nil {
		return err
	}

	scoreString := fmt.Sprintf("You scored %v points", scene.Score)
	scoreTextSurface, _ := baseFont.RenderUTF8Solid(scoreString, sdl.Color{R: 255, G: 255, B: 255})
	scoreTextTexture, _ := renderer.CreateTextureFromSurface(scoreTextSurface)

	scene.ScoreText.Surface = scoreTextSurface
	scene.ScoreText.Texture = scoreTextTexture

	playButtonSurface, _ := baseFont.RenderUTF8Solid("Press SPACE to play again", sdl.Color{R: 255, G: 255, B: 255})
	playButtonTexture, _ := renderer.CreateTextureFromSurface(playButtonSurface)

	scene.PlayButton.Surface = playButtonSurface
	scene.PlayButton.Texture = playButtonTexture

	baseFont.Close()

	return nil
}

func (scene *ScoreScene) Draw(renderer *sdl.Renderer) error {
	screenW, _, _ := renderer.GetOutputSize()
	scoreCenter := (screenW - scene.ScoreText.Surface.W) / 2.0
	playButtonCenter := (screenW - scene.PlayButton.Surface.W) / 2.0

	err := renderer.Copy(
		scene.ScoreText.Texture,
		nil,
		&sdl.Rect{X: scoreCenter, Y: 300, W: scene.ScoreText.Surface.W, H: scene.ScoreText.Surface.H})
	if err != nil {
		return err
	}

	err = renderer.Copy(
		scene.PlayButton.Texture,
		nil,
		&sdl.Rect{X: playButtonCenter, Y: 500, W: scene.PlayButton.Surface.W, H: scene.PlayButton.Surface.H})
	if err != nil {
		return err
	}

	return nil
}

func (scene *ScoreScene) Update(renderer *sdl.Renderer) error {
	return nil
}

func (scene *ScoreScene) Unload() error {
	scene.ScoreText.Surface.Free()
	err := scene.ScoreText.Texture.Destroy()
	if err != nil {
		return err
	}

	return nil
}
