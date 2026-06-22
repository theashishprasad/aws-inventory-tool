package main

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/theashishprasad/aws-inventory-tool/models"
)

func main() {
	data, err := os.ReadFile("sample/inventory.json")
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	var inventory models.Inventory

	err = json.Unmarshal(data, &inventory)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	fmt.Println("AWS Inventory Report")
	fmt.Println()

	fmt.Printf("Instances : %d\n", inventory.InstanceCount)
	fmt.Printf("Region    : %s\n", inventory.Region)
}
