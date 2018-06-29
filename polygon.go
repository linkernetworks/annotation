package annotation

type Point int

type PolygonAnnotation struct {
	Label  string  `bson:"label" json:"label"`
	Points []Point `bson:"points" json:"points"`
}

func AnyPolygonAnnotation(annots []Annotation) bool {
	for _, annot := range annots {
		if annot.Polygon != nil && len(annot.Polygon.Points) > 0 {
			return true
		}
	}
	return false
}
