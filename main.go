package main

import (
	"github.com/horm-database/go-horm/horm"
	_ "go.uber.org/automaxprocs"
)

func main() {
	TestMySQL()
	TestRedis()
}

// init 配置全局统一接入协议执行器
func init() {
	horm.SetGlobalClient("ws_test.app1.server1.service1")

	//horm.SetGlobalClient("", horm.WithAppID(10002),
	//	horm.WithSecret("S959223456"),
	//	horm.WithTarget("ip://127.0.0.1:8180"),
	//	horm.WithTimeout(332320),
	//	horm.WithCaller("app.server.service"))

	//horm.SetGlobalProxy("10002", "S959223456", "polaris://rpc.server.access.goapi", 332320)
}
