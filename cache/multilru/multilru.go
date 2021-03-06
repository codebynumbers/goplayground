package multilru

import (
	"github.com/majek/goplayground/cache/lrucache"
	"hash"
	"hash/crc32"
	"time"
)

type MultiLRUCache struct {
	buckets uint
	cache   []*lrucache.LRUCache
	hash    hash.Hash
}


// Using this constructor is almost always wrong. Use NewMultiLRUCache instead.
func (m *MultiLRUCache) Init(buckets, bucket_capacity uint) {
	m.buckets = buckets
	m.cache = make([]*lrucache.LRUCache, buckets)
	for i := uint(0); i < buckets; i++ {
		m.cache[i] = lrucache.NewLRUCache(bucket_capacity)
	}
}

func NewMultiLRUCache(buckets, bucket_capacity uint) *MultiLRUCache {
	m := &MultiLRUCache{}
	m.Init(buckets, bucket_capacity)
	return m
}

func (m *MultiLRUCache) bucketNo(key string) uint {
	// Arbitrary choice. Any fast hash will do.
	return uint(crc32.ChecksumIEEE([]byte(key))) % m.buckets
}

func (m *MultiLRUCache) Set(key string, value interface{}, expire time.Time) {
	m.cache[m.bucketNo(key)].Set(key, value, expire)
}

func (m *MultiLRUCache) SetNow(key string, value interface{}, expire time.Time, now time.Time) {
	m.cache[m.bucketNo(key)].SetNow(key, value, expire, now)
}

func (m *MultiLRUCache) Get(key string) (value interface{}, ok bool) {
	return m.cache[m.bucketNo(key)].Get(key)
}

func (m *MultiLRUCache) GetQuiet(key string) (value interface{}, ok bool) {
	return m.cache[m.bucketNo(key)].Get(key)
}

func (m *MultiLRUCache) GetNotStale(key string) (value interface{}, ok bool) {
	return m.cache[m.bucketNo(key)].GetNotStale(key)
}

func (m *MultiLRUCache) GetNotStaleNow(key string, now time.Time) (value interface{}, ok bool) {
	return m.cache[m.bucketNo(key)].GetNotStaleNow(key, now)
}

func (m *MultiLRUCache) Del(key string) (value interface{}, ok bool) {
	return m.cache[m.bucketNo(key)].Del(key)
}

func (m *MultiLRUCache) Clear() int {
	var s int
	for _, c := range m.cache {
		s += c.Clear()
	}
	return s
}

func (m *MultiLRUCache) Len() int {
	var s int
	for _, c := range m.cache {
		s += c.Len()
	}
	return s
}

func (m *MultiLRUCache) Capacity() int {
	var s int
	for _, c := range m.cache {
		s += c.Capacity()
	}
	return s
}

func (m *MultiLRUCache) Expire() int {
	var s int
	for _, c := range m.cache {
		s += c.Expire()
	}
	return s
}

func (m *MultiLRUCache) ExpireNow(now time.Time) int {
	var s int
	for _, c := range m.cache {
		s += c.ExpireNow(now)
	}
	return s
}
