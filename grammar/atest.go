package main

import "fmt"

func main() {
	ST()
}
func ST() {
	var i *int
	i = new(int)
	fmt.Println(&i, i)

	*i = 1
	fmt.Println(&i, i, *i)
}
