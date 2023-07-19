package linkedlist

// prepend
// append
// lookup
// insert
// delete

type Node[T any] struct {
	next  *Node[T]
	value T
}

type LinkedList[T any] struct {
	head *Node[T]
	tail *Node[T]
}

func New[T any]() *LinkedList[T] {
	return &LinkedList[T]{}
}

func (list *LinkedList[T]) Prepend(value T) *LinkedList[T] {
	node := &Node[T]{nil, value}
	if list.head != nil {
		node.next = list.head
		list.head = node
	} else {
		// make head
		list.head = node
		list.tail = node
	}

	return list
}

func (list *LinkedList[T]) Append(value T) *LinkedList[T] {
	// create node
	node := &Node[T]{nil, value}
	// if first node in list
	if list.head == nil {
		list.head = node
		list.tail = node
	} else {
		list.tail.next = node
		list.tail = list.tail.next
	}

	return list
}

func (list *LinkedList[T]) Lookup() T {
	return T
}

func (list *LinkedList[T]) Insert() T {
	return T
}

func (list *LinkedList[T]) Delete() T {
	return T
}
