package info

import "sync"

type DollsMutex struct {
	mu    sync.Mutex
	Dolls map[string]*DollProfile
}

func (dm *DollsMutex) Write(doll *DollProfile) {
	dm.mu.Lock()
	defer dm.mu.Unlock()

	dm.Dolls[doll.Name] = doll
}

func (dm *DollsMutex) Read(field string) (*DollProfile, bool) {
	dm.mu.Lock()
	defer dm.mu.Unlock()

	v, ok := dm.Dolls[field]

	return v, ok
}
