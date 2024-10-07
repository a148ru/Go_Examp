package mygr

import (
	"fmt"
	"time"
)

func f() {
	//time.Sleep(1 * time.Second)
	fmt.Println("Горутина запущена")

}

func Mygr() {
	go f()
	time.Sleep(2 * time.Second)
	fmt.Println("Основная функция")
}
