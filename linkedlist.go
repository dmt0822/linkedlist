package linkedlist

import "fmt"

type Node struct {
	val  int
	next *Node
}
type MyLinkedList struct {
	head *Node
}

func Constructor() MyLinkedList {
	return MyLinkedList{}
}

func (this *MyLinkedList) Get(targetIndex int) int {
	if targetIndex < 0 {
		return -1
	}

	currentNode := this.head

	for currentIndex := 0; currentIndex <= targetIndex; currentIndex++ {
		if isEmptyNode(currentNode) {
			return -1
		}
		if currentIndex == targetIndex {
			return currentNode.val
		}
		if currentIndex == targetIndex-1 {
			if currentNode.next != nil {
				return currentNode.next.val
			}
		}
		if currentNode.next == nil {
			return -1
		}
		currentNode = currentNode.next
	}
	return -1
}

func (this *MyLinkedList) AddAtHead(val int) {
	if isEmptyNode(this.head) {
		this.head = &Node{val, nil}
		return
	}

	this.head = &Node{val, this.head}
}

func (this *MyLinkedList) AddAtTail(val int) {
	this.add(val, this.head)
}

func (this *MyLinkedList) add(val int, currentNode *Node) {
	if isEmptyNode(currentNode) {
		this.head = &Node{val, nil}
		return
	}

	if currentNode.next == nil {
		currentNode.next = &Node{val, nil}
		return
	}

	this.add(val, currentNode.next)
}

func isEmptyNode(node *Node) bool {
	emptyNode := Node{}

	return node == &emptyNode || node == nil
}

/** Add a node of value val before the index-th node in the linked list.
 * If index equals to the length of linked list, the node will be appended to the end of linked list.
 * If index is greater than the length, the node will not be inserted.
 */
func (this *MyLinkedList) AddAtIndex(targetIndex int, val int) {
	currentNode := this.head

	if isEmptyNode(currentNode) {
		this.head = &Node{val, nil}
		return
	}

	for currentIndex := 0; currentIndex <= targetIndex; currentIndex++ {
		if currentIndex == targetIndex {
			this.head = &Node{val, currentNode}
		}
		if currentIndex == targetIndex-1 {
			currentNext := currentNode.next
			newNode := &Node{val, currentNext}
			currentNode.next = newNode
			return
		}
		if currentNode.next == nil {
			return
		}
		currentNode = currentNode.next
	}
}

func (this *MyLinkedList) DeleteAtIndex(targetIndex int) {
	currentNode := this.head

	for currentIndex := 0; currentIndex <= targetIndex; currentIndex++ {
		if isEmptyNode(currentNode) {
			return
		}
		if currentNode.next == nil {
			return
		}
		if currentIndex == targetIndex-1 {
			if currentNode.next.next != nil {
				currentNode.next = currentNode.next.next
				return
			}

			currentNode.next = nil
			return
		}
		if currentIndex == targetIndex {
			this.head = currentNode.next
			return
		}
		currentNode = currentNode.next
	}
}

func (this *MyLinkedList) String() {
	this.print(this.head)
}

func (this *MyLinkedList) print(node *Node) {
	empty := &MyLinkedList{}
	if node == empty.head {
		return
	}
	fmt.Println(node)
	if node.next != empty.head {
		this.print(node.next)
	}
}
