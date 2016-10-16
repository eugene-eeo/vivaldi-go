package vivaldi

import "math"
import "errors"

const cc = 0.25
const ce = 0.5
const InitialError = 10

type Context struct {
	Vec   *HVector
	Error float64
}

func NewContext() *Context {
	return NewContextFromValues(
		NewHVector(0, 0, 0),
		InitialError,
	)
}

func NewContextFromValues(vec *HVector, error float64) *Context {
	return &Context{
		Vec:   vec,
		Error: error,
	}
}

func (ctx *Context) Update(rtt float64, peer *Context) *HVector {
	w := ctx.Error / (ctx.Error + peer.Error) // w = e_i / (e_i + e_j)
	ab := ctx.Vec.Sub(peer.Vec)               // x_i - x_j
	re := rtt - ab.Magnitude()                // rtt - |x_i - x_j|
	es := math.Abs(re) / rtt                  // e_s = ||x_i - x_j| - rtt| / rtt
	ctx.Error = es*ce*w + ctx.Error*(1-ce*w)  // e_i = e_s*c_e*w + e_i*(1 - c_e*w)
	// ∂ = c_c * w
	d := cc * w
	// x_i = x_i + ∂*(rtt - |x_i - x_j|)*u(x_i - x_j)
	ctx.Vec = ctx.Vec.Add(ab.Unit().Scale(d * re))
	return ctx.Vec
}

func (ctx *Context) EstimateRTT(vec *HVector) (float64, error) {
	if ctx.Vec.AtOrigin() || vec.AtOrigin() {
		return 0.0, errors.New("cannot estimate RTT for uninintialised vectors")
	}
	return ctx.Vec.Sub(vec).Magnitude(), nil
}
