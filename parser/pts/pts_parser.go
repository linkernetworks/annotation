package pts

import (
	"bufio"
	"io"
	"log"
	"os"
	"regexp"
	"strconv"
)

type Point struct {
	X int `json:"x"`
	Y int `json:"y"`
}

func ParsePts(reader io.Reader) ([]Point, error) {
	var err error
	var points []Point
	var scanner = bufio.NewScanner(reader)
	var annoIndex = 0
	for scanner.Scan() {
		r, _ := regexp.Compile("([+-]?[0-9]*[.]?[0-9]+) ([+-]?[0-9]*[.]?[0-9]+)")

		result := r.FindStringSubmatch(scanner.Text())

		if len(result) < 3 {
			continue
		}

		xx, _ := strconv.ParseFloat(result[1], 64)
		yy, _ := strconv.ParseFloat(result[2], 64)

		x := int(xx)
		y := int(yy)
		points = append(points, Point{
			X: x,
			Y: y,
		})

		annoIndex++
	}

	if err = scanner.Err(); err != nil {
		log.Println(err)
		return points, err
	}
	return points, nil

}

func ParsePtsFile(filename string) ([]Point, error) {
	f, err := os.Open(filename)
	if nil != err {
		return []Point{}, err
	}
	defer f.Close()
	return ParsePts(f)
}
