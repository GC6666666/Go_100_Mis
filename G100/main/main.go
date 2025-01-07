package main

import (
	"fmt"
	"strings"
)

func main() {
	str := "xo123456xo"
	trimedStr := strings.TrimSuffix(str, "xo")
	fmt.Println(trimedStr)
}
