package cache2Q

import (
	"sync"
	"time"
)

type CacheItem struct {
	sync.RWMutex

	key           interface{}
	value         interface{}
	accessedOn    time.Time             //最后一次被访问的时间
	createdOn     time.Time             //被创建的时间
	aboutToExpire func(key interface{}) //被删除时的回调函数
	accessCount   int64
	next          *CacheItem
	prev          *CacheItem
}

func (item *CacheItem) Key() interface{} {
	item.RLock()
	defer item.RUnlock()
	return item.key
}

func (item *CacheItem) Value() interface{} {
	item.RLock()
	defer item.RUnlock()
	return item.value
}

func (item *CacheItem) SetAboutToExpireFunc(f func(key interface{})) {
	item.Lock()
	defer item.Unlock()
	item.aboutToExpire = f
}

func (item *CacheItem) CreateedTime() time.Time {
	return item.createdOn //不会改变,所以不用上锁
}

func (item *CacheItem) AccessedTime() time.Time {
	item.RLock()
	defer item.RUnlock()
	return item.accessedOn
}
func (item *CacheItem) AccessCount() int64 {
	item.RLock()
	defer item.RUnlock()
	return item.accessCount
}
func (item *CacheItem) KeepAlive() {
	item.Lock()
	defer item.Unlock()

	item.accessedOn = time.Now()
	item.accessCount++
}
func NewCacheItem(key, value interface{}) *CacheItem { //产生一个新节点
	t := time.Now()
	return &CacheItem{
		key:           key,
		value:         value,
		createdOn:     t,
		accessedOn:    t,
		aboutToExpire: nil,
		next:          nil,
		prev:          nil,
	}
}
