package vivaldi

import (
	"math"
	"math/rand"
)

type HVector struct {
	X, Y, H float64
}

func NewHVector(x, y, height float64) *HVector {
	return &HVector{
		X: x,
		Y: y,
		H: height,
	}
}

func (a *HVector) Add(b *HVector) *HVector {
	return NewHVector(
		a.X+b.X,
		a.Y+b.Y,
		math.Abs(a.H+b.H),
	)
}

func (a *HVector) Sub(b *HVector) *HVector {
	return NewHVector(
		a.X-b.X,
		a.Y-b.Y,
		math.Abs(a.H+b.H),
	)
}

func (a *HVector) Magnitude() float64 {
	return math.Sqrt(a.X*a.X+a.Y*a.Y) + a.H
}

func (a *HVector) Scale(b float64) *HVector {
	return NewHVector(
		b*a.X,
		b*a.Y,
		b*a.H,
	)
}

func (a *HVector) Unit() *HVector {
	magnitude := a.Magnitude()
	if magnitude == 0 {
		return NewHVector(
			rand.Float64(),
			rand.Float64(),
			rand.Float64(),
		).Unit()
	}
	return a.Scale(1 / magnitude)
}

func (a *HVector) AtOrigin() bool {
	return a.X == 0 && a.Y == 0 && a.H == 0
}
