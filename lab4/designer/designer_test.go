package designer

import (
	"bytes"
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestDesigner(t *testing.T) {
	data := `{"type": "rectangle", "color": "green", "left_top": {"x": 1, "y": 1}, "right_bottom": {"x": 3, "y": 3}}
{"type": "triangle", "color": "red", "vertex_a": {"x": 1, "y": 1}, "vertex_b": {"x": 3, "y": 3}, "vertex_c": {"x": 1.5, "y": 4.5}}
{"type": "ellipse", "color": "pink", "center": {"x": 3.14, "y": 6.28}, "vertical_radius": 3.7, "horizontal_radius": 15}
{"type": "regular polygon", "color": "black", "center": {"x": 3.14, "y": 6.28}, "radius": 3.14, "vertices": 7}`

	expectedDescriptions := []string{
		`Green rectangle: A=(1.0, 1.0), B=(3.0, 3.0)`,
		`Red triangle: A=(1.0, 1.0), B=(3.0, 3.0), C=(1.5, 4.5)`,
		`Pink ellipse: O=(3.1, 6.3), HR=15.0, VR=3.7`,
		`Black 7-sided regular polygon: O=(3.1, 6.3), R=3.1`,
	}

	designer := New()

	buf := bytes.NewBufferString(data)
	draft, err := designer.CreateDraft(buf)
	assert.Nil(t, err)

	assert.Equal(t, len(expectedDescriptions), len(draft))
	for index, shape := range draft {
		desc := fmt.Sprint(shape)
		assert.Equal(t, expectedDescriptions[index], desc)
	}
}
func TestDesignerWrongData(t *testing.T) {
	data := `{`

	designer := New()

	buf := bytes.NewBufferString(data)
	_, err := designer.CreateDraft(buf)
	assert.NotNil(t, err)
}
