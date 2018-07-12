package labelme

import (
	"encoding/hex"
	"encoding/json"
	"fmt"
	"log"

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
	// Labelme spec not allow point or line (two points) input shape.
	// We will skip it.
	if len(s.Points) > 2 {
		if s.FillColor != nil {
			l.FillColor = *s.FillColor
			s.FillColor = nil
		}

		if s.LineColor != nil {
			l.LineColor = *s.LineColor
			s.LineColor = nil
		}
		l.Shapes = append(l.Shapes, s)
	}

}

func (l *LabelmeJSON) JSON() ([]byte, error) {
	return json.MarshalIndent(l, "  ", " ")
}

func PolygonAnnotationToShapeWithColorString(ann annotation.PolygonAnnotation, lineColor string, fillColor string) Shape {
	line := colorStringToIntArray(lineColor)
	fill := colorStringToIntArray(fillColor)
	return PolygonAnnotationToShape(ann, &line, &fill)
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

func RectAnnotationToShapeWithColorString(ann annotation.RectAnnotation, lineColor string, fillColor string) Shape {
	line := colorStringToIntArray(lineColor)
	fill := colorStringToIntArray(fillColor)
	return RectAnnotationToShape(ann, &line, &fill)
}

func RectAnnotationToShape(ann annotation.RectAnnotation, lineColor *[4]int, fillColor *[4]int) Shape {
	s := Shape{
		Label:     ann.Label,
		FillColor: fillColor,
		LineColor: lineColor,
	}
	s.Points = append(s.Points, [2]int{ann.X, ann.Y})
	s.Points = append(s.Points, [2]int{ann.X + ann.Width, ann.Y})
	s.Points = append(s.Points, [2]int{ann.X + ann.Width, ann.Y + ann.Height})
	s.Points = append(s.Points, [2]int{ann.X, ann.Y + ann.Height})
	return s
}

func colorStringToIntArray(s string) [4]int {
	var ret [4]int
	data, err := hex.DecodeString(s[1:])
	if err != nil {
		log.Printf("colorStringToIntArray: %v\n", err)
		return ret
	}
	for k, v := range data {
		fmt.Printf("% d\n", v)
		ret[k] = int(v)
	}

	return ret
}
