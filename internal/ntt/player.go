package ntt

import (
	"math"

	rl "github.com/gen2brain/raylib-go/raylib"
)

const base_speed = 100

type Player struct {
	Shape        rl.Rectangle
	Rotation     float32
	Weapon       Weapon
	WeaponOffset rl.Vector2
	Camera       *rl.Camera2D
}

var DC = rl.NewColor(40, 140, 50, 100)

func (p *Player) Center() rl.Vector2 {
	return rl.Vector2{
		X: p.Shape.X + p.Shape.Width/2,
		Y: p.Shape.Y + p.Shape.Height/2,
	}
}

func getSpeed() float32 {
	if rl.IsKeyDown(rl.KeyLeftShift) {
		return base_speed + 50
	} else {
		return base_speed
	}
}

func (p *Player) Update(dt float32) {
	mouse_pos := rl.GetScreenToWorld2D(rl.GetMousePosition(), *p.Camera)

	delta := rl.Vector2Subtract(p.Center(), mouse_pos)

	p.Rotation = float32(math.Atan2(float64(delta.Y), float64(delta.X))) * rl.Rad2deg


	if rl.IsKeyPressed(rl.MouseLeftButton) {
		p.Weapon.Attack()
	}

	speed := getSpeed()

	if rl.IsKeyDown(rl.KeyW) {
		p.Shape.Y -= speed * dt
	}
	if rl.IsKeyDown(rl.KeyA) {
		p.Shape.X -= speed * dt
	}
	if rl.IsKeyDown(rl.KeyS) {
		p.Shape.Y += speed * dt
	}
	if rl.IsKeyDown(rl.KeyD) {
		p.Shape.X += speed * dt
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
	rl.DrawLineV(p.Center(), mouse_pos, rl.Red)

	rl.DrawRectangleRec(p.Shape, DC)
}
