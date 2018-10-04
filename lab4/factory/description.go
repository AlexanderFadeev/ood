package factory

import (
	"encoding/json"

	"ood/lab4/color"
	"ood/lab4/point"
	"ood/lab4/shape"

	"github.com/pkg/errors"
)

type shapeDescription struct {
	Type   *shapeType   `json:"type"`
	Color  *color.Color `json:"color"`
	Center *point.Point `json:"center"`
	rectangleDescription
	triangleDescription
	ellipseDescription
	regularPolygonDescription
}

func newDescription(descriptionStr string) (*shapeDescription, error) {
	var desc shapeDescription
	err := json.Unmarshal([]byte(descriptionStr), &desc)
	if err != nil {
		return nil, errors.Wrap(err, "Failed to unmarshal description")
	}

	return &desc, nil
}

func (sd *shapeDescription) toShape() (shape.Shape, error) {
	if sd.Type == nil {
		return nil, errors.New("Missing shape type")
	}
	if sd.Color == nil {
		return nil, errors.New("Missing shape color")
	}

	switch *sd.Type {
	case shapeTypeRectangle:
		return sd.rectangleDescription.toShape(*sd.Color)
	case shapeTypeTriangle:
		return sd.triangleDescription.toShape(*sd.Color)
	case shapeTypeEllipse:
		return sd.ellipseDescription.toShape(*sd.Color, sd.Center)
	case shapeTypeRegularPolygon:
		return sd.regularPolygonDescription.toShape(*sd.Color, sd.Center)
	default:
		return nil, errors.New("Missing shape type")
	}
}

type rectangleDescription struct {
	LeftTop     *point.Point `json:"left_top"`
	RightBottom *point.Point `json:"right_bottom"`
}

func (rd *rectangleDescription) toShape(color color.Color) (shape.Shape, error) {
	if rd.LeftTop == nil {
		return nil, errors.New("Missing rectangle left top vertex")
	}
	if rd.RightBottom == nil {
		return nil, errors.New("Missing rectangle right bottom vertex")
	}

	return shape.NewRectangle(*rd.LeftTop, *rd.RightBottom, color), nil
}

type triangleDescription struct {
	VertexA *point.Point `json:"vertex_a"`
	VertexB *point.Point `json:"vertex_b"`
	VertexC *point.Point `json:"vertex_c"`
}

func (td *triangleDescription) toShape(color color.Color) (shape.Shape, error) {
	if td.VertexA == nil {
		return nil, errors.New("Missing triangle vertex A")
	}
	if td.VertexB == nil {
		return nil, errors.New("Missing triangle vertex B")
	}
	if td.VertexC == nil {
		return nil, errors.New("Missing triangle vertex C")
	}

	return shape.NewTriangle(*td.VertexA, *td.VertexB, *td.VertexC, color), nil
}

type ellipseDescription struct {
	VerticalRadius   *float64 `json:"vertical_radius"`
	HorizontalRadius *float64 `json:"horizontal_radius"`
}

func (ed *ellipseDescription) toShape(color color.Color, center *point.Point) (shape.Shape, error) {
	if center == nil {
		return nil, errors.New("Missing ellipse center")
	}
	if ed.VerticalRadius == nil {
		return nil, errors.New("Missing ellipse vertical radius")
	}
	if ed.HorizontalRadius == nil {
		return nil, errors.New("Missing ellipse horizontal radius")
	}

	return shape.NewEllipse(*center, *ed.HorizontalRadius, *ed.VerticalRadius, color)
}

type regularPolygonDescription struct {
	Radius   *float64 `json:"radius"`
	Vertices *uint    `json:"vertices"`
}

func (rpd *regularPolygonDescription) toShape(color color.Color, center *point.Point) (shape.Shape, error) {
	if center == nil {
		return nil, errors.New("Missing regular polygon center")
	}
	if rpd.Radius == nil {
		return nil, errors.New("Missing regular polygon radius")
	}
	if rpd.Vertices == nil {
		return nil, errors.New("Missing regular polygon vertices count")
	}

	return shape.NewRegularPolygon(*rpd.Vertices, *center, *rpd.Radius, color)
}
