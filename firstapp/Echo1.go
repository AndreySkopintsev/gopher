package main

import (
	"fmt"
	"os"
)

func main() {

	for i := 1; i < len(os.Args); i++ {
		num := i
		index := "Index "
		result := fmt.Sprint(index, num)
		fmt.Println(result + " Value:" + os.Args[i])
	}
}
