package audio

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/oakmound/oak/v2/oakerr"
)

func TestErrorChannel(t *testing.T) {
	err := oakerr.ExistingElement{}
	err2 := <-errChannel(err)
	assert.Equal(t, err, err2)
}
