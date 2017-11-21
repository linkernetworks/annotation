package annotation

import "bitbucket.org/linkernetworks/cv-tracker/server/parser/pts"
import "strconv"
import "math"
import "image"

func Min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func Max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

type PointAnnotation struct {
	Label string `bson:"label" json:"label"`
	X     int    `bson:"x" json:"x"`
	Y     int    `bson:"y" json:"y"`
}

type RectAnnotation struct {
	Label  string `bson:"label" json:"label"`
	X      int    `bson:"x" json:"x,omitempty"`
	Y      int    `bson:"y" json:"y,omitempty"`
	Width  int    `bson:"width" json:"width,omitempty"`
	Height int    `bson:"height" json:"height,omitempty"`
}

func (ra *RectAnnotation) Rectangle() image.Rectangle {
	min := image.Point{ra.X, ra.Y}
	max := image.Point{ra.X + ra.Width, ra.Y + ra.Height}
	return image.Rectangle{Min: min, Max: max}
}

type Annotation struct {
	Id   int    `bson:"id" json:"id"`
	Type string `bson:"type" json:"type"`

	// Exclusive fields
	Rect  *RectAnnotation  `bson:"rect,omitempty" json:"rect,omitempty"`
	Point *PointAnnotation `bson:"point,omitempty" json:"point,omitempty"`
}

type AnnotationCollection []Annotation

func (ac AnnotationCollection) PointAnnotations() []PointAnnotation {
	var ps []PointAnnotation
	for _, a := range ac {
		if a.Point != nil {
			ps = append(ps, *a.Point)
		}
	}
	return ps
}

func (ac AnnotationCollection) RectAnnotations() []RectAnnotation {
	var rs []RectAnnotation
	for _, a := range ac {
		if a.Rect != nil {
			rs = append(rs, *a.Rect)
		}
	}
	return rs
}

func FindPointAnnotationRect(label string, annots AnnotationCollection, padding int, image image.Image) RectAnnotation {
	bounds := image.Bounds()
	minx := math.MaxInt64
	miny := math.MaxInt64
	maxx := 0
	maxy := 0
	for _, annot := range annots {
		if annot.Point == nil {
			continue
		}
		minx = Min(minx, annot.Point.X)
		miny = Min(minx, annot.Point.Y)

		maxx = Max(maxx, annot.Point.X)
		maxy = Max(maxy, annot.Point.Y)
	}
	return RectAnnotation{
		Label:  label,
		X:      Max(minx-padding, 0),
		Y:      Max(miny-padding, 0),
		Width:  Min(maxx-minx+padding, bounds.Max.X),
		Height: Min(maxy-miny+padding, bounds.Max.Y),
	}
}

func PointToAnnotation(points []pts.Point) []Annotation {
	var annotations []Annotation
	index := 0
	for _, v := range points {
		annotations = append(annotations, Annotation{
			Id:   index,
			Type: "point",
			Point: &PointAnnotation{
				Label: strconv.Itoa(index),
				X:     v.X,
				Y:     v.Y,
			},
		})
		index++
	}
	return annotations
}

func AnyRectAnnotation(annots []Annotation) bool {
	for _, a := range annots {
		if a.Rect != nil {
			return true
		}
	}
	return false
}

func AnyPointAnnotation(annots []Annotation) bool {
	for _, annot := range annots {
		if annot.Point != nil {
			return true
		}
	}
	return false
}
