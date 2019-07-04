// flag.go
package main

import(
	"flag"
	"fmt"
)

func main(){
	var name string

	flag.StringVar(&name, "opt", "","Usage")

	flag.Parse()

	if name == "" {
		name = "Hello World"
	}else{
		name = "hello " + name
	}

	fmt.Println(name)
}