package main

import (
	"fmt"
	"io/ioutil"
	"os"
)

func main() {

	read, err := ioutil.ReadFile(os.Args[1])

	if err != nil {
		fmt.Println("Cannot read file: ", err)
		os.Exit(1)
	}
	fmt.Println(string(read))
}
