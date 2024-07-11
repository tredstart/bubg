package ntt

import (
	"log"

	rl "github.com/gen2brain/raylib-go/raylib"
)

type Sprite struct {
	Pos         rl.Vector2
	Rotation    float32
	Scale       float32
	Tint        rl.Color
	Texture     rl.Texture2D
	Origin      rl.Vector2
	TextureRect rl.Rectangle
}

func (s *Sprite) Render() {
	dest := rl.Rectangle{
		X:      s.Pos.X,
		Y:      s.Pos.Y,
		Width:  s.TextureRect.Width,
		Height: s.TextureRect.Height,
	}
    log.Println(s.Texture.ID, " : ", dest)
	rl.DrawTexturePro(s.Texture, s.TextureRect, dest, s.Origin, s.Rotation, s.Tint)
}
