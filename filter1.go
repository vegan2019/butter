/*	Copyright (c) 2020, Serhat Şevki Dinçer.
	This Source Code Form is subject to the terms of the Mozilla Public
	License, v. 2.0. If a copy of the MPL was not distributed with this
	file, You can obtain one at http://mozilla.org/MPL/2.0/.
*/

package butter

// First order filter
type filter1 struct {
	s, a1, b0, b1 float64
}

func (f *filter1) Next(u float64) float64 {
	t := u - f.a1*f.s // Direct Form 2
	y := f.b0*t + f.b1*f.s
	f.s = t // shift memory
	return y
}

func (f *filter1) NextS(u, y []float64) {
	n := len(u)
	if n > len(y) {
		n = len(y)
	}
	for i := 0; i < n; i++ {
		y[i] = f.Next(u[i])
	}
}

func (f *filter1) Reset(u, y float64) {
	f.s = (y - f.b0*u) / (f.b1 - f.b0*f.a1) // division is safe
}

// NewLowPass1 creates first order Low-Pass filter
func NewLowPass1(wc float64) Filter {
	if !valid(wc) {
		return nil
	}
	var f filter1
	wat := prewarp(wc) // wa * dt

	a0 := wat + 2
	f.a1 = (wat - 2) / a0
	f.b0 = wat / a0
	f.b1 = f.b0
	f.s = 0
	return &f
}

// NewHighPass1 creates first order High-Pass filter
func NewHighPass1(wc float64) Filter {
	if !valid(wc) {
		return nil
	}
	var f filter1
	wat := prewarp(wc) // wa * dt

	a0 := wat + 2
	f.a1 = (wat - 2) / a0
	f.b0 = 2 / a0
	f.b1 = -f.b0
	f.s = 0
	return &f
}
