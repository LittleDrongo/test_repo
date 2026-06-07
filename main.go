package main

import (
	"fmt"

	"github.com/LittleDrongo/fmn-lib/console/cmd"
)

func main() {
	username := cmd.Input("Введите имя")
	fmt.Println("Hello" + username)
	fmt.Println("dsadasa")
}
