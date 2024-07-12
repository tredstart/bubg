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

	gunt := rl.LoadTexture("assets/sprites/test_gun.png")
	ntt.BULLET_TEXTURE = rl.LoadTexture("assets/sprites/bullet.png")
	defer rl.UnloadTexture(gunt)
	source_rect := rl.Rectangle{
		Width:  float32(gunt.Width),
		Height: float32(gunt.Height),
	}

	world := ntt.World{}
	defer world.Unload()
	player_pos := world.CurrentMap.LoadMap("assets/maps/test")
	world.Player = ntt.Player{
		Shape: ntt.NewRect(
			player_pos,
			ntt.PLAYER_WIDTH, ntt.PLAYER_HEIGHT, 0,
			rl.Red,
		),
		Stats: ntt.Stats{
			MaxHealth:     100,
			CurrentHealth: 100,
		},
	}

	smg := &ntt.Weapon{
		Texture: ntt.Sprite{
			Texture:     gunt,
			Scale:       1,
			Tint:        rl.RayWhite,
			TextureRect: source_rect,
		},
		RateOfFire:     ntt.NewTimer(0.01),
		BulletVelocity: 700,
		ReloadTime:     ntt.NewTimer(2),
		Ammo:           100,
		AmmoCapacity:   100,
	}
    smg.ReloadTime.Callback = smg.Reload
	pistol := &ntt.Weapon{
		Texture: ntt.Sprite{
			Texture:     gunt,
			Scale:       1,
			Tint:        rl.RayWhite,
			TextureRect: source_rect,
		},
		RateOfFire:     ntt.NewTimer(0.3),
		BulletVelocity: 500,
		ReloadTime:     ntt.NewTimer(1),
		Ammo:           8,
		AmmoCapacity:   8,
	}

    pistol.ReloadTime.Callback = pistol.Reload

	world.Player.Inventory = ntt.Inventory{}
	world.Player.Weapons[0] = pistol
	world.Player.Weapons[1] = smg

	camera := rl.Camera2D{}
	camera.Zoom = 1.0
	camera.Offset = rl.Vector2{X: float32(SCREEN_WIDTH) / 2, Y: float32(SCREEN_HEIGHT) / 2}
	world.Player.Camera = &camera

	world.Player.World = &world

	rl.SetTargetFPS(60)

	for !rl.WindowShouldClose() {
		dt := rl.GetFrameTime()
		world.Update(dt)
		camera.Target = world.Player.Shape.Origin()
		rl.BeginDrawing()
		rl.ClearBackground(rl.RayWhite)
		{
			rl.BeginMode2D(camera)
			{
				world.Render()
			}
			rl.EndMode2D()
			world.Player.Display()
			world.Player.Inventory.Display()

		}
		rl.EndDrawing()
	}
}
