package main

import (
	"fmt"
	"os"
)

type point struct {
	x, y int
}

func main() {

	p := point{1, 2}

	fmt.Printf("struct1: %v\n", p)             // 打印结构体 struct1: {1 2}
	fmt.Printf("struct1: %+v\n", p)            // 打印结构体，包含字段名 struct1: {x:1 y:2}
	fmt.Printf("struct1: %#v\n", p)            // struct1: main.point{x:1, y:2}
	fmt.Printf("type: %T\n", p)                // 打印类型 type: main.point
	fmt.Printf("bool: %t\n", true)             // 打印布尔型 bool: true
	fmt.Printf("int: %d\n", 123)               // 打印整数类型 int: 123
	fmt.Printf("bin: %b\n", 14)                // 打印二进制类型 bin: 1110
	fmt.Printf("char: %c\n", 98)               // 打印char类型  char: b
	fmt.Printf("hex: %x\n", 456)               // 打印16进制类型 hex: 1c8
	fmt.Printf("float: %f\n", 78.9)            // 打印浮点型 float: 78.900000
	fmt.Printf("float:%e\n", 123400000.0)      // 打印科学计数法 float:1.234000e+08
	fmt.Printf("float:%E\n", 123400000.0)      // 打印科学计数法 float:1.234000E+08
	fmt.Printf("str:%s\n", "\"string\"")       // 打印基本字符串 str:"string"
	fmt.Printf("str:%q\n", "\"string\"")       // 打印带双引号字符串 str:"\"string\""
	fmt.Printf("str: %x\n", "hex this")        // 打印字符串的十六进制  str: 6865782074686973
	fmt.Printf("pointer: %p\n", &p)            // 打印指针的地址
	fmt.Printf("|%5d|%5d|\n", 12, 345)         // 打印字符串用固定宽度 |   12|  345|
	fmt.Printf("|%6.3f|%6.3f|\n", 1.2, 3.45)   // 打印float类型 | 1.200| 3.450|
	fmt.Printf("|%-6.3f|%-6.3f|\n", 1.2, 3.45) // 左对齐打印 |1.200 |3.450 |
	fmt.Printf("|%6s|%6s|\n", "foo", "b")      // 打印字符串用固定格式 |   foo|     b|
	fmt.Printf("|%-6s|%-6s|\n", "foo", "b")    // 左对齐打印字符串 |foo   |b     |

	s := fmt.Sprintf("sprintf: a %s", "string")
	fmt.Println(s)

	fmt.Fprintf(os.Stderr, "io: an %s\n", "error")
}
