package linkedlist

import (
	"fmt"
	"strings"
)

type ListNode struct {
	Val  int
	Next *ListNode
}

func New(val int, next *ListNode) *ListNode {
	return &ListNode{val, next}
}

func Get(head *ListNode, targetIndex int) int {
	if targetIndex < 0 {
		return -1
	}

	if head.IsEmpty() {
		return -1
	}

	currentNode := head

	for currentIndex := 0; currentIndex <= targetIndex; currentIndex++ {
		if currentNode.IsEmpty() {
			return -1
		}
		if currentIndex == targetIndex {
			return currentNode.Val
		}
		if currentIndex == targetIndex-1 {
			if currentNode.Next != nil {
				return currentNode.Next.Val
			}
		}
		if currentNode.Next == nil {
			return -1
		}
		currentNode = currentNode.Next
	}
	return -1
}

func AddAtHead(head *ListNode, val int) *ListNode {
	return New(val, head)
}

func AddAtTail(head *ListNode, val int) *ListNode {
	if head.IsEmpty() {
		return &ListNode{val, nil}
	}
	currentNode := head
	newNode := New(val, nil)

	for currentNode != nil {
		if currentNode.Next == nil {
			currentNode.Next = newNode
			break
		}
		currentNode = currentNode.Next
	}
	return head
}

func AddAtIndex(head *ListNode, targetIndex int, val int) *ListNode {
	if head.IsEmpty() {
		return New(val, nil)
	}

	if targetIndex == 0 {
		return New(val, head)
	}

	newHead := head
	currentNode := head
	var previousNode *ListNode

	for currentIndex := 0; currentIndex <= targetIndex; currentIndex++ {
		if currentNode == nil {
			return newHead
		}
		if currentIndex == targetIndex {
			newNode := New(val, currentNode)
			if !(previousNode.IsEmpty()) {
				previousNode.Next = newNode
			}
			return newHead
		}
		previousNode = currentNode
		currentNode = currentNode.Next
	}
	return newHead
}

func DeleteAtIndex(head *ListNode, targetIndex int) *ListNode {
	if head.IsEmpty() {
		return head
	}
	currentNode := head
	newHead := head
	var previousNode *ListNode

	for currentIndex := 0; currentIndex <= targetIndex; currentIndex++ {
		if currentNode.IsEmpty() {
			return newHead
		}
		if currentIndex == targetIndex {
			if !previousNode.IsEmpty() {
				previousNode.Next = currentNode.Next
			} else {
				newHead = currentNode.Next
			}
		}
		previousNode = currentNode
		currentNode = currentNode.Next
	}
	return newHead
}

func RemoveNthFromEnd(head *ListNode, numFromEnd int) *ListNode {
	fast := head
	slow := head
	delay := numFromEnd + 1
	currentIndex := 0
	newHead := head

	for fast != nil {
		fast = fast.Next
		if currentIndex >= delay {
			slow = slow.Next
		}
		currentIndex++
		if fast == nil {
			if currentIndex < delay {
				newHead = slow.Next
			} else if slow.Next != nil {
				slow.Next = slow.Next.Next
			} else {
				slow.Next = nil
			}
		}
	}

	return newHead
}

func ReverseList(head *ListNode) *ListNode {
	var previous *ListNode

	for head != nil {
		next := head.Next
		head.Next = previous
		previous = head
		head = next
	}

	return previous
}

func RemoveElements(head *ListNode, val int) *ListNode {
	if head.IsEmpty() {
		return head
	}
	var newHead, tailNode, nextNode *ListNode
	currentNode := head
	for currentNode != nil {
		nextNode = currentNode.Next
		if currentNode.Val != val {
			currentNode.Next = nil
			if newHead == nil {
				newHead = currentNode
			} else {
				tailNode.Next = currentNode

			}
			tailNode = currentNode
		}
		currentNode = nextNode
	}
	return newHead
}

func OddEvenList(head *ListNode) *ListNode {
	if head.IsEmpty() || head.Next == nil {
		return head
	}
	var oddHead, oddTail, evenHead, evenTail, nextNode *ListNode
	currentNode := head
	index := 0
	for currentNode != nil {
		nextNode = currentNode.Next
		if index == 0 || index%2 == 0 {
			if evenHead == nil {
				evenHead = currentNode
				evenTail = evenHead
			} else {
				evenTail.Next = currentNode
			}
			evenTail = currentNode
			evenTail.Next = nil
		} else {
			if oddHead == nil {
				oddHead = currentNode
				oddHead.Next = nil
			} else {
				oddTail.Next = currentNode
			}
			oddTail = currentNode
			oddTail.Next = nil
		}
		currentNode = nextNode
		index++
	}

	evenTail.Next = oddHead
	return evenHead
}

func (l *ListNode) IsEmpty() bool {
	return l == nil || l == &ListNode{}
}

func (l *ListNode) ToString() string {
	out := ""
	for l != nil {
		out += fmt.Sprintf("%v ", l)
		l = l.Next
	}
	return strings.Trim(out, " ")
}
