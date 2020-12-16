package main

import (
	"fmt"
	"log"

	"github.com/BlinkTheThings/gohidapi/pkg/hidapi"
)

func main() {
	major, minor, patch := hidapi.Version()
	fmt.Printf("gohidapi test/example tool. Compiled with hidapi version %d.%d.%d.\n", major, minor, patch)

	ret := hidapi.Init()
	if ret != 0 {
		log.Fatal("Error initializing hidapi.")
	}

	defer hidapi.Exit()

	devices := hidapi.Enumerate(0, 0)
	for _, dev := range devices {
		fmt.Printf("Device found\n  type: %04x %04x\n  path: %s\n", dev.VendorID, dev.ProductID, dev.Path)
		fmt.Printf("  Release:      %x\n", dev.Release)
		fmt.Printf("  Interface:    %d\n", dev.Interface)
		fmt.Printf("  Usage (page): 0x%x (0x%x)\n", dev.Usage, dev.UsagePage)
		fmt.Println()
	}
}
