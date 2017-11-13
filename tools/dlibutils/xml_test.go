package dlibutils

import (
	"encoding/xml"
	"testing"
)

import (
	_ "gopkg.in/mgo.v2"
	_ "gopkg.in/mgo.v2/bson"
)

func check(t *testing.T, err error) {
	if err != nil {
		t.Fatal(err)
	}
}

func TestGenerateShapePredictorTrainingXml(t *testing.T) {
	val := Dataset{
		Name: "Face Recognition Training",
		Comment: `These are images from the PASCAL VOC 2011 dataset.
   The face landmarks are from dlib's shape_predictor_68_face_landmarks.dat
   landmarking model.  The model uses the 68 landmark scheme used by the iBUG
   300-W dataset`,
		Images: Images{
			Images: []Image{
				Image{
					File: "test.jpg",
					Box: Box{
						Top: 20, Left: 20, Width: 100, Height: 100,
						Parts: []Part{
							Part{"00", 201, 107},
							Part{"01", 201, 103},
						},
					},
				},
			},
		},
	}
	out, err := xml.MarshalIndent(val, "  ", " ")
	check(t, err)
	t.Logf("xml: %s", out)
}
