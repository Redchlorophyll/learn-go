package main

import (
	"flag"
	"fmt"
)

func main() {
	var user *string = flag.String("host", "root", "put your host database here...")
	var pass *string = flag.String("password", "root", "put your password here")

	flag.Parse()

	fmt.Println(user, pass)
}
