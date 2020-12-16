package main

import (
	"fmt"

	"github.com/BlinkTheThings/gohidapi/pkg/hidapi"
)

func main() {
	fmt.Printf("gohidapi test/example tool. Compiled with hidapi version %s.\n", hidapi.VersionStr())
}
