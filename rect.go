package annotation

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

func (ra RectAnnotation) Canon() RectAnnotation {
	rect := image.Rect(ra.X, ra.Y, ra.X+ra.Width, ra.Y+ra.Height)
	rect = rect.Canon()
	ra.X = rect.Min.X
	ra.Y = rect.Min.Y
	ra.Width = rect.Dx()
	ra.Height = rect.Dy()
	return ra
}

// FixBounds returns a new fixed rectangle annotation
func (ra RectAnnotation) FixBounds(size image.Point) RectAnnotation {
	ra.X = Max(ra.X, 0)
	ra.Y = Max(ra.Y, 0)
	ra.Width = Min(ra.X+ra.Width, size.X) - ra.X
	ra.Height = Min(ra.Y+ra.Height, size.Y) - ra.Y
	return ra
}

func AnyRectAnnotation(annots []Annotation) bool {
	for _, a := range annots {
		if a.Rect != nil {
			return true
		}
	}
	return false
}
