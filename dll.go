package data_structures

import (
	"errors"
	"strconv"
)

type DLItem struct {
	value int
	next  *DLItem
	prev  *DLItem
	list  *DLList
}

type DLList struct {
	head *DLItem
	tail *DLItem
}

func (list *DLList) InsertBeginning(value int) {
	var i = DLItem{
		value: value,
		next:  list.head,
		list:  list,
	}

	if list.head != nil {
		list.head.prev = &i
	} else {
		list.tail = &i
	}
	list.head = &i

}

func (list *DLList) InsertAfter(node *DLItem, value int) error {
	if node == nil {
		return errors.New("Node item can be nil")
	}

	var i = DLItem{
		value: value,
		prev:  node,
		list:  list,
	}

	if (node == list.tail) {
		list.tail = &i
	}

	var iter = list.head
	for iter != nil {
		if iter == node {
			if node.next != nil {
				node.next.prev = &i
			}
			i.next = node.next
			node.next = &i
			return nil
		}
		iter = iter.next
	}

	return errors.New("Current node not found in List")
}

func (list *DLList) RemoveBeginning()  {
	if list.head != nil {
		list.head = list.head.next
		if list.head != nil {
			list.head.prev = nil
		} else {
			list.tail = nil
		}
	}

}

func (list *DLList) RemoveAfter(node *DLItem)  {
	if node.next != nil {
		if list.tail == node.next {
			list.tail = node
		}

		node.next = node.next.next

	}
}

func (list *DLList) Traverse(visitor IVisitor) {
	var iter = list.head
	for iter != nil {
		visitor(iter.value)
		iter = iter.next
	}
}

func (list *DLList) String() string {
	var res = ""
	list.Traverse(func(value int) {
		res += strconv.Itoa(value) + "<->"
	})
	return res + "nil"
}
