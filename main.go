package main

import (
	"fmt"

	"github.com/theashishprasad/aws-inventory-tool/inventory"
)

func main() {
	inventoryData, err := inventory.LoadInventory("sample/inventory.json")
	if err != nil {
		fmt.Println("Error: ", err)
		return
	}

	fmt.Println("AWS Inventory Report")
	fmt.Println()

	fmt.Printf("Instances : %d\n", inventoryData.InstanceCount)
	fmt.Printf("Region    : %s\n", inventoryData.Region)
}
