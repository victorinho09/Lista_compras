package main

import "fmt"

func debug(s string, v ...interface{}) {
	line := fmt.Sprintf(s, v...)
	fmt.Println(line)
}

func debugErr(s string, v ...interface{}) {
	line := fmt.Sprintf(s, v...)
	fmt.Println("Error: ************* " + line)
}