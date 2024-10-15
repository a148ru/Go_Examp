package main

import (
	"fmt"
	"os"
)

type person struct {
	name string
	age  int32
	// weight float64
}

func main() {
	filename := "data.txt"
	writeData(filename)
	p := readData(filename)

	fmt.Println(p.name, p.age)
}

func writeData(filename string) {
	// Начальные данные
	var name string = "Tom"
	var age int32 = 24

	file, err := os.Create(filename)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	defer file.Close()

	fmt.Fprintln(file, name)
	fmt.Fprintln(file, age)
}

func readData(filename string) *person {
	// var name string
	// var age int32

	var p *person = new(person)

	file, err := os.Open(filename)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	defer file.Close()

	fmt.Fscanln(file, &p.name)
	fmt.Fscanln(file, &p.age)

	return p
}
