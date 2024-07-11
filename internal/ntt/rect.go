package ntt

import (

	rl "github.com/gen2brain/raylib-go/raylib"
)

type Rect struct {
	Center rl.Vector2
	Width  float32
	Height float32
	Color  rl.Color

	tl rl.Vector2
	dl rl.Vector2
	dr rl.Vector2
	tr rl.Vector2

	Filled bool
}

func NewRect(origin rl.Vector2, width, height, rotation float32, color rl.Color) Rect {
	rect := Rect{
		Width:  width,
		Height: height,
		Color:  color,
	}
	rect.Rotate(rotation)
	rect.Move(origin)
	return rect
}

func (r *Rect) Vertices() []rl.Vector2 {
	return []rl.Vector2{
		r.tl,
		r.dl,
		r.dr,
		r.tr,
	}
}

func (r *Rect) Origin() rl.Vector2 {
	return r.Center
}

func (r *Rect) Rotate(deg float32) {
	r.tl = RotatePoint(r.tl, r.Center, deg)
	r.tr = RotatePoint(r.tr, r.Center, deg)
	r.dl = RotatePoint(r.dl, r.Center, deg)
	r.dr = RotatePoint(r.dr, r.Center, deg)
}

func (r *Rect) HalfSize() rl.Vector2 {
	return rl.Vector2{
		X: r.Width / 2,
		Y: r.Height / 2,
	}
}

func (r *Rect) Move(new_origin rl.Vector2) {
    r.Center = new_origin
	half_size := r.HalfSize()

	r.tl.X = r.Center.X - half_size.X
	r.tr.X = r.Center.X + half_size.X
	r.dl.X = r.Center.X - half_size.X
	r.dr.X = r.Center.X + half_size.X

	r.tl.Y = r.Center.Y - half_size.Y
	r.tr.Y = r.Center.Y - half_size.Y
	r.dl.Y = r.Center.Y + half_size.Y
	r.dr.Y = r.Center.Y + half_size.Y
}

func (r *Rect) Render() {
	if !r.Filled {
		rl.DrawLineV(r.tl, r.dl, r.Color)
		rl.DrawLineV(r.dl, r.dr, r.Color)
		rl.DrawLineV(r.dr, r.tr, r.Color)
		rl.DrawLineV(r.tr, r.tl, r.Color)
		rl.DrawPixelV(r.Center, r.Color)
	} else {
		rl.DrawRectangleRec(BB(r), r.Color)
	}
}
