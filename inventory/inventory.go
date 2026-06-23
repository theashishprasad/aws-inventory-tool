package inventory

import (
	"encoding/json"
	"os"

	"github.com/theashishprasad/aws-inventory-tool/models"
)

func LoadInventory(file string) (models.Inventory, error) {
	var inventory models.Inventory

	data, err := os.ReadFile(file)
	if err != nil {
		return inventory, err
	}

	err = json.Unmarshal(data, &inventory)
	if err != nil {
		return inventory, err
	}

	return inventory, nil
}
