package labelme

import (
	"encoding/json"

	"github.com/linkernetworks/annotation"
)

type LabelME struct {
	Shapes    []Shape `json:"shapes"`
	LineColor [4]int  `json:"line_color"`
	FillColor [4]int  `json:"fill_color"`
	ImagePath string  `json:"imagePath"`
	ImageData string  `json:"imageData"`
}

type Shape struct {
	File      string  `xml:"file,attr"`
	Label     string  `json:"label"`
	LineColor *[4]int `json:"line_color"`
	FillColor *[4]int `json:"fill_color"`
	Points    [2]int  `json:"points"`
}

func NewLabelME(imgPath string, label string, lineColor [4]int, fillColor [4]int, imgData []byte) *LabelME {
	return &LabelME{
		ImagePath: imgPath,
		LineColor: lineColor,
		FillColor: fillColor,
		ImageData: string(imgData),
	}
}

func (l *LabelME) AddAnnotation(ann annotation.PolygonAnnotation) {
	s := Shape{Lebel: ann.Label}
	for _, v := range ann.Points {
		s.Points = append(s.Points, []int{})
	}
	l.Shapes = append(l.Shapes, s)
}

func (l *LabelME) JSON() ([]byte, error) {
	return json.MarshalIndent(l, "  ", " ")
}
