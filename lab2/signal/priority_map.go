package signal

import "github.com/AlexanderFadeev/ood/lab2/priority_map"

type iterateFunc func(key uint, value *connection)

func (i iterateFunc) exec(key interface{}, value interface{}) {
	i(key.(uint), value.(*connection))
}

type uintToConnectionPriorityMap interface {
	set(key uint, value *connection, priority uint)
	get(key uint) (*connection, bool)
	delete(key uint)
	iterate(iterateFunc iterateFunc)
}

func makeUintToConnectionPriorityMap() uintToConnectionPriorityMap {
	return &uintToConnectionPriorityMapImpl{
		impl: priority_map.Make(),
	}
}

type uintToConnectionPriorityMapImpl struct {
	impl priority_map.PriorityMap
}

func (m *uintToConnectionPriorityMapImpl) set(key uint, value *connection, priority uint) {
	m.impl.Set(key, value, priority)
}

func (m *uintToConnectionPriorityMapImpl) get(key uint) (*connection, bool) {
	val, ok := m.impl.Get(key)
	if !ok {
		return nil, false
	}

	return val.(*connection), ok
}

func (m *uintToConnectionPriorityMapImpl) delete(key uint) {
	m.impl.Delete(key)
}

func (m *uintToConnectionPriorityMapImpl) iterate(fn iterateFunc) {
	m.impl.Iterate(fn.exec)
}
