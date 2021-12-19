package main

import (
	"Space_Invaders_Go/scenes"
	"fmt"
	"github.com/veandco/go-sdl2/sdl"
	"github.com/veandco/go-sdl2/ttf"
	"os"
)

var window *sdl.Window
var renderer *sdl.Renderer

func initializeDependencies() {
	// Initialize sdl
	err := sdl.Init(sdl.INIT_EVERYTHING)
	if err != nil {
		fmt.Println("initializing sdl: ", err)
		return
	}

	// Initialize ttf
	err = ttf.Init()
	if err != nil {
		fmt.Println("initializing ttf: ", err)
		return
	}
}

func createWindowAndRenderer() {
	window, err := sdl.CreateWindow(
		"Space Invaders Go",
		sdl.WINDOWPOS_UNDEFINED,
		sdl.WINDOWPOS_UNDEFINED,
		700,
		900,
		sdl.WINDOW_OPENGL)
	if err != nil {
		fmt.Println("creating window: ", err)
		os.Exit(1)
	}

	renderer, err = sdl.CreateRenderer(window, -1, sdl.RENDERER_ACCELERATED)
	if err != nil {
		fmt.Println("creating renderer: ", err)
		os.Exit(1)
	}
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

func main() {
	initializeDependencies()
	createWindowAndRenderer()
	loader := scenes.NewSceneLoader(renderer)

	for {
		pollQuitEvent()

		err := renderer.SetDrawColor(0, 0, 0, 255)
		if err != nil {
			fmt.Println("setting draw color: ", err)
			return
		}

		err = renderer.Clear()
		if err != nil {
			fmt.Println("clearing renderer: ", err)
			return
		}

		err = loader.PollKeyEvents(renderer)
		if err != nil {
			fmt.Println(err)
			return
		}

		err = loader.DrawScene(renderer)
		if err != nil {
			fmt.Println("drawing scene: ", err)
			return
		}

		renderer.Present()
	}
}
