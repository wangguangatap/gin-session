package session

import (
	"fmt"
	uuid "github.com/satori/go.uuid"
	"sync"
)

type MemorySessionManager struct {
	SessionMap map[string]*MemorySession
	rwlock sync.RWMutex
}

/*
*构造函数
*/
func NewMemorySessionManager()*MemorySessionManager{
	return &MemorySessionManager{
		SessionMap: make(map[string]*MemorySession,1024),
	}
}

/*
* Init初始化，再memory中没用
*/
func(m *MemorySessionManager) Init(Addr string,options ...string)error{
	return nil
}

/*
*CreateSession方法
*/
func(m *MemorySessionManager)CreateSession() *MemorySession{
	m.rwlock.Lock()
	defer m.rwlock.Unlock()
	sessionId:=uuid.NewV4()
	return NewMemorySession(sessionId.String())
}

/*
*GetSession
 */
func(m *MemorySessionManager)GetSession(sessionId string)(*MemorySession,error){
	m.rwlock.RLock()
	defer m.rwlock.RUnlock()
	data,ok:=m.SessionMap[sessionId]
	if !ok{
		return nil,fmt.Errorf("%v不存在",sessionId)
	}
	return data,nil
}


