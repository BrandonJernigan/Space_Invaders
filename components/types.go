package components

import "github.com/veandco/go-sdl2/sdl"

type Vector struct {
	X int32
	Y int32
}

type Size struct {
	W int32
	H int32
}

type Renderer interface {
	Load(renderer *sdl.Renderer) error
	Draw(renderer *sdl.Renderer) error
	Update() error
	Unload() error
	CheckActive() bool
}

type Collider interface {
	OnCollision()
}

type GameObject struct {
	Texture  *sdl.Texture
	Position Vector
	Size     Size
	Active   bool
}
