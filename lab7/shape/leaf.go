package shape

type leaf struct{}

func (leaf) GetGroup() Group {
	return nil
}
