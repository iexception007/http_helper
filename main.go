package main

import "fmt"

func main() {


	fmt.Printf("%d:%s\n", 100, HttpCode[100])
	fmt.Printf("%d:%s\n", 200, HttpCode[200])
	fmt.Printf("%d:%s\n", 300, HttpCode[300])
	fmt.Printf("%d:%s\n", 500, HttpCode[500])

	for k, v := range HttpCode {
		fmt.Printf("%d:%s\n", k ,v)
	}
}