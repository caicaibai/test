package main

import "fmt"

func main() {

	var res string

	res = "111111"

	fmt.Println(res + "\n")

	r := &res

	*r = "aaaaa"

	fmt.Println(res)

}
