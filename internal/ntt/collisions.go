package ntt

import (
	"math"

	rl "github.com/gen2brain/raylib-go/raylib"
)

func ResolvePlayerTiles(player *Player, tiles Tiles) {
	for _, tile := range tiles {
		if Collides(player.Shape, tile.Shape) {
			shift_x, shift_y := Overlap(BB(tile.Shape), BB(player.Shape))
			if math.Abs(float64(shift_y)) > math.Abs(float64(shift_x)) {
				shift_y = 0
			} else {
				shift_x = 0
			}

			origin := player.Shape.Origin

			origin.X += shift_x
			origin.Y += shift_y
			player.Shape.Move(origin)
			player.Shape.Rotation = player.rotation
		}
	}
}

func BulletCollidesTiles(bullet *Bullet, tiles Tiles) bool {
	for _, tile := range tiles {
		if Collides(bullet.Texture.Hitbox(), tile.Shape) {
			return true
		}
	}

	return false
}

func ResolvePlayerDetectWeapon(player *Player, weapons []*Weapon) {
	for _, weapon := range weapons {
		if weapon != nil {
			if rl.CheckCollisionCircleRec(weapon.Texture.Center(), weapon.Texture.TextureRect.Width, BB(player.Shape)) {
				player.DetectedWeapon = weapon
				break
			} else {
				player.DetectedWeapon = nil
			}
		}
	}
}
