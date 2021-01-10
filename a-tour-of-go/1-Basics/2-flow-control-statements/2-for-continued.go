package main

import "fmt"

func main() {
	sum := 1
	for ; sum < 1000; { // the init and post statements are optional
		sum += sum
	}
	fmt.Println(sum)
}
