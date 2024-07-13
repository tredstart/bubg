package mods

import (
	rl "github.com/gen2brain/raylib-go/raylib"
	"github.com/tredstart/bubg/internal/ntt"
)

const (
	max_capacity = 2
	reload_time  = 1.5
)

type DoubleMag struct {
	Texture rl.Texture2D
}

func (d *DoubleMag) Icon() rl.Texture2D {
	return d.Texture
}

func (d *DoubleMag) Description() string {
	return "Doubles magazine capacity, but increases reload time"
}

func (d *DoubleMag) Type() ntt.ModType {
	return ntt.Magazine
}

func (d *DoubleMag) Mod(w *ntt.Weapon) {
	w.AmmoCapacity *= max_capacity
	w.ReloadTime.End *= reload_time
}
