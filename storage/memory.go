package storage

import (
	"encoding/json"
	"os"
	"sync"
)

const storageFile = "urls.json"

type MemoryStore struct {
	data map[string]string
	mu   sync.RWMutex
}

func NewMemoryStore() *MemoryStore {
	store := &MemoryStore{
		data: make(map[string]string),
	}
	store.loadFromDisk()
	return store
}

func (m *MemoryStore) Save(code string, url string) {
	m.mu.Lock()
	defer m.mu.Unlock()
	m.data[code] = url
	m.saveToDisk()
}

func (m *MemoryStore) Get(code string) (string, bool) {
	m.mu.RLock()
	defer m.mu.RUnlock()
	url, exists := m.data[code]
	return url, exists
}

func (m *MemoryStore) GetAll() map[string]string {
	m.mu.RLock()
	defer m.mu.RUnlock()
	copy := make(map[string]string)
	for k, v := range m.data {
		copy[k] = v
	}
	return copy
}

func (m *MemoryStore) Delete(code string) bool {
	m.mu.Lock()
	defer m.mu.Unlock()
	_, exists := m.data[code]
	if !exists {
		return false
	}
	delete(m.data, code)
	m.saveToDisk()
	return true
}

func (m *MemoryStore) saveToDisk() {
	file, err := os.Create(storageFile)
	if err != nil {
		return
	}
	defer file.Close()
	json.NewEncoder(file).Encode(m.data)
}

func (m *MemoryStore) loadFromDisk() {
	file, err := os.Open(storageFile)
	if err != nil {
		return
	}
	defer file.Close()
	json.NewDecoder(file).Decode(&m.data)
}
