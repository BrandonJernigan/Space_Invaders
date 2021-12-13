package main

import (
	"Space_Invaders_Go/components"
	"Space_Invaders_Go/utilities"
	"fmt"
	"github.com/veandco/go-sdl2/sdl"
	"os"
)

var gameObjects []components.Updater
var shieldObject *components.GameObject
var playerObject *components.Player
var winState bool
var enemyCount int
var enemyValue int
var score int

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
	playerObject = player

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

func checkEnemyCollisions(renderer *sdl.Renderer) {
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

						tex, _ := utilities.LoadTexture(renderer, "sprites/enemy-explosion.bmp")
						renderer.Copy(
							tex,
							&sdl.Rect{X: 0, Y: 0, W: 64, H: 64},
							&sdl.Rect{X: int32(enemyPosition.X), Y: int32(enemyPosition.Y), W: 64, H: 64})

						enemyCount--
						score += enemyValue
						return
					}
				}
			}
		}
	}
}

func checkShieldCollisions(renderer *sdl.Renderer) {
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

	for _, enemy := range components.Enemies {
		enemyPosition := enemy.Object.Position
		shieldPosition := shieldObject.Position

		if enemy.CheckActive() && shieldObject.CheckActive() {
			// check y position
			if enemyPosition.Y >= (shieldPosition.Y-20) && enemyPosition.Y <= (shieldPosition.Y+60) {
				// check x position
				if enemyPosition.X >= shieldPosition.X && enemyPosition.X <= (shieldPosition.X+32) {
					enemy.OnCollision()
					shieldObject.OnCollision()

					tex, _ := utilities.LoadTexture(renderer, "sprites/enemy-explosion.bmp")
					renderer.Copy(
						tex,
						&sdl.Rect{X: 0, Y: 0, W: 64, H: 64},
						&sdl.Rect{X: int32(enemyPosition.X), Y: int32(enemyPosition.Y), W: 64, H: 64})

					enemyCount--
					return
				}
			}
		}
	}
}

func checkPlayerCollisions(renderer *sdl.Renderer) {
	for _, enemy := range components.Enemies {
		enemyPosition := enemy.Object.Position
		playerPosition := playerObject.Object.Position

		if enemy.CheckActive() && playerObject.CheckActive() {
			// check y position
			if enemyPosition.Y >= playerPosition.Y {
				// check x position
				if enemyPosition.X >= playerPosition.X && enemyPosition.X <= playerPosition.X+32 {
					tex, _ := utilities.LoadTexture(renderer, "sprites/player-explosion.bmp")
					playerObject.OnCollision()
					renderer.Copy(
						tex,
						&sdl.Rect{X: 0, Y: 0, W: 64, H: 64},
						&sdl.Rect{X: int32(enemyPosition.X), Y: int32(enemyPosition.Y), W: 64, H: 64})
				}
			}
		}
	}
}

func checkEnemyPositions() bool {
	for _, enemy := range components.Enemies {
		enemyPosition := enemy.Object.Position
		if enemyPosition.Y-64 >= screenHeight {
			winState = false
			return true
		}
	}
	return false
}

func checkGameEndCondition() bool {
	if enemyCount == 0 {
		winState = true
		return true
	} else if !playerObject.CheckActive() {
		winState = false
		return true
	}
	return false
}

func pollStartEvent() bool {
	keys := sdl.GetKeyboardState()
	if keys[sdl.SCANCODE_SPACE] == 1 {
		return true
	}

	return false
}

func pollQuitEvent() {
	for event := sdl.PollEvent(); event != nil; event = sdl.PollEvent() {
		switch event.(type) {
		case *sdl.QuitEvent:
			sdl.Quit()
			os.Exit(0)
			return
		}
	}
}

func mainMenuLoop(renderer *sdl.Renderer) {
	menuLoop := true
	for menuLoop {
		pollQuitEvent()
		menuLoop = !pollStartEvent()

		renderer.SetDrawColor(0, 0, 0, 255)
		renderer.Clear()

		renderer.Present()
	}
}

func mainGameLoop(renderer *sdl.Renderer) {
	gameLoop := true
	for gameLoop {
		pollQuitEvent()

		renderer.SetDrawColor(0, 0, 0, 255)
		renderer.Clear()

		updateGameObjects(renderer)
		checkShieldCollisions(renderer)
		checkEnemyCollisions(renderer)
		checkPlayerCollisions(renderer)

		gameLoop = !checkGameEndCondition()
		gameLoop = !checkEnemyPositions()

		renderer.Present()
	}
}

func endGameLoop(renderer *sdl.Renderer) {
	endLoop := true
	for endLoop {
		pollQuitEvent()

		if winState {
			renderer.SetDrawColor(0, 0, 0, 255)
		} else {
			renderer.SetDrawColor(0, 255, 0, 255)
		}
		renderer.Clear()

		renderer.Present()
	}
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
	enemyValue = 50

	mainMenuLoop(renderer)
	mainGameLoop(renderer)
	endGameLoop(renderer)
}
