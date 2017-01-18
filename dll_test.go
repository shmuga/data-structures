package data_structures

import (
	"strings"
	"testing"
)

func TestDLList_InsertBeginning(t *testing.T) {
	var l DLList
	l.InsertBeginning(1)
	assertEqual(t, l.head, l.tail, "List and tail now the same")
	assertEqual(t, l.head.value, 1, "Insert is not working")
	assert(t, l.head.next == nil, l.head.next)

	l.InsertBeginning(2)
	assertNotEqual(t, l.head, l.tail, "List and tail now not the same")
	assertEqual(t, l.head.value, 2, "Second Insert is not working")
	assertEqual(t, l.head.next.prev.value, l.head.value, "Second Insert is not working")
	assertEqual(t, strings.Compare(l.String(), "2<->1<->nil"), 0, l.String())
}

func TestDLList_InsertAfter(t *testing.T) {
	var l DLList
	l.InsertBeginning(1)
	l.InsertAfter(l.head, 2)
	assertEqual(t, l.tail.value, 2, "Tail should have value 2")
	assertEqual(t, l.head.next.value, 2, "Insert is not working")
	l.InsertAfter(l.head.next, 3)
	assertEqual(t, l.tail.value, 3, "Tail should have value 3")
	assertEqual(t, l.head.next.next.value, 3, "Insert is not working")
	assertEqual(t, strings.Compare(l.String(), "1<->2<->3<->nil"), 0, l.String())
}


func TestDLList_InsertAfter_EmptyNode(t *testing.T) {
	var l DLList
	l.InsertBeginning(1)
	var err = l.InsertAfter(nil, 3)
	assert(t, err != nil, "No error thrown")
}

func TestDLList_InsertAfter_UnknownNode(t *testing.T) {
	var l DLList
	var n DLItem
	l.InsertBeginning(1)
	var err = l.InsertAfter(&n, 3)
	assert(t, err != nil, "No error thrown")
}

func TestDLList_RemoveBeginning(t *testing.T) {
	var l DLList
	l.InsertBeginning(1)
	l.InsertBeginning(2)
	l.InsertBeginning(3)
	assertEqual(t, strings.Compare(l.String(), "3<->2<->1<->nil"), 0, l.String())
	l.RemoveBeginning()
	assertEqual(t, l.head.value, 2, "")
	assert(t, l.head.prev == nil, "")
	l.RemoveBeginning()
	assertEqual(t, l.head.value, 1, "")
	assert(t, l.head.prev == nil, "")
	l.RemoveBeginning()
	assert(t, l.tail == nil, "Tail is not set to nil")
	assertEqual(t, strings.Compare(l.String(), "nil"), 0, l.String())
}


func TestDLList_RemoveAfter_Head(t *testing.T) {
	var l DLList
	l.InsertBeginning(1)
	l.InsertBeginning(2)

	l.RemoveAfter(l.head)
	assertEqual(t, l.head.value, 2, "Value is not correct")
}

func TestDLList_RemoveAfter_Tail(t *testing.T) {
	var l DLList
	l.InsertBeginning(1)
	l.InsertBeginning(2)

	l.RemoveAfter(l.tail)
	assert(t, l.tail.next == nil, "Value is not correct")
	assert(t, l.tail.value == 1, "Value is not correct")
}

func TestDLList_RemoveAfter(t *testing.T) {
	var l DLList
	l.InsertBeginning(1)
	l.InsertBeginning(2)
	l.InsertBeginning(3)

	l.RemoveAfter(l.head.next)
	assertEqual(t, strings.Compare(l.String(), "3<->2<->nil"), 0, l.String())

}

func TestDLList_RemoveAfter_RemoveTail(t *testing.T) {
	var l DLList
	l.InsertBeginning(1)
	l.InsertBeginning(2)
	l.InsertBeginning(3)
	old := l.tail.prev
	l.RemoveAfter(l.tail.prev)
	assertEqual(t, l.tail, old, "Tail is not replced")
	assertEqual(t, strings.Compare(l.String(), "3<->2<->nil"), 0, l.String())

}
