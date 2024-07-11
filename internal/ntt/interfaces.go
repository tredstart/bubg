package ntt

import rl "github.com/gen2brain/raylib-go/raylib"

type Shape interface {
	Vertices() []rl.Vector2
	Origin() rl.Vector2
}

type Node interface {
	SetOrigin(rl.Vector2)
	Rotate(float32)
	Update(float32)
}

type Weapon interface {
	Node
	Renderable
	Attack(*World)
}

type Renderable interface {
	Render()
}
