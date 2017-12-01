package annotation

import (
	"github.com/stretchr/testify/assert"
	"image"
	"testing"
)

func TestRectCanonicalize(t *testing.T) {
	r := RectAnnotation{
		Label:  "test",
		X:      100,
		Y:      100,
		Width:  -100,
		Height: -90,
	}
	r = r.Canon()
	assert.Equal(t, 0, r.X)
	assert.Equal(t, 10, r.Y)
	assert.Equal(t, 100, r.Width)
	assert.Equal(t, 90, r.Height)
}

func TestRectFixBounds(t *testing.T) {
	r := RectAnnotation{
		Label:  "test",
		X:      100,
		Y:      100,
		Width:  100,
		Height: 110,
	}
	r2 := r.FixBounds(image.Point{X: 180, Y: 180})
	assert.Equal(t, 80, r2.Width)
	assert.Equal(t, 80, r2.Height)
}
