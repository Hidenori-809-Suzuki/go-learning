package main

import (
	"fmt"
	// "os/user"
	// "time"
)

func foo(){
	xi := 1
	fmt.Println(xi)
}

func main() {
	var i int = 1
	fmt.Println("Hello World", i, /*time.Now()*/)
	// fmt.Println(user.Current())
	foo()
}
