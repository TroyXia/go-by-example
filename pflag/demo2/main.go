package main

import (
	"fmt"
	"github.com/spf13/pflag"
	"os"
)

func main() {
	flagset := pflag.NewFlagSet("generic", pflag.ExitOnError)
	var name = flagset.StringP("name", "n", "hxia", "name for application")
	var age = flagset.IntP("age", "a", 18, "age for application")

	flagset.Lookup("age").NoOptDefVal = "18"

	flagset.SortFlags = true
	flagset.Parse(os.Args[1:])

	a, err := flagset.GetInt("age")
	if err != nil {
		fmt.Printf("failed to get age: %v\n", err)
	}

	fmt.Println(*name, *age, a)
}
