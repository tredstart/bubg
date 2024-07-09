package ntt

import (
	rl "github.com/gen2brain/raylib-go/raylib"
)

const player_speed = 300

type Player struct {
	Shape    Rect
	Rotation float32
	Camera   *rl.Camera2D
}

var DC = rl.NewColor(40, 140, 50, 100)

func (p *Player) Origin() rl.Vector2 {
	return rl.Vector2{
		X: p.Shape.X + p.Shape.Width/2,
		Y: p.Shape.Y + p.Shape.Height/2,
	}
}

func (p *Player) Update(dt float32) {
	mouse_pos := rl.GetScreenToWorld2D(rl.GetMousePosition(), *p.Camera)

	p.Rotation = LookAt(mouse_pos, p.Origin())

	if rl.IsKeyDown(rl.KeyW) {
		p.Shape.Y -= player_speed * dt
	}
	if rl.IsKeyDown(rl.KeyA) {
		p.Shape.X -= player_speed * dt
	}
	if rl.IsKeyDown(rl.KeyS) {
		p.Shape.Y += player_speed * dt
	}
	if rl.IsKeyDown(rl.KeyD) {
		p.Shape.X += player_speed * dt
	}
}

func (p *Player) Render() {
	dest := rl.Rectangle{
		X:      p.Shape.Width/2 + p.Shape.X,
		Y:      p.Shape.Height/2 + p.Shape.Y,
		Width:  p.Shape.Width,
		Height: p.Shape.Height,
	}
	origin := rl.Vector2{X: p.Shape.Width / 2, Y: p.Shape.Height / 2}
	rl.DrawRectanglePro(dest, origin, p.Rotation, rl.Red)
	mouse_pos := rl.GetScreenToWorld2D(rl.GetMousePosition(), *p.Camera)
	rl.DrawLineV(p.Origin(), mouse_pos, rl.Red)

	rl.DrawRectangleRec(p.Shape.Rectangle, DC)
}
