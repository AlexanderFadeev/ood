package signal

type value struct {
	v interface{}
}

func (v *value) set(newValue interface{}) error {
	v.v = newValue
	return nil
}

func (v *value) get() interface{} {
	return v.v
}
