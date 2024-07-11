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
}

func (w *World) Update(dt float32) {
	w.Player.Update(dt)
	for _, bullet := range w.Bullets {
		bullet.Update(dt)
	}

	// collision update

	ResolvePlayerTiles(&w.Player, w.CurrentMap)
	bullets := []*Bullet{}
	for _, bullet := range w.Bullets {
		if !BulletCollidesTiles(bullet, w.CurrentMap) {
			bullets = append(bullets, bullet)
		}
	}

    w.Bullets = bullets
}

func (w *World) Render() {
	w.Player.Render()
	w.CurrentMap.Render()
	for _, bullet := range w.Bullets {
		bullet.Render()
	}
}

func (w *World) Unload() {
	rl.UnloadTexture(BULLET_TEXTURE)
}
