package units

import "math"

var (
	Angle = UnitOptionQuantity("angle")

	// revolution
	Revolution = NewUnit("revolution", "rev", Angle, SI)

	// metric
	Radian      = NewUnit("radian", "rad", Angle, SI)
	ExaRadian   = Exa(Radian)
	PetaRadian  = Peta(Radian)
	TeraRadian  = Tera(Radian)
	GigaRadian  = Giga(Radian)
	MegaRadian  = Mega(Radian)
	KiloRadian  = Kilo(Radian)
	HectoRadian = Hecto(Radian)
	DecaRadian  = Deca(Radian)
	DeciRadian  = Deci(Radian)
	CentiRadian = Centi(Radian)
	MilliRadian = Milli(Radian)
	MicroRadian = Micro(Radian)
	NanoRadian  = Nano(Radian)
	PicoRadian  = Pico(Radian)
	FemtoRadian = Femto(Radian)
	AttoRadian  = Atto(Radian)

	// degree
	Degree    = NewUnit("degree", "deg", Angle, SI, UnitOptionAliases("°"))
	Arcminute = NewUnit("arcminute", "arcmin", Angle, UnitOptionAliases("MOA"))
	Arcsecond = NewUnit("arcsecond", "arcsec", Angle)

	// gradian
	Gradian = NewUnit("gradian", "grad", Angle)
)

func init() {
	NewRatioConversion(Revolution, Radian, 2.0*math.Pi)
	NewRatioConversion(Revolution, Degree, 360)
	NewRatioConversion(Degree, Arcminute, 60)
	NewRatioConversion(Arcminute, Arcsecond, 60)
	NewRatioConversion(Revolution, Gradian, 400)
}
