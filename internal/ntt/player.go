package ntt

import (
	"math"

	rl "github.com/gen2brain/raylib-go/raylib"
)

const (
	player_speed  = 300
	PLAYER_WIDTH  = 50
	PLAYER_HEIGHT = 50
)

type Stats struct {
	MaxHealth     float32
	CurrentHealth float32
}

type Player struct {
	Shape        Rect
	Camera       *rl.Camera2D
	rotation     float32
	World        *World
	activeWeapon uint8
	Velocity     rl.Vector2
	Direction    rl.Vector2

	Inventory
	Stats

	DetectedWeapon *Weapon
}

func (p *Player) Update(dt float32) {
	for _, weapon := range p.Weapons {
		if weapon != nil {
			p.World.Weapons[weapon.ID] = nil
		}
	}
	// FIXME: active hud should be a game state and should pause the game
	mouse_pos := rl.GetScreenToWorld2D(rl.GetMousePosition(), *p.Camera)
	p.rotation = LookAt(mouse_pos, p.Shape.Origin())

	p.Direction.X = 0
	p.Direction.Y = 0

	origin := p.Shape.Origin()

	if rl.IsKeyDown(rl.KeyW) {
		p.Direction.Y = -1
	}
	if rl.IsKeyDown(rl.KeyA) {
		p.Direction.X = -1
	}
	if rl.IsKeyDown(rl.KeyS) {
		p.Direction.Y = 1
	}
	if rl.IsKeyDown(rl.KeyD) {
		p.Direction.X = 1
	}

	p.Velocity.X = p.Direction.X * player_speed
	p.Velocity.Y = p.Direction.Y * player_speed

	if rl.IsKeyPressed(rl.KeyOne) {
		p.activeWeapon = 0
	}
	if rl.IsKeyPressed(rl.KeyTwo) {
		p.activeWeapon = 1
	}
	if rl.IsKeyPressed(rl.KeyThree) {
		p.activeWeapon = 2
	}
	current_weapon := p.CurrentWeapon()

	if current_weapon != nil {

		if rl.IsMouseButtonDown(rl.MouseLeftButton) {
			if current_weapon.Attack(p.World) {
				recoil := rl.Vector2{
					X: current_weapon.Recoil * float32(math.Cos(float64(p.rotation*rl.Deg2rad))),
					Y: current_weapon.Recoil * float32(math.Sin(float64(p.rotation*rl.Deg2rad))),
				}
				p.Velocity.X += recoil.X
				p.Velocity.Y += recoil.Y
			}

		}
	}

	origin.X += p.Velocity.X * dt
	origin.Y += p.Velocity.Y * dt

	if current_weapon != nil {
		current_weapon.SetOrigin(origin)
		current_weapon.Rotate(p.rotation)
		current_weapon.Update(dt)
	}

	p.Shape.Move(origin)
	p.Shape.Rotate(p.rotation)
	if p.DetectedWeapon != nil && rl.IsKeyPressed(rl.KeyF) {
		p.activeHUD = !p.activeHUD
	}
	if rl.IsKeyPressed(rl.KeyTab) {
		p.activeHUD = !p.activeHUD
	}

}

func (p *Player) EquipWeapon(inventory_id int) {
	tmp_weapon := p.Inventory.Weapons[inventory_id]
	if tmp_weapon != nil {
		p.DropWeapon(inventory_id)
	}
	p.Inventory.Weapons[inventory_id] = p.DetectedWeapon
	p.DetectedWeapon.Detectable = false
	p.DetectedWeapon.Reload()
}

func (p *Player) DropWeapon(inventory_id int) {
	tmp_weapon := p.Inventory.Weapons[inventory_id]
	if tmp_weapon != nil {
		tmp_weapon.Detectable = true
		tmp_weapon.SetOrigin(p.Shape.Center)
		tmp_weapon.Texture.Origin = rl.Vector2{X: 0, Y: 0}
		tmp_weapon.Rotate(0)
		p.World.Weapons[tmp_weapon.ID] = tmp_weapon
	}
	p.Inventory.Weapons[inventory_id] = nil
}

func (p *Player) CurrentWeapon() *Weapon {
	return p.Weapons[p.activeWeapon]
}

func (p *Player) Render() {
	p.Shape.Render()
	mouse_pos := rl.GetScreenToWorld2D(rl.GetMousePosition(), *p.Camera)
	rl.DrawLineV(p.Shape.Origin(), mouse_pos, rl.Red)
	rl.DrawRectangleRec(BB(&p.Shape), rl.NewColor(0, 179, 69, 80))
	current_weapon := p.CurrentWeapon()
	if current_weapon != nil {
		p.CurrentWeapon().Render()
	}
}

func (p *Player) Display() {
	back_health := rl.Rectangle{
		X:      10,
		Y:      10,
		Width:  p.MaxHealth + 4,
		Height: 10,
	}

	health := rl.Rectangle{
		X:      12,
		Y:      12,
		Width:  p.CurrentHealth,
		Height: 8,
	}

	rl.DrawRectangleRec(back_health, rl.LightGray)
	rl.DrawRectangleRec(health, rl.Red)

	current_weapon := p.CurrentWeapon()
	if current_weapon != nil {
		back_ammo := rl.Rectangle{
			X:      10,
			Y:      30,
			Height: 7,
			Width:  float32(current_weapon.AmmoCapacity) + 4,
		}
		ammo := rl.Rectangle{
			X:      12,
			Y:      32,
			Height: 3,
			Width:  float32(current_weapon.Ammo),
		}
		rl.DrawRectangleRec(back_ammo, rl.LightGray)
		rl.DrawRectangleRec(ammo, rl.Yellow)
	}

}
