package main

import (
	"fmt"
	"os"

	"github.com/theashishprasad/aws-inventory-tool/inventory"
)

func main() {
	args := os.Args
	if len(args) != 2 {
		fmt.Println("Usage: go run main.go <aws-region>")
		return
	}

	inventoryData, err := inventory.LoadInventory(args[1])
	if err != nil {
		fmt.Println("Error: ", err)
		return
	}

	fmt.Println("AWS Inventory Report")
	fmt.Println()

	fmt.Printf("Region        : %s\n", inventoryData.Region)
	fmt.Printf("EC2 Instances : %d\n", inventoryData.InstanceCount)
	fmt.Printf("S3 Buckets    : %d\n", inventoryData.BucketCount)
}
