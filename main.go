package main

import (
	"flag"
	"fmt"
)

func main() {
	var port int
	flag.IntVar(&port, "p", 8000, "specify port to use. defaults to 8000")
	flag.Parse()
	
	fmt.Printf("port = %d\n", port)
	fmt.Printf("other args: %+v\n", flag.Args())
}
