package main

import (
	"fmt"
	"log"

	"github.com/BlinkTheThings/gohidapi/pkg/hidapi"
)

func main() {
	fmt.Printf("gohidapi test/example tool. Compiled with hidapi version %s.\n", hidapi.VersionStr())

	major, minor, patch := hidapi.Version()
	fmt.Printf("gohidapi test/example tool. Compiled with hidapi version %d.%d.%d.\n", major, minor, patch)

	ret := hidapi.Init()
	if ret != 0 {
		log.Fatal("Error initializing hidapi.")
	}

	defer hidapi.Exit()
}
