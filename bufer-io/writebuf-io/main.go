package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	rows := []string{
		"Hello Go!",
		"Welcome to Golang",
	}

	file, err := os.Create("../some.dat")
	writer := bufio.NewWriter(file)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	defer file.Close()

	for _, row := range rows {
		writer.WriteString(row)
		writer.WriteString("\n")
	}
	writer.Flush()
}
