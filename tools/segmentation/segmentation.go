package segmentation

import (
	"image"
	"image/color"

	"github.com/linkernetworks/annotation"
	"github.com/llgcode/draw2d/draw2dimg"
	"github.com/llgcode/draw2d/draw2dkit"
)

var fillColorPicker []color.Color

func init() {
	fillColorPicker = append(fillColorPicker, color.RGBA{0x7f, 0xff, 0x00, 0xff}) //chartreuse
	fillColorPicker = append(fillColorPicker, color.RGBA{0x44, 0x44, 0x44, 0xff}) //chartreuse
	fillColorPicker = append(fillColorPicker, color.RGBA{0xdc, 0x14, 0x3c, 0xff}) //crimson
	fillColorPicker = append(fillColorPicker, color.RGBA{0xff, 0x00, 0xff, 0xff}) //fuchsia
	fillColorPicker = append(fillColorPicker, color.RGBA{0xad, 0xff, 0x2f, 0xff}) //greenyellow
	fillColorPicker = append(fillColorPicker, color.RGBA{0x4b, 0x00, 0x82, 0xff}) //indigo
}

type Point struct {
	X float64
	Y float64
}

type Object []Point

type SegmentationImage struct {
	DestImage    *image.RGBA
	GraphContext *draw2dimg.GraphicContext
	Objects      []Object
}

func NewSegmentationImage(img image.Image) *SegmentationImage {
	s := new(SegmentationImage)
	s.DestImage = image.NewRGBA(image.Rect(0, 0, img.Bounds().Dx(), img.Bounds().Dy()))
	s.GraphContext = draw2dimg.NewGraphicContext(s.DestImage)

	//Black as base image
	s.GraphContext.SetFillColor(color.RGBA{0x00, 0x00, 0x00, 0xff})
	draw2dkit.Rectangle(s.GraphContext, 0, 0, float64(img.Bounds().Dx()), float64(img.Bounds().Dy()))
	s.GraphContext.Fill()

	s.GraphContext.SetLineWidth(5)
	return s
}

func (s *SegmentationImage) AddPolygonAnnotation(polygon annotation.PolygonAnnotation) {
	var obj Object
	for i := 0; i < len(polygon.Points); {
		pt := Point{X: float64(polygon.Points[i]), Y: float64(polygon.Points[i+1])}
		obj = append(obj, pt)
		i = i + 2
	}
	s.Objects = append(s.Objects, obj)
}

// Draw all objects with the same color.
func (s *SegmentationImage) DrawSegmentationClassImage(file string) error {
	s.drawObjects(s.Objects, false)
	return draw2dimg.SaveToPngFile(file, s.DestImage)
}

// Draw all objects with the different colors.
func (s *SegmentationImage) DrawSegmentationLabelImage(file string) error {
	s.drawObjects(s.Objects, true)
	return draw2dimg.SaveToPngFile(file, s.DestImage)
}

func (s *SegmentationImage) drawObjects(objs []Object, separatedObj bool) {
	setGraphColor(s.GraphContext, 0)
	for i, obj := range objs {
		s.GraphContext.BeginPath() // Initialize a new path
		for j, pt := range obj {
			if j == 0 {
				s.GraphContext.MoveTo(pt.X, pt.Y)
			} else {
				s.GraphContext.LineTo(pt.X, pt.Y)
			}
		}
		//draw line to first point to close object
		s.GraphContext.LineTo(obj[0].X, obj[0].Y)
		s.GraphContext.Close()
		if separatedObj {
			if i >= len(fillColorPicker) {
				i = (i % len(fillColorPicker))
			}
			setGraphColor(s.GraphContext, i)
		}
		s.GraphContext.FillStroke()
	}
}

// set color with color preset table
func setGraphColor(gc *draw2dimg.GraphicContext, index int) {
	gc.SetStrokeColor(fillColorPicker[index])
	gc.SetFillColor(fillColorPicker[index])
}
