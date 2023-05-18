package hw04lrucache

type List interface {
	Len() int
	Front() *ListItem
	Back() *ListItem
	PushFront(v interface{}) *ListItem
	PushBack(v interface{}) *ListItem
	Remove(i *ListItem)
	MoveToFront(i *ListItem)
}

func NewList() List {
	return new(list)
}

type ListItem struct {
	Value interface{}
	Next  *ListItem
	Prev  *ListItem
}

type list struct {
	size int
	head *ListItem
	tail *ListItem
}

func Newlist() *list {
	return &list{0, nil, nil}
}

func (l list) Len() int {
	return l.size
}

func (l list) Front() *ListItem {
	return l.head
}

func (l list) Back() *ListItem {
	return l.tail
}

func (l *list) PushFront(v interface{}) *ListItem {
	node := ListItem{v, nil, nil}
	if l.head == nil {
		l.head = &node
		l.tail = &node
	} else {
		l.head.Prev = &node
		node.Next = l.head
		l.head = &node
	}
	l.size++
	return &node
}

func (l *list) PushBack(v interface{}) *ListItem {
	node := ListItem{v, nil, nil}
	if l.tail == nil {
		l.head = &node
		l.tail = &node
	} else {
		l.tail.Next = &node
		node.Prev = l.tail
		l.tail = &node
	}
	l.size++
	return &node
}

func (l *list) Remove(item *ListItem) {
	if item == l.head {
		if l.size != 1 {
			l.head = l.head.Next
			l.head.Prev = nil
		} else {
			l.head = nil
		}
		l.size--
		return
	}
	if item == l.tail {
		if l.size != 1 {
			l.tail = l.tail.Prev
			l.tail.Next = nil
		} else {
			l.tail = nil
		}
		l.size--
		return
	}
	ptr := l.head
	for ptr != item {
		ptr = ptr.Next
	}
	ptr.Prev.Next = ptr.Next
	ptr.Next.Prev = ptr.Prev
	l.size--
}

func (l *list) MoveToFront(item *ListItem) {
	if item == l.head {
		return
	} else {
		l.Remove(item)
		l.PushFront(item.Value)
	}
}
