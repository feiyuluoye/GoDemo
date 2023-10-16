package main

import (
	"fmt"
	"godemo.com/practices"
)

func message(s string) {
	fmt.Printf("============%s============", s)
}

func main() {

	practices.VarMain()
	practices.PointerMain()
	message("string main")
	practices.StringMain()

}
