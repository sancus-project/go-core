package context

import (
	"context"
	"sync"
)

type contextKey struct {
	name string
}

var (
	list []*contextKey
	mu   sync.Mutex
)

func (k *contextKey) String() string {
	return "context value " + k.name
}

// Creates new key for context.WithValue()
func NewContextKey(name string) *contextKey {
	mu.Lock()
	defer mu.Unlock()

	k := &contextKey{name}

	list = append(list, k)
	return k
}

func Keys() []string {
	var v []string

	mu.Lock()
	defer mu.Unlock()

	if l := len(list); l > 0 {
		v = make([]string, 0, l)

		for i, k := range list {
			v[i] = k.name
		}
	}

	return v
}

func Values(ctx context.Context) map[string]interface{} {
	var m map[string]interface{}

	mu.Lock()
	defer mu.Unlock()

	if l := len(list); l > 0 {
		m = make(map[string]interface{}, l)

		for _, k := range list {
			m[k.name] = ctx.Value(k)
		}
	}

	return m
}
