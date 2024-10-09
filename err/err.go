package err

import (
	"errors"
	"fmt"
)

func get_err() error {
	return errors.New(" Some Error Occurred")
}

func MainFunc() {
	if e := get_err(); e != nil {
		fmt.Println(e)
	}
}
