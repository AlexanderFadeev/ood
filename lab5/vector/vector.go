package vector

type Vector []interface{}

func (v *Vector) Insert(value interface{}, position int) {
	*v = append(*v, nil)
	copy((*v)[position+1:], (*v)[position:])
	(*v)[position] = value
}

func (v *Vector) Push(value interface{}) {
	*v = append(*v, value)
}

func (v *Vector) Delete(position int) {
	copy((*v)[position:], (*v)[position+1:])
	*v = (*v)[:len(*v)-1]
}

func (v *Vector) Pop() interface{} {
	var val interface{}
	*v, val = (*v)[:len(*v)-1], (*v)[len(*v)-1]
	return val
}
