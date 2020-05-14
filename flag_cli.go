package main

import (
	"fmt"
	"github.com/jessevdk/go-flags"
)

var opts struct {
	Name    string `short:"n" long:"name" default:"World" description:"A name to say hello to."`
	Spanish bool   `short:"s" long:"spanish" description:"Use Spanish Language"`
}

func main() {
	_, err := flags.Parse(&opts)
	if err != nil {
		fmt.Println(err)
	}
	if opts.Spanish == true {
		fmt.Printf("Hola %s!\n", opts.Name)
	} else {
		fmt.Printf("Hello %s!\n", opts.Name)
	}
}
