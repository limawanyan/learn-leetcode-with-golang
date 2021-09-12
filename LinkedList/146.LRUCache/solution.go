package _46_LRUCache

// 双向链表
type DLinkedNode struct {
	key,value int
	prev,next *DLinkedNode
}

// LRU大小、容量、map映射、伪头部、伪尾部
type LRUCache struct {
	size int
	capacity int
	cache map[int]*DLinkedNode
	head,tail *DLinkedNode
}

// 初始化节点数据
func initDLinkedNode(key, value int) *DLinkedNode {
	return &DLinkedNode{
		key:key,
		value:value,
	}
}

// 初始化LRU
func Constructor(capacity int) LRUCache {
	l := LRUCache{
		cache: map[int]*DLinkedNode{},
		head: initDLinkedNode(0,0),
		tail: initDLinkedNode(0,0),
		capacity: capacity,
	}
	l.head.next = l.tail
	l.tail.prev = l.head
	return l
}

// 添加新节点到双向链表头部
func (this *LRUCache) addToHead(node *DLinkedNode)  {
	node.prev = this.head
	node.next = this.head.next
	this.head.next.prev = node
	this.head.next = node
}

// 移除节点
func (this *LRUCache) removeNode(node *DLinkedNode) {
	node.prev.next = node.next
	node.next.prev = node.prev
}

// 移除节点并将节点添加到头部
func (this *LRUCache) moveToHead(node *DLinkedNode)  {
	this.removeNode(node)
	this.addToHead(node)
}

// 移除尾节点
func (this *LRUCache) removeTail() *DLinkedNode {
	node := this.tail.prev
	this.removeNode(node)
	return node
}

// 获取节点值
func (this *LRUCache) Get(key int) int {
	// 查询是否存在节点key
	if _,ok := this.cache[key];!ok{
		return -1
	}
	// 存在则取出节点并更新节点位置
	node := this.cache[key]
	this.moveToHead(node)
	return node.value
}

// 更新节点
func (this *LRUCache) Put(key int,value int)  {
	// 不存在节点，执行插入操作
	if _,ok := this.cache[key];!ok{
		// 组装新节点数据
		node := initDLinkedNode(key,value)
		// 存入map
		this.cache[key] = node
		// 添加到头部
		this.addToHead(node)
		// 链表长度++
		this.size++
		// 如果超出链表容量执行淘汰策略
		if this.size > this.capacity{
			// 移除尾节点（最久未使用节点）
			removed := this.removeTail()
			// 删除map映射
			delete(this.cache,removed.key)
			// 链表长度--
			this.size--
		}
		// 已存在节点，执行更新操作
	}else{
		node := this.cache[key]
		// 更新节点值
		node.value = value
		// 节点移动到头部
		this.moveToHead(node)
	}
}