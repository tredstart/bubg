package ntt

import rl "github.com/gen2brain/raylib-go/raylib"

type Inventory struct {
	activeHUD bool
	Weapons   [3]*Weapon
	Hovered   int
}

func (i *Inventory) Update(player *Player) {
	if i.Hovered != -1 {
		if rl.IsMouseButtonPressed(rl.MouseButtonLeft) && player.DetectedWeapon != nil {
			// make a proper equip here
			player.EquipWeapon(i.Hovered)
            player.activeHUD = false
		}
		if rl.IsMouseButtonPressed(rl.MouseButtonRight) {
			// drop a weapon
			player.DropWeapon(i.Hovered)
            player.activeHUD = false
		}
	}
	if rl.IsKeyPressed(rl.KeyTab) || rl.IsKeyPressed(rl.KeyF) {
		player.activeHUD = !player.activeHUD
	}
}

func (i *Inventory) Display(screen_width, screen_height int) {
	if i.activeHUD {
		rect := rl.Rectangle{
			X:      30,
			Y:      float32(screen_height) / 5,
			Height: float32(screen_height) / 2,
			Width:  float32(screen_width) / 5,
		}

		i.Hovered = -1
		mouse_pos := rl.GetMousePosition()

		for index, weapon := range i.Weapons {
			if rl.CheckCollisionPointRec(mouse_pos, rect) {
				i.Hovered = index
			}
			alpha := 70
			if i.Hovered == index {
				alpha = 90
			}
			rl.DrawRectangleRec(rect, rl.NewColor(10, 10, 10, uint8(alpha)))
			if weapon != nil {
				weapon.Display(rect)
			}
				rect.X += rect.Width + 10
		}

	}
}
