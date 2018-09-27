package shape

type Rectangle struct {
	leftTop     Point
	rightBottom Point
}

func NewRectangle(leftTop, rightBottom Point) Rectangle {
	if leftTop.x > rightBottom.x {
		leftTop.x, rightBottom.x = rightBottom.x, leftTop.x
	}
	if leftTop.y > rightBottom.y {
		leftTop.y, rightBottom.y = rightBottom.y, leftTop.y
	}

	return Rectangle{
		leftTop:     leftTop,
		rightBottom: rightBottom,
	}
}

func (r Rectangle) GetLeftTop() Point {
	return r.leftTop
}

func (r Rectangle) GetRightBottom() Point {
	return r.rightBottom
}
