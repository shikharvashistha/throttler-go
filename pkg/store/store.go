package store

import "github.com/shikharvashistha/throttler-go/pkg/store/keyvalue"

type store struct {
	kv keyvalue.KV
}

func (s *store) KV() keyvalue.KV {
	return s.kv
}
