package segmentation

import (
	"image"
	_ "image/jpeg"
	"log"
	"os"
	"testing"

	"github.com/linkernetworks/annotation"
)

const sourceFile string = "../../test_img/Aaron_Eckhart_0001.jpg"
const targetLabelFile string = "./obj.jpg"
const targetClassFile string = "./class.jpg"

func prepareTestPolygon() []annotation.PolygonAnnotation {
	return []annotation.PolygonAnnotation{
		annotation.PolygonAnnotation{
			Label:  "aaa",
			Points: []annotation.Point{50, 50, 100, 50, 100, 100},
		},
		annotation.PolygonAnnotation{
			Label:  "bbb",
			Points: []annotation.Point{150, 50, 110, 50, 110, 100},
		},
	}
}
func TestSegementationLabel(t *testing.T) {
	reader, err := os.Open(sourceFile)
	if err != nil {
		log.Fatal(err)
	}
	defer reader.Close()

	m, _, err := image.Decode(reader)
	if err != nil {
		log.Fatal(err)
	}

	anns := prepareTestPolygon()

	seg := NewSegmentationImage(m)
	for _, ann := range anns {
		seg.AddPolygonAnnotation(ann)
	}
	if err := seg.DrawSegmentationLabelImage(targetLabelFile); err != nil {
		t.Fatal(err)
	}
	os.Remove(targetLabelFile)
}

func TestSegementationObj(t *testing.T) {
	reader, err := os.Open(sourceFile)
	if err != nil {
		log.Fatal(err)
	}
	defer reader.Close()

	m, _, err := image.Decode(reader)
	if err != nil {
		log.Fatal(err)
	}

	anns := prepareTestPolygon()

	seg := NewSegmentationImage(m)
	for _, ann := range anns {
		seg.AddPolygonAnnotation(ann)
	}
	if err := seg.DrawSegmentationClassImage(targetClassFile); err != nil {
		t.Fatal(err)
	}

	os.Remove(targetClassFile)
}
