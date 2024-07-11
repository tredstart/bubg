package ntt

import (
	rl "github.com/gen2brain/raylib-go/raylib"
)

const (
	bullet_texture = "assets/sprites/bullet.png"
)

type Bullet struct {
	Texture  Sprite
	Hitbox   Rect
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

		Hitbox: NewRect(pos, float32(BULLET_TEXTURE.Width), float32(BULLET_TEXTURE.Height),
			rotation,
			rl.NewColor(255, 0, 79, 69)),
	}

	return bullet
}

func (b *Bullet) SetVelocity(vel rl.Vector2) {
	b.Velocity = vel
}

func (b *Bullet) Update(dt float32) {

	origin := b.Hitbox.Origin()
	origin.X += b.Velocity.X * dt
	origin.Y += b.Velocity.Y * dt
	b.Hitbox.Move(origin)
	b.Hitbox.Rotate(b.Texture.Rotation)
	b.Texture.Pos = origin
}

func (b *Bullet) Render() {
	b.Texture.Render()
	b.Hitbox.Render()
}
