package leetcode

import (
	"sort"
	"time"
)

type LFUCache struct {
	Size     int
	Capacity int
	Cache    map[int]*CacheNode
}

type CacheNode struct {
	Key      int
	Value    int
	LastTime int64
	HitCount int
}

func Constructor(capacity int) LFUCache {
	cache := make(map[int]*CacheNode, 0)
	lfuCache := LFUCache{
		Size:     0,
		Capacity: capacity,
		Cache:    cache,
	}

	return lfuCache
}

func initCacheNode(key int, value int) *CacheNode {
	return &CacheNode{
		Key:      key,
		Value:    value,
		LastTime: time.Now().UnixNano(),
		HitCount: 1,
	}
}

func (this *CacheNode) addHitCount() {
	this.HitCount++
	this.LastTime = time.Now().UnixNano()
}

func (this *CacheNode) setValue(value int) {
	this.Value = value
}

func (this *LFUCache) AddNode(key int, value int) {
	cacheNode := initCacheNode(key, value)
	this.Cache[key] = cacheNode
}

func (this *LFUCache) RemoveNode(key int) {
	// 删除map
	delete(this.Cache, key)
}

func (this *LFUCache) Remove() {
	cacheArr := make([]*CacheNode, 0)
	for _, value := range this.Cache {
		cacheArr = append(cacheArr, value)
	}
	if len(cacheArr) == 0 {
		return
	}
	// 先比较hitcount
	sort.Slice(cacheArr, func(i, j int) bool {
		return cacheArr[i].HitCount <= cacheArr[j].HitCount
	})
	tmp := cacheArr[0].HitCount
	i := 0
	for ; i < len(cacheArr); i++ {
		if cacheArr[i].HitCount != tmp {
			break
		}
	}
	cacheArr = cacheArr[0:i]
	// 再比较时间戳
	sort.Slice(cacheArr, func(i, j int) bool {
		return cacheArr[i].LastTime <= cacheArr[j].LastTime
	})

	removeKey := cacheArr[0].Key
	this.RemoveNode(removeKey)
}

func (this *LFUCache) Get(key int) int {
	if this.Capacity <= 0 {
		return -1
	}
	if cacheNode, ok := this.Cache[key]; ok == true {
		cacheNode.addHitCount()
		return cacheNode.Value
	} else {
		return -1
	}
}

func (this *LFUCache) Put(key int, value int) {
	if cacheNode, ok := this.Cache[key]; ok == true {
		cacheNode.setValue(value)
		cacheNode.addHitCount()
	} else {
		// 移除
		if this.Size >= this.Capacity {
			this.Remove()
			this.Size--
		}
		this.AddNode(key, value)
		this.Size++
	}
}
