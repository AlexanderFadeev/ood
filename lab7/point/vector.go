package point

type Vector Point

func NewVector(from, to Point) Vector {
	return Vector{
		X: to.X - from.X,
		Y: to.Y - from.Y,
	}
}

func (v Vector) Negative() Vector {
	return Vector{
		X: -v.X,
		Y: -v.Y,
	}
}

func (v Vector) PairwiseProduct(v2 Vector) Vector {
	return Vector{
		X: v.X * v2.X,
		Y: v.Y * v2.Y,
	}
}

func (v Vector) PairwiseDivision(v2 Vector) Vector {
	return Vector{
		X: v.X / v2.X,
		Y: v.Y / v2.Y,
	}
}
