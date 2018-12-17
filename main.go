package main

import (
	"fmt"
	"strings"

	flag "github.com/ogier/pflag"
)

var (
	user string
)

func main() {
	flag.Parse()
	users := strings.Split(user, ",")
	fmt.Printf("Searching users: %s\n", users)
	result := getUsers(user)
	fmt.Println(`Name: `, result.Name)
	fmt.Println(`Location: `, result.Location)
}

func init() {
	flag.StringVarP(&user, "user", "u", "", "Search Users")
}
