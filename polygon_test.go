package annotation

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestAnyPolygonAnnotationInVaild(t *testing.T) {
	annots := []Annotation{}

	// generate a annotation with no rect should return false
	annots = append(annots, Annotation{})
	notFound := AnyPointAnnotation(annots)
	assert.False(t, notFound, "Annotation has no polygon")
}