package main

type Node2 struct {
	Key   int
	Value int
	Prev  *Node2
	Next  *Node2
}

type LRUCache struct {
	Capacity  int
	Cache     map[int]*Node2
	FirstNode *Node2
	LastNode  *Node2
}

func Constructor(capacity int) LRUCache {
	return LRUCache{
		Capacity:  capacity,
		Cache:     make(map[int]*Node2),
		FirstNode: nil,
		LastNode:  nil,
	}
}

func (this *LRUCache) Get(key int) int {
	v, ok := this.Cache[key]
	if ok {
		if v == this.FirstNode && v == this.LastNode {

		} else if v == this.FirstNode {
			next := v.Next
			next.Prev = v.Prev
			this.LastNode.Next = v
			v.Prev = this.LastNode
			this.FirstNode = next
			this.LastNode = v
		} else if v == this.LastNode {

		} else {
			prev := v.Prev
			next := v.Next
			prev.Next = next
			next.Prev = prev
			this.LastNode.Next = v
			v.Prev = this.LastNode
			this.LastNode = v
		}

		return v.Value
	}

	return -1
}

func (this *LRUCache) Put(key int, value int) {
	node := &Node2{
		Key:   key,
		Value: value,
	}

	v, ok := this.Cache[key]
	if ok {
		if v == this.FirstNode && v == this.LastNode {

		} else if v == this.FirstNode {
			next := v.Next
			next.Prev = v.Prev
			this.LastNode.Next = v
			v.Prev = this.LastNode
			this.FirstNode = next
			this.LastNode = v
		} else if v == this.LastNode {

		} else {
			prev := v.Prev
			next := v.Next
			prev.Next = next
			next.Prev = prev
			this.LastNode.Next = v
			v.Prev = this.LastNode
			this.LastNode = v
		}
		v.Value = value
		node = v
	} else if this.Capacity == 1 {
		first := this.FirstNode
		if first != nil {
			delete(this.Cache, first.Key)
		}
		this.FirstNode = node
		this.LastNode = node
	} else if len(this.Cache) < this.Capacity {
		if this.FirstNode == nil && this.LastNode == nil {
			this.FirstNode = node
		} else {
			this.LastNode.Next = node
			node.Prev = this.LastNode
		}
		this.LastNode = node
	} else {
		first := this.FirstNode
		second := first.Next
		delete(this.Cache, first.Key)
		first.Next = nil
		second.Prev = nil
		this.FirstNode = second

		this.LastNode.Next = node
		node.Prev = this.LastNode
		this.LastNode = node
	}

	this.Cache[key] = node
}

/**
 * Your LRUCache object will be instantiated and called as such:
 * obj := Constructor(capacity);
 * param_1 := obj.Get(key);
 * obj.Put(key,value);
 */
