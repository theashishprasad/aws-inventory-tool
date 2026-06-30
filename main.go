package main

import (
	"fmt"
	"os"

	"github.com/theashishprasad/aws-inventory-tool/inventory"
	"github.com/theashishprasad/aws-inventory-tool/report"
)

func main() {
	args := os.Args
	if len(args) != 2 {
		fmt.Println("Usage: go run main.go <aws-region>")
		return
	}

	inventoryData, err := inventory.LoadInventory(args[1])
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	reportData := report.GenerateReport(inventoryData)
	fmt.Println(reportData)

	err = inventory.ExportInventory("inventory.json", inventoryData)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	fmt.Println("\nInventory exported to inventory.json")
}