package main

import "fmt"

//Used for error checking and handling in the future
func errCheck(err error) {
	if err != nil {
		fmt.Println(err)
		panic(err)
	}
}
