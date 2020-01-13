package session

import "sync"

type MemorySessionManager struct {
	SessionMap map[string]Session
	rwlock sync.RWMutex
}

