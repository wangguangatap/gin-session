package session

import (
	"fmt"
	"sync"
)

/*
 * 将session写入内存
 */

type MemorySession struct {
	SessionId string
	Data map[string]interface{}
	rwlock sync.RWMutex
}

/*
* 构造函数创建实例化
*/
func CreateMemorySession(id string)*MemorySession{
	return &MemorySession{
		SessionId: id,
		Data:      make(map[string]interface{},4),
	}
}


/*
* set方法设置session
* @param  id string
* @param  value interface{}
* @return
*/
func(m *MemorySession)Set(key string,value interface{}){
	m.rwlock.Lock()
	defer m.rwlock.Unlock()
	m.Data[key]=value
}

/*
* get放获取session
* @param id string
* @return
*/
func(m *MemorySession)Get(key string)(interface{},error){
	m.rwlock.RLock()
	defer m.rwlock.RUnlock()
	value,ok:=m.Data[key]
	if !ok{
		return nil,fmt.Errorf("%v的session不存在",key)
	}
	return value,nil
}

/*
* delete放法删除session
* @param id string
* @return
 */
func(m *MemorySession)Delete(key string){
	m.rwlock.Lock()
	defer m.rwlock.Unlock()
	delete(m.Data,key)
}


/*
* save方法保存session，再memory中没有用
* @param id string
* @return
 */
func(m *MemorySession)Save()error{
	return nil
}


