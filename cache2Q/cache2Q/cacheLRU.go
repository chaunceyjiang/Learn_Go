package cache2Q

import (
	"log"
	"sync"
	"time"
)

type CacheLRU struct {
	sync.RWMutex

	selfHead *CacheItem  //自身链表头部
	selfTail *CacheItem  //自身尾部
	fifoHead *CacheFIFO  //fifo链表头部
	name     string      //链表名称
	length   int64       //链表长度
	counter  int64       //计算器
	items    sync.Map    //map 用来快速访问
	logger   *log.Logger //日志

	loadData func(key interface{}, args ...interface{}) *CacheItem //访问数据时的回调函数

	addedItem func(item *CacheItem) //添加节点时的回调函数

	aboutToDeleteItem func(item *CacheItem) //从该链表删除时的回调函数

}

func (lru *CacheLRU) log(v ...interface{}) {
	if lru.logger == nil {
		return
	}
	lru.logger.Println(v)
}
func (lru *CacheLRU) SetLoadDataFunc(f func(ke interface{}, args ...interface{}) *CacheItem) {
	lru.Lock()
	defer lru.Unlock()

	lru.loadData = f
}

func (lru *CacheLRU) SetAddedItemFunc(f func(item *CacheItem)) {
	lru.Lock()
	defer lru.Unlock()

	lru.addedItem = f
}

func (lru *CacheLRU) SetAboutToDeleteItemFunc(f func(item *CacheItem)) {
	lru.Lock()
	defer lru.Unlock()

	lru.aboutToDeleteItem = f
}
func (lru *CacheLRU) expirationCheck() {
	lru.Lock()
	defer lru.Unlock()
	count := lru.counter
	lens := lru.length
	for lens < count {
		lru.log("链表超过指定长度,尾部上移", "Key", lru.selfHead.key, lru.selfHead.createdOn)
		aboutToExpire := lru.selfHead.aboutToExpire
		if aboutToExpire != nil {
			aboutToExpire(lru.selfHead.key)
		}
		aboutToDeleteItem := lru.aboutToDeleteItem
		if aboutToDeleteItem != nil {
			aboutToDeleteItem(lru.selfHead)
		}
		lru.selfTail = lru.selfTail.prev
		lru.counter--
		count = lru.counter
	}
}
func (lru *CacheLRU) addInternal(key interface{}, item *CacheItem) {
	lru.RLock()
	addedItem := lru.addedItem
	lru.RUnlock()
	if addedItem != nil {
		addedItem(item)
	}
	item.Lock()
	item.accessCount++
	item.accessedOn = time.Now()
	item.next = lru.selfHead
	item.Unlock()

	lru.Lock()
	lru.selfHead = item

	if lru.selfTail == nil {
		lru.selfTail = lru.selfHead
	}

	lru.counter++
	lru.Unlock()
	lru.items.Store(key, item)

	lru.log("成功添加到LRU Key", item.key, "Value", item.value, "created", item.createdOn)
	lru.expirationCheck()
}
func (lru *CacheLRU) Delete(key interface{}) {
	r, _ := lru.items.Load(key)
	item,ok := r.(*CacheItem)
	if !ok{
		return
	}
	lru.RLock()
	aboutToDeleteItem := lru.aboutToDeleteItem
	aboutToExpire := item.aboutToExpire
	lru.RUnlock()
	if aboutToExpire != nil {
		aboutToExpire(key)
	}
	if aboutToDeleteItem != nil {
		aboutToDeleteItem(item)
	}
	item.Lock()
	item.prev.next = item.next
	item.Unlock()
	lru.items.Delete(key)
}

func (lru *CacheLRU) Value(key interface{}, args ...interface{}) (item *CacheItem, err error) {
	r, ok := lru.items.Load(key)
	lru.RLock()
	loadData := lru.loadData
	lru.RUnlock()
	err = nil
	if ok {
		lru.log("key 存储在LRU中",key)
		item = r.(*CacheItem)
		item.KeepAlive()
		return
	}
	if loadData != nil {
		item=loadData(key,args...)
		if item!=nil{
			lru.Add(key,item.value)
		}
		return
	}
	item, err = lru.fifoHead.Value(key, args)
	return
}

func (lru *CacheLRU) Flush() {
	lru.Lock()
	defer lru.Unlock()
	lru.counter = 0
	var m sync.Map
	lru.items = m
	lru.selfHead = nil
	lru.selfTail = nil
	lru.fifoHead.Flush()
}
func (lru *CacheLRU)Add(key,value interface{})  {
	lru.fifoHead.Add(key,value)
}
func (lru *CacheLRU)GetFIFO() *CacheFIFO  {
	return lru.fifoHead
}