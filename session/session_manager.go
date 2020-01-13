package session

/*
定义sessionManager接口
Init()				初始化
CreateSession()		创建session
GetSession()		获取session
*/
type SessionManager interface {
	Init(Addr string,options ...string)error
	CreateSession() (Session,error)
	GetSession(sessionId string)(Session,error)
}