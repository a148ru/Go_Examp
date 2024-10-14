package main

import (
	"fmt"
	"io"
	"os"
)

func main() {
	text := "Test text\n"
	file, err := os.Create("test.txt")
	if err != nil {
		fmt.Println("Ошибка создания файла")
		fmt.Println(err)
		os.Exit(-1)
	}
	//defer file.Close()

	// Запись строки в файл
	_, err = file.WriteString(text)
	if err != nil {
		fmt.Println("Ошибка записи в файл", file.Name())
		fmt.Println(err)
		os.Exit(-1)
	}

	// Запись не текстовой бинарной информации
	data := []byte("Test string in bytes\n")
	_, err = file.Write(data)
	if err != nil {
		fmt.Println("Ощибка добавления данных")
		fmt.Println(err)
		os.Exit(-1)
	}
	fmt.Println("Write Done.")

	file.Close()

	// Чтение данных из файла

	file, err = os.Open("test.txt")
	if err != nil {
		fmt.Println("Ошибка открытия файла на чтение")
		fmt.Println(err)
		os.Exit(-1)
	}

	defer file.Close()

	read_data := make([]byte, 64)

	for {
		n, err := file.Read(read_data)
		if err == io.EOF {
			break
		}
		fmt.Print(string(read_data[:n]))
	}

	// вывод данных из файла в консоль с помощью io.Copy
	fmt.Println("------ io.Copy------")

	file.Seek(0, io.SeekStart) // Вернуть указатель на начало файла
	n, err := io.Copy(os.Stdout, file)
	if err != nil && err != io.EOF {
		fmt.Println(err)
	}
	fmt.Println("Write", n)
}
