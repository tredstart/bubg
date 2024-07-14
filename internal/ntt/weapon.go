package ntt

import (
	"log"
	"math"

	rl "github.com/gen2brain/raylib-go/raylib"
)

type Weapon struct {
	AmmoCapacity uint32
	Ammo         uint32

	// Texture width is used for detection
	Texture        Sprite
	RateOfFire     Timer
	BulletVelocity float32
	ReloadTime     Timer
	Icon           rl.Texture2D
	Description    string

	// attachments

	Mods []Modifier

	BaseDamage float32
	Recoil     float32

	// NOTE: maybe temporary?
	Detectable bool
}

func (g *Weapon) EquipMod(mod Modifier) {
	g.Mods[int(mod.Type())%len(g.Mods)] = mod
	mod.Mod(g)
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
	if g.Detectable {
		rl.DrawCircleV(g.Texture.Center(), g.Texture.TextureRect.Width, rl.NewColor(0, 30, 150, 50))
	}
}

func (g *Weapon) Attack(world *World) bool {
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
		return true
	}
	return false
}

func (g *Weapon) Display(rect rl.Rectangle) {
	rl.DrawTexturePro(
		g.Icon,
		rl.Rectangle{
			X:      0,
			Y:      0,
			Width:  float32(g.Icon.Width),
			Height: float32(g.Icon.Height),
		},
		rl.Rectangle{
			X:      rect.X + rect.Width/2 - 40,
			Y:      rect.Y + rect.Height/2 - 35,
			Width:  80,
			Height: 50,
		},
		rl.Vector2{
			X: 0,
			Y: 0,
		},
		0,
		rl.RayWhite,
	)
	rl.DrawText(g.Description, int32(rect.X)+g.Icon.Width, int32(rect.Y)+g.Icon.Height/2, 16, rl.RayWhite)
	for i, mod := range g.Mods {
		if mod != nil {
			log.Println("mod : ", i)
			icon := g.Icon
			rl.DrawTexture(icon,
				int32(rect.X+10),
				int32(rect.Y+rect.Height-float32(icon.Height)-10),
				rl.RayWhite,
			)
		} else {
			log.Println("no mod", i)
			rl.DrawText("no mod", int32(rect.X)+80*int32(i), rect.ToInt32().Y+rect.ToInt32().Height-40, 15, rl.Red)
		}
	}
}
