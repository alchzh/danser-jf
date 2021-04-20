package color

import (
	"github.com/go-gl/mathgl/mgl32"
	"github.com/wieku/danser-go/app/bmath"
	"github.com/wieku/danser-go/framework/math/math32"
	"math"
)

type Color struct {
	R, G, B, A float32
}

func NewRGBA(r, g, b, a float32) Color {
	return Color{r, g, b, a}
}

func NewRGB(r, g, b float32) Color {
	return NewRGBA(r, g, b, 1.0)
}

func NewIRGBA(r, g, b, a uint8) Color {
	return Color{float32(r) / 255, float32(g) / 255, float32(b) / 255, float32(a) / 255}
}

func NewIRGB(r, g, b uint8) Color {
	return NewIRGBA(r, g, b, 255)
}

func NewLA(lightness, a float32) Color {
	return NewRGBA(lightness, lightness, lightness, a)
}

func NewL(lightness float32) Color {
	return NewLA(lightness, 1.0)
}

func NewHSVA(h, s, v, a float32) Color {
	r, g, b := HSVToRGB(h, s, v)
	return NewRGBA(r, g, b, a)
}

func NewHSV(h, s, v float32) Color {
	return NewHSVA(h, s, v, 1.0)
}

func NewHSLuvA(h, s, l float64, a float32) Color {
	r, g, b := HSLuvToRGB(h, s, l)
	return NewRGBA(r, g, b, a)
}

func NewHSLuv(h, s, l float64) Color {
	return NewHSLuvA(h, s, l, 1.0)
}

func (c Color) GetHue() float32 {
	h1, _, _ := RGBToHSV(c.R, c.G, c.B)
	return h1
}

func (c Color) Mix(c1 Color, t float32) Color {
	t = bmath.ClampF32(t, 0.0, 1.0)
	return NewRGBA(
		c.R+(c1.R-c.R)*t,
		c.G+(c1.G-c.G)*t,
		c.B+(c1.B-c.B)*t,
		c.A+(c1.A-c.A)*t,
	)
}

func (c Color) Shift(h, s, v float32) Color {
	h1, s1, v1 := RGBToHSV(c.R, c.G, c.B)

	hR := math32.Mod(h1+h, 360)
	if hR < 0 {
		hR += 360
	}

	sR := bmath.ClampF32(s1+s, 0, 1)
	vR := bmath.ClampF32(v1+v, 0, 1)

	return NewHSVA(hR, sR, vR, c.A)
}

func (c Color) Shade(amount float32) Color {
	if amount < 0 {
		return c.Darken(-amount)
	}

	return c.Lighten(amount)
}

func (c Color) Shade2(amount float32) Color {
	if amount < 0 {
		return c.Darken(-amount)
	}

	return c.Lighten2(amount)
}

func (c Color) ShadeHSLuv(amount float32) Color {
	if amount < 0 {
		return c.DarkenHSLuv(-amount)
	}

	return c.LightenHSLuv(amount)
}

func (c Color) Darken(amount float32) Color {
	scale := math32.Max(1.0, 1.0+amount)
	return NewRGBA(c.R/scale, c.G/scale, c.B/scale, c.A)
}

func (c Color) Lighten(amount float32) Color {
	scale := math32.Max(1.0, 1.0+amount)
	return NewRGBA(c.R*scale, c.G*scale, c.B*scale, c.A)
}

func (c Color) Lighten2(amount float32) Color {
	amount *= 0.5
	scale := 1.0 + 0.5*amount

	return NewRGBA(
		math32.Min(1.0, c.R*scale+amount),
		math32.Min(1.0, c.G*scale+amount),
		math32.Min(1.0, c.B*scale+amount),
		c.A)
}

func (c Color) DarkenHSLuv(amount float32) Color {
	scale := float64(math32.Min(0, 1.0-amount))
	h, s, l := RGBToHSLuv(c.R, c.G, c.B)

	return NewHSLuvA(
		h,
		s,
		l*scale,
		c.A)
}

func (c Color) LightenHSLuv(amount float32) Color {
	h, s, l := RGBToHSLuv(c.R, c.G, c.B)

	return NewHSLuvA(
		h,
		s,
		// Adjustment taken from http://colorspace.r-forge.r-project.org/reference/lighten.html
		100 - (100 - l) * float64(1 - amount),
		c.A)
}

func (c Color) PackInt() uint32 {
	rI := uint32(c.R * 255)
	gI := uint32(c.G * 255)
	bI := uint32(c.B * 255)
	aI := uint32(c.A * 255)

	return aI<<24 | bI<<16 | gI<<8 | rI
}

func (c Color) PackFloat() float32 {
	return math.Float32frombits(c.PackInt())
}

func (c Color) ToVec4() mgl32.Vec4 {
	return mgl32.Vec4{c.R, c.G, c.B, c.A}
}

func (c Color) ToArray() []float32 {
	return []float32{c.R, c.G, c.B, c.A}
}
