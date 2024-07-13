package ntt

import rl "github.com/gen2brain/raylib-go/raylib"

type Shape interface {
	Vertices() []rl.Vector2
	Origin() rl.Vector2
}

type Renderable interface {
	Render()
}

type ModType uint

// TODO: I need to capitalize on this idea and make 
// special way to map enum to a slot in the weapon
const (
	Magazine ModType = iota
	Muzzle
	Scope
	Underbarrel
	Rail
)

type Modifier interface {
	Mod(*Weapon)
	Icon() rl.Texture2D
    Description() string
    Type() ModType
}
