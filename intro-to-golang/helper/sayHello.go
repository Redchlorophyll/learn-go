package helper

import "fmt"

func SayHello(name string) {
	fmt.Println("hello", name)
}

var Version = "1.0.0" // can be imported
var tmpVersion = "2.0.0" // can't be exported