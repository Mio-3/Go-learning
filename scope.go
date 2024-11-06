package main

import (
	"fmt"
)

func Scope() {
	// ①スコープ
	age := 10
	fmt.Println(age)

	if true {
		fmt.Println(age)
		age += 20
		fmt.Println(age)
	}

	fmt.Println(age)

	num := 0
	fmt.Println(num)
	i_slice := []int{1, 2, 3}

	for _, i := range i_slice {
		num += i
	}

	fmt.Println(num)

	// ②スコープ
	age2 := 10
	fmt.Println(age2)

	if true {
		age2 := 20
		fmt.Println(age2)

		age2 += 30
		fmt.Println(age2)
	}
	fmt.Println(age2)

	num2 := 0
	fmt.Println(num2)
	i_slice2 := []int{1, 2, 3}
	for _, i := range i_slice2 {
		num2 := 0
		num2 += i
	}
	fmt.Println(num2)
}