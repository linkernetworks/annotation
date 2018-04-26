package yolocsv

import (
	"bytes"
	"encoding/csv"
	"fmt"
	"image"

	"github.com/linkernetworks/annotation"
)

//YoloCSV -
///notebook/chrisjan/project/images/nokia3/jpg/59f95e6e1c35e40001a92492_0.jpg	0.569393	0.659381	0.114002	0.063752	BTS
//center x, center y, width and height
//x,y,w,h都是相對於圖片尺寸width, height的百分比
type YoloCSV struct {
	//Data - Each row contains as follow:
	// FileName string
	// CenterX  string
	// CenterY  string
	// Width    string
	// Height   string
	// Label    string
	Data [][]string
}

//AddImage - Add image to csv
func (y *YoloCSV) AddImage(file string, annots annotation.AnnotationCollection, image image.Image) {
	rects := annots.RectAnnotations()
	imgRec := image.Bounds()

	imgWidth := imgRec.Max.X - imgRec.Min.X
	imgHeight := imgRec.Max.Y - imgRec.Min.Y

	for _, rect := range rects {
		fmt.Println(rect.Width, imgWidth, float32(rect.Width)/float32(imgWidth))
		fmt.Println(float32(annotation.Min(rect.Width, imgWidth)) / float32(imgWidth))
		fmt.Println(imgHeight)

		var record []string
		//FileName
		record = append(record, file)
		//CenterX
		record = append(record, fmt.Sprintf("%f", float32(annotation.Min(rect.X+rect.Width, imgWidth))/float32(2)/float32(imgWidth)))
		//CenterY
		record = append(record, fmt.Sprintf("%f", float32(annotation.Min(rect.Y+rect.Height, imgHeight))/float32(2)/float32(imgHeight)))
		//Width
		record = append(record, fmt.Sprintf("%f", float32(annotation.Min(rect.Width, imgWidth))/float32(imgWidth)))
		//Height
		record = append(record, fmt.Sprintf("%f", float32(annotation.Min(rect.Height, imgHeight))/float32(imgHeight)))
		//Label
		record = append(record, rect.Label)

		y.Data = append(y.Data, record)
	}
}

//CSV -
func (y *YoloCSV) CSV() ([]byte, error) {
	var b bytes.Buffer
	w := csv.NewWriter(&b)
	w.WriteAll(y.Data) // calls Flush internally
	return b.Bytes(), nil
}
