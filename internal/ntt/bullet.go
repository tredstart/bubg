package ntt

import (
	rl "github.com/gen2brain/raylib-go/raylib"
)

const (
	bullet_texture = "assets/sprites/bullet.png"
)

type Bullet struct {
	Texture  Sprite
	Velocity rl.Vector2
}

func NewBullet(pos rl.Vector2, rotation float32) *Bullet {
	bullet := &Bullet{
		Texture: Sprite{
			Texture: BULLET_TEXTURE,
			TextureRect: rl.Rectangle{
				Width:  float32(BULLET_TEXTURE.Width),
				Height: float32(BULLET_TEXTURE.Height),
			},
			Scale:    1,
			Tint:     rl.RayWhite,
			Rotation: rotation,
			Pos:      pos,
			Origin: rl.Vector2{
				X: float32(BULLET_TEXTURE.Width) / 2,
				Y: float32(BULLET_TEXTURE.Height) / 2,
			},
		},
	}

	return bullet
}

func (b *Bullet) SetVelocity(vel rl.Vector2) {
	b.Velocity = vel
}

func (b *Bullet) Update(dt float32) {
	origin := b.Texture.Pos
	origin.X += b.Velocity.X * dt
	origin.Y += b.Velocity.Y * dt
	b.Texture.Pos = origin
}

func (b *Bullet) Render() {
	b.Texture.Render()
}
