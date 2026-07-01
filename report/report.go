package report

import (
	"fmt"

	"github.com/theashishprasad/aws-inventory-tool/models"
)

// GenerateReport formats an AWS inventory report for console output.
func GenerateReport(i models.Inventory) string {
	return fmt.Sprintf(`AWS Inventory Report

Region        : %s
EC2 Instances : %d
S3 Buckets    : %d`,
		i.Region,
		i.InstanceCount,
		i.BucketCount,
	)
}
