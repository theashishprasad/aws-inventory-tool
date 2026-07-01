package models

// Inventory represents AWS infrastructure inventory.
type Inventory struct {
	Region        string `json:"region"`
	InstanceCount int    `json:"instances"`
	BucketCount   int    `json:"buckets"`
}
