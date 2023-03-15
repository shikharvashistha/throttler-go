package store

import (
	"github.com/shikharvashistha/throttler-go/pkg/store/keyvalue"
	"gorm.io/gorm"
)

type Store interface {
	KV() keyvalue.KV
}

func NewStore(db *gorm.DB) Store {
	return &store{
		kv: keyvalue.NewKVStore(),
	}
}
