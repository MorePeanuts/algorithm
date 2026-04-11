// Package leetcode0460 solves LeetCode 460. LFU Cache
package leetcode0460

type node struct {
	key, val, freq int
	prev, next     *node
}

type LinkedList struct {
	head, tail *node
	len        int
}

func (l *LinkedList) Push(n *node) {
	if l.len == 0 {
		l.head = n
		l.tail = n
		l.head.prev = nil
		l.head.next = nil
	} else {
		n.prev = l.tail
		l.tail.next = n
		l.tail = n
		l.tail.next = nil
	}
	l.len++
}

func (l *LinkedList) Pop() *node {
	var res *node
	if l.len == 0 {
		res = nil
	} else {
		res = l.head
		l.head = l.head.next
		res.next = nil
		if l.head != nil {
			l.head.prev = nil
		} else {
			l.tail = nil
		}
	}
	l.len--
	return res
}

// Remove method assumes by default that n is a node in l.
func (l *LinkedList) Remove(n *node) *node {
	if l.head == n {
		l.Pop()
		return n
	}
	n.prev.next = n.next
	if n.next != nil {
		n.next.prev = n.prev
	} else {
		l.tail = n.prev
	}
	n.prev, n.next = nil, nil
	l.len--
	return n
}

type LFUCache struct {
	hashtable map[int]*node
	queues    []*LinkedList
	capacity  int
}

func Constructor(capacity int) LFUCache {
	return LFUCache{map[int]*node{}, []*LinkedList{}, capacity}
}

func (this *LFUCache) Get(key int) int {
	n, exist := this.hashtable[key]
	if !exist {
		return -1
	}
	this.visitNode(n)
	return n.val
}

func (this *LFUCache) Put(key int, value int) {
	n, exist := this.hashtable[key]
	if exist {
		n.val = value
		this.visitNode(n)
	} else {
		n = &node{key, value, 0, nil, nil}

		// If capacity is about to be exceeded
		if len(this.hashtable) == this.capacity {
			temp := this.popNode()
			delete(this.hashtable, temp.key)
		}
		this.hashtable[key] = n

		// Add the new node to the end of the first queue
		if n.freq > len(this.queues)-1 {
			this.queues = append(this.queues, &LinkedList{})
		}
		this.queues[n.freq].Push(n)
	}
}

func (this *LFUCache) visitNode(n *node) {
	n = this.queues[n.freq].Remove(n)
	n.freq++
	if n.freq > len(this.queues)-1 {
		this.queues = append(this.queues, &LinkedList{})
	}
	this.queues[n.freq].Push(n)
}

func (this *LFUCache) popNode() *node {
	var res *node
	for i, queue := range this.queues {
		if queue.len != 0 {
			res = queue.Pop()
			if i == len(this.queues)-1 && queue.len == 0 {
				this.queues = this.queues[:len(this.queues)-1]
			}
			break
		}
	}
	return res
}
