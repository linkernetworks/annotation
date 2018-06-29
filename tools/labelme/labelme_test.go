package labelme

import "testing"

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
				588,
				178
			  ],
			  [
				588,
				428
			  ],
			  [
				347,
				464
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
		"imagePath": "../Pictures/60168783_p0.png",
		"imageData": "deadbeaf",
	  }`

	l := NewLabelME("../Pictures/60168783_p0.png", "aaa", [4]int{0, 255, 0, 128}, [4]int{255, 0, 0, 128}, []byte("deadbeaf"))
	anns := []annotation.PointAnnotation{}
	anns = append(anns, annotation.PointAnnotation{Label: "aaa", })
	l.AddAnnotations(anns []annotation.PointAnnotation)
}
