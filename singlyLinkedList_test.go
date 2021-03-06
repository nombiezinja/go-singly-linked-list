package go_singly_linked_list

import (
	"fmt"
	"reflect"
	"testing"
)

func TestFront(t *testing.T) {
	s := New()
	if s.Front() != nil {
		t.Error("Expected nil, got ", s.Front())
	}

	n1 := &Node{Value: "test1"}
	s.Append(n1)
	if s.Front() != n1 {
		t.Error("Expected ", n1, ", got ", s.Front())
	}

	n2 := &Node{Value: "test2"}
	s.Append(n2)
	if s.Front() != n1 {
		t.Error("Expected ", n1, ", got ", s.Front())
	}
}

func TestBack(t *testing.T) {
	s := New()
	if s.Back() != nil {
		t.Error("Expected nil, got ", s.Back())
	}

	n1 := &Node{Value: "test1"}
	s.Append(n1)
	if s.Back() != n1 {
		t.Error("Expected ", n1, ", got ", s.Back())
	}

	n2 := &Node{Value: "test2"}
	s.Append(n2)
	if s.Back() != n2 {
		t.Error("Expected ", n2, ", got ", s.Back())
	}
}

func TestAppend(t *testing.T) {
	s := New()
	n1 := &Node{Value: "test1"}
	n2 := &Node{Value: "test2"}
	n3 := &Node{Value: "test3"}

	if s.Length() != 0 {
		t.Error("Expected 0, got ", s.Length())
	}
	if s.front != nil {
		t.Error("Expected nil, got ", s.front)
	}

	s.Append(n1)
	if s.Length() != 1 {
		t.Error("Expected 1, got ", s.Length())
	}
	if s.front != n1 {
		t.Error("Expected ", n1, ", got ", s.front)
	}

	s.Append(n2)
	if s.Length() != 2 {
		t.Error("Expected 2, got ", s.Length())
	}
	if s.front.next != n2 {
		t.Error("Expected ", n2, ", got ", s.front.next)
	}

	s.Append(n3)
	if s.Length() != 3 {
		t.Error("Expected 3, got ", s.Length())
	}
	if s.front.next.next != n3 {
		t.Error("Expected ", n3, ", got ", s.front.next.next)
	}
}

func TestPrepend(t *testing.T) {
	s := New()
	n1 := &Node{Value: "test1"}
	n2 := &Node{Value: "test2"}

	s.Prepend(n1)
	if s.Length() != 1 {
		t.Error("Expected 1, got ", s.Length())
	}
	if s.front != n1 {
		t.Error("Expected ", n1, ", got ", s.front)
	}

	s.Prepend(n2)
	if s.Length() != 2 {
		t.Error("Expected 2, got ", s.Length())
	}
	if s.front != n2 {
		t.Error("Expected ", n2, ", got ", s.front)
	}
	if s.front.next != n1 {
		t.Error("Expected ", n1, ", got ", s.front.next)
	}
}

var appendToSlice ValueAlterer = func(slice interface{}, value interface{}) interface{} {
	assertedSlice, ok := slice.([]string)
	if !ok {
		assertedSlice = []string{}
	}
	slice = append(assertedSlice, value.(string))
	return slice
}

