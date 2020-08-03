package main

import (
	"fmt"

	"github.com/jaypipes/ghw"
)

func main() {
	pci, err := ghw.PCI()
	if err != nil {
		fmt.Printf("Error getting PCI info: %v", err)
	}
	fmt.Printf("host PCI devices:\n")
	fmt.Println("====================================================")
	devices := pci.ListDevices()
	if len(devices) == 0 {
		fmt.Printf("error: could not retrieve PCI devices\n")
		return
	}
	fmt.Printf("%v ", pci.GetDevice("0000:00:1f.6"))

	for _, device := range devices {
		vendor := device.Vendor
		vendorName := vendor.Name
		if len(vendor.Name) > 20 {
			vendorName = string([]byte(vendorName)[0:17]) + "..."
		}
		product := device.Product
		productName := product.Name
		if len(product.Name) > 40 {
			productName = string([]byte(productName)[0:37]) + "..."
		}
		fmt.Printf("%-12s\t%-20s\t%-40s\n", device.Address, vendorName, productName)
	}
}
