package main

import (
	// 导入 automaxprocs 包，可以在程序启动时自动设置 GOMAXPROCS 配置，
	// 使其与 Linux 容器的 CPU 配额相匹配。
	// 这避免了在容器中运行时，因默认 GOMAXPROCS 值不合适导致的性能问题，
	// 确保 Go 程序能够充分利用可用的 CPU 资源，避免 CPU 浪费。
	"github.com/go-basic-frames/cmd/fg-apiserver/app"
	_ "go.uber.org/automaxprocs"
	"os"
)

// Go 程序默认入口函数，阅读项目代码的入口
func main() {
	// 创建项目
	command := app.NewFastGOCommand()

	// 执行命令并处理错误
	if err := command.Execute(); err != nil {
		// 错误则退出程序，且其他程序（bash 脚本）可以根据返回的状态码判断服务运行状态
		os.Exit(1)
	}

}
