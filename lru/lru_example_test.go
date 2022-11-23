package lru_test

import (
	"fmt"

	"go.acim.net/cog/lru"
)

func ExampleCache_Get() {
	type value struct {
		foo string
		bar uint
	}

	in := &value{foo: "foo", bar: 1}

	cache := lru.NewCache[string, *value](1)

	cache.Set("baz", in)

	got, ok := cache.Get("baz")
	if !ok {
		panic("want value")
	}

	fmt.Printf("foo = %s, bar = %d", got.foo, got.bar)

	// Output:
	// foo = foo, bar = 1
}
