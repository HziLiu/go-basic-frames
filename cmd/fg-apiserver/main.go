package main

import (
	"flag"
	"fmt"
)

// Go 程序默认入口函数，阅读项目代码的入口
func main() {
	// 解析命令行参数
	option1 := flag.String("option1", "default_value", "Description of option 1")
	option2 := flag.Int("option2", 0, "Description of option 2")
	flag.Parse()

	// 执行简单的业务逻辑
	fmt.Println("Option 1 value:", *option1)
	fmt.Println("Option 2 value:", *option2)

	// 添加业务逻辑代码
}
