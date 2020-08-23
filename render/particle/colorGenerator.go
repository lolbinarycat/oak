package particle

import (
	"image/color"

	"github.com/oakmound/oak/v2/alg"
	"github.com/oakmound/oak/v2/shape"

	"github.com/200sc/go-dist/intrange"
)

// A ColorGenerator generates ColorParticles
type ColorGenerator struct {
	BaseGenerator
	StartColor, StartColorRand color.Color
	EndColor, EndColorRand     color.Color
	// The size, in pixel radius, of spawned particles
	Size    intrange.Range
	EndSize intrange.Range
	//
	// Some sort of particle type, for rendering triangles or squares or circles...
	Shape shape.Shape
}

// NewColorGenerator returns a new color generator with some applied options.
func NewColorGenerator(options ...func(Generator)) Generator {
	g := new(ColorGenerator)
	g.setDefaults()

	for _, opt := range options {
		opt(g)
	}

	return g
}

func (cg *ColorGenerator) setDefaults() {
	cg.BaseGenerator.setDefaults()
	cg.StartColor = color.RGBA{0, 0, 0, 0}
	cg.StartColorRand = color.RGBA{0, 0, 0, 0}
	cg.EndColor = color.RGBA{0, 0, 0, 0}
	cg.EndColorRand = color.RGBA{0, 0, 0, 0}
	cg.Size = intrange.Constant(1)
	cg.EndSize = intrange.Constant(1)
	cg.Shape = shape.Square
}

// Generate creates a source using this generator
func (cg *ColorGenerator) Generate(layer int) *Source {
	// Convert rotation from degrees to radians
	if cg.Rotation != nil {
		cg.Rotation = cg.Rotation.Mult(alg.DegToRad)
	}
	return NewSource(cg, layer)
}

// GenerateParticle creates a particle from a generator
func (cg *ColorGenerator) GenerateParticle(bp *baseParticle) Particle {
	return &ColorParticle{
		baseParticle: bp,
		startColor:   randColor(cg.StartColor, cg.StartColorRand),
		endColor:     randColor(cg.EndColor, cg.EndColorRand),
		size:         float64(cg.Size.Poll()),
		endSize:      float64(cg.EndSize.Poll()),
	}
}

// GetParticleSize on a color generator returns that the particles
// are per-particle specificially sized
func (cg *ColorGenerator) GetParticleSize() (float64, float64, bool) {
	return 0, 0, true
}

// Coloration
//

// SetStartColor lets cg have its color be set
func (cg *ColorGenerator) SetStartColor(sc, scr color.Color) {
	cg.StartColor = sc
	cg.StartColorRand = scr
}

// SetEndColor lets cg have its end color be set
func (cg *ColorGenerator) SetEndColor(ec, ecr color.Color) {
	cg.EndColor = ec
	cg.EndColorRand = ecr
}

//
// Sizing
//

// A Sizeable is a generator that can have some size set to it
type Sizeable interface {
	SetSize(i intrange.Range)
	SetEndSize(i intrange.Range)
}

// Size is an option to set a Sizeable size
func Size(i intrange.Range) func(Generator) {
	return func(g Generator) {
		if g2, ok := g.(Sizeable); ok {
			g2.SetSize(i)
		}
	}
}

// EndSize sets the end size of a Sizeable
func EndSize(i intrange.Range) func(Generator) {
	return func(g Generator) {
		if g2, ok := g.(Sizeable); ok {
			g2.SetEndSize(i)
		}
	}
}

// SetSize satisfies Sizeable
func (cg *ColorGenerator) SetSize(i intrange.Range) {
	cg.Size = i
}

// SetEndSize stasfies Sizeable
func (cg *ColorGenerator) SetEndSize(i intrange.Range) {
	cg.EndSize = i
}

//
// Shaping
//

// SetShape satisfies Shapeable
func (cg *ColorGenerator) SetShape(sf shape.Shape) {
	cg.Shape = sf
}
