package ntt

import (
	"math"

	rl "github.com/gen2brain/raylib-go/raylib"
)

type Weapon struct {
	AmmoCapacity uint8
	Ammo         uint8

	Texture        Sprite
	RateOfFire     Timer
	BulletVelocity float32
	ReloadTime     Timer
	Icon           rl.Texture2D
	Description    string

	// attachments
	Magazine    float32
	Scope       float32
	Muzzle      float32
	Underbarrel float32
	Rail        float32
	// Bullet BulletType

	// maybe?
	// Slots Attachment
}

func (g *Weapon) Reload() {
	g.Ammo = g.AmmoCapacity
}

func (g *Weapon) Rotate(deg float32) {
	g.Texture.Rotation = deg
}

func (g *Weapon) Update(float32) {
	g.Texture.Origin = WeaponOffset(g.Texture.Texture)
	g.RateOfFire.Tick()
	g.ReloadTime.Tick()

	if g.Ammo <= 0 {
		g.ReloadTime.Start()
	}
}

func (g *Weapon) SetOrigin(origin rl.Vector2) {
	g.Texture.Pos = origin
}

func (g *Weapon) Render() {
	g.Texture.Render()
}

func (g *Weapon) Attack(world *World) {
	if g.RateOfFire.Finished && g.Ammo > 0 {
		g.Ammo -= 1
		g.RateOfFire.Start()
		offset := WeaponOffset(g.Texture.Texture)
		origin := rl.Vector2{
			X: g.Texture.Pos.X - offset.X,
			Y: g.Texture.Pos.Y - offset.Y + DEFAULT_WEAPON_MARGIN,
		}
		bullet_pos := RotatePoint(origin, g.Texture.Pos, g.Texture.Rotation)
		bullet := NewBullet(bullet_pos, g.Texture.Rotation)

		velocity := rl.Vector2{
			X: g.BulletVelocity * -float32(math.Cos(float64(g.Texture.Rotation)*rl.Deg2rad)),
			Y: g.BulletVelocity * -float32(math.Sin(float64(g.Texture.Rotation)*rl.Deg2rad)),
		}
		bullet.SetVelocity(velocity)
		world.Bullets = append(world.Bullets, bullet)
	}
}
