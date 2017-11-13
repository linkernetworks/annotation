package pascalvoc

import (
	"fmt"
	"testing"
)

func TestVOCXML(t *testing.T) {

	v := &Voc{}
	v.Data = VocAnnotation{
		FileName: "test1.jpg",
		ImageSize: Size{
			Height: 480,
			Width:  360,
			Depth:  9,
		},
		Segmented: 0,
		Objects: []Object{
			Object{
				Name:     "",
				Diffcult: 0,
				BoundingBox: BoundBox{
					Xmin: 120,
					Xmax: 133,
					Ymin: 211,
					Ymax: 255,
				},
			},
		},
	}
	out, err := v.XML()
	if err != nil {
		t.Error(err)
	}
	fmt.Println(string(out))
}
