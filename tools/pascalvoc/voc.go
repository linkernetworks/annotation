package pascalvoc

import (
	"encoding/xml"
	"image"
	"os"

	"bitbucket.org/linkernetworks/cv-tracker/src/annotation"
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
    <diffcult>0</diffcult>
    <bndbox>
      <xmin>0</xmin>
      <ymin>0</ymin>
      <xmax>150</xmax>
      <ymax>150</ymax>
    </bndbox>
  </object>
</annotation>
*/

//VocAnnotation - Pascal VOC for annotation
type VocAnnotation struct {
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
	Diffcult    int      `xml:"diffcult"`
	BoundingBox BoundBox `xml:"bndbox"`
}

//BoundBox -
type BoundBox struct {
	Xmin int `xml:"xmin"`
	Ymin int `xml:"ymin"`
	Xmax int `xml:"xmax"`
	Ymax int `xml:"ymax"`
}

//Voc -
type Voc struct {
	Data VocAnnotation `xml:"annotation"`
}

//PrintVoc : Just print pascal data to stdout
//Just command line call it and pipe to file
func (*Voc) PrintVoc(a VocAnnotation) {
	enc2 := xml.NewEncoder(os.Stdout)
	enc2.Indent("  ", "    ")
	enc2.Encode(a)
}

//AddImage -
func (v *Voc) AddImage(file string, annots annotation.AnnotationCollection, image image.Image) {
	//FIXME Hardcode the label name to BTS for demo.
	rects := annots.RectAnnotations()
	imgRec := image.Bounds()
	XBoundary := imgRec.Max.X - imgRec.Min.X
	YBoundary := imgRec.Max.Y - imgRec.Min.Y
	var objs []Object

	for _, rect := range rects {
		var xMaxValue, yMaxValue int
		// if the rect max value is greater than image width.
		// it shoud be given the image width
		if (rect.X + rect.Width) > XBoundary {
			xMaxValue = XBoundary
		} else {
			xMaxValue = rect.X + rect.Width
		}
		// if the rect max value is greater than image height.
		// it shoud be given the image height
		if (rect.Y + rect.Height) > YBoundary {
			xMaxValue = YBoundary
		} else {
			yMaxValue = rect.Y + rect.Height
		}

		obj := Object{
			Name:     rect.Label,
			Diffcult: 0,
			BoundingBox: BoundBox{
				Xmin: rect.X,
				Xmax: xMaxValue,
				Ymin: rect.Y,
				Ymax: yMaxValue,
			}}
		objs = append(objs, obj)
	}
	v.Data = VocAnnotation{
		FileName: file,
		ImageSize: Size{
			Height: YBoundary,
			Width:  XBoundary,
			Depth:  0,
		},
		Segmented: 0,
		Objects:   objs,
	}
}

//XML -
func (v *Voc) XML() ([]byte, error) {
	return xml.MarshalIndent(v.Data, "  ", " ")

}
