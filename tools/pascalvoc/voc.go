package pascalvoc

import (
	"encoding/xml"
	"image"
	"io"

	"github.com/linkernetworks/annotation"
)

/*
<annotation>
  <filename>3_A_1245.png</filename>
  <size>
    <width>1920</width>
    <height>1200</height>
    <depth>3</depth>
  </size>
  <segmented>0</segmented>
  <object>
    <name>car</name>
    <difficult>0</difficult>
    <bndbox>
      <xmin>0</xmin>
      <ymin>0</ymin>
      <xmax>150</xmax>
      <ymax>150</ymax>
    </bndbox>
  </object>
</annotation>
*/

//Voc -
type Voc struct {
	Data VocAnnotation
	size image.Point `xml:-`
}

//VocAnnotation - Pascal VOC for annotation
type VocAnnotation struct {
	XMLName   string   `xml:"annotation"`
	FileName  string   `xml:"filename"`
	ImageSize Size     `xml:"size"`
	Segmented int      `xml:"segmented"`
	Objects   []Object `xml:"object"`
}

//Size -
type Size struct {
	Width  int `xml:"width"`
	Height int `xml:"height"`
	Depth  int `xml:"depth"`
}

//Object -
type Object struct {
	Name        string   `xml:"name"`
	Difficult   int      `xml:"difficult"`
	BoundingBox BoundBox `xml:"bndbox"`
}

//BoundBox -
type BoundBox struct {
	Xmin int `xml:"xmin"`
	Ymin int `xml:"ymin"`
	Xmax int `xml:"xmax"`
	Ymax int `xml:"ymax"`
}

func NewVocXml(filename string, width int, height int, depth int) *Voc {
	return &Voc{
		size: image.Point{X: width, Y: height},
		Data: VocAnnotation{
			FileName:  filename,
			Segmented: 0,
			ImageSize: Size{
				Width:  width,
				Height: height,
				Depth:  depth,
			},
		},
	}
}

func RectAnnotationToObject(rect annotation.RectAnnotation) Object {
	return Object{
		Name:      rect.Label,
		Difficult: 0,
		BoundingBox: BoundBox{
			Xmin: rect.X,
			Xmax: rect.X + rect.Width,
			Ymin: rect.Y,
			Ymax: rect.Y + rect.Height,
		},
	}
}

func (v *Voc) RectAnnotations() []annotation.RectAnnotation {
	var rects []annotation.RectAnnotation
	for _, obj := range v.Data.Objects {
		var rect = annotation.RectAnnotation{
			Label:  obj.Name,
			X:      obj.BoundingBox.Xmin,
			Y:      obj.BoundingBox.Ymin,
			Width:  obj.BoundingBox.Xmax - obj.BoundingBox.Xmin,
			Height: obj.BoundingBox.Ymax - obj.BoundingBox.Ymin,
		}
		rects = append(rects, rect.Canon())
	}
	return rects
}

func (v *Voc) AnnotationCollection() annotation.AnnotationCollection {
	var annots []annotation.Annotation
	var rects = v.RectAnnotations()
	for idx, rect := range rects {
		var r = rect.Canon()
		annots = append(annots, annotation.Annotation{
			Id:   idx,
			Type: "rect",
			Rect: &r,
		})
	}
	return annots
}

func (v *Voc) AddObject(o Object) {
	v.Data.Objects = append(v.Data.Objects, o)
}

func (v *Voc) AddRectAnnotation(rect annotation.RectAnnotation) {
	var obj = RectAnnotationToObject(rect.FixBounds(v.size).Canon())
	v.AddObject(obj)
}

// AddObject addes the object from annotation collections
func (v *Voc) AddAnnotations(rects []annotation.RectAnnotation) {
	for _, rect := range rects {
		// if the rect max value is greater than image height.
		// it shoud be given the image height
		var obj = RectAnnotationToObject(rect.FixBounds(v.size).Canon())
		v.AddObject(obj)
	}
}

func Parse(reader io.Reader) (*Voc, error) {
	var voc Voc
	var decoder = xml.NewDecoder(reader)
	var err = decoder.Decode(&voc.Data)
	voc.size = image.Point{X: voc.Data.ImageSize.Width, Y: voc.Data.ImageSize.Height}
	return &voc, err
}

//XML -
func (v *Voc) XML() ([]byte, error) {
	return xml.MarshalIndent(v.Data, "", "  ")
}
