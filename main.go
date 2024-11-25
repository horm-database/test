package main

import (
	_ "go.uber.org/automaxprocs"

	"github.com/horm-database/go-horm/horm"
)

func main() {
	TestMySQL()
}

// init 配置全局统一接入协议执行器
func init() {
	horm.SetGlobalClient("default/app.server.service1",
		horm.WithAppID(10002),
		horm.WithSecret("S959223456"),
		horm.WithTarget("ip://127.0.0.1:8180"),
		horm.WithTimeout(332320),
		horm.WithCaller("app.server.service"))

	//horm.SetGlobalProxy("10002", "S959223456", "polaris://rpc.server.access.goapi", 332320)
}
