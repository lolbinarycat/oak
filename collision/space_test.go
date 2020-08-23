package collision

import (
	"testing"

	"github.com/oakmound/oak/v2/alg/floatgeom"

	"github.com/oakmound/oak/v2/physics"
	"github.com/stretchr/testify/assert"
)

func TestSpaceFuncs(t *testing.T) {
	Clear()
	s := NewUnassignedSpace(10, 10, 10, 10)
	assert.NotEmpty(t, s.String())

	// Getters
	cx, cy := s.GetCenter()
	assert.Equal(t, cx, float64(15))
	assert.Equal(t, cy, float64(15))
	x, y := s.GetPos()
	assert.Equal(t, x, float64(10))
	assert.Equal(t, y, float64(10))

	// Positional comparison
	s2 := NewUnassignedSpace(20, 20, 10, 10)
	assert.True(t, s2.Above(s) < 0)
	assert.True(t, s2.Below(s) > 0)
	assert.True(t, s2.LeftOf(s) < 0)
	assert.True(t, s2.RightOf(s) > 0)

	// Containment
	assert.False(t, s2.Contains(s))
	s3 := NewUnassignedSpace(5, 5, 20, 20)
	assert.True(t, s3.Contains(s))
	s4 := NewUnassignedSpace(15, 15, 10, 10)

	// Overlap
	xover, yover := s4.Overlap(s)
	assert.Equal(t, xover, -5.0)
	assert.Equal(t, yover, -5.0)
	xover, yover = s.Overlap(s4)
	assert.Equal(t, xover, 5.0)
	assert.Equal(t, yover, 5.0)
	xover, yover = s.Overlap(s2)
	assert.Equal(t, xover, 0.0)
	assert.Equal(t, yover, 0.0)
	ov := s.OverlapVector(s4)
	assert.Equal(t, ov, physics.NewVector(5, 5))
	spaces := s.SubtractRect(1, 1, 8, 8)
	assert.Equal(t, len(spaces), 4)
}

func TestNewRect(t *testing.T) {
	s := NewUnassignedSpace(0, 0, 0, 0)
	assert.Equal(t, 1.0, s.GetW())
	assert.Equal(t, 1.0, s.GetH())
	s = NewUnassignedSpace(0, 0, -10, -10)
	assert.Equal(t, 10.0, s.GetW())
	assert.Equal(t, 10.0, s.GetH())
	s = NewRectSpace(floatgeom.NewRect3WH(0, 0, 0, 10, 10, 0), 0, 0)
	assert.Equal(t, 10.0, s.GetW())
	assert.Equal(t, 10.0, s.GetH())
}
