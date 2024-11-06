package main

import (
	"fmt"
)

func Lint() {
	fmt.Println(123)

	fmt.Println(1.23)

	fmt.Println("abc")

	fmt.Println(`abc
	def
	ghi
	`)
	
	fmt.Println(true)

	fmt.Println(nil)
}