func TestInsertInAscendingOrder(t *testing.T) {
	s := New()
	n1 := &Node{Key: 1}
	n2 := &Node{Key: 2}
	n5 := &Node{Key: 5}
	n4 := &Node{Key: 4, Value: []string{"hello"}}
	n4ExpectedValue := []string{"hello", "world"}
	s.Append(n1)
	s.Append(n2)
	s.Append(n4)
	s.Append(n5)

	s.InsertInAscendingOrder(3, "hello", appendToSlice)
	s.InsertInAscendingOrder(4, "world", appendToSlice)
	s.InsertInAscendingOrder(8, "hello", appendToSlice)

	fmt.Println("1", s.Front())
	fmt.Println("2", s.Front().Next())
	fmt.Println("3", s.Front().Next().Next())
	fmt.Println("4", s.Front().Next().Next().Next())
	fmt.Println("5", s.Front().Next().Next().Next().Next())
	fmt.Println("6", s.Front().Next().Next().Next().Next().Next())

	if s.Length() != 6 {
		t.Errorf("Expected length to be 6, but instead got %v", s.Length())
	}
	if s.Front().Key != 1 {
		t.Errorf("Expected node at position 1 to have Key 1, but instead got %v", s.Front())
	}
	if s.Front().Next().Key != 2 {
		t.Errorf("Expected node at position 2 to have Key 2, but instead got %v", s.Front().Next().Key)
	}
	if s.Front().Next().Next().Key != 3 {
		t.Errorf("Expected node at position 3 to have Key 3, but instead got %v", s.Front().Next().Next().Key)
	}
	if !reflect.DeepEqual(s.Front().Next().Next().Value, []string{"hello"}) {
		t.Errorf("Expected node at position 3 to have value hello, but instead got %v", s.Front().Next().Next().Value)
	}
	if !reflect.DeepEqual(s.Front().Next().Next().Next().Value, n4ExpectedValue) {
		t.Errorf("Expected node at position 4 to have value hello world, but instead got %v", s.Front().Next().Next().Next().Value)
	}
	if s.Front().Next().Next().Next().Next().Key != 5 {
		t.Errorf("Expected node at position 5 to have Key 5, but instead got %v", s.Front().Next().Next().Next().Next().Key)
	}
	if s.Front().Next().Next().Next().Next().Next().Key != 8 {
		t.Errorf("Expected node at position 6 to have Key 8, but instead got %v", s.Front().Next().Next().Next().Next().Next().Key)
	}
}
func TestInsertBefore(t *testing.T) {
	s := New()
	n1 := &Node{Value: "test1"}
	n2 := &Node{Value: "test2"}

	s.Append(n1)
	s.InsertBefore(n2, n1)
	if s.Length() != 2 {
		t.Error("Expected 2, got ", s.Length())
	}
	if s.front != n2 {
		t.Error("Expected ", n2, ", got ", s.front)
	}
	if s.front.next != n1 {
		t.Error("Expected ", n1, ", got ", s.front.next)
	}

	n3 := &Node{Value: "test3"}
	s.InsertBefore(n3, n1)
	if s.Length() != 3 {
		t.Error("Expected 3, got ", s.Length())
	}
	if s.front.next != n3 {
		t.Error("Expected ", n3, ", got ", s.front.next)
	}
	if s.front.next.next != n1 {
		t.Error("Expected ", n1, ", got ", s.front.next.next)
	}

	n4 := &Node{Value: "test4"}
	s.InsertBefore(n4, n1)
	if s.Length() != 4 {
		t.Error("Expected 4, got ", s.Length())
	}
	if s.front.next.next != n4 {
		t.Error("Expected ", n4, ", got ", s.front.next.next)
	}
	if s.front.next.next.next != n1 {
		t.Error("Expected ", n1, ", got ", s.front.next.next.next)
	}
}

func TestInsertAfter(t *testing.T) {
	s := New()
	n1 := &Node{Value: "test1"}
	n2 := &Node{Value: "test2"}

	s.Append(n1)
	s.InsertAfter(n2, n1)
	v := s.front.next
	if s.Length() != 2 {
		t.Error("Expected 2, got ", s.Length())
	}
	if v != n2 {
		t.Error("Expected ", n2, ", got ", v)
	}

	n3 := &Node{Value: "test3"}
	s.InsertAfter(n3, n2)
	v = s.front.next.next
	if s.Length() != 3 {
		t.Error("Expected 3, got ", s.Length())
	}
	if v != n3 {
		t.Error("Expected ", n3, ", got ", v)
	}
}

// TestRemoveOne tests removing the only node from a list
func TestRemoveOne(t *testing.T) {
	s := New()
	n1 := &Node{Value: "test1"}

	s.Append(n1)
	s.Remove(n1)

	if s.Length() != 0 {
		t.Error("Expected 0, got ", s.Length())
	}
	if s.front != nil {
		t.Error("Expected nil, got ", s.front)
	}
}

// TestRemoveTwofront tests removing the front node of a list with 2 nodes
func TestRemoveTwofront(t *testing.T) {
	s := New()
	n1 := &Node{Value: "test1"}
	n2 := &Node{Value: "test2"}

	s.Append(n1)
	s.Append(n2)
	s.Remove(n1)

	if s.Length() != 1 {
		t.Error("Expected 1, got ", s.Length())
	}
	if s.front != n2 {
		t.Error("Expected ", n2, ", got ", s.front)
	}
}

// TestRemoveTwoLast tests removing the last node of a list with 2 nodes
func TestRemoveTwoLast(t *testing.T) {
	s := New()
	n1 := &Node{Value: "test1"}
	n2 := &Node{Value: "test2"}

	s.Append(n1)
	s.Append(n2)
	s.Remove(n2)

	if s.front.next != nil {
		t.Error("Expected nil, got ", s.front.next)
	}
}

// TestRemoveThreeLast tests removing the last node of a list with 3 nodes
func TestRemoveThreeLast(t *testing.T) {
	s := New()
	n1 := &Node{Value: "test1"}
	n2 := &Node{Value: "test2"}
	n3 := &Node{Value: "test3"}

	s.Append(n1)
	s.Append(n2)
	s.Append(n3)
	s.Remove(n3)

	if s.Length() != 2 {
		t.Error("Expected 2, got ", s.Length())
	}
	if s.front.next.next != nil {
		t.Error("Expected nil, got ", s.front.next.next)
	}
}

// TestRemoveNonExistingNode tests removing a node not found in the list
func TestRemoveNonExistingNode(t *testing.T) {
	s := New()
	n1 := &Node{Value: "test1"}
	n2 := &Node{Value: "test2"}

	s.Append(n1)
	s.Remove(n2)

	if s.Length() != 1 {
		t.Error("Expected 1, got ", s.Length())
	}
	if s.front != n1 {
		t.Error("Expected ", n1, ", got ", s.front)
	}
}

