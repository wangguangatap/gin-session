package session

/*
定义session接口
所有session必须实现
Get()	获取单条session value值
Set()   设置单条session value值
Del()   删除单条session value值
Save()  保存session
*/
type Session interface {
	Get(key string)(interface{},error)
	Set(key string,value interface{})
	Delete(key string)
	Save()error
}
