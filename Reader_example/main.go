package main

import (
	"fmt"
	"io"
)

type phoneReader string

func (ph phoneReader) Read(p []byte) (int, error) {
	count := 0
	for i := 0; i < len(ph); i++ {
		if ph[i] >= '0' && ph[i] <= '9' {
			p[count] = ph[i]
			count++
		}
	}
	return count, io.EOF
}

func main() {
	phone1 := phoneReader("+7(809)900 20 20")
	phone2 := phoneReader("+7-950-324-99-22")

	buffer := make([]byte, len(phone1))
	phone1.Read(buffer)
	fmt.Println(string(buffer))

	buffer = make([]byte, len(phone2))
	phone2.Read(buffer)
	fmt.Println(string(buffer))
}
