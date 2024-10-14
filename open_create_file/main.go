package main

import (
	"fmt"
	"os"
)

func main() {
	file, err := os.Create("test.txt")
	if err != nil {
		fmt.Println("Ошибка создания файла")
		os.Exit(-1)
	}
	defer file.Close()
	fmt.Println(file.Name(), "Создан")
}
