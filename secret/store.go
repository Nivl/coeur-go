package secret

import (
	"sync"

	"github.com/google/uuid"
)

var globalStore = store{}

// Redacted represents the string displayed when
// printing the secret.
const Redacted = "REDACTED"

type store struct {
	secrets sync.Map
}

func (s *store) Register(secret interface{}) string {
	uid := uuid.New().String()
	s.secrets.Store(uid, secret)
	return uid
}

func (s *store) Get(key string) (interface{}, bool) {
	return s.secrets.Load(key)
}

func (s *store) Remove(key string) {
	s.secrets.Delete(key)
}
