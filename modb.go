package modb

import (
	"fmt"
	"runtime/debug"

	"github.com/manelmontilla/modc"
)

// CallMeMaybe It's hard to look right at you baby, but here's my number, so
// call me, maybe? Hey, I just met you, and this is crazy.
func CallMeMaybe() string {
	modcV := modc.CallMeMaybe()
	fmt.Printf("got modc version: %s\n", modcV)
	v, _ := debug.ReadBuildInfo()
	return v.Main.Version
}
