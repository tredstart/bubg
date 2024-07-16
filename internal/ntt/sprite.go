package ntt

import (
	rl "github.com/gen2brain/raylib-go/raylib"
)

type Sprite struct {
	// the actual position of the Sprite in the world
	Pos      rl.Vector2
	Rotation float32
	Scale    float32
	Tint     rl.Color
	Texture  rl.Texture2D
	// pivot point relative to the sprite
	Origin rl.Vector2
	// TextureRect is used to pick a specifict rectangle on the spritesheet
	TextureRect rl.Rectangle
}

// Calculates the center point of the sprite
func (s *Sprite) Center() rl.Vector2 {
	return rl.Vector2{
		X: s.Pos.X + s.TextureRect.Width/2,
		Y: s.Pos.Y + s.TextureRect.Height/2,
	}
}

var css_order = [4][2]int8{
	{-1, -1},
	{1, -1},
	{1, 1},
	{-1, 1},
}

func CssOrder() [4][2]int8 {
    return css_order
}

func (s *Sprite) Hitbox() Polygon {
	poly := Polygon{
		Origin: s.Center(),
	}

	half_width := s.TextureRect.Width / 2
	half_height := s.TextureRect.Height / 2

	poly.Vertices = make([]rl.Vector2, 4)
	for i := range poly.Vertices {
		poly.Vertices[i] = RotatePoint(
			rl.Vector2{
				X: poly.Origin.X + half_width*float32(css_order[i][0]),
				Y: poly.Origin.Y + half_height*float32(css_order[i][1]),
			},
			poly.Origin,
			s.Rotation,
		)
	}

	return poly
}

func (s *Sprite) Render() {
	dest := rl.Rectangle{
		X:      s.Pos.X,
		Y:      s.Pos.Y,
		Width:  s.TextureRect.Width,
		Height: s.TextureRect.Height,
	}
	rl.DrawTexturePro(s.Texture, s.TextureRect, dest, s.Origin, s.Rotation, s.Tint)
}
