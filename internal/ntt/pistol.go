package ntt

import (
	"math"

	rl "github.com/gen2brain/raylib-go/raylib"
)

const (
	pistol_bullet_velocity = 1500.0
)

type Pistol struct {
	Texture Sprite
}

func (g *Pistol) Rotate(deg float32) {
	g.Texture.Rotation = deg
}

func (g *Pistol) Update(float32) {
	g.Texture.Origin = WeaponOffset(g.Texture.Texture)
}

func (g *Pistol) SetOrigin(origin rl.Vector2) {
	g.Texture.Pos = origin
}

func (g *Pistol) Render() {
	g.Texture.Render()
}

func (g *Pistol) Attack(world *World) {
	bullet := NewBullet(g.Texture.Pos, g.Texture.Rotation)

	velocity := rl.Vector2{
		X: pistol_bullet_velocity * -float32(math.Cos(float64(g.Texture.Rotation) * rl.Deg2rad)),
		Y: pistol_bullet_velocity * -float32(math.Sin(float64(g.Texture.Rotation) * rl.Deg2rad)),
	}
	bullet.SetVelocity(velocity)
	world.Bullets = append(world.Bullets, bullet)
}
