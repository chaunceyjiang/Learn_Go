package cache2Q

import (
	"sync"
	"log"
	"os"
)

var mutex sync.RWMutex
func NewCache(lruName string,LRUlength,FIFOlength int64) *CacheLRU {
	var m,n sync.Map
	var l *CacheLRU
	var f *CacheFIFO
	mutex.Lock()
	logger:=log.New(os.Stdout,"Cache2Q : ",log.LUTC|log.Ldate|log.Lshortfile)
	f=&CacheFIFO{
		name:lruName,
		selfHead:nil,
		selfTail:nil,
		length:FIFOlength,
		logger:logger,
		items:n,
		transCount:2,
	}
	l=&CacheLRU{
		name:lruName,
		selfTail:nil,
		selfHead:nil,
		length:LRUlength,
		items:m,
		fifoHead:f,
		logger:logger,
	}
	l.fifoHead.lruHead=l
	mutex.Unlock()
	return l
}