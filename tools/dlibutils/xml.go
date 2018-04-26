package dlibutils

import (
	"encoding/xml"
	"image"
	"strconv"

	"github.com/linkernetworks/annotation"
)

// import "bitbucket.org/linkernetworks/aurora/server/parser/pts"

type Dataset struct {
	XMLName xml.Name `xml:"dataset"`
	Name    string   `xml:"Name"`
	Comment string   `xml:"comment"`
	Images  Images   `xml:"images"`
}

type Images struct {
	Images []Image `xml:"image"`
}

type Image struct {
	File string `xml:"file,attr"`
	Box  Box    `xml:"box"`
}

type Box struct {
	Top    int    `xml:"top,attr"`
	Left   int    `xml:"left,attr"`
	Width  int    `xml:"width,attr"`
	Height int    `xml:"height,attr"`
	Parts  []Part `xml:"part"`
}

type Part struct {
	Name string `xml:"name,attr"`
	X    int    `xml:"x,attr"`
	Y    int    `xml:"y,attr"`
}

func NewDataset(name string) Dataset {
	return Dataset{
		Name:   name,
		Images: Images{},
	}
}

func PointAnnotationsToParts(ps []annotation.PointAnnotation) []Part {
	parts := []Part{}
	for i, p := range ps {
		parts = append(parts, Part{
			Name: strconv.Itoa(i),
			X:    p.X,
			Y:    p.X,
		})
	}
	return parts
}

func (x *Dataset) AddImage(file string, annots annotation.AnnotationCollection, image image.Image) {
	rect := annotation.FindPointAnnotationRect("box", annots, 0, image)
	x.Images.Images = append(x.Images.Images, Image{
		File: file,
		Box: Box{
			Top:    rect.X,
			Left:   rect.Y,
			Width:  rect.Width,
			Height: rect.Height,
			Parts:  PointAnnotationsToParts(annots.PointAnnotations()),
		},
	})
}

func (x *Dataset) XML() ([]byte, error) {
	return xml.MarshalIndent(x, "  ", " ")
}
