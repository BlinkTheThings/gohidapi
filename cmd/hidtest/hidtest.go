package main

import (
	"fmt"

	"github.com/BlinkTheThings/gohidapi/pkg/hidapi"
)

func main() {
	fmt.Printf("gohidapi test/example tool. Compiled with hidapi version %s.\n", hidapi.VersionStr())

	major, minor, patch := hidapi.Version()
	fmt.Printf("gohidapi test/example tool. Compiled with hidapi version %d.%d.%d.\n", major, minor, patch)
}
