package inventory

import (
	"encoding/json"
	"os"
	"testing"

	"github.com/theashishprasad/aws-inventory-tool/models"
)

func TestExportInventory(t *testing.T) {
	tests := []struct {
		name      string
		inventory models.Inventory
	}{
		{
			name: "Empty Inventory",
			inventory: models.Inventory{
				Region:        "ap-south-1",
				InstanceCount: 0,
				BucketCount:   0,
			},
		},
		{
			name: "Single Resources",
			inventory: models.Inventory{
				Region:        "us-east-1",
				InstanceCount: 1,
				BucketCount:   1,
			},
		},
		{
			name: "Multiple Resources",
			inventory: models.Inventory{
				Region:        "eu-west-1",
				InstanceCount: 12,
				BucketCount:   8,
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			fileName := "test_inventory.json"

			err := ExportInventory(fileName, tt.inventory)
			if err != nil {
				t.Fatalf("ExportInventory() error = %v", err)
			}

			defer os.Remove(fileName)

			data, err := os.ReadFile(fileName)
			if err != nil {
				t.Fatalf("ReadFile() error = %v", err)
			}

			var got models.Inventory

			err = json.Unmarshal(data, &got)
			if err != nil {
				t.Fatalf("Unmarshal() error = %v", err)
			}

			if got != tt.inventory {
				t.Errorf("got %+v, want %+v", got, tt.inventory)
			}
		})
	}
}