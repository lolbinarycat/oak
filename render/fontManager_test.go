package render

import (
	"image"
	"testing"

	"github.com/oakmound/oak/v2/oakerr"
	"github.com/stretchr/testify/assert"
)

func TestFontManager(t *testing.T) {

	initTestFont()

	fm := NewFontManager()

	fm.Get("def")
	// That may or may not be nil depending on if this is being run in a -coverprofile test
	// or not. Todo: fiddle with fonts and fix it
	//assert.NotNil(t, f)
	f := fm.Get("other")
	assert.Nil(t, f)

	fg := FontGenerator{
		RawFile: luxisrTTF,
		Color:   image.Black,
	}

	err := fm.NewFont("other", fg)
	assert.Nil(t, err)

	f = fm.Get("other")
	assert.NotNil(t, f)

	err = fm.NewFont("def", fg)
	assert.NotNil(t, err)
	if exists, ok := err.(oakerr.ExistingElement); ok {
		assert.True(t, exists.Overwritten)
	}
}
