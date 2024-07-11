package ntt

import (
	"math"

	rl "github.com/gen2brain/raylib-go/raylib"
)

const DEFAULT_WEAPON_MARGIN = 5

func Collides(shape1, shape2 Shape) bool {
	sep1 := FindMinSeparation(shape1, shape2)
	sep2 := FindMinSeparation(shape2, shape1)
	return sep1 <= 0 && sep2 <= 0
}

func WeaponOffset(sprite rl.Texture2D) rl.Vector2 {
	return rl.Vector2{
		X: float32(PLAYER_WIDTH/2 + sprite.Width/2),
		Y: float32(PLAYER_HEIGHT/2 + sprite.Height),
	}
}

func Overlap(a, b rl.Rectangle) (float32, float32) {
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

func BB(s Shape) rl.Rectangle {
	var min_x, max_x, min_y, max_y float32
	min_x = math.MaxFloat32
	min_y = math.MaxFloat32
	max_x = -math.MaxFloat32
	max_y = -math.MaxFloat32

	for _, v := range s.Vertices() {
		if v.X < min_x {
			min_x = v.X
		}
		if v.X > max_x {
			max_x = v.X
		}
		if v.Y < min_y {
			min_y = v.Y
		}
		if v.Y > max_y {
			max_y = v.Y
		}
	}

	return rl.Rectangle{
		X:      min_x,
		Y:      min_y,
		Width:  max_x - min_x,
		Height: max_y - min_y,
	}
}

func FindMinSeparation(shape1, shape2 Shape) float64 {
	separation := -math.MaxFloat64
	vertices1 := shape1.Vertices()
	vertices2 := shape2.Vertices()

	for i, va := range vertices1 {
		normal := Normal(va, vertices1[(i+1)%len(vertices1)])
		min_sep := math.MaxFloat64
		for _, vb := range vertices2 {
			projection := rl.Vector2DotProduct(rl.Vector2Subtract(vb, va), normal)
			min_sep = math.Min(min_sep, float64(projection))
		}

		if min_sep > separation {
			separation = min_sep
		}
	}
	return separation
}

// Gets a normal for the Shape edge
func Normal(v1, v2 rl.Vector2) rl.Vector2 {
	edge := rl.Vector2Subtract(v2, v1)
	length := rl.Vector2Length(edge)
	return rl.Vector2{X: -edge.Y / length, Y: edge.X / length}
}

func LookAt(object, target rl.Vector2) float32 {
	delta := rl.Vector2Subtract(target, object)
	return float32(math.Atan2(float64(delta.Y), float64(delta.X))) * rl.Rad2deg
}

func RotatePoint(point, pivot rl.Vector2, angle float32) rl.Vector2 {
	s := float32(math.Sin(float64(angle * rl.Deg2rad)))
	c := float32(math.Cos(float64(angle * rl.Deg2rad)))

	// Translate point to origin
	point.X -= pivot.X
	point.Y -= pivot.Y

	// Rotate point
	xnew := point.X*c - point.Y*s
	ynew := point.X*s + point.Y*c

	// Translate point back
	point.X = xnew + pivot.X
	point.Y = ynew + pivot.Y

	return point
}
