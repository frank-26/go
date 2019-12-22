// Copyright © 2016 Alan A. A. Donovan & Brian W. Kernighan.
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/

// See page 6.
//!+

// Echo2 prints its command-line arguments.
package main

import (
	"fmt"
	"os"
)

func main() {
	s, sep := "", ""
	for _, arg := range os.Args[1:] {
		s += sep + arg
		sep = " "
	}
	fmt.Println(s)
}

//!-

/*
每次循环迭代，range产生一对值；索引以及在该索引处的元素值。
这个例子不需要索引，但range的语法要求, 要处理元素, 必须处理索引。一种思路是把索引赋值给一个临时变量,
如temp, 然后忽略它的值，但Go语言不允许使用无用的局部变量（local variables），因为这会导致编译错误。
空标识符_（blank identifier）应运而生
*/

/*
声明一个变量有好几种方式：
s := "" //短变量声明，最简洁，但只能用在函数内部，而不能用于包变量
var s string //依赖于字符串的默认初始化零值机制，被初始化为""
var s = "" //用得很少，除非同时声明多个变量。
var s string = "" //形式显式地标明变量的类型，当变量类型与初值类型相同时，类型冗余，但如果两者类型不同，变量类型就必须了

*/