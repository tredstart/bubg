package ntt

import (
	"math"

	rl "github.com/gen2brain/raylib-go/raylib"
)

type Polygon struct {
	Origin    rl.Vector2
	Radius    float32
	Color     rl.Color
	angleStep float64
	Rotation  float32

	Vertices []rl.Vector2

	Filled bool
}

func NewPolygon(origin rl.Vector2, n uint8, side, rotation float32, color rl.Color) Polygon {
	polygon := Polygon{
		Color:     color,
		Radius:    side / float32(2*math.Sin(math.Pi/float64(n))),
		Vertices:  make([]rl.Vector2, n),
		angleStep: 2 * math.Pi / float64(n),
		Rotation:  rotation,
	}
	polygon.Move(origin)
	return polygon
}

func (p *Polygon) Move(new_origin rl.Vector2) {
	p.Origin = new_origin
	p.Rotation -= 45
	for i := range p.Vertices {
		angle := p.angleStep*float64(i) + float64((p.Rotation)*rl.Deg2rad)
		p.Vertices[i].X = p.Origin.X + p.Radius*float32(math.Cos(angle))
		p.Vertices[i].Y = p.Origin.Y + p.Radius*float32(math.Sin(angle))
	}
}

func (p *Polygon) Render() {
	if !p.Filled {
		rl.DrawPolyLines(p.Origin, int32(len(p.Vertices)), p.Radius, p.Rotation, p.Color)
		rl.DrawPixelV(p.Origin, p.Color)
	} else {
		rl.DrawPoly(p.Origin, int32(len(p.Vertices)), p.Radius, p.Rotation, p.Color)
	}

}
