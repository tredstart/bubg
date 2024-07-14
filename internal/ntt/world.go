package ntt

import (
	rl "github.com/gen2brain/raylib-go/raylib"
)

var (
	BULLET_TEXTURE rl.Texture2D
)

type World struct {
	Player     Player
	CurrentMap Tiles
	Bullets    []*Bullet
	Weapons    []*Weapon
}

func (w *World) Update(dt float32) {
	if !w.Player.activeHUD {
		w.Player.Update(dt)
		for _, bullet := range w.Bullets {
			bullet.Update(dt)
		}

		// collision update

		ResolvePlayerTiles(&w.Player, w.CurrentMap)
		ResolvePlayerDetectWeapon(&w.Player, w.Weapons)
		bullets := []*Bullet{}
		for _, bullet := range w.Bullets {
			if !BulletCollidesTiles(bullet, w.CurrentMap) {
				bullets = append(bullets, bullet)
			}
		}

		w.Bullets = bullets
	} else {
		w.Player.Inventory.Update(&w.Player)
	}
}

func (w *World) Render() {
	w.Player.Render()
	w.CurrentMap.Render()
	for _, bullet := range w.Bullets {
		bullet.Render()
	}
	for _, weapon := range w.Weapons {
		weapon.Render()
	}
}
