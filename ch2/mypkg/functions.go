package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {
	var result int64
	for _, arg := range os.Args[1:] {
		t, err := strconv.ParseInt(arg, 10, 64)
		result = Sum(result, t)
		if err != nil {
			fmt.Fprintf(os.Stderr, "Ошибка: %v\n", err)
			os.Exit(1)
		}
	}

	fmt.Printf("Результат: %d", result)
}
func Sum(a, b int64) int64 {
	return a + b
}
