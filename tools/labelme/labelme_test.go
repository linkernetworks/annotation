package labelme

import (
	"encoding/json"
	"testing"

	"github.com/linkernetworks/annotation"
	"github.com/stretchr/testify/assert"
)

func TestLabelMe(t *testing.T) {
	expectedJSON := `{
		"shapes": [
		  {
			"label": "aaa",
			"line_color": null,
			"fill_color": null,
			"points": [
			  [
				197,
				210
			  ],
			  [
				411,
				187
			  ],
			  [
				247,
				323
			  ]
			]
		  }
		],
		"lineColor": [
		  0,
		  255,
		  0,
		  128
		],
		"fillColor": [
		  255,
		  0,
		  0,
		  128
		],
		"imagePath": "../pictures/60168783_p0.png",
		"imageData": "deadbeaf"
	  }`

	l := &LabelmeJSON{
		ImagePath: "../pictures/60168783_p0.png",
		ImageData: "deadbeaf",
		LineColor: [4]int{0, 255, 0, 128},
		FillColor: [4]int{255, 0, 0, 128},
	}

	ann := annotation.PolygonAnnotation{
		Label:  "aaa",
		Points: []annotation.Point{197, 210, 411, 187, 247, 323},
	}
	s := PolygonAnnotationToShape(ann, nil, nil)
	l.AddShape(s)

	ret, err := l.JSON()
	assert.NoError(t, err)

	compareLabel := LabelmeJSON{}
	err = json.Unmarshal([]byte(expectedJSON), &compareLabel)
	assert.NoError(t, err)
	compareString, err := json.MarshalIndent(compareLabel, "  ", " ")
	assert.NoError(t, err)
	t.Log(string(compareString))
	t.Log(string(ret))
	assert.Equal(t, compareString, ret)
}
