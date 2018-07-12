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

func TestLabelMeRect(t *testing.T) {
	expectedJSON := `{
		"shapes": [
		  {
			"label": "aaa",
			"line_color": null,
			"fill_color": null,
			"points": [
			  [
				100,
				100
			  ],
			  [
				150,
				100
			  ],
			  [
				150,
				150
			  ],
			  [
				100,
				150
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

	ann := annotation.RectAnnotation{
		Label:  "aaa",
		X:      100,
		Y:      100,
		Width:  50,
		Height: 50,
	}
	s := RectAnnotationToShape(ann, nil, nil)
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

func TestLabelMeRectNagtiveHeight(t *testing.T) {
	expectedJSON := `{
		"shapes": [
		  {
			"label": "aaa",
			"line_color": null,
			"fill_color": null,
			"points": [
			  [
				150,
				150
			  ],
			  [
				100,
				150
			  ],
			  [
				100,
				100
			  ],
			  [
				150,
				100
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

	ann := annotation.RectAnnotation{
		Label:  "aaa",
		X:      150,
		Y:      150,
		Width:  -50,
		Height: -50,
	}
	s := RectAnnotationToShape(ann, nil, nil)
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

func TestColorToInt(t *testing.T) {
	// "#D0021B" only RGB
	expect := [4]int{208, 2, 27, 0}
	ans := colorStringToIntArray("#D0021B")
	assert.Equal(t, expect, ans)

	// "#D0021B0A" include RGBA
	expect = [4]int{208, 2, 27, 10}
	ans = colorStringToIntArray("#D0021B0A")
	assert.Equal(t, expect, ans)

	// low case: RGBA lower case
	expect = [4]int{208, 2, 27, 10}
	ans = colorStringToIntArray("#d0021b0a")
	assert.Equal(t, expect, ans)

	// odd-digit hexadecimal will return empty
	expect = [4]int{0, 0, 0, 0}
	ans = colorStringToIntArray("#d0021ba")
	assert.Equal(t, expect, ans)
}
