package modb

import (
	"runtime/debug"
)

const modName = "github.com/manelmontilla/modb"

// CallMeMaybe It's hard to look right at you baby, but here's my number, so
// call me, maybe? Hey, I just met you, and this is crazy.
func CallMeMaybe() string {
	return myVersion()
}

func myVersion() string {
	v, _ := debug.ReadBuildInfo()
	version := "unknown"
	for _, m := range v.Deps {
		if m.Path == modName {
			version = m.Version
			break
		}
	}
	return version
}
