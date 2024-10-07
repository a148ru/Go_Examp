package mychan

import (
	"fmt"
)

func strlen(s string, c chan int) {
	c <- len(s) // результат в канал
}

func Run() {
	// Пример каналов
	c := make(chan int) // создаем канал int
	go strlen("Salutations", c)
	go strlen("World", c)
	x, y := <-c, <-c // из канала в переменные
	fmt.Println(x, y, x+y)
}
