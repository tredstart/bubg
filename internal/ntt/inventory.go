package ntt

import rl "github.com/gen2brain/raylib-go/raylib"

type Inventory struct {
	activeHUD bool
	Weapons   [3]*Weapon
}

func (i *Inventory) Display() {
	if i.activeHUD {
		rect1 := rl.Rectangle{
			X:      30,
			Y:      70,
			Height: 150,
			Width:  500,
		}
		rect2 := rl.Rectangle{
			X:      30,
			Y:      230,
			Height: 150,
			Width:  500,
		}
		rect3 := rl.Rectangle{
			X:      30,
			Y:      390,
			Height: 150,
			Width:  500,
		}

		rl.DrawRectangleRec(rect1, rl.NewColor(10, 10, 10, 70))
		rl.DrawRectangleRec(rect2, rl.NewColor(10, 10, 10, 70))
		rl.DrawRectangleRec(rect3, rl.NewColor(10, 10, 10, 70))
	}
}
