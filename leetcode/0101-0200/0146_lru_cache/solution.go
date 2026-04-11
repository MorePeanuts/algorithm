// Package leetcode0146 solves LeetCode 146. LRU Cache
package leetcode0146

type node struct {
	key  int
	val  int
	next *node
	prev *node
}

type LRUCache struct {
	hashtable map[int]*node
	head      *node
	tail      *node
	capacity  int
}

func Constructor(capacity int) LRUCache {
	return LRUCache{hashtable: map[int]*node{}, capacity: capacity}
}

func (this *LRUCache) Get(key int) int {
	n, exist := this.hashtable[key]
	if !exist {
		return -1
	}
	this.moveToTail(n)
	return n.val
}

func (this *LRUCache) Put(key int, value int) {
	n, exist := this.hashtable[key]
	if exist {
		n.val = value
		this.moveToTail(n)
	} else {
		n = &node{key, value, nil, this.tail}
		this.hashtable[key] = n

		// Add the new node to the end of the queue
		if this.tail != nil {
			// The queue is not empty
			this.tail.next = n
		} else {
			// The queue is empty
			this.head = n
		}
		this.tail = n

		// If the queue length exceeds the capacity, remove the node at the head of the queue.
		if len(this.hashtable) > this.capacity {
			delete(this.hashtable, this.head.key)
			this.head.next, this.head = nil, this.head.next
			this.head.prev = nil
		}
	}
}

func (this *LRUCache) moveToTail(n *node) {
	if this.tail != n {
		if this.head == n {
			// If n is at the head of the queue, the head pointer needs to be reset.
			this.head = this.head.next
		}
		n.next.prev = n.prev
		if n.prev != nil {
			n.prev.next = n.next
		}
		this.tail.next, n.prev = n, this.tail
		this.tail, n.next = n, nil
	}
}

/**
 * Your LRUCache object will be instantiated and called as such:
 * obj := Constructor(capacity);
 * param_1 := obj.Get(key);
 * obj.Put(key,value);
 */
