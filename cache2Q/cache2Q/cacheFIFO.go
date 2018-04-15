package cache2Q

import (
	"log"
	"sync"
)

type CacheFIFO struct {
	sync.RWMutex

	selfHead   *CacheItem  //自身链表头部
	selfTail   *CacheItem  //自身尾部
	lruHead    *CacheLRU   //lru链表头部
	name       string      //链表名称
	length     int64       //链表长度
	counter    int64       //计算器
	items      sync.Map    //map 用来快速访问
	logger     *log.Logger //日志
	transCount int64
	loadData   func(key interface{}, args ...interface{}) *CacheItem //访问数据时的回调函数

	addedItem func(item *CacheItem) //添加节点时的回调函数

	aboutToDeleteItem func(item *CacheItem) //从该链表删除时的回调函数

	fifoTransferToLRU func(item *CacheItem)
}

func (fifo *CacheFIFO) log(v ...interface{}) {
	if fifo.logger == nil {
		return
	}
	fifo.logger.Println(v)
}
func (fifo *CacheFIFO) SetLogger(logger *log.Logger) {
	fifo.Lock()
	defer fifo.Unlock()

	fifo.logger = logger
}

//若没有找到，则添加
func (fifo *CacheFIFO) SetLoadDataFunc(f func(ke interface{}, args ...interface{}) *CacheItem) {
	fifo.Lock()
	defer fifo.Unlock()

	fifo.loadData = f
}
func (fifo *CacheFIFO) SetAddedItemFunc(f func(item *CacheItem)) {
	fifo.Lock()
	defer fifo.Unlock()

	fifo.addedItem = f
}

func (fifo *CacheFIFO) SetAboutToDeleteItemFunc(f func(item *CacheItem)) {
	fifo.Lock()
	defer fifo.Unlock()

	fifo.aboutToDeleteItem = f
}
func (fifo *CacheFIFO) SetFifoTransferToLRU(f func(item *CacheItem)) {
	fifo.Lock()
	defer fifo.Unlock()

	fifo.fifoTransferToLRU = f
}
func (fifo *CacheFIFO) Length() int64 {
	return fifo.length
}
func (fifo *CacheFIFO) Count() int64 {
	fifo.RLock()
	defer fifo.RUnlock()

	return fifo.counter
}

func (fifo *CacheFIFO) Foreach(f func(k, v interface{}) bool) {
	fifo.items.Range(f)
}

func (fifo *CacheFIFO) expirationCheck() {
	fifo.Lock()
	defer fifo.Unlock()
	count := fifo.counter
	lens := fifo.length
	for lens < count {
		fifo.log("链表超过指定长度,头部下移", "Key", fifo.selfHead.key, fifo.selfHead.createdOn)
		aboutToExpire := fifo.selfHead.aboutToExpire
		if aboutToExpire != nil {
			aboutToExpire(fifo.selfHead.key)
		}
		fifo.selfHead = fifo.selfHead.next
		fifo.counter--
		count = fifo.counter
	}
}

func (fifo *CacheFIFO) Add(key, value interface{}) {
	fifo.Lock()
	item := NewCacheItem(key, value)

	addedItem := fifo.addedItem
	if addedItem != nil {
		addedItem(item)
	}
	item.prev = fifo.selfTail
	if fifo.selfTail!=nil{
		fifo.selfTail.next = item
	}
	fifo.selfTail = item
	if fifo.selfHead==nil{
		fifo.selfHead=fifo.selfTail
	}
	fifo.counter++
	fifo.Unlock()
	fifo.items.Store(key, item)
	fifo.log("成功添加到", fifo.name, " FIFO, Key", item.key, "Value", item.value, "created", item.createdOn)
	fifo.expirationCheck()

}

func (fifo *CacheFIFO) Value(key interface{}, args ...interface{}) (*CacheItem, error) {
	r, _ := fifo.items.Load(key)
	fifo.RLock()
	loadData := fifo.loadData
	item,ok:= r.(*CacheItem)
	if !ok{
		fifo.RUnlock()
		return nil,ErrKeyNotFound
	}
	count := item.accessCount
	fifo.RUnlock()

	if ok { //如果有，就添加到LRU队列中,并且从FIFO队列中删除
		if count >= fifo.transCount {
			fifo.Lock()
			fifo.counter--
			fifo.Unlock()
			item.Lock()
			//从FIFO队列中删除
			if item.prev!=nil{
				item.prev.next = item.next
			}
			item.Unlock()

			fifo.items.Delete(key)
			//添加到LRU队列中

			if fifo.fifoTransferToLRU != nil {
				fifo.fifoTransferToLRU(item)
			}
			fifo.lruHead.addInternal(key, item)
			//上面有问题,没有对LRU队列加锁?
			fifo.log("Key:", item.key, "从 ", fifo.name, "迁移到 ", fifo.lruHead.name, "LRU 队列里！")
		}
		item.KeepAlive()
		return item, nil
	}
	if loadData != nil {
		item = loadData(key, args...) //自定义添加
		if item != nil {
			fifo.Add(key, item.value)
			return item, nil
		}
		return nil, ErrKeyNotFoundAndLoadable
	}
	return nil, ErrKeyNotFound
}
func (fifo *CacheFIFO) Delete(key interface{}) {
	r, _ := fifo.items.Load(key)
	item := r.(*CacheItem)
	fifo.Lock()
	aboutToDeleteItem := fifo.aboutToDeleteItem
	aboutToExpire := item.aboutToExpire
	if aboutToExpire != nil {
		aboutToExpire(key)
	}
	if aboutToDeleteItem != nil {
		aboutToDeleteItem(item)
	}
	item.prev.next = item.next
	fifo.Unlock()
	fifo.items.Delete(key)
}
func (fifo *CacheFIFO) Flush() {
	fifo.Lock()
	defer fifo.Unlock()
	fifo.log("清空FIFO队列....", fifo.name)
	var m sync.Map
	fifo.items = m
	fifo.selfTail = nil
	fifo.selfHead = nil
	fifo.counter = 0
}
