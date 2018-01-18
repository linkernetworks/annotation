package pascalvoc

import (
	"bytes"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestParse(t *testing.T) {
	testxml := `<annotation>
  <filename>test1.jpg</filename>
  <size>
    <width>480</width>
    <height>360</height>
    <depth>16</depth>
  </size>
  <segmented>0</segmented>
  <object>
    <name>car</name>
    <difficult>0</difficult>
    <bndbox>
      <xmin>120</xmin>
      <ymin>211</ymin>
      <xmax>133</xmax>
      <ymax>255</ymax>
    </bndbox>
  </object>
</annotation>`

	voc, err := Parse(bytes.NewReader([]byte(testxml)))
	assert.NoError(t, err)
	xml, err := voc.XML()
	assert.NoError(t, err)
	assert.Equal(t, testxml, string(xml))
}

func TestVOCXML(t *testing.T) {
	v := NewVocXml("test1.jpg", 480, 360, 16)
	v.AddObject(Object{
		Name:      "car",
		Difficult: 0,
		BoundingBox: BoundBox{
			Xmin: 120,
			Xmax: 133,
			Ymin: 211,
			Ymax: 255,
		},
	})
	out, err := v.XML()
	if err != nil {
		t.Error(err)
	}

	expectedXml := `<annotation>
  <filename>test1.jpg</filename>
  <size>
    <width>480</width>
    <height>360</height>
    <depth>16</depth>
  </size>
  <segmented>0</segmented>
  <object>
    <name>car</name>
    <difficult>0</difficult>
    <bndbox>
      <xmin>120</xmin>
      <ymin>211</ymin>
      <xmax>133</xmax>
      <ymax>255</ymax>
    </bndbox>
  </object>
</annotation>`

	t.Log(string(out))
	assert.Equal(t, expectedXml, string(out))
}
