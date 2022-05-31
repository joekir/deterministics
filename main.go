package main

import (
	"fmt"
	"os"

	d "github.com/joekir/deterministics/lib"
)

func main() {
	args := os.Args[1:]
	if len(args) > 1 {
		fmt.Println("Only accepts single argument for password variable, otherwise defaults to 'hunter2' if no argument passed.")
		os.Exit(1)
	}

	var password string
	if len(args) == 1 {
		password = args[0]
	} else {
		password = "hunter2"
	}

	err := d.DeriveKeys(password, "outdir/test.priv", "outdir/test.pub")
	if err != nil {
		panic(err)
	}
	fmt.Println("KeyGen Complete.")
}
