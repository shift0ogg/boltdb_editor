package models

// BucketInfo represents a bucket in BoltDB
type BucketInfo struct {
	Name string `json:"name"`
}

// KeyValue represents a key-value pair
type KeyValue struct {
	Key   string `json:"key"`
	Value string `json:"value"`
}

// DatabaseInfo represents database metadata
type DatabaseInfo struct {
	Path    string       `json:"path"`
	Buckets []BucketInfo `json:"buckets"`
}
