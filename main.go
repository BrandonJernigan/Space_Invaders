package main

import (
	"Space_Invaders_Go/components"
	"fmt"
	"github.com/veandco/go-sdl2/sdl"
)

var gameObjects []components.Updater
var shieldObject *components.GameObject
var enemyCount int

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
	createEnemyFleet(renderer)

	player, err := components.NewPlayer(renderer)
	if err != nil {
		panic(fmt.Errorf("creating new player: %v", err))
		return
	}
	gameObjects = append(gameObjects, player)

	shield, err := components.NewShield(renderer)
	if err != nil {
		return
	}
	shieldObject = shield
	gameObjects = append(gameObjects, shield)
}

func createEnemyFleet(renderer *sdl.Renderer) {
	enemyType := [6]string{"one", "two", "three", "four", "five", "six"}
	baseYPosition := 64 * 5

	for i := 0; i < 6; i++ {
		yPosition := baseYPosition - (64 * i)

		for j := 0; j < 9; j++ {
			xPosition := j * 64

			enemy, err := components.NewEnemy(
				renderer,
				enemyType[i],
				components.Vector{X: float64(xPosition), Y: float64(yPosition)})
			if err != nil {
				panic(fmt.Errorf("creating new enemy: %v", err))
				return
			}

			gameObjects = append(gameObjects, enemy)
		}
	}
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

func checkEnemyCollisions() {
	for _, bullet := range components.PlayerBullets {
		bulletPosition := bullet.Object.Position

		for _, enemy := range components.Enemies {
			enemyPosition := enemy.Object.Position

			if bullet.CheckActive() && enemy.CheckActive() {
				// check y position
				if bulletPosition.Y >= (enemyPosition.Y-20) && bulletPosition.Y <= (enemyPosition.Y+20) {
					// check x position
					if bulletPosition.X >= enemyPosition.X && bulletPosition.X <= (enemyPosition.X+32) {
						bullet.OnCollision()
						enemy.OnCollision()
						enemyCount--
						return
					}
				}
			}
		}
	}
}

func checkShieldCollisions() {
	for _, bullet := range components.PlayerBullets {
		bulletPosition := bullet.Object.Position

		if bullet.CheckActive() && shieldObject.CheckActive() {
			// check y position
			if bulletPosition.Y >= (shieldObject.Position.Y-20) && bulletPosition.Y <= (shieldObject.Position.Y+60) {
				// check x position
				if bulletPosition.X >= shieldObject.Position.X && bulletPosition.X <= (shieldObject.Position.X+88) {
					bullet.OnCollision()
					return
				}
			}
		}
	}
}

func checkWinCondition() {
	if enemyCount == 0 {
		fmt.Println("win")
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

	enemyCount = len(components.Enemies)

	gameLoop := true

	for gameLoop {
		gameLoop = pollQuitEvent()

		renderer.SetDrawColor(0, 0, 0, 255)
		renderer.Clear()

		updateGameObjects(renderer)
		checkShieldCollisions()
		checkEnemyCollisions()
		checkWinCondition()

		renderer.Present()
	}
}
