package main

import (
	"math/rand"

	rl "github.com/gen2brain/raylib-go/raylib"
	"github.com/tredstart/bubg/internal/ntt"
	"github.com/tredstart/bubg/internal/ntt/mods"
)

func main() {
	rl.InitWindow(0, 0, "Box's Unkbenownst Backgrounds")
	defer rl.CloseWindow()

	rl.ToggleFullscreen()

	SCREEN_WIDTH := rl.GetScreenWidth()
	SCREEN_HEIGHT := rl.GetScreenHeight()

	gunt := rl.LoadTexture("assets/sprites/test_gun.png")
	ntt.BULLET_TEXTURE = rl.LoadTexture("assets/sprites/bullet.png")
	defer rl.UnloadTexture(gunt)
	defer rl.UnloadTexture(ntt.BULLET_TEXTURE)
	gunt_source_rect := rl.Rectangle{
		Width:  float32(gunt.Width),
		Height: float32(gunt.Height),
	}

	dm := mods.DoubleMag{
		Texture: ntt.BULLET_TEXTURE,
	}

	world := ntt.World{}
	spawn_data := world.CurrentMap.LoadMap("assets/maps/test")
	world.Player = ntt.Player{
		Shape: ntt.NewRect(
			spawn_data.PlayerPos,
			ntt.PLAYER_WIDTH, ntt.PLAYER_HEIGHT, 0,
			rl.Red,
		),
		Stats: ntt.Stats{
			MaxHealth:     100,
			CurrentHealth: 100,
		},
	}

	for _, point := range spawn_data.SpawnPoints {
		if rand.Intn(2) == 0 {
			world.Weapons = append(world.Weapons, &ntt.Weapon{
				Texture: ntt.Sprite{
					Pos:         point,
					Texture:     gunt,
					Scale:       1,
					Tint:        rl.RayWhite,
					TextureRect: gunt_source_rect,
				},
				Icon:         gunt,
				AmmoCapacity: uint32(rand.Intn(100) + 1),
				// FIXME: this should be redone as something normal/take the number of mods available
				// NOTE: also they should be randomly filled with some of the mods
				Mods:       make([]ntt.Modifier, rand.Intn(5)+1),
				Detectable: true,
			})
		}
	}

	smg := &ntt.Weapon{
		Texture: ntt.Sprite{
			Texture:     gunt,
			Scale:       1,
			Tint:        rl.RayWhite,
			TextureRect: gunt_source_rect,
		},
		RateOfFire:     ntt.NewTimer(0.01),
		BulletVelocity: 700,
		ReloadTime:     ntt.NewTimer(2),
		AmmoCapacity:   100,
		Icon:           gunt,
		Description:    "RATATATA",
		Mods:           make([]ntt.Modifier, 4),
        Recoil: 200,
	}

	smg.ReloadTime.Callback = smg.Reload
    smg.Reload()

	pistol := &ntt.Weapon{
		Texture: ntt.Sprite{
			Texture:     gunt,
			Scale:       1,
			Tint:        rl.RayWhite,
			TextureRect: gunt_source_rect,
		},
		RateOfFire:     ntt.NewTimer(0.3),
		BulletVelocity: 500,
		ReloadTime:     ntt.NewTimer(1),
		AmmoCapacity:   8,
		Icon:           gunt,
		Description:    "It's not small, \nit's just cold out here",
		Mods:           make([]ntt.Modifier, 3),
        Recoil: 100,
	}

	pistol.ReloadTime.Callback = pistol.Reload
	pistol.EquipMod(&dm)
	pistol.Reload()

	rifle := &ntt.Weapon{
		Texture: ntt.Sprite{
			Texture:     gunt,
			Scale:       1,
			Tint:        rl.RayWhite,
			TextureRect: gunt_source_rect,
		},
		RateOfFire:     ntt.NewTimer(1),
		BulletVelocity: 1500,
		ReloadTime:     ntt.NewTimer(1.3),
		Ammo:           1,
		AmmoCapacity:   1,
		Description:    "Faithful railgun. Maybe.",
		Icon:           gunt,
		Mods:           make([]ntt.Modifier, 1),
	}

	rifle.ReloadTime.Callback = rifle.Reload

	world.Player.Inventory = ntt.Inventory{}
	world.Player.Weapons[0] = pistol
	world.Player.Weapons[1] = smg
	// world.Player.Weapons[2] = rifle

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
		rl.DrawFPS(int32(SCREEN_WIDTH)-150, 50)
		{
			rl.BeginMode2D(camera)
			{
				world.Render()
			}
			rl.EndMode2D()
			world.Player.Display()
			world.Player.Inventory.Display(SCREEN_WIDTH, SCREEN_HEIGHT)

		}
		rl.EndDrawing()
	}
}
