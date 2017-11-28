package annotation

type PointAnnotation struct {
	Label string `bson:"label" json:"label"`
	X     int    `bson:"x" json:"x"`
	Y     int    `bson:"y" json:"y"`
}

func AnyPointAnnotation(annots []Annotation) bool {
	for _, annot := range annots {
		if annot.Point != nil {
			return true
		}
	}
	return false
}
