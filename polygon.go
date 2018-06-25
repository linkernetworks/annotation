package annotation

type Point float64

type PolygibAnnotation struct {
	Label  string  `bson:"label" json:"label"`
	Points []Point `bson:"points" json:"points"`
}

func AnyPolygonAnnotation(annots []Annotation) bool {
	for _, annot := range annots {
		if len(annot.Polygon.Points) >0 {
			return true
		}
	}
	return false
}
