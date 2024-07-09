package ntt

import rl "github.com/gen2brain/raylib-go/raylib"

type Gun struct {
	Texture Sprite
}

func (g *Gun) Rotate(deg float32) {
	g.Texture.Rotation = deg
}

func (g *Gun) Update(float32) {
	g.Texture.Origin = WeaponOffset(g.Texture.Texture)
}

func (g *Gun) SetOrigin(origin rl.Vector2) {
	g.Texture.Pos = origin
}

func (g *Gun) Render() {
	g.Texture.Render()
}

func (g *Gun) Attack() {
	println("pew pew")
}
