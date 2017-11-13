package pascalvoc

import (
	"bitbucket.org/linkernetworks/cv-tracker/src/annotation"
	"encoding/xml"
	"image"
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

//AddImage addes the annotation object
func (v *Voc) AddImage(file string, annots annotation.AnnotationCollection, image image.Image) {
	rects := annots.RectAnnotations()
	imgRec := image.Bounds()

	imgWidth := imgRec.Max.X - imgRec.Min.X
	imgHeight := imgRec.Max.Y - imgRec.Min.Y
	var objs []Object

	for _, rect := range rects {
		// if the rect max value is greater than image height.
		// it shoud be given the image height
		obj := Object{
			Name:     rect.Label,
			Diffcult: 0,
			BoundingBox: BoundBox{
				Xmin: max(rect.X, 0),
				Xmax: min(rect.X+rect.Width, imgWidth),
				Ymin: max(rect.Y, 0),
				Ymax: min(rect.Y+rect.Height, imgHeight),
			}}
		objs = append(objs, obj)
	}

	v.Data = VocAnnotation{
		FileName: file,
		ImageSize: Size{
			Width:  imgWidth,
			Height: imgHeight,
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
