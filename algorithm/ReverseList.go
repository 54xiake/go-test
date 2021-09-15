package main

// go实现单链表翻转
import "fmt"

type ListNode struct {
	data interface{}
	Next *ListNode
}

//反转单链表
func reverseList(head *ListNode) *ListNode {
	cur := head
	var pre *ListNode
	for cur != nil {
		cur.Next, pre, cur = pre, cur, cur.Next
	}
	return pre
}

func CreateNode(node *ListNode, max int) {
	cur := node
	for i := 1; i < max; i++ {
		cur.Next = &ListNode{}
		cur.Next.data = i
		cur = cur.Next
	}
}

// 打印链表的方法
func PrintNode(info string, node *ListNode) {
	fmt.Print(info)
	for cur := node; cur != nil; cur = cur.Next {
		fmt.Print(cur.data, " ")
	}
	fmt.Println()
}

func main() {
	var head = new(ListNode)
	CreateNode(head, 10)
	PrintNode("前：", head)
	yyy := reverseList(head)
	PrintNode("后：", yyy)

}
