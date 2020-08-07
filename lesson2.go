package main

import "fmt"

func main() {

	var res string

	res = "111111"

	fmt.Println(res + "\n")

	r := &res

	*r = "aaaaa"

	var a int

	fmt.Println(res)

	fmt.Println(a)

}
