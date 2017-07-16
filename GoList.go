package GoList

import "errors"

type Node struct {
	Value interface{}
	Next *Node
}

type List struct {
	Head *Node
}

func (this *List) Front() *Node {
	return this.Head
}

func (this *List) PushBack(value interface{}) {
	if this.Head == nil {
		this.Head = &Node{Value:value, Next:nil}
		return
	}
	node := this.Head
	for {
		if node.Next != nil {
			node = node.Next
		}else {
			node.Next = &Node{Value:value, Next:nil}
			return
		}
	}
}

func (this *List) PushAt(pos int, value interface{}) error {
	if pos > this.GetLength() || pos < 0 {
		return errors.New("Index out of range")
	}
	if pos == 0 {
		this.PushFront(value)
		return nil
	}
	cur := 1
	node := this.Head
	for {
		if pos == cur {
			nx := node.Next
			node.Next = &Node{Value:value, Next:nx}
			return nil
		}
		node = node.Next
		cur++
	}
}

func (this *List) PushFront(value interface{}) {
	if this.Head == nil {
		this.Head = &Node{Value:value, Next:nil}
		return
	}
	node := this.Head
	this.Head = &Node{Value:value, Next:node}
}

func (this *List) GetLength() int {
	len := 0
	node := this.Head
	for {
		if node == nil {
			return len
		}
		node = node.Next
		len ++
	}
}

func (this *List) GetAt(pos int) (interface{}, error) {
	cur := 0
	node := this.Head
	for {
		if node == nil {
			return nil, errors.New("Index out of range")
		}
		if cur == pos {
			return node.Value, nil
		}
		node = node.Next
		cur ++
	}
}

func (this *List) SetAt(pos int, value interface{}) (interface{}, error) {
	if pos < 0 || pos >= this.GetLength() {
		return nil, errors.New("Index out of range")
	}
	cur := 0
	node := this.Head
	for {
		if cur == pos {
			prev_value := node.Value
			node.Value = value
			return prev_value, nil
		}
		node = node.Next
		cur ++
	}
}

func (this *List) PopBack() (interface{}, error) {
	if this.Head == nil {
		return nil, errors.New("Empty list")
	}
	if this.Head.Next == nil {
		val := this.Head.Value
		this.Head = nil
		return val, nil
	}
	node := this.Head
	for {
		if node.Next.Next == nil {
			val := node.Next.Value
			node.Next = nil
			return val, nil
		}
		node = node.Next
	}
}

func (this *List) PopAt(pos int) (interface{}, error) {
	if pos >= this.GetLength() || pos < 0 {
		return nil, errors.New("Index out of range")
	}
	if pos == 0 {
		return this.PopFront()
	}
	node := this.Head
	cur := 1
	for {
		if cur == pos {
			val := node.Next.Value
			node.Next = node.Next.Next
			return val, nil
		}
		node = node.Next
		cur ++
	}
}

func (this *List) PopFront() (interface{}, error) {
	if this.Head == nil {
		return nil, errors.New("Empty list")
	}
	val := this.Head.Value
	this.Head = this.Head.Next
	return val, nil
}

func (this *List) Contains(value interface{}) bool {
	for node := this.Front(); node != nil; node = node.Next {
		if node.Value == value {
			return true
		}
	}
	return false
}
