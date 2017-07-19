package main

import (
	d "./deterministics"
	"fmt"
)

func main() {
	err := d.DeriveKeys("password", "outdir/test.priv", "outdir/test.pub")
	if err != nil {
		panic(err)
	}
	fmt.Println("KeyGen Complete.")
}
