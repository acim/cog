package lru

import (
	"testing"
)

func TestRemove(t *testing.T) {
	t.Parallel()

	l := newList[string, string]()
	el := l.prepend("foo", "bar")

	if el.key != "foo" || el.value != "bar" { //nolint:goconst
		t.Fatal("key and value doesn't match")
	}

	l.remove(el)

	if l.size != 0 {
		t.Errorf("want size 0")
	}
}

func TestToTop(t *testing.T) {
	t.Parallel()

	l := newList[string, string]()

	l.prepend("foo", "bar")
	l.prepend("bar", "baz")

	el := l.last()

	l.toTop(el)

	el = l.last()
	if el.key != "bar" || el.value != "baz" {
		t.Errorf("last() = (%q, %q); want ('bar', 'baz')", el.key, el.value)
	}
}

func TestLast(t *testing.T) {
	t.Parallel()

	l := newList[string, string]()
	el := l.last()

	if el != nil {
		t.Errorf("last() = %v; want nil", el)
	}

	l.prepend("foo", "bar")
	l.prepend("bar", "baz")

	if l.size != 2 {
		t.Fatal("want size 2")
	}

	el = l.last()
	if el.key != "foo" || el.value != "bar" {
		t.Errorf("last() = (%q, %q); want ('foo', 'bar')", el.key, el.value)
	}
}
