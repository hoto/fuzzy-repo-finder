package main

import (
	"fmt"
	flag "github.com/ogier/pflag"
)

var (
	user string
)

func main() {
	fmt.Println("Hello, World")
	flag.Parse()
	fmt.Println(user)
}

func init() {
	flag.StringVarP(&user, "user", "u", "", "Search Users")
}
