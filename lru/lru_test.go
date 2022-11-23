package lru_test

import (
	"reflect"
	"sync"
	"testing"

	"go.acim.net/cog"
	"go.acim.net/cog/lru"
)

var _ cog.Cache[string, any] = (*lru.Cache[string, any])(nil)

var cache *lru.Cache[int, int] //nolint:gochecknoglobals

func TestConcurrency(t *testing.T) {
	t.Parallel()

	size := uint(5)
	cache := lru.NewCache[int, string](size)

	var wg sync.WaitGroup

	for k, v := range []string{"a", "b", "c", "d", "e", "f", "g", "h", "i", "j"} {
		wg.Add(1)

		go func(k int, v string) {
			defer wg.Done()

			cache.Set(k, v)
		}(k, v)
	}

	wg.Wait()

	count := uint(0)

	for _, k := range []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9} {
		if _, ok := cache.Get(k); ok {
			count++
		}
	}

	if count != size {
		t.Errorf("cache size %d; want %d", count, size)
	}
}

func TestGet(t *testing.T) {
	t.Parallel()

	type value struct {
		foo string
		bar uint
	}

	cache := lru.NewCache[string, *value](1)

	if _, ok := cache.Get("foo"); ok {
		t.Error("Get() = ok; want nok")
	}

	in := &value{foo: "foo", bar: 1}

	cache.Set("baz", in)

	got, ok := cache.Get("baz")
	if !ok {
		t.Error("Get() = nok; want ok")
	}

	if !reflect.DeepEqual(got, in) {
		t.Errorf("Get() = %v; want %v", got, in)
	}

	if _, ok = cache.Get("bar"); ok {
		t.Error("Get() = ok; want nok")
	}
}

func TestSet_SameKey(t *testing.T) {
	t.Parallel()

	cache := lru.NewCache[string, string](2)

	cache.Set("foo", "foo")
	cache.Set("foo", "bar")
	cache.Set("bar", "foo")
	cache.Set("bar", "bar")
	cache.Set("baz", "foo")
	cache.Set("baz", "bar")

	cache.Get("baz")

	got, ok := cache.Get("bar")
	if !ok {
		t.Error("got nok; want ok")
	}

	if got != "bar" {
		t.Errorf("Get() = %q; want 'bar'", got)
	}
}

func BenchmarkCache_Set(b *testing.B) {
	b.ReportAllocs()

	var wg sync.WaitGroup

	for n := 0; n < b.N; n++ {
		wg.Add(1)

		go func(n int) {
			defer wg.Done()
			cache.Set(n, n)
		}(n)
	}

	wg.Wait()
}

func BenchmarkCache_Get(b *testing.B) {
	b.ReportAllocs()

	var wg sync.WaitGroup

	for n := 0; n < b.N; n++ {
		wg.Add(1)

		go func(n int) {
			defer wg.Done()
			cache.Get(n)
		}(n)
	}

	wg.Wait()
}

func init() { //nolint:gochecknoinits
	cache = lru.NewCache[int, int](1500000)
}
