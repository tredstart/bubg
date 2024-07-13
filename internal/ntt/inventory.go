package ntt

import rl "github.com/gen2brain/raylib-go/raylib"

type Inventory struct {
	activeHUD bool
	Weapons   [3]*Weapon
}

func (i *Inventory) Display(screen_width, screen_height int) {
	if i.activeHUD {
		rect := rl.Rectangle{
			X:      30,
			Y:      float32(screen_height) / 5,
			Height: float32(screen_height) / 2,
			Width:  float32(screen_width) / 5,
		}

		for _, weapon := range i.Weapons {
			rl.DrawRectangleRec(rect, rl.NewColor(10, 10, 10, 70))
			if weapon != nil {
				weapon.Display(rect)
				rect.X += rect.Width + 10
			}
		}

	}
}
