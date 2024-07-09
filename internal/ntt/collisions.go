package ntt

import (
	"math"

	rl "github.com/gen2brain/raylib-go/raylib"
)

func overlap(a, b rl.Rectangle) (float32, float32) {
	p_x := b.X
	c_x := a.X
	p_y := b.Y
	c_y := a.Y
	p_width := b.Width
	c_width := a.Width
	p_height := b.Height
	c_height := a.Height

	var shift_x, shift_y float32

	if (c_x + c_width/2) < (p_x + p_width/2) {
		shift_x = (c_x + c_width) - p_x
	} else {
		shift_x = c_x - (p_x + p_width)
	}
	if (c_y + c_height/2) < (p_y + p_height/2) {
		shift_y = (c_y + c_height) - p_y
	} else {
		shift_y = c_y - (p_y + p_height)
	}

	return shift_x, shift_y
}

func Resolve(player *Player, tiles Tiles) {
	playerTiles(player, tiles)
}

func playerTiles(player *Player, tiles Tiles) {
	for _, tile := range tiles {
		if Collides(&player.Shape, &tile.Shape) {
			shift_x, shift_y := Overlap(BB(&tile.Shape), BB(&player.Shape))
			if math.Abs(float64(shift_y)) > math.Abs(float64(shift_x)) {
				shift_y = 0
			} else {
				shift_x = 0
			}

			origin := player.Shape.Origin()

			origin.X += shift_x
			origin.Y += shift_y
			player.Shape.Move(origin)
            player.Shape.Rotate(player.rotation)
		}
	}
}
