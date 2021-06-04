package context

type contextKey struct {
	name string
}

func (k *contextKey) String() string {
	return "context value " + k.name
}

// Creates new key for context.WithValue()
func NewContextKey(name string) *contextKey {
	return &contextKey{name}
}
