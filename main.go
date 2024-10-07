package main

import (
	"fmt"
	"main/mychan"
	"main/mygr"
)

func main() {
	fmt.Println("Test Go apps!")
	mygr.Mygr()
	mychan.Run()

}
