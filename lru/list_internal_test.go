package lru

import "testing"

func TestEmptyList(t *testing.T) {
	t.Parallel()

	l := newList[string, string]()
	el := l.last()

	if el != nil {
		t.Errorf("last() = %v; want nil", el)
	}
}
