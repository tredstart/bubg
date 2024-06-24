package main

import (
	rl "github.com/gen2brain/raylib-go/raylib"
	"github.com/tredstart/bubg/internal/ntt"
)

func main() {
	rl.InitWindow(1080, 720, "Box's Unkbenownst Backgrounds")
	defer rl.CloseWindow()

	SCREEN_WIDTH := rl.GetScreenWidth()
	SCREEN_HEIGHT := rl.GetScreenHeight()

	player := ntt.Player{
		Shape:    rl.NewRectangle(100, 100, 50, 50),
		Rotation: 0,
	}
	player.Weapon = &ntt.Fist{}

	camera := rl.Camera2D{}
	camera.Zoom = 1.0
	camera.Offset = rl.Vector2{X: float32(SCREEN_WIDTH) / 2, Y: float32(SCREEN_HEIGHT) / 2}
	player.Camera = &camera

	tiles := ntt.Tiles{}

	tiles.LoadMap("assets/maps/test")

	rl.SetTargetFPS(60)

	for !rl.WindowShouldClose() {
		dt := rl.GetFrameTime()
		player.Update(dt)
		camera.Target = player.Center()
		rl.BeginDrawing()
		rl.ClearBackground(rl.RayWhite)
		{
			rl.BeginMode2D(camera)
			{
				player.Render()
			tiles.Render()
			}
			rl.EndMode2D()

		}
		rl.EndDrawing()
	}
}
