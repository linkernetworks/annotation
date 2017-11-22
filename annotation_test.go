package annotation

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestAnyRectAnnotationVaild(t *testing.T) {
	annots := []Annotation{}
	annots = append(annots, Annotation{})

	// generate a annotation with a rect should return true
	var r RectAnnotation
	var p PointAnnotation

	r.Label = "label"
	r.X = 5
	r.Y = 5
	r.Width = 5
	r.Height = 5
	annots = append(annots, Annotation{
		Id:    5,
		Type:  "test",
		Rect:  &r,
		Point: &p,
	})

	found := AnyPointAnnotation(annots)
	assert.True(t, found, "Annotation has a point")
}

func TestAnyRectAnnotationInVaild(t *testing.T) {
	annots := []Annotation{}

	// generate a annotation with no rect should return false
	annots = append(annots, Annotation{})
	notFound := AnyRectAnnotation(annots)
	assert.False(t, notFound, "Annotation has no rect")
}

func TestAnyPointAnnotationVaild(t *testing.T) {
	annots := []Annotation{}
	annots = append(annots, Annotation{})

	// generate a annotation with a rect should return true
	var r RectAnnotation
	var p PointAnnotation

	p.Label = "label"
	p.X = 5
	p.Y = 6
	annots = append(annots, Annotation{
		Id:    5,
		Type:  "test",
		Rect:  &r,
		Point: &p,
	})

	found := AnyPointAnnotation(annots)
	assert.True(t, found, "Annotation has a rect")
}

func TestAnyPointAnnotationInvalid(t *testing.T) {
	annots := []Annotation{}

	// generate a annotation with no point should return false
	annots = append(annots, Annotation{})
	notFound := AnyPointAnnotation(annots)
	assert.False(t, notFound, "Annotation has no point")
}
