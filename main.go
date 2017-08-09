package main

import (
	"bufio"
	"fmt"
	"os"
)

func line_counter(arg string) (int64, error) {

	counter := int64(0)
	file, err := os.Open(arg)
	if err != nil {
		return 0, err
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		counter++
	}
	return counter, scanner.Err()
}

func main() {

	counter, err := line_counter(os.Args[1])

	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("Количество строк в файле:", counter)
}