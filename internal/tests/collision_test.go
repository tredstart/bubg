package tests

import (
	"testing"

	rl "github.com/gen2brain/raylib-go/raylib"
	"github.com/tredstart/bubg/internal/ntt"
)

func TestCollidePolygonangles(t *testing.T) {

	rect1 := ntt.NewPolygon(rl.Vector2{X: 4, Y: 4}, 4, 4, 0, rl.Color{})
	rect2 := ntt.NewPolygon(rl.Vector2{X: 5, Y: 5}, 4, 4, 0, rl.Color{})

	ok := ntt.Collides(rect1, rect2)
	if !ok {
		t.Errorf("Polygons do not collide but supposed to: \n%v\n%v", rect1, rect2)
	}
}

func TestPolygonsDontCollide(t *testing.T) {

	rect1 := ntt.NewPolygon(rl.Vector2{X: 0, Y: 0}, 4, 4, 0, rl.Color{})
	rect2 := ntt.NewPolygon(rl.Vector2{X: 5, Y: 5}, 4, 4, 0, rl.Color{})

	ok := ntt.Collides(rect1, rect2)
	if ok {
		t.Errorf("Polygons collide but not supposed to: \n%v\n%v", rect1, rect2)
	}
}

func TestPolygonRotatedCollide(t *testing.T) {

	rect1 := ntt.NewPolygon(rl.Vector2{X: 4, Y: 4}, 4, 4, 45, rl.Color{})
	rect2 := ntt.NewPolygon(rl.Vector2{X: 4, Y: 8}, 4, 4, 0, rl.Color{})

	ok := ntt.Collides(rect1, rect2)
	if !ok {
		t.Errorf("Polygons not collide but supposed to: \n%v\n%v", rect1, rect2)
	}
}
func TestPolygonRotatedNotCollide(t *testing.T) {

	rect1 := ntt.NewPolygon(rl.Vector2{X: 4, Y: 3}, 4, 4, 45, rl.Color{})
	rect2 := ntt.NewPolygon(rl.Vector2{X: 4, Y: 8}, 4, 4, 0, rl.Color{})

	ok := ntt.Collides(rect1, rect2)
	if ok {
		t.Errorf("Polygons collide but not supposed to: \n%v\n%v", rect1, rect2)
	}
}
