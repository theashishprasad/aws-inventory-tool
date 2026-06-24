package inventory

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/ec2"
	"github.com/theashishprasad/aws-inventory-tool/models"
)

func LoadInventory() (models.Inventory, error) {
	var inventory models.Inventory

	ctx := context.Background()

	cfg, err := config.LoadDefaultConfig(ctx)
	if err != nil {
		return inventory, err
	}

	client := ec2.NewFromConfig(cfg)

	instances, err := client.DescribeInstances(ctx, &ec2.DescribeInstancesInput{})
	if err != nil {
		return inventory, err
	}

	count := 0

	for _, r := range instances.Reservations {
		count += len(r.Instances)
	}

	inventory.InstanceCount = count

	return inventory, nil
}
