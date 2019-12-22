// Copyright © 2016 Alan A. A. Donovan & Brian W. Kernighan.
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/

// See page 4.
//!+

// Echo1 prints its command-line arguments.
package main

import (
	"fmt"
	"os"
)

func main() {
	var s, sep string
	// for循环三个部分不需括号包围。大括号强制要求, 左大括号必须和post语句在同一行。
	for i := 1; i < len(os.Args); i++ {//自增语句i++给i加1；这和i += 1以及i = i + 1都是等价的。对应的还有i--给i减1。它们是语句，而不像C系的其它语言那样是表达式。
		fmt.Println(os.Args[i])
		s += sep + os.Args[i]
		sep = " *** "
	}
	fmt.Println(s)
}

//!-

// go run ch1/echo1/main.go  a b // a *** b

/*
for initialization; condition; post {
    // zero or more statements
}

1. initialization语句是可选的，在循环开始前执行。initalization如果存在，必须是一条简单语句（simple statement），即，短变量声明、自增语句、赋值语句或函数调用。
2. condition是一个布尔表达式（boolean expression），其值在每次循环迭代开始时计算。如果为true则执行循环体语句。
3. post语句在循环体执行结束后执行，之后再次对condition求值。condition值为false时，循环结束。

// a traditional "while" loop
for condition {
    // ...
}

// a traditional infinite loop
for {
    // 这就变成一个无限循环，尽管如此，还可以用其他方式终止循环, 如一条break或return语句。
}
*/