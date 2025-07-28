package main

import (
	"fmt"
	//"time"
	"flag"
	"github.com/Norrun/CLI-Habit-Thingemabob/src/core"
	"os"
)

func main() {
	newCmd := flag.NewFlagSet("new", flag.ExitOnError)
	argL := len(os.Args)
	namef := newCmd.String("n", "some-hsbit", "sets the name of the habit")
	descf := newCmd.String("d", "something todo", "sets the habit description")
	_ = argL
	switch os.Args[1] {
	case "new":
		newCmd.Parse(os.Args[2:])

		str, err := core.New(*namef, *descf)
		if core.Exists(err) {
			fmt.Fprintf(os.Stderr, "%e", err)
			os.Exit(err.Code())
		}
		fmt.Println(str)
	case "check":
		res, err := core.Check(os.Args[2], false)
		if core.Exists(err) {
			fmt.Fprintf(os.Stderr, "%e", err)
		}
		fmt.Println(res)
	case "uncheck":
		res, err := core.Check(os.Args[2], true)
		if core.Exists(err) {
			fmt.Fprintf(os.Stderr, "%e", err)
		}
		fmt.Println(res)

	}
}
