// Package List implements a singly linked list
package go_singly_linked_list

// List is a singly linked list implementation
type List struct {
	front *Node

	length int
}

// Init initializes an empty list
func (s *List) Init() *List {
	s.length = 0
	return s
}

// New returns an initialized list
func New() *List {
	return new(List).Init()
}

// Front returns the first node in list s
func (s *List) Front() *Node {
	return s.front
}

// Back returns the last node in list s
func (s *List) Back() *Node {
	currentNode := s.front
	for currentNode != nil && currentNode.next != nil {
		currentNode = currentNode.next
	}
	return currentNode
}

// Append appends node n to list s
func (s *List) Append(n *Node) {
	if s.front == nil {
		s.front = n
	} else {
		currentNode := s.front

		for currentNode.next != nil {
			currentNode = currentNode.next
		}

		currentNode.next = n
	}

	s.length++
}

// Prepend prepends node n to list s
func (s *List) Prepend(n *Node) {
	if s.front == nil {
		s.front = n
	} else {
		n.next = s.front
		s.front = n
	}

	s.length++
}

// InsertBefore inserts node insert before node before in list s
func (s *List) InsertBefore(insert *Node, before *Node) {
	if s.front == before {
		insert.next = before
		s.front = insert
		s.length++
	} else {
		currentNode := s.front
		for currentNode != nil && currentNode.next != nil && currentNode.next != before {
			currentNode = currentNode.next
		}

		if currentNode.next == before {
			insert.next = before
			currentNode.next = insert
			s.length++
		}
	}
}

// InsertAfter inserts node insert after node after in list s
func (s *List) InsertAfter(insert *Node, after *Node) {
	currentNode := s.front
	for currentNode != nil && currentNode != after && currentNode.next != nil {
		currentNode = currentNode.next
	}

	if currentNode == after {
		insert.next = after.next
		after.next = insert
		s.length++
	}
}

type ValueAlterer func(toAlter interface{}, alterWith interface{}) interface{}

// InsertInOrder inserts new node such that list is always in ascending order of the
// float64 keys, then returns the node
func (s *List) InsertInAscendingOrder(key float64, value interface{}, valueAlterer ValueAlterer) {
	currentNode := s.front
	toInsert := &Node{Key: key, Value: value}
	if currentNode == nil {
		s.front = toInsert
	} else {
		for currentNode.Key < key {
			// If node at current iteration is smaller than key, move on
			currentNode = currentNode.next
			// if key already exists in list, alter the value with anonymous function passed as argument
		}
		if currentNode.Key == key {
			currentNode.Value = valueAlterer(currentNode.Value, value)
		} else {
			// if node at current iteration is bigger than key, insert before this node
			s.InsertBefore(toInsert, currentNode)
		}
	}
}

// Remove removes node n from list s
func (s *List) Remove(n *Node) {
	if s.front == n {
		s.front = n.next
		s.length--
	} else {
		currentNode := s.front

		// search for node n
		for currentNode != nil && currentNode.next != nil && currentNode.next != n {
			currentNode = currentNode.next
		}

		// see if current's next node is n
		// if it's not n, then node n wasn't found in list s
		if currentNode.next == n {
			currentNode.next = currentNode.next.next
			s.length--
		}
	}
}

// RemoveBefore removes node before node before
func (s *List) RemoveBefore(before *Node) {
	if s.front != nil && s.front != before {
		if s.front.next == before {
			s.front = before
		} else {
			currentNode := s.front
			for currentNode.next.next != nil && currentNode.next.next != before {
				currentNode = currentNode.next
			}
			if currentNode.next.next == before {
				currentNode.next = before
			}
		}
	}
}

// RemoveAfter removes node after node after
func (s *List) RemoveAfter(after *Node) {
	if s.front != nil && s.front.next != nil {
		currentNode := s.front
		for currentNode != after && currentNode.next != nil {
			currentNode = currentNode.next
		}

		if currentNode == after {
			currentNode.next = currentNode.next.next
		}
	}
}

// GetAtPos returns the node at index in list s
func (s *List) GetAtPos(index int) *Node {
	currentNode := s.front
	count := 0
	for count < index && currentNode != nil && currentNode.next != nil {
		currentNode = currentNode.next
		count++
	}

	if count == index {
		return currentNode
	} else {
		return nil
	}
}

// Find returns the node with matching value or nil if not found
func (s *List) Find(value interface{}) *Node {
	currentNode := s.front
	for currentNode != nil && currentNode.Value != value && currentNode.next != nil {
		currentNode = currentNode.next
	}

	return currentNode
}

// Find returns the node with matching key or nil if not found
func (s *List) FindByKey(key float64) *Node {
	currentNode := s.front
	for currentNode != nil && currentNode.Key != key && currentNode.next != nil {
		currentNode = currentNode.next
	}

	return currentNode
}

// Length returns the amount of nodes in list s
func (s *List) Length() int {
	return s.length
}
