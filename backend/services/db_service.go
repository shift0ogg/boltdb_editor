package services

import (
	"boltdb-editor/backend/models"
	"fmt"

	"github.com/asdine/storm/v3"
	"go.etcd.io/bbolt"
)

// DBService manages BoltDB operations
type DBService struct {
	db *storm.DB
}

// NewDBService creates a new DBService
func NewDBService() *DBService {
	return &DBService{}
}

// OpenDatabase opens a BoltDB database
func (s *DBService) OpenDatabase(path string) error {
	if s.db != nil {
		s.db.Close()
	}
	db, err := storm.Open(path)
	if err != nil {
		return err
	}
	s.db = db
	return nil
}

// CloseDatabase closes the current database
func (s *DBService) CloseDatabase() error {
	if s.db != nil {
		return s.db.Close()
	}
	return nil
}

// ListBuckets lists all buckets
func (s *DBService) ListBuckets() ([]models.BucketInfo, error) {
	if s.db == nil {
		return nil, fmt.Errorf("no database opened")
	}
	var buckets []models.BucketInfo
	err := s.db.Bolt.View(func(tx *bbolt.Tx) error {
		return tx.ForEach(func(name []byte, _ *bbolt.Bucket) error {
			buckets = append(buckets, models.BucketInfo{Name: string(name)})
			return nil
		})
	})
	return buckets, err
}

// CreateBucket creates a new bucket
func (s *DBService) CreateBucket(name string) error {
	if s.db == nil {
		return fmt.Errorf("no database opened")
	}
	return s.db.Bolt.Update(func(tx *bbolt.Tx) error {
		_, err := tx.CreateBucketIfNotExists([]byte(name))
		return err
	})
}

// DeleteBucket deletes a bucket
func (s *DBService) DeleteBucket(name string) error {
	if s.db == nil {
		return fmt.Errorf("no database opened")
	}
	return s.db.Bolt.Update(func(tx *bbolt.Tx) error {
		return tx.DeleteBucket([]byte(name))
	})
}

// ListKeys lists all keys in a bucket
func (s *DBService) ListKeys(bucket string) ([]models.KeyValue, error) {
	if s.db == nil {
		return nil, fmt.Errorf("no database opened")
	}
	var keys []models.KeyValue
	err := s.db.Bolt.View(func(tx *bbolt.Tx) error {
		b := tx.Bucket([]byte(bucket))
		if b == nil {
			return fmt.Errorf("bucket not found")
		}
		return b.ForEach(func(k, v []byte) error {
			keys = append(keys, models.KeyValue{Key: string(k), Value: string(v)})
			return nil
		})
	})
	return keys, err
}

// GetValue gets a value by key
func (s *DBService) GetValue(bucket, key string) (string, error) {
	if s.db == nil {
		return "", fmt.Errorf("no database opened")
	}
	var value []byte
	err := s.db.Bolt.View(func(tx *bbolt.Tx) error {
		b := tx.Bucket([]byte(bucket))
		if b == nil {
			return fmt.Errorf("bucket not found")
		}
		value = b.Get([]byte(key))
		return nil
	})
	if err != nil {
		return "", err
	}
	if value == nil {
		return "", fmt.Errorf("key not found")
	}
	return string(value), nil
}

// SetValue sets a value for a key
func (s *DBService) SetValue(bucket, key, value string) error {
	if s.db == nil {
		return fmt.Errorf("no database opened")
	}
	return s.db.Bolt.Update(func(tx *bbolt.Tx) error {
		b := tx.Bucket([]byte(bucket))
		if b == nil {
			return fmt.Errorf("bucket not found")
		}
		return b.Put([]byte(key), []byte(value))
	})
}

// DeleteKey deletes a key
func (s *DBService) DeleteKey(bucket, key string) error {
	if s.db == nil {
		return fmt.Errorf("no database opened")
	}
	return s.db.Bolt.Update(func(tx *bbolt.Tx) error {
		b := tx.Bucket([]byte(bucket))
		if b == nil {
			return fmt.Errorf("bucket not found")
		}
		return b.Delete([]byte(key))
	})
}
