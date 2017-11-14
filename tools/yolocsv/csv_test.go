package yolocsv

import (
	"fmt"
	"image"
	_ "image/jpeg"
	"log"
	"os"
	"testing"

	"bitbucket.org/linkernetworks/cv-tracker/src/annotation"
)

func TestYoloCSV(t *testing.T) {
	c := new(YoloCSV)
	c.Data = [][]string{
		{"f1", "0.2", "0.3", "0.4", "0.5", "label1"},
		{"f2", "0.21", "0.31", "0.41", "0.51", "label2"},
	}
	byt, _ := c.CSV()
	fmt.Println(string(byt))
}

func TestAnnotationCSV(t *testing.T) {
	r1 := annotation.RectAnnotation{
		Label:  "l1",
		X:      20,
		Y:      30,
		Width:  100,
		Height: 90,
	}

	r2 := annotation.RectAnnotation{
		Label:  "l2",
		X:      30,
		Y:      40,
		Width:  140, //max 150
		Height: 80,  //max 103
	}

	ann1 := annotation.Annotation{
		Id:   123,
		Type: "test1",
		Rect: &r1,
	}
	ann2 := annotation.Annotation{
		Id:   234,
		Type: "test2",
		Rect: &r2,
	}
	var annCols annotation.AnnotationCollection
	annCols = append(annCols, ann1)
	annCols = append(annCols, ann2)

	reader, err := os.Open("../../../../tests/fixtures/lfw/Aaron_Eckhart/Aaron_Eckhart_0001.jpg")
	if err != nil {
		log.Fatal(err)
	}
	defer reader.Close()

	m, _, err := image.Decode(reader)
	if err != nil {
		log.Fatal(err)
	}

	c := new(YoloCSV)
	c.AddImage("test1", annCols, m)
	byt, _ := c.CSV()
	fmt.Println(string(byt))
}
