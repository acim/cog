// Package cog contains thread safe generic cache implementations.
package cog

// Cache defines common cache methods.
type Cache[K comparable, V any] interface {
	// Set sets new key with provided value.
	Set(key K, value V)
	// Get gets value for a key.
	Get(key K) (V, bool)
}
