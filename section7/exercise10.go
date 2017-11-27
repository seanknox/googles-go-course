package main

import "fmt"

func main() {
	fmt.Println(true && true)  // print "true"
	fmt.Println(true && false) // print "false"
	fmt.Println(true || true)  // print "true"
	fmt.Println(true || false) // print "false"
	fmt.Println(!true)         // print "true"

}
