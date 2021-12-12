package main

import (
	"Space_Invaders_Go/components"
	"fmt"
	"github.com/veandco/go-sdl2/sdl"
)

var gameObjects []components.Updater

const (
	screenWidth  = 600
	screenHeight = 900
)

func createWindow() *sdl.Window {
	w, err := sdl.CreateWindow(
		"Space Invaders",
		sdl.WINDOWPOS_UNDEFINED,
		sdl.WINDOWPOS_UNDEFINED,
		screenWidth,
		screenHeight,
		sdl.WINDOW_OPENGL)

	if err != nil {
		panic(fmt.Errorf("creating window: %v", err))
		return nil
	}

	return w
}

func createRenderer(window *sdl.Window) *sdl.Renderer {
	r, err := sdl.CreateRenderer(window, -1, sdl.RENDERER_PRESENTVSYNC)
	if err != nil {
		panic(fmt.Errorf("creating renderer: %v", err))
		return nil
	}

	return r
}

func createGameObjects(renderer *sdl.Renderer) {
	player, err := components.NewPlayer(renderer)
	if err != nil {
		panic(fmt.Errorf("creating new player: %v", err))
		return
	}

	enemy, err := components.NewEnemy(
		renderer,
		"one",
		components.Vector{X: 0, Y: 0})
	if err != nil {
		panic(fmt.Errorf("creating new enemy: %v", err))
		return
	}

	gameObjects = append(gameObjects, player)
	gameObjects = append(gameObjects, enemy)
}

func updateGameObjects(renderer *sdl.Renderer) {
	for _, object := range gameObjects {
		if object.CheckActive() {
			err := object.OnUpdate()
			if err != nil {
				fmt.Println("updating object: ", err)
				return
			}

			err = object.OnDraw(renderer)
			if err != nil {
				fmt.Println("drawing object: ", err)
				return
			}
		}
	}
}

func pollQuitEvent() bool {
	for event := sdl.PollEvent(); event != nil; event = sdl.PollEvent() {
		switch event.(type) {
		case *sdl.QuitEvent:
			return false
		}
	}

	return true
}

func main() {
	if err := sdl.Init(sdl.INIT_EVERYTHING); err != nil {
		panic(fmt.Errorf("initializing sdl: %v", err))
		return
	}
	defer sdl.Quit()

	window := createWindow()
	defer window.Destroy()

	renderer := createRenderer(window)
	defer renderer.Destroy()

	createGameObjects(renderer)

	gameLoop := true

	for gameLoop {
		gameLoop = pollQuitEvent()

		renderer.SetDrawColor(0, 0, 0, 255)
		renderer.Clear()

		updateGameObjects(renderer)

		renderer.Present()
	}
}
