// Copyright © 2016 Alan A. A. Donovan & Brian W. Kernighan.
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/

// See page 8.
//!+

// Dup1 prints the text of each line that appears more than
// once in the standard input, preceded by its count.
package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	counts := make(map[string]int)
	input := bufio.NewScanner(os.Stdin)
	for input.Scan() {
		counts[input.Text()]++
	}
	// NOTE: ignoring potential errors from input.Err()
	for line, n := range counts {
		if n > 1 {
			fmt.Printf("%d\t%s\n", n, line)
		}
	}
}

//!-

/*
1. 正如for循环一样，if语句条件两边也不加括号，但是主体部分需要加。
2. map存储了键/值（key/value）的集合，对集合元素，提供常数时间的存、取或测试操作
3. 类似于C或其它语言里的printf函数.fmt.Printf函数对一些表达式产生格式化输出。该函数的首个参数是个格式字符串，指定后续参数被如何格式化。各个参数的格式取决于“转换字符”（conversion character），形式为百分号后跟一个字母。
	%d          十进制整数
	%x, %o, %b  十六进制，八进制，二进制整数。
	%f, %g, %e  浮点数： 3.141593 3.141592653589793 3.141593e+00
	%t          布尔：true或false
	%c          字符（rune） (Unicode码点)
	%s          字符串
	%q          带双引号的字符串"abc"或带单引号的字符'c'
	%v          变量的自然形式（natural format）
	%T          变量的类型
	%%          字面上的百分号标志（无操作数）

*/