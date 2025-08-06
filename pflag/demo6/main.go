package main

import (
	"fmt"
	"github.com/spf13/pflag"
)

func main() {
	var (
		name    string
		age     int
		debug   bool
		timeout int
	)

	pflag.StringVar(&name, "name", "demo", "name of the person")
	pflag.BoolVar(&debug, "debug", false, "debug mode")
	pflag.IntVar(&age, "age", 0, "age value")
	pflag.IntVar(&timeout, "timeout", 0, "timeout value")

	pflag.Parse()

	// 位置很重要
	pflag.VisitAll(func(f *pflag.Flag) {
		fmt.Printf("%s: %v\n", f.Name, f.Value)
		if f.Changed {
			fmt.Printf("%s changed: %v\n", f.Name, f.Value)
		}
	})
}
