package ntt

import rl "github.com/gen2brain/raylib-go/raylib"

type Rect struct {
	rl.Rectangle
}

func (r *Rect) Vertices() []rl.Vector2 {
	return []rl.Vector2{
		{X: r.X, Y: r.Y},
		{X: r.X, Y: r.Y + r.Height},
		{X: r.X + r.Width, Y: r.Y + r.Height},
		{X: r.X + r.Width, Y: r.Y},
	}
}

func (r *Rect) Origin() rl.Vector2 {
	return rl.Vector2{
		X: r.X + r.Width/2,
		Y: r.Y + r.Height/2,
	}
}

// func (r *Rect) Rotate(deg float32) {
//     RotatePoint()
// }
//
// func (r *Rect) Render() {
//     vertices := r.Vertices()
//     rl.DrawLineV()
//     rl.DrawLineV()
//     rl.DrawLineV()
//     rl.DrawLineV()
// }
