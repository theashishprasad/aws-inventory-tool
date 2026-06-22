package models

type Inventory struct {
	InstanceCount int    `json:"instance_count"`
	Region        string `json:"region"`
}
