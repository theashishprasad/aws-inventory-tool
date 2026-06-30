package report

import (
	"testing"

	"github.com/theashishprasad/aws-inventory-tool/models"
)

func TestGenerateReport(t *testing.T) {
	tests := []struct {
		name      string
		inventory models.Inventory
		want      string
	}{
		{
			name: "Empty Inventory",
			inventory: models.Inventory{
				Region:        "ap-south-1",
				InstanceCount: 0,
				BucketCount:   0,
			},
			want: `AWS Inventory Report

Region        : ap-south-1
EC2 Instances : 0
S3 Buckets    : 0`,
		},
		{
			name: "Single Resources",
			inventory: models.Inventory{
				Region:        "us-east-1",
				InstanceCount: 1,
				BucketCount:   1,
			},
			want: `AWS Inventory Report

Region        : us-east-1
EC2 Instances : 1
S3 Buckets    : 1`,
		},
		{
			name: "Multiple Resources",
			inventory: models.Inventory{
				Region:        "eu-west-1",
				InstanceCount: 12,
				BucketCount:   8,
			},
			want: `AWS Inventory Report

Region        : eu-west-1
EC2 Instances : 12
S3 Buckets    : 8`,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := GenerateReport(tt.inventory)

			if got != tt.want {
				t.Errorf("got:\n%s\n\nwant:\n%s", got, tt.want)
			}
		})
	}
}