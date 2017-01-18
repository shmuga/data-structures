package data_structures

import (
	"strings"
	"testing"
)

func TestSLList_InsertBeginning(t *testing.T) {
	var l SLList
	l.InsertBeginning(123)
	assertEqual(t, l.head.value, 123, "First values is not set")
	l.InsertBeginning(111)
	assertEqual(t, l.head.next.value, 123, "Head is shifted. It's missing")
	assertEqual(t, l.head.value, 111, "Head is no replaced with new value.")
}

func TestSLList_Traverse(t *testing.T) {
	var l SLList
	l.InsertBeginning(1)
	l.InsertBeginning(2)
	l.InsertBeginning(3)
	l.InsertBeginning(0)
	l.InsertBeginning(5)
	assertEqual(t, strings.Compare(l.String(), "5->0->3->2->1->nil"), 0, l.String())
}

func TestSLList_InsertAfter(t *testing.T) {
	var l SLList
	l.InsertBeginning(1)
	l.InsertAfter(l.head, 2)
	l.InsertAfter(l.head.next, 3)

	assertEqual(t, l.head.next.value, 2, "Second variable value is not 2")
	assertEqual(t, l.head.next.next.value, 3, "Third variable value is not 3")
	assertEqual(t, strings.Compare(l.String(), "1->2->3->nil"), 0, l.String())
}

func TestSLList_InsertAfter_UnknownNode(t *testing.T) {
	var l SLList
	var l1 SLList
	l.InsertBeginning(1)
	l1.InsertBeginning(1)

	var n = SLItem{value: 1, next: nil, list: &l1}

	var err = l.InsertAfter(&n, 1)
	assertNotEqual(t, err, nil, "Error not thrown for unkown node")
}

func TestSLList_InsertAfter_EmptyNode(t *testing.T) {
	var l SLList
	l.InsertBeginning(1)
	var err = l.InsertAfter(nil, 1)
	assertNotEqual(t, err, nil, "No error throw on insert after empty node")
}

func TestSLList_RemoveBeginning(t *testing.T) {
	var l SLList
	l.InsertBeginning(1)
	l.InsertBeginning(2)

	var err = l.RemoveBeginning()
	assertEqual(t, err, nil, "Not nill during removing from not empty list")
	assertEqual(t, l.head.value, 1, "Value is not correct after deletion")
	assert(t, strings.Compare(l.String(), "1->nil") == 0, l.String())
}

func TestSLList_RemoveBeginning_EmptyList(t *testing.T) {
	var l SLList
	var err = l.RemoveBeginning()
	assertEqual(t, err, nil, "Error")
	//assertEqual(t, err, nil, err)
}

func TestSLList_RemoveAfter(t *testing.T) {
	var l SLList
	l.InsertBeginning(1)
	l.InsertBeginning(2)

	l.RemoveAfter(l.head)
	l.RemoveAfter(l.head)
	l.RemoveAfter(l.head)
}
