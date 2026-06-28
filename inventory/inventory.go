package inventory

import (
	"context"
	"encoding/json"
	"os"
	"sync"

	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/ec2"
	"github.com/aws/aws-sdk-go-v2/service/s3"
	"github.com/theashishprasad/aws-inventory-tool/models"
)

func LoadInventory(region string) (models.Inventory, error) {
	var wg sync.WaitGroup
	var inventory models.Inventory
	var ec2Err error
	var s3Err error

	ctx := context.Background()

	cfg, err := config.LoadDefaultConfig(ctx, config.WithRegion(region))
	if err != nil {
		return inventory, err
	}

	wg.Add(1)

	go func() {
		defer wg.Done()

		ec2Client := ec2.NewFromConfig(cfg)

		instances, err := ec2Client.DescribeInstances(ctx, &ec2.DescribeInstancesInput{})
		if err != nil {
			ec2Err = err
			return
		}

		instanceCount := 0

		for _, r := range instances.Reservations {
			instanceCount += len(r.Instances)
		}

		inventory.InstanceCount = instanceCount
	}()

	wg.Add(1)

	go func() {
		defer wg.Done()

		s3Client := s3.NewFromConfig(cfg)
		buckets, err := s3Client.ListBuckets(ctx, &s3.ListBucketsInput{})
		if err != nil {
			s3Err = err
			return
		}

		inventory.BucketCount = len(buckets.Buckets)
	}()

	wg.Wait()

	if ec2Err != nil {
		return inventory, ec2Err
	}

	if s3Err != nil {
		return inventory, s3Err
	}

	inventory.Region = cfg.Region

	return inventory, nil
}

func ExportInventory(fileName string, inventory models.Inventory) error {
	jsonData, err := json.MarshalIndent(inventory, "", "	")
	if err != nil {
		return err
	}

	err = os.WriteFile(fileName, jsonData, 0644)
	if err != nil {
		return err
	}

	return nil
}
