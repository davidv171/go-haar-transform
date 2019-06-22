package main

import "fmt"

//Used for error checking and handling in the future
func ErrCheck(err error) {
	if err != nil {
		fmt.Println(err)
		panic(err)
	}
}
