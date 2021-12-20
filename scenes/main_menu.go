package scenes

import (
	"github.com/veandco/go-sdl2/sdl"
	"github.com/veandco/go-sdl2/ttf"
)

type MainMenu struct {
	Title      *sdl.Texture
	PlayButton Text
}

const (
	titleSprite  = "sprites/title.bmp"
	fontBase     = "fonts/Barlow_Condensed/BarlowCondensed-Regular.ttf"
	baseFontSize = 40
)

func NewMainMenu() *MainMenu {
	return &MainMenu{}
}

func (menu *MainMenu) Load(renderer *sdl.Renderer) error {
	img, err := sdl.LoadBMP(titleSprite)
	if err != nil {
		return err
	}

	tex, err := renderer.CreateTextureFromSurface(img)
	if err != nil {
		return err
	}

	menu.Title = tex
	img.Free()

	baseFont, err := ttf.OpenFont(fontBase, baseFontSize)
	if err != nil {
		return err
	}

	playButtonSurface, _ := baseFont.RenderUTF8Solid("Press SPACE to start", sdl.Color{R: 255, G: 255, B: 255})
	playButtonTexture, _ := renderer.CreateTextureFromSurface(playButtonSurface)

	menu.PlayButton.Surface = playButtonSurface
	menu.PlayButton.Texture = playButtonTexture

	baseFont.Close()

	return nil
}

func (menu *MainMenu) Draw(renderer *sdl.Renderer) error {
	screenW, _, _ := renderer.GetOutputSize()
	playButtonCenter := (screenW - menu.PlayButton.Surface.W) / 2.0

	err := renderer.Copy(
		menu.Title,
		nil,
		&sdl.Rect{X: 76, Y: 100, W: 548, H: 353})
	if err != nil {
		return err
	}

	err = renderer.Copy(
		menu.PlayButton.Texture,
		nil,
		&sdl.Rect{
			X: playButtonCenter,
			Y: 600,
			W: menu.PlayButton.Surface.W,
			H: menu.PlayButton.Surface.H})
	if err != nil {
		return err
	}

	return nil
}

func (menu *MainMenu) Update(renderer *sdl.Renderer) error {
	return nil
}

func (menu *MainMenu) Unload() error {
	err := menu.Title.Destroy()
	if err != nil {
		return err
	}

	return nil
}
