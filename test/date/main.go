package main

import (
	"fmt"
	"time"
)

func main() {
	for i := 1; i < 10; i++ {
		t := time.Now().AddDate(0, 0, -i)
		fmt.Println(t.Format("2006-01-02"))
	}
}
