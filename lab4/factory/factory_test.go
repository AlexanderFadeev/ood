package factory

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFactoryInvalidDescription(t *testing.T) {
	table := []string{
		`{"oops"`,
		`{}`,
		`{"type": "plumbus"}`,
		`{"type": 42}`,
		`{"type": "plumbus", "color": "green"}`,
		`{"type": "ellipse"}`,
		`{"type": "ellipse", "color": 42}`,
		`{"type": "ellipse", "color": "invalid"}`,
		`{"type": "ellipse", "color": "green"}`,
		`{"type": "ellipse", "color": "yellow", "center": {"x": 5, "y": 5}}`,
		`{"type": "ellipse", "color": "green", "center": {"x": 5, "y": 5}, "vertical_radius": 3}`,
		`{"type": "rectangle", "color": "blue"}`,
		`{"type": "rectangle", "color": "blue", "left_top": {"x": 1, "y": 2}}`,
		`{"type": "triangle", "color": "green"}`,
		`{"type": "triangle", "color": "green", "vertex_a": {"x": 1, "y": 2}}`,
		`{"type": "triangle", "color": "green", "vertex_a": {"x": 1, "y": 2}, "vertex_b": {"x": 1, "y": 2}}`,
		`{"type": "regular polygon", "color": "green"}`,
		`{"type": "regular polygon", "color": "green", "center": {"x": 5, "y": 5}}`,
		`{"type": "regular polygon", "color": "green", "center": {"x": 5, "y": 5}, "radius": 3}`,
	}

	factory := New()
	for _, desc := range table {
		_, err := factory.CreateShape(desc)
		assert.NotNilf(t, err, "Expected factory to fail at desc `%s`", desc)
	}
}

func TestFactoryValidDescription(t *testing.T) {
	table := []struct{ jsonDesc, stringDesc string }{
		{`{"type": "rectangle", "color": "green", "left_top": {"x": 1, "y": 1}, "right_bottom": {"x": 3, "y": 3}}`,
			`Green rectangle: A=(1.0, 1.0), B=(3.0, 3.0)`,
		},
		{`{"type": "triangle", "color": "red", "vertex_a": {"x": 1, "y": 1}, "vertex_b": {"x": 3, "y": 3},
			"vertex_c": {"x": 1.5, "y": 4.5}}`, `Red triangle: A=(1.0, 1.0), B=(3.0, 3.0), C=(1.5, 4.5)`,
		},
		{`{"type": "ellipse", "color": "pink", "center": {"x": 3.14, "y": 6.28}, "vertical_radius": 3.7,
			"horizontal_radius": 15}`, `Pink ellipse: O=(3.1, 6.3), HR=15.0, VR=3.7`,
		},
		{`{"type": "regular polygon", "color": "black", "center": {"x": 3.14, "y": 6.28}, "radius": 3.14,
			"vertices": 7}`, `Black 7-sided regular polygon: O=(3.1, 6.3), R=3.1`,
		},
	}

	factory := New()
	for _, row := range table {
		shape, err := factory.CreateShape(row.jsonDesc)
		if !assert.Nil(t, err) {
			continue
		}
		assert.Equal(t, row.stringDesc, shape.String())
	}
}
