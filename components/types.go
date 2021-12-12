package components

import "github.com/veandco/go-sdl2/sdl"

type Vector struct {
	X float64
	Y float64
}

type Size struct {
	W int32
	H int32
}

type Updater interface {
	CheckActive() bool
	OnUpdate() error
	OnDraw(renderer *sdl.Renderer) error
}
type Collider interface {
	OnCollision() error
}

type GameObject struct {
	texture  *sdl.Texture
	Position Vector
	Size     Size
	Active   bool
}
