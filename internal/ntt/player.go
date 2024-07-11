package ntt

import (
	rl "github.com/gen2brain/raylib-go/raylib"
)

const (
	player_speed  = 300
	PLAYER_WIDTH  = 50
	PLAYER_HEIGHT = 50
)

type Player struct {
	Shape        Rect
	Camera       *rl.Camera2D
	rotation     float32
	World        *World
	activeWeapon uint8

	Inventory
}

func (p *Player) Update(dt float32) {
	mouse_pos := rl.GetScreenToWorld2D(rl.GetMousePosition(), *p.Camera)
	p.rotation = LookAt(mouse_pos, p.Shape.Origin())

	origin := p.Shape.Origin()

	if rl.IsKeyDown(rl.KeyW) {
		origin.Y -= player_speed * dt
	}
	if rl.IsKeyDown(rl.KeyA) {
		origin.X -= player_speed * dt
	}
	if rl.IsKeyDown(rl.KeyS) {
		origin.Y += player_speed * dt
	}
	if rl.IsKeyDown(rl.KeyD) {
		origin.X += player_speed * dt
	}
	if rl.IsKeyPressed(rl.KeyOne) {
		p.activeWeapon = 0
	}
	if rl.IsKeyPressed(rl.KeyTwo) {
		p.activeWeapon = 1
	}
	current_weapon := p.CurrentWeapon()

	p.Shape.Move(origin)
	p.Shape.Rotate(p.rotation)
	current_weapon.SetOrigin(origin)
	current_weapon.Rotate(p.rotation)
	current_weapon.Update(dt)

	if rl.IsMouseButtonDown(rl.MouseLeftButton) {
		current_weapon.Attack(p.World)
	}

}

func (p *Player) CurrentWeapon() Weapon {

	return p.Weapons[p.activeWeapon]
}

func (p *Player) Render() {
	p.Shape.Render()
	mouse_pos := rl.GetScreenToWorld2D(rl.GetMousePosition(), *p.Camera)
	rl.DrawLineV(p.Shape.Origin(), mouse_pos, rl.Red)
	rl.DrawRectangleRec(BB(&p.Shape), rl.NewColor(0, 179, 69, 80))
	p.CurrentWeapon().Render()
}
