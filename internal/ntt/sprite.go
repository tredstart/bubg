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

func (s *Sprite) Center() rl.Vector2 {
	return rl.Vector2{
		X: s.Pos.X + s.TextureRect.Width/2,
		Y: s.Pos.Y + s.TextureRect.Height/2,
	}
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
