package annotation

import "image"

type RectAnnotation struct {
	Label  string `bson:"label" json:"label"`
	X      int    `bson:"x" json:"x"`
	Y      int    `bson:"y" json:"y"`
	Width  int    `bson:"width" json:"width"`
	Height int    `bson:"height" json:"height"`
}

func (ra *RectAnnotation) Rectangle() image.Rectangle {
	min := image.Point{ra.X, ra.Y}
	max := image.Point{ra.X + ra.Width, ra.Y + ra.Height}
	return image.Rectangle{Min: min, Max: max}
}

func AnyRectAnnotation(annots []Annotation) bool {
	for _, a := range annots {
		if a.Rect != nil {
			return true
		}
	}
	return false
}
