// Copyright © 2016 Alan A. A. Donovan & Brian W. Kernighan.
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/

// See page 1.

// Helloworld is our first Go program.
//!+
package main // Go语言的代码通过包（package）组织，包类似于其它语言里的库（libraries）或者模块（modules）。

import "fmt" //一系列导入（import）的包

func main() {
	fmt.Println("Hello, 世界")
}

//!-

// $ go build helloworld.go 编译该程序，并保存以备将来用
