package shape

import (
	"ood/lab5/vector"
)

type shapes struct {
	vector.Vector
}

func (s *shapes) GetShapesCount() int {
	return len(s.Vector)
}

func (s *shapes) GetShape(index int) Shape {
	return s.Vector[index].(Shape)
}

func (s *shapes) InsertShape(shape Shape, index int) bool {
	if s.checkIfIncludedInShape(shape) {
		return false
	}
	s.Insert(shape, index)
	return true
}

func (s *shapes) RemoveShape(index int) {
	s.Delete(index)
}

func (s *shapes) checkIfIncludedInShape(shape Shape) bool {
	group := shape.GetGroup()
	if group == nil {
		return false
	}

	if s == group.getShapes() {
		return true
	}

	for index := 0; index < group.GetShapesCount(); index++ {
		if s.checkIfIncludedInShape(group.GetShape(index)) {
			return true
		}
	}

	return false
}
