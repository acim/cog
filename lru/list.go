package lru

type el[K comparable, V any] struct {
	next, prev *el[K, V]
	key        K
	value      V
}

type list[K comparable, V any] struct {
	root el[K, V]
	size uint
}

func newList[K comparable, V any]() *list[K, V] {
	l := &list[K, V]{size: 0} //nolint:exhaustruct
	l.root.next = &l.root
	l.root.prev = &l.root

	return l
}

func (l *list[K, V]) remove(e *el[K, V]) {
	e.prev.next = e.next
	e.next.prev = e.prev
	e.next = nil
	e.prev = nil
	l.size--
}

func (l *list[K, V]) prepend(k K, v V) *el[K, V] {
	e := &el[K, V]{key: k, value: v} //nolint:exhaustruct
	e.prev = &l.root
	e.next = l.root.next
	e.prev.next = e
	e.next.prev = e
	l.size++

	return e
}

func (l *list[K, V]) toTop(e *el[K, V]) {
	if l.root.next == e {
		return
	}

	e.prev.next = e.next
	e.next.prev = e.prev
	e.prev = &l.root
	e.next = l.root.next
	e.prev.next = e
	e.next.prev = e
}

func (l *list[K, V]) last() *el[K, V] {
	if l.size == 0 {
		return nil
	}

	return l.root.prev
}
