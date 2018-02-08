package main

import "fmt"

type other struct{}

type something struct {
	*other
}

func main() {
	fmt.Println(something{&other{}})
}
