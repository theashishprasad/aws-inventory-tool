package main

import (
	"fmt"

	"github.com/theashishprasad/aws-inventory-tool/inventory"
)

func main() {
	inventoryData, err := inventory.LoadInventory()
	if err != nil {
		fmt.Println("Error: ", err)
		return
	}

	fmt.Println("AWS Inventory Report")
	fmt.Println()

	fmt.Printf("EC2 Instances : %d\n", inventoryData.InstanceCount)
	fmt.Printf("S3 Buckets    : %d\n", inventoryData.BucketCount)
}
