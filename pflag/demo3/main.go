package main

import (
	"fmt"
	"github.com/spf13/pflag"
	"os"
	"strings"
)

func normalizeFunc(fs *pflag.FlagSet, name string) pflag.NormalizedName {
	// --my-flag == --my_flag == --my.flag
	from := []string{".", "_"}
	to := "-"

	for _, step := range from {
		name = strings.Replace(name, step, to, -1)
	}

	return pflag.NormalizedName(name)
}

func main() {
	flagset := pflag.NewFlagSet("generic", pflag.ExitOnError)
	var name = flagset.StringP("name", "n", "hxia", "name for application")
	var age = flagset.IntP("age", "a", 18, "age for application")
	var newName = flagset.String("new-name", "huyun", "new name for application")

	flagset.Lookup("age").NoOptDefVal = "18"

	flagset.SortFlags = true
	flagset.SetNormalizeFunc(normalizeFunc)

	flagset.Parse(os.Args[1:])

	a, err := flagset.GetInt("age")
	if err != nil {
		fmt.Printf("failed to get age: %v\n", err)
	}

	fmt.Println(*name, *age, a, *newName)
}
