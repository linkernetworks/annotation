package labelme

import (
	"encoding/json"

	"github.com/linkernetworks/annotation"
)

type LabelmeJSON struct {
	Shapes    []Shape `bson:"shapes" json:"shapes"`
	LineColor [4]int  `bson:"lineColor" json:"lineColor"`
	FillColor [4]int  `bson:"fillColor" json:"fillColor"`
	ImagePath string  `bson:"imagePath" json:"imagePath"`
	ImageData string  `bson:"imageData" json:"imageData"`
}

type Shape struct {
	Label     string  `bson:"label" json:"label"`
	LineColor *[4]int `bson:"line_color" json:"line_color"`
	FillColor *[4]int `bson:"fill_color" json:"fill_color"`
	Points    []Point `bson:"points" json:"points"`
}

type Point [2]int

func (l *LabelmeJSON) AddShape(s Shape) {
	l.Shapes = append(l.Shapes, s)
}

func (l *LabelmeJSON) JSON() ([]byte, error) {
	return json.MarshalIndent(l, "  ", " ")
}

func PolygonAnnotationToShape(ann annotation.PolygonAnnotation, lineColor *[4]int, fillColor *[4]int) Shape {
	s := Shape{
		Label:     ann.Label,
		FillColor: fillColor,
		LineColor: lineColor,
	}
	for i := 0; i < len(ann.Points); {
		s.Points = append(s.Points, [2]int{int(ann.Points[i]), int(ann.Points[i+1])})
		i = i + 2
	}
	return s
}
