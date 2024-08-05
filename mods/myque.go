package tuls

// type Node struct {
// 	queue Position
// 	next  *Node
// }

// type LinkedList struct {
// 	head *Node
// }

// func (list *LinkedList) PushBack(data Position) {
// 	newNode := &Node{queue: data, next: nil}
// 	if list.head == nil {
// 		list.head = newNode
// 		return
// 	}

// 	current := list.head
// 	for current.next != nil {
// 		current = current.next
// 	}
// 	current.next = newNode
// }

// func (list *LinkedList) PushFront(data Position) {
// 	if list.head == nil {
// 		newNode := &Node{queue: Position{
// 			Y: 0,
// 			X: 0,
// 		}, next: nil}
// 		list.head = newNode
// 		return
// 	}
// 	newNode := &Node{queue: Position{
// 		Y: 0,
// 		X: 0,
// 	}, next: list.head}
// 	list.head = newNode

// }

// func (list *LinkedList) PopFront() *Node {
// 	if list.head != nil {
// 		k := list.head
// 		list.head = list.head.next
// 		return k
// 	}
// 	return nil
// }

// func (list *LinkedList) Front() *Node {
// 	return list.head
// }

// func (list *LinkedList) PopBack() *Node {
// 	if list.head == nil {
// 		return nil
// 	}
// 	k := list.head
// 	if list.head.next == nil {
// 		list.head = nil
// 		return k
// 	}

// 	var current *Node = list.head
// 	for current.next.next != nil {
// 		current = current.next
// 	}
// 	current.next = nil
// 	return k

// }

// func (list *LinkedList) countNodes() (count int) {
// 	current := list.head
// 	for current != nil {
// 		current = current.next
// 		count++
// 	}
// 	return
// }

// func (list *LinkedList) At(index int) *Node {
// 	var count int = 0
// 	var current *Node = list.head

// 	count = list.countNodes()

// 	if index < 0 || index > count {
// 		return nil
// 	}

// 	current = list.head
// 	for i := 0; i < index; i++ {
// 		current = current.next
// 	}
// 	return current
// }

// func (list *LinkedList) Print() {
// 	var current *Node = list.head
// 	for current != nil {
// 		fmt.Printf("%d -> ", current.queue)
// 		current = current.next
// 	}
// 	fmt.Println()
// }

type QueElem struct {
	queue Position
	next  *QueElem
}

type QueList struct {
	firstElem *QueElem
}

func (que *QueList) PushBack(pos Position) {
	if que.firstElem == nil {
		newElem := QueElem{queue: pos}
		que.firstElem = &newElem
		return
	}

	current := que.firstElem
	for current.next != nil {
		current = current.next
	}
	newElem := QueElem{queue: pos}
	current.next = &newElem
}

func (que *QueList) PushFront(pos Position) {
	newElem := QueElem{queue: pos}
	current := que.firstElem
	que.firstElem = &newElem
	que.firstElem.next = current
}

func (que QueList) Back() *QueElem {
	current := que.firstElem
	for current.next != nil {
		current = current.next
	}
	return current
}

func (que QueList) Front() *QueElem {
	return que.firstElem
}

func (que QueList) At(index int) *QueElem {
	current := que.firstElem

	for i := 0; i < index; i++ {
		current = current.next
	}

	return current
}

func (que QueList) PopBack() *QueElem {
	if que.firstElem == nil {
		return nil
	}

	if que.firstElem.next == nil {
		return que.firstElem
	}

	current := que.firstElem.next
	prev := que.firstElem
	for current.next != nil {
		prev = current
		current = current.next
	}
	prev.next = nil
	return current
}

func (que QueList) PopFront() *QueElem {
	current := que.firstElem
	que.firstElem = que.firstElem.next
	return current
}
