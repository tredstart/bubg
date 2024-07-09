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

	tiles := ntt.Tiles{}

	player_pos := tiles.LoadMap("assets/maps/test")
	player := ntt.Player{
		Shape: ntt.NewRect(
			player_pos,
			ntt.PLAYER_WIDTH, ntt.PLAYER_HEIGHT, 0,
			rl.Red,
		),
	}

	gunt := rl.LoadTexture("assets/sprites/test_gun.png")
	defer rl.UnloadTexture(gunt)
	source_rect := rl.Rectangle{
		Width:  float32(gunt.Width),
		Height: float32(gunt.Height),
	}

	player.Weapon = &ntt.Gun{
		Texture: ntt.Sprite{
			Texture:     gunt,
			Scale:       1,
			Tint:        rl.RayWhite,
			TextureRect: source_rect,
		},
	}

	camera := rl.Camera2D{}
	camera.Zoom = 1.0
	camera.Offset = rl.Vector2{X: float32(SCREEN_WIDTH) / 2, Y: float32(SCREEN_HEIGHT) / 2}
	player.Camera = &camera

	rl.SetTargetFPS(60)

	for !rl.WindowShouldClose() {
		dt := rl.GetFrameTime()
		player.Update(dt)
		ntt.Resolve(&player, tiles)
		camera.Target = player.Shape.Origin()
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
