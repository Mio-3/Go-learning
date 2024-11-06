package main

import (
	"fmt"
	"reflect"
	"strconv"
)

func DateType() {
	i := 123
	f := 1.23
	v := "abc"
	b := true

	fmt.Println(i,reflect.TypeOf(i))
	fmt.Println(f,reflect.TypeOf(f))
	fmt.Println(v,reflect.TypeOf(v))
	fmt.Println(b,reflect.TypeOf(b))

	var f2 = float64(i)
	var i2 = int(f)

	fmt.Println(f2,reflect.TypeOf(f2))
	fmt.Println(i2,reflect.TypeOf(i2))


	var bool2 bool = i != 0 // true
	var bool3 bool = f != 0 // true
	fmt.Println(bool2, reflect.TypeOf(bool2))
	fmt.Println(bool3, reflect.TypeOf(bool3))

	s2 := strconv.FormatInt(int64(i), 10)
	s3 := strconv.FormatFloat(float64(f), 'E', -1, 64)
	fmt.Println(s2, reflect.TypeOf(s2))
	fmt.Println(s3, reflect.TypeOf(s3))

	s4 := strconv.FormatBool(b) // "true"
	fmt.Println(s4, reflect.TypeOf(s4))

	i3, err := strconv.ParseInt("123", 10, 0) // 123
	fmt.Println(i3, reflect.TypeOf(i3), err, reflect.TypeOf(err))

	f4, err := strconv.ParseFloat("1.23", 64) // 1.23
	fmt.Println(f4, reflect.TypeOf(f4), err, reflect.TypeOf(err))

	b4, err := strconv.ParseBool("true") // true
	fmt.Println(b4, reflect.TypeOf(b4), err, reflect.TypeOf(err))
}