func TestRemoveBeforeEmpty(t *testing.T) {
	s := New()
	n1 := &Node{Value: "test1"}

	s.RemoveBefore(n1)
	v := s.front
	if v != nil {
		t.Error("Expected nil, got ", v)
	}
}

func TestRemoveBeforeOneNode(t *testing.T) {
	s := New()
	n1 := &Node{Value: "test1"}

	s.Append(n1)
	s.RemoveBefore(n1)

	v := s.front
	if v != n1 {
		t.Error("Expected ", n1, ", got ", v)
	}
}

func TestRemoveBeforeTwoNodes(t *testing.T) {
	s := New()
	n1 := &Node{Value: "test1"}
	n2 := &Node{Value: "test2"}

	s.Append(n1)
	s.Append(n2)
	s.RemoveBefore(n2)

	v := s.front
	if v != n2 {
		t.Error("Expected ", n2, ", got ", v)
	}
}

func TestRemoveBeforeThreeNodes(t *testing.T) {
	s := New()
	n1 := &Node{Value: "test1"}
	n2 := &Node{Value: "test2"}
	n3 := &Node{Value: "test3"}

	s.Append(n1)
	s.Append(n2)
	s.Append(n3)
	s.RemoveBefore(n3)

	v := s.front.next
	if v != n3 {
		t.Error("Expected ", n3, ", got ", v)
	}
}

func TestRemoveBeforeFourNodes(t *testing.T) {
	s := New()
	n1 := &Node{Value: "test1"}
	n2 := &Node{Value: "test2"}
	n3 := &Node{Value: "test3"}
	n4 := &Node{Value: "test4"}

	s.Append(n1)
	s.Append(n2)
	s.Append(n3)
	s.Append(n4)
	s.RemoveBefore(n4)

	v := s.front.next.next
	if v != n4 {
		t.Error("Expected ", n4, ", got ", v)
	}
}

func TestRemoveAfterEmpty(t *testing.T) {
	s := New()
	n1 := &Node{Value: "test1"}

	s.RemoveAfter(n1)

	v := s.front
	if v != nil {
		t.Error("Expected nil, got ", v)
	}
}

func TestRemoveAfterOneNode(t *testing.T) {
	s := New()
	n1 := &Node{Value: "test1"}

	s.Append(n1)
	s.RemoveAfter(n1)

	v := s.front.next
	if v != nil {
		t.Error("Expected nil, got ", v)
	}
}

func TestRemoveAfterTwoNodes(t *testing.T) {
	s := New()
	n1 := &Node{Value: "test1"}
	n2 := &Node{Value: "test2"}

	s.Append(n1)
	s.Append(n2)
	s.RemoveAfter(n1)

	v := s.front.next
	if v != nil {
		t.Error("Expected nil, got ", v)
	}
}

func TestRemoveAfterThreeNodes(t *testing.T) {
	s := New()
	n1 := &Node{Value: "test1"}
	n2 := &Node{Value: "test2"}
	n3 := &Node{Value: "test3"}

	s.Append(n1)
	s.Append(n2)
	s.Append(n3)
	s.RemoveAfter(n2)

	v := s.front.next.next
	if v != nil {
		t.Error("Expected nil, got ", v)
	}
}

func TestGetAtPos(t *testing.T) {
	s := New()

	v := s.GetAtPos(3)
	if v != nil {
		t.Error("Expected nil, got ", v)
	}

	n1 := &Node{Value: "test1"}
	s.Append(n1)
	v = s.GetAtPos(0)
	if v != n1 {
		t.Error("Expected ", n1, ", got ", v)
	}

	n2 := &Node{Value: "test2"}
	s.Append(n2)
	v = s.GetAtPos(1)
	if v != n2 {
		t.Error("Expected ", n2, ", got ", v)
	}
}

func TestFind(t *testing.T) {
	s := New()

	v := s.Find("test")
	if v != nil {
		t.Error("Expected nil, got ", v)
	}

	n1 := &Node{Value: "test1"}
	s.Append(n1)
	v = s.Find("test1")
	if v != n1 {
		t.Error("Expected ", n1, ", got ", v)
	}

	n2 := &Node{Value: "test2"}
	s.Append(n2)
	v = s.Find("test2")
	if v != n2 {
		t.Error("Expected ", n2, ", got ", v)
	}
}

func TestFindByKey(t *testing.T) {
	s := New()

	v := s.FindByKey(0.1)
	if v != nil {
		t.Error("Expected nil, got ", v)
	}

	n1 := &Node{Key: 0.1}
	s.Append(n1)
	v = s.FindByKey(0.1)
	if v != n1 {
		t.Error("Expected ", n1, ", got ", v)
	}

	n2 := &Node{Key: 0.3}
	s.Append(n2)
	v = s.Find(0.3)
	if v != n2 {
		t.Error("Expected ", n2, ", got ", v)
	}
}
