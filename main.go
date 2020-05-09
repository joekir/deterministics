package main

import (
	"fmt"
	d "github.com/joekir/deterministics/lib"
)

func main() {
	err := d.DeriveKeys("password", "outdir/test.priv", "outdir/test.pub")
	if err != nil {
		panic(err)
	}
	fmt.Println("KeyGen Complete.")
}
