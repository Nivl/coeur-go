// Package secret contains structs and methods to store secrets
package secret

// String represents a secret containing a string
type String struct {
	key string
}

// NewString creates a new secret string.
func NewString(secret string) String {
	// no need to register a new secret if the value is empty
	if secret == "" {
		return String{}
	}

	key := globalStore.Register(secret)
	return String{key: key}
}

// Value returns the secret value as a string.
func (s String) Value() string {
	if s.key == "" {
		return ""
	}

	secret, ok := globalStore.Get(s.key)
	if !ok {
		return ""
	}
	return secret.(string)
}

// Free removes the secret from the memory
// The secret won't be usable anymore
func (s String) Free() {
	if s.key == "" {
		return
	}

	globalStore.Remove(s.key)
}

// String implements fmt.Stringer and redacts the sensitive value.
func (s String) String() string {
	return Redacted
}

// GoString implements fmt.GoStringer
func (s String) GoString() string {
	return Redacted
}

// MarshalJSON implements json.MarshalJSON
func (s String) MarshalJSON() ([]byte, error) {
	return []byte(`"` + Redacted + `"`), nil
}
