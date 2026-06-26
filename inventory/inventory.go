package inventory

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/ec2"
	"github.com/aws/aws-sdk-go-v2/service/s3"
	"github.com/theashishprasad/aws-inventory-tool/models"
)

func LoadInventory(region string) (models.Inventory, error) {
	var inventory models.Inventory

	ctx := context.Background()

	cfg, err := config.LoadDefaultConfig(ctx, config.WithRegion(region))
	if err != nil {
		return inventory, err
	}

	ec2Client := ec2.NewFromConfig(cfg)

	instances, err := ec2Client.DescribeInstances(ctx, &ec2.DescribeInstancesInput{})
	if err != nil {
		return inventory, err
	}

	instanceCount := 0

	for _, r := range instances.Reservations {
		instanceCount += len(r.Instances)
	}

	inventory.InstanceCount = instanceCount

	s3Client := s3.NewFromConfig(cfg)
	buckets, err := s3Client.ListBuckets(ctx, &s3.ListBucketsInput{})
	if err != nil {
		return inventory, err
	}

	inventory.BucketCount = len(buckets.Buckets)

	inventory.Region = cfg.Region

	return inventory, nil
}
