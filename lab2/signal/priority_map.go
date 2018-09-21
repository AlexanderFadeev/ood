package signal

import "ood/lab2/priority_map"

type iterateFunc func(key uint, value Slot)

func (i iterateFunc) exec(key interface{}, value interface{}) {
	i(key.(uint), value.(Slot))
}

type uintToSlotPriorityMap interface {
	set(key uint, value Slot, priority uint)
	get(key uint) (Slot, bool)
	delete(key uint)
	iterate(iterateFunc iterateFunc)
}

func makeUintToSlotPriorityMap() uintToSlotPriorityMap {
	return &uintToSlotPriorityMapImpl{
		impl: priority_map.Make(),
	}
}

type uintToSlotPriorityMapImpl struct {
	impl priority_map.PriorityMap
}

func (m *uintToSlotPriorityMapImpl) set(key uint, value Slot, priority uint) {
	m.impl.Insert(key, value, priority)
}

func (m *uintToSlotPriorityMapImpl) get(key uint) (Slot, bool) {
	val, ok := m.impl.Get(key)
	if !ok {
		return nil, false
	}

	return val.(Slot), ok
}

func (m *uintToSlotPriorityMapImpl) delete(key uint) {
	m.impl.Delete(key)
}

func (m *uintToSlotPriorityMapImpl) iterate(fn iterateFunc) {
	m.impl.Iterate(fn.exec)
}
