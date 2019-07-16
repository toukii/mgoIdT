package main

import (
	"fmt"
	"os"
	"time"
	"unicode"
)

func main() {
	id := "5cf64fecb28d3aad532c2819"
	if len(os.Args) > 1 {
		id = os.Args[1]
	}
	ids := id[:8]
	sum := int64(0)
	for _, it := range ids {
		n := getNum(it)
		sum = sum<<4 + n
	}
	fmt.Println(time.Unix(sum, 0))
}

func getNum(r rune) int64 {
	if unicode.IsNumber(r) {
		return int64(r) - 48
	}
	return int64(r) - 97 + 10
}
