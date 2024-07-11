package tests

import (
	"testing"

	rl "github.com/gen2brain/raylib-go/raylib"
	"github.com/tredstart/bubg/internal/ntt"
)

func TestCollideRectangles(t *testing.T) {
	rect1 := ntt.Rect{
		Origin: rl.Vector2{X: 4, Y: 4},
		Width:  4,
		Height: 4,
	}
	rect2 := ntt.Rect{
		Origin: rl.Vector2{X: 5, Y: 5},
		Width:  4,
		Height: 4,
	}

	ok := ntt.Collides(&rect1, &rect2)
	if !ok {
		t.Errorf("Rects do not collide but supposed to: \n%v\n%v", rect1, rect2)
	}
}

func TestRectsDontCollide(t *testing.T) {
	rect1 := ntt.Rect{
		Origin: rl.Vector2{X: 0, Y: 0},
		Width:  4,
		Height: 4,
	}
	rect2 := ntt.Rect{
		Origin: rl.Vector2{X: 5, Y: 5},
		Width:  4,
		Height: 4,
	}

	ok := ntt.Collides(&rect1, &rect2)
	if ok {
		t.Errorf("Rects collide but not supposed to: \n%v\n%v", rect1, rect2)
	}
}

func TestRectRotatedCollide(t *testing.T) {

	rect1 := ntt.Rect{
		Origin:   rl.Vector2{X: 4, Y: 4},
		Width:    4,
		Height:   4,
		Rotation: 45,
	}
	rect2 := ntt.Rect{
		Origin: rl.Vector2{X: 4, Y: 8},
		Width:  4,
		Height: 4,
	}

	ok := ntt.Collides(&rect1, &rect2)
	if !ok {
		t.Errorf("Rects not collide but supposed to: \n%v\n%v", rect1, rect2)
	}
}
func TestRectRotatedNotCollide(t *testing.T) {

	rect1 := ntt.Rect{
		Origin: rl.Vector2{X: 4, Y: 3},
		Width:  4,
		Height: 4,
	}
	rect2 := ntt.Rect{
		Origin: rl.Vector2{X: 4, Y: 8},
		Width:  4,
		Height: 4,
	}

	ok := ntt.Collides(&rect1, &rect2)
	if ok {
		t.Errorf("Rects collide but not supposed to: \n%v\n%v", rect1, rect2)
	}
}

func TestObjectsBB(t *testing.T) {

	rect := ntt.Rect{
		Origin: rl.Vector2{X: 0, Y: 0},
		Width:  4,
		Height: 4,
	}

	bb := ntt.BB(&rect)
	if bb.X != -2 || bb.Y != -2 {
		t.Errorf("bb x wrong %v", bb)
	}
	if bb.Width != 4 || bb.Height != 4 {
		t.Errorf("bb width/hegiht is wrong %v", bb)
	}
}
