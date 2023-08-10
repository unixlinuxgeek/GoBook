// Измените программу так, чтобы она выводила также os.Args[0] имя выполняемой программы
package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {
	var s, sep string

	for i := 0; i < len(os.Args); i++ {
		s += strconv.Itoa(i) + " \n" + sep + os.Args[i]
	}
	fmt.Println(s)
}
