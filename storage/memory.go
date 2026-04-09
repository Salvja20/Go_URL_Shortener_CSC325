package storage

import "sync"

type MemoryStore struct {
	data map[string]string
	mu   sync.RWMutex
}

func NewMemoryStore() *MemoryStore {
	return &MemoryStore{
		data: make(map[string]string),
	}
}

func (m *MemoryStore) Save(code string, url string) {
	m.mu.Lock()
	defer m.mu.Unlock()

	m.data[code] = url
}

func (m *MemoryStore) Get(code string) (string, bool) {
	m.mu.RLock()
	defer m.mu.RUnlock()

	url, exists := m.data[code]
	return url, exists
}