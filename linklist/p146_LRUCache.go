package linklist

import "fmt"

// 146. LRU 缓存 medium 2023-12-06 535
// 实现lru缓存

// !!!! 数据结构设计原因：
// 要求1. 要让 put 和 get 方法的时间复杂度为 O(1) | 快速找某个 key 是否已存在并得到对应的 val  ---》 哈希表
// 要求2. cache 中的元素必须有时序 && 支持在任意位置快速插入和删除元素 --》 双向链表 （单链表 链表尾部删除O(n)
// ---> 双向链表+哈希表

// https://leetcode.cn/problems/lru-cache/description/
// struct LRUCache 包括四个成员size,capacity,map[int]*DLinkedNode,双链表的伪头部伪尾部
// 要实现三个方法Constructor初始化、Get、Put
// 双链表结构体包括key,value,prev,next
// 双链表 可以实现moveToHead、removeNode、addToHead、removeTail等辅助实现

func Problem146() {
	lru := Constructor(2)
	lru.Put(1, 1)
	lru.Put(2, 2)
	fmt.Println(lru.Get(1))
	lru.Put(3, 3)
	fmt.Println(lru.Get(2))
	lru.Put(4, 4)
}

type LRUCache struct {
	size     int
	capacity int
	cache    map[int]*DLinkedNode
	// 双向链表使用伪头部伪尾部标记界限 不用判断头尾
	head, tail *DLinkedNode
}
type DLinkedNode struct {
	key, value int
	prev, next *DLinkedNode
}

func Constructor(capacity int) LRUCache {
	l := LRUCache{
		size:     0,
		capacity: capacity,
		cache:    map[int]*DLinkedNode{},
		head:     &DLinkedNode{0, 0, nil, nil},
		tail:     &DLinkedNode{0, 0, nil, nil},
	}
	// 初始化 伪头部伪尾部要互相指向
	// 注意是l
	l.head.next = l.tail
	l.tail.prev = l.head
	return l
}

// Get 如果有缓存将双链表结点移到队头 否则返回-1
func (this *LRUCache) Get(key int) int {
	if _, ok := this.cache[key]; !ok {
		return -1
	}
	node := this.cache[key]
	this.moveToHead(node)
	return node.value
}

// Put 如果key已存在，更新value并moveToHead
// key不存在，插入新的key,value 若容量已满 淘汰队尾结点
func (this *LRUCache) Put(key, value int) {
	if _, ok := this.cache[key]; ok {
		node := this.cache[key]
		node.value = value
		this.moveToHead(node)
	} else {
		// 创建新节点
		node := &DLinkedNode{key, value, nil, nil}
		// map中添加
		this.cache[key] = node
		// 双向链表中添加
		this.addToHead(node)
		// size++
		this.size++
		if this.size > this.capacity {
			// 淘汰最久未使用结点 (map、DLinkList)
			// 写错：先删map中的键 否则this.tail.prev.key是错的
			delete(this.cache, this.tail.prev.key)
			this.removeNode(this.tail.prev)
			this.size--
		}
	}
}
func (this *LRUCache) moveToHead(node *DLinkedNode) {
	this.removeNode(node)
	this.addToHead(node)
}
func (this *LRUCache) removeNode(node *DLinkedNode) {
	node.prev.next = node.next
	node.next.prev = node.prev
}
func (this *LRUCache) addToHead(node *DLinkedNode) {
	node.prev = this.head
	node.next = this.head.next
	this.head.next.prev = node
	this.head.next = node
}

//func (this *LRUCache) removeTail() *DLinkedNode {
//	node := this.tail.prev
//	this.removeNode(node)
//	return node
//}
