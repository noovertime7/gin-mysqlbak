package main

import (
	"fmt"
	"github.com/gorhill/cronexpr"
	"time"
)

func main() {
	expr, err := cronexpr.Parse("*/5 * * * * * * ") // 如果表达式解析错误将返回一个错误
	if err != nil {
		fmt.Println(err)
		return
	}
	nextTime := expr.Next(time.Now())
	fmt.Println(nextTime)
}
