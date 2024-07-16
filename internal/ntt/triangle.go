package ntt

import rl "github.com/gen2brain/raylib-go/raylib"

type Triangle struct {
	Center rl.Vector2
	Color  rl.Color
	Side   float32

	t rl.Vector2
	l rl.Vector2
	r rl.Vector2
}

func (t *Triangle) Vertices() []rl.Vector2 {
	return []rl.Vector2{
		t.t,
		t.l,
		t.r,
	}
}

func (t *Triangle) Origin() rl.Vector2 {
	return t.Center
}

func (t *Triangle) Render() {
	rl.DrawLineV(t.t, t.l, t.Color)
	rl.DrawLineV(t.l, t.r, t.Color)
	rl.DrawLineV(t.r, t.t, t.Color)
}

func (t *Triangle) Rotate(deg float32) {
    t.t = RotatePoint(t.t, t.Center, deg)
    t.l = RotatePoint(t.l, t.Center, deg)
    t.r = RotatePoint(t.r, t.Center, deg)
}

func (t *Triangle) Move(origin rl.Vector2) {
}

func NewTriangle(origin rl.Vector2, side, rotation float32, color rl.Color) Triangle {
	tri := Triangle{
		Side:  side,
		Color: color,
	}
	tri.Rotate(rotation)
	tri.Move(origin)
	return tri
}
