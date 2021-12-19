package scenes

import (
	"fmt"
	"github.com/veandco/go-sdl2/sdl"
	"os"
)

type SceneLoader struct {
	ActiveScene Scene
	GameRunning bool
}

func NewSceneLoader(renderer *sdl.Renderer) *SceneLoader {
	loader := &SceneLoader{}
	loader.GameRunning = false
	err := loader.LoadMainMenu(renderer)
	if err != nil {
		fmt.Println("creating scene loader: ", err)
		os.Exit(1)
	}
	return loader
}

func (loader *SceneLoader) LoadMainMenu(renderer *sdl.Renderer) error {
	menu := NewMainMenu()
	err := menu.Load(renderer)
	if err != nil {
		return err
	}

	loader.ActiveScene = menu
	return nil
}

func (loader *SceneLoader) LoadGameScene(renderer *sdl.Renderer) error {
	game := NewGameScene()
	loader.GameRunning = true
	err := game.Load(renderer)
	if err != nil {
		return err
	}

	loader.ActiveScene = game
	return nil
}

func (loader *SceneLoader) PollKeyEvents(renderer *sdl.Renderer) error {
	if !loader.GameRunning {
		keys := sdl.GetKeyboardState()
		if keys[sdl.SCANCODE_SPACE] == 1 {
			err := loader.ActiveScene.Unload()
			if err != nil {
				return err
			}

			err = loader.LoadGameScene(renderer)
			if err != nil {
				return err
			}
		}
	}
	return nil
}

func (loader *SceneLoader) DrawScene(renderer *sdl.Renderer) error {
	err := loader.ActiveScene.Update()
	if err != nil {
		return err
	}

	err = loader.ActiveScene.Draw(renderer)
	if err != nil {
		return err
	}
	return nil
}
