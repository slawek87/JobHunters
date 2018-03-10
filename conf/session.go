package conf

import (
	"github.com/astaxie/beego/session"
	_ "github.com/astaxie/beego/session/redis"
)

var GlobalSessions *session.Manager

func SessionInit() {
	ManagerConf := session.ManagerConfig{
		CookieName:     "gosessionid",
		Gclifetime:     3600,
		ProviderConfig: "127.0.0.1:6379,1"}

	GlobalSessions, err := session.NewManager("redis", &ManagerConf)

	if err != nil {
		panic(err)
	}

	go GlobalSessions.GC()
}
