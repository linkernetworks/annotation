package annotation

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestAnyRectAnnotationVaild(t *testing.T) {
	annots := AnnotationCollection{}

	// generate a annotation with a rect should return true
	var r RectAnnotation

	r.Label = "label"
	r.X = 5
	r.Y = 5
	r.Width = 5
	r.Height = 5
	annots = append(annots, Annotation{
		Id:   5,
		Type: "test",
		Automated: true,
		Rect: &r,
	})

	found := AnyRectAnnotation(annots)
	assert.True(t, found, "Annotation has a point")
	assert.Equal(t, 1, len(annots.RectAnnotations()))
	assert.Equal(t, 0, len(annots.PointAnnotations()))
	assert.Equal(t, 0, len(annots.PolygonAnnotations()))
}

func TestAnyRectAnnotationInVaild(t *testing.T) {
	annots := AnnotationCollection{}

	// generate a annotation with no rect should return false
	annots = append(annots, Annotation{})
	notFound := AnyRectAnnotation(annots)
	assert.False(t, notFound, "Annotation has no rect")
}

func TestAnyPointAnnotationVaild(t *testing.T) {
	annots := AnnotationCollection{}

	// generate a annotation with a rect should return true
	var p PointAnnotation

	p.Label = "label"
	p.X = 5
	p.Y = 6
	annots = append(annots, Annotation{
		Id:    5,
		Type:  "test",
		Automated: true,
		Point: &p,
	})

	found := AnyPointAnnotation(annots)
	assert.True(t, found, "Annotation has a rect")

	assert.Equal(t, 0, len(annots.RectAnnotations()))
	assert.Equal(t, 1, len(annots.PointAnnotations()))
	assert.Equal(t, 0, len(annots.PolygonAnnotations()))
}

func TestAnyPointAnnotationValid(t *testing.T) {
	annots := AnnotationCollection{}

	// generate a annotation with no point should return false
	annots = append(annots, Annotation{})
	notFound := AnyPointAnnotation(annots)
	assert.False(t, notFound, "Annotation has no point")
}

func TestAnyPolygonAnnotationValid(t *testing.T) {
	annots := AnnotationCollection{}

	// generate a polygon annotation
	poly := PolygonAnnotation{
		Label:  "test1",
		Points: []Point{200, 300},
	}

	annots = append(annots, Annotation{
		Id:      1,
		Type:    "test",
		Label:   "test label",
		Automated: true,
		Polygon: &poly,
	})
	found := AnyPolygonAnnotation(annots)
	assert.True(t, found, "Annotation has polygon")
	assert.Equal(t, 0, len(annots.RectAnnotations()))
	assert.Equal(t, 0, len(annots.PointAnnotations()))
	assert.Equal(t, 1, len(annots.PolygonAnnotations()))
}

func TestAnyPolygonAnnotationInVaild(t *testing.T) {
	annots := AnnotationCollection{}

	// generate a annotation with no rect should return false
	annots = append(annots, Annotation{})
	notFound := AnyPolygonAnnotation(annots)
	assert.False(t, notFound, "Annotation has no polygon")
}
