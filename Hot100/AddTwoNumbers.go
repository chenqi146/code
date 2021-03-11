package main

type ListNode struct {
	Val  int
	Next *ListNode
}

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {

	temp := 0
	l := new(ListNode)
	t := l
	for true {
		if l1 == nil {
			l1 = new(ListNode)
			l1.Val = 0
		}
		if l2 == nil {
			l2 = new(ListNode)
			l2.Val = 0
		}
		i := l1.Val + l2.Val
		if temp >= 10 {
			i += 1
		}
		temp = i
		l.Val = i % 10
		l1 = l1.Next
		l2 = l2.Next
		if temp >= 10 && l2 == nil {
			l2 = new(ListNode)
			l2.Val = 0
		}
		if temp >= 10 && l1 == nil {
			l1 = new(ListNode)
			l1.Val = 0
		}
		if l1 == nil && l2 == nil {
			break
		}
		l.Next = new(ListNode)
		l = l.Next
	}

	return t
}

func addTwoNumbers1(l1 *ListNode, l2 *ListNode) *ListNode {

	// 记录上一个节点的相加余下的值
	temp := 0
	// 哨兵节点
	l := new(ListNode)
	// 需要操作的节点
	t := l
	for l1 != nil || l2 != nil || temp != 0 {

		l.Next = new(ListNode)
		l = l.Next
		if l1 != nil {
			temp += l1.Val
			l1 = l1.Next
		}

		if l2 != nil {
			temp += l2.Val
			l2 = l2.Next
		}

		l.Val = temp % 10
		temp /= 10

	}

	return t.Next
}

func main() {

	node := new(ListNode)
	listNode := new(ListNode)
	t := new(ListNode)
	dd := new(ListNode)
	dd1 := new(ListNode)
	dd2 := new(ListNode)
	dd3 := new(ListNode)
	dd3.Val = 9
	dd2.Val = 9
	dd2.Next = dd3
	dd1.Val = 9
	dd1.Next = dd2
	dd.Val = 9
	dd.Next = dd1
	node.Val = 9
	node.Next = dd
	listNode.Val = 9
	listNode.Next = node
	t.Val = 9
	t.Next = listNode

	node2 := new(ListNode)
	listNode2 := new(ListNode)
	t2 := new(ListNode)
	ddw := new(ListNode)
	ddw.Val = 9
	node2.Val = 9
	node2.Next = ddw
	listNode2.Val = 9
	listNode2.Next = node2
	t2.Val = 9
	t2.Next = listNode2

	println(addTwoNumbers1(t, t2))

}
