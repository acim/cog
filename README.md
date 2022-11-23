# cog

[![pipeline](https://github.com/acim/cog/actions/workflows/pipeline.yaml/badge.svg)](https://github.com/acim/cog/actions/workflows/pipeline.yaml)
[![Go Reference](https://pkg.go.dev/badge/go.acim.net/cog.svg)](https://pkg.go.dev/go.acim.net/cog)
[![Go Report](https://goreportcard.com/badge/go.acim.net/cog)](https://goreportcard.com/report/go.acim.net/cog)
![Go Coverage](https://img.shields.io/badge/coverage-100%25-brightgreen?style=flat&logo=go)
[![License](https://img.shields.io/badge/license-BSD--2--Clause--Patent-orange.svg)](https://github.com/acim/cog/blob/main/LICENSE)

Thread safe generic cache implementations in Go, fully tested and optimized for best performance.

## [LRU Cache](<https://en.wikipedia.org/wiki/Cache_replacement_policies#Least_recently_used_(LRU)>) Example

```go
package main

import (
	"fmt"

	"go.acim.net/cog/lru"
)

func main() {
	cache := lru.NewCache[string, string](10)

	cache.Set("foo", "bar")

	v, ok := cache.Get("foo")
	if !ok {
		panic("want value")
	}

	fmt.Printf("foo = %s\n", v)
}
```
