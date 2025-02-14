package model

import (
	"database/sql/driver"
	"fmt"
)

type StorageType uint8

const (
	StorageTypeLocal StorageType = iota
	StorageTypeS3
)

func (s *StorageType) Value() (driver.Value, error) {
	if s == nil {
		return 0, nil
	}
	return uint8(*s), nil
}

func (s *StorageType) Scan(value any) error {
	if value == nil {
		return nil
	}
	if v, ok := value.(uint8); ok {
		*s = StorageType(v)
		return nil
	}
	return fmt.Errorf("sql/driver: unsupported value %v (type %T) converting to StorageType", value, value)
}
