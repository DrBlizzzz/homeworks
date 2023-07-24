package hw04lrucache

import (
	"fmt"
)

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
	Head *ListItem
	Tail *ListItem
	Size int
}

func (l list) Len() int {
	return l.Size
}

func (l list) Front() *ListItem {
	return l.Head
}

func (l list) Back() *ListItem {
	return l.Tail
}

func (l *list) PushFront(v interface{}) *ListItem {
	if l.Head == nil {
		node := ListItem{v, nil, nil}
		l.Head = &node
		l.Tail = &node
	} else {
		node := ListItem{v, nil, l.Head}
		l.Head.Prev = &node
		l.Head = &node
	}
	l.Size++
	return l.Head
}

func (l *list) PushBack(v interface{}) *ListItem {
	if l.Tail == nil {
		node := ListItem{v, nil, nil}
		l.Head = &node
		l.Tail = &node
	} else {
		node := ListItem{v, l.Tail, nil}
		l.Tail.Next = &node
		l.Tail = &node
	}
	l.Size++
	return l.Head
}

func (l list) FindElement(v interface{}) int {
	return 0
}

func (l *list) Remove(i *ListItem) {
	fmt.Println("Removing in progress")
	l.Size--
}

func (l *list) MoveToFront(i *ListItem) {
	fmt.Println("MoveToFront in progress")
}
