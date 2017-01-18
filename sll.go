package data_structures

import (
	"errors"
	"strconv"
)

type SLList struct {
	head *SLItem
}

type SLItem struct {
	value int
	next  *SLItem
	list  *SLList
}

func (list *SLList) InsertBeginning(value int) {
	var i = SLItem{value: value, next: list.head, list: list}
	list.head = &i
}

func (list *SLList) InsertAfter(node *SLItem, value int) error {
	if node == nil {
		return errors.New("Node to insert after cant be nil")
	}

	var newNode = SLItem{value: value, next: node.next, list: list}
	var curr = list.head
	for curr != nil {
		if curr == node {
			curr.next = &newNode
			return nil
		}
		curr = curr.next
	}
	return errors.New("`Can't` find this node in list.")
}

func (list *SLList) RemoveBeginning() error {
	if list.head == nil {
		return nil
	}

	list.head = list.head.next
	return nil
}

func (list *SLList) RemoveAfter(node *SLItem) {
	if node.next == nil {
		return
	}
	node.next = node.next.next
	return
}

func (list *SLList) Traverse(visitor IVisitor) {
	var iter = list.head
	for iter != nil {
		visitor(iter.value)
		iter = iter.next
	}
}

func (list *SLList) String() string {
	var res = ""
	list.Traverse(func(value int) {
		res += strconv.Itoa(value) + "->"
	})
	return res + "nil"
}
