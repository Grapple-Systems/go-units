package units

import (
	"math"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestAngleSystems(t *testing.T) {
	SI := "metric"
	// radian
	assert.Equal(t, SI, Radian.System())
	assert.Equal(t, SI, ExaRadian.System())
	assert.Equal(t, SI, PetaRadian.System())
	assert.Equal(t, SI, TeraRadian.System())
	assert.Equal(t, SI, GigaRadian.System())
	assert.Equal(t, SI, MegaRadian.System())
	assert.Equal(t, SI, KiloRadian.System())
	assert.Equal(t, SI, HectoRadian.System())
	assert.Equal(t, SI, DecaRadian.System())
	assert.Equal(t, SI, DeciRadian.System())
	assert.Equal(t, SI, CentiRadian.System())
	assert.Equal(t, SI, MilliRadian.System())
	assert.Equal(t, SI, MicroRadian.System())
	assert.Equal(t, SI, NanoRadian.System())
	assert.Equal(t, SI, PicoRadian.System())
	assert.Equal(t, SI, FemtoRadian.System())
	assert.Equal(t, SI, AttoRadian.System())
	// degree
	assert.Equal(t, SI, Degree.System())
}

func TestAngleConversion(t *testing.T) {
	revolution := NewValue(1, Revolution)

	value, err := revolution.Convert(Radian)
	if assert.NoError(t, err) {
		assert.Equal(t, NewValue(2*math.Pi, Radian), value)
	}
	value, err = revolution.Convert(Degree)
	if assert.NoError(t, err) {
		assert.Equal(t, NewValue(360, Degree), value)
	}
	value, err = revolution.Convert(Arcminute)
	if assert.NoError(t, err) {
		assert.Equal(t, NewValue(360*60, Arcminute), value)
	}
	value, err = revolution.Convert(Arcsecond)
	if assert.NoError(t, err) {
		assert.Equal(t, NewValue(360*60*60, Arcsecond), value)
	}
	value, err = revolution.Convert(Gradian)
	if assert.NoError(t, err) {
		assert.Equal(t, NewValue(400, Gradian), value)
	}
}
