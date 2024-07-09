package ntt

import rl "github.com/gen2brain/raylib-go/raylib"

type Shape interface {
    Vertices() []rl.Vector2
    Origin() rl.Vector2
}
