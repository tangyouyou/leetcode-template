package leetcode

//运用你所掌握的数据结构，设计和实现一个  LRU (最近最少使用) 缓存机制 。
//实现 LRUCache 类：
type LRUCache struct {
	size int
	capacity int
	cache map[int]*DLinkedNode
	head *DLinkedNode
	tail *DLinkedNode
}

// 定义双向链表
type DLinkedNode struct{
	key int
	value int
	prev *DLinkedNode
	next *DLinkedNode
}

func initDLinkedNode(key, value int) *DLinkedNode {
	return &DLinkedNode{
		key : key,
		value : value,
	}
}

func Constructor(capacity int) LRUCache {
	l := LRUCache{
		cache : map[int]*DLinkedNode{},
		head : initDLinkedNode(0, 0),
		tail : initDLinkedNode(0, 0),
		capacity : capacity,
	}
	// 构建链表环
	l.head.next = l.tail
	l.tail.next = l.head

	return l
}

func (this *LRUCache) addToHead(node *DLinkedNode) {
	node.prev = this.head
	node.next = this.head.next
	this.head.next.prev = node
	this.head.next = node
}

func (this *LRUCache) removeNode(node *DLinkedNode) {
	node.prev.next = node.next
	node.next.prev = node.prev
}

func (this *LRUCache) moveToHead(node *DLinkedNode) {
	this.removeNode(node)
	this.addToHead(node)
}

func (this *LRUCache) remvoeTail() *DLinkedNode {
	node := this.tail.prev
	this.removeNode(node)
	return node
}

func (this *LRUCache) Get(key int) int {
	if _,ok := this.cache[key]; !ok {
		return -1
	}
	node := this.cache[key]
	this.moveToHead(node)

	return node.value
}


func (this *LRUCache) Put(key int, value int)  {
	if _,ok := this.cache[key]; !ok {
		node := initDLinkedNode(key, value)
		this.cache[key] = node
		this.addToHead(node)
		this.size++
		// 容量已满
		if this.size > this.capacity {
			// 删除双向链表
			removed := this.remvoeTail()
			// 删除 map
			delete(this.cache, removed.key)
			this.size--
		}
	} else {
		node := this.cache[key]
		node.value = value
		this.moveToHead(node)
	}
}