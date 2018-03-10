package conf

import (
	"github.com/astaxie/beego/session"
	_ "github.com/astaxie/beego/session/redis"
)

func SessionInit() {
	ManagerConf := &session.ManagerConfig{
		Gclifetime:     3600,
		ProviderConfig: "127.0.0.1:6379"}

	GlobalSessions, err := session.NewManager("redis", ManagerConf)

	if err != nil {
		panic(err)
	}

	go GlobalSessions.GC()
}
