package main

import "fmt"

type DNode struct {
	data int
	next *DNode
	prev *DNode
}

type DoubleLinkList struct {
	length int
	head   *DNode
	tail   *DNode
}

func CreateDNode(data int) *DNode {
	return &DNode{
		data,
		nil,
		nil,
	}
}

func CreateLinkList() *DoubleLinkList {
	return &DoubleLinkList{
		0,
		&DNode{
			0,
			nil,
			nil,
		},
		&DNode{
			0,
			nil,
			nil,
		},
	}
}

//初始化链表
func (l *DoubleLinkList) InitList() {
	l.length = 0
	l.head = nil
	l.tail = nil
}

//获得链表长度
func (l *DoubleLinkList) GetListLength() int {
	return l.length
}

//正向打印链表
func (l *DoubleLinkList) PrintList() {
	pre := l.head.next
	for i := 0; i < l.length; i++ {
		fmt.Printf("%d ", pre.data)
		pre = pre.next
	}
	fmt.Printf("\n")
}

//反向打印链表
func (l *DoubleLinkList) PrintListReverse() {
	pre := l.tail.prev
	for i := 0; i < l.length; i++ {
		fmt.Printf("%d ", pre.data)
		pre = pre.prev
	}
	fmt.Printf("\n")
}

//尾插法创建链表
func (l *DoubleLinkList) CreateDoubleLinkList(arr []int, n int) {
	p := l.head
	for i := 0; i < n; i++ {
		r := CreateDNode(arr[i])
		p.next = r
		r.prev = p
		p = r
	}
	p.next = l.tail
	l.tail.prev = p
	l.length = n
}

//获得链表中的值的索引值
func (l *DoubleLinkList) LocateElem(data int) int {
	pre := l.head.next
	indexOfData := 0
	for pre != l.tail && pre.data != data {
		pre = pre.next
		indexOfData++
	}
	if pre == l.tail {
		return -1
	} else {
		return indexOfData
	}
}

//取得索引值为index的节点的值
func (l *DoubleLinkList) GetElem(index int) int {
	i := 0
	pre := l.head.next
	for pre != nil && i < index {
		pre = pre.next
		i++
	}
	if pre == l.tail {
		return -1
	} else {
		return pre.data
	}
}

//插入节点
func (l *DoubleLinkList) InsertList(index, data int) bool {
	if index < 1 || index > l.length {
		return false
	}
	s := CreateDNode(data)
	p := l.head.next
	for i := 0; i < index; i++ {
		p = p.next
		i++
	}
	s.next = p.next
	s.next.prev = s
	p.next = s
	s.prev = p
	l.length++
	return true
}

//删除节点
func (l *DoubleLinkList) DeleteList(index int) bool {
	if index < 1 || index > l.length {
		return false
	}
	p := l.head.next
	for i := 0; i < index; i++ {
		p = p.next
		i++
	}
	s := p.next
	p.next = s.next
	s.next.prev = p
	p = nil
	l.length--
	return true
}

func main() {
	arr := []int{1, 2, 3, 4, 5, 6, 7, 8}
	list := CreateLinkList()
	list.CreateDoubleLinkList(arr, 8)
	fmt.Printf("%d\n", list.GetListLength())
	list.PrintList()
	list.PrintListReverse()
	fmt.Printf("%d\n", list.LocateElem(6))
	fmt.Printf("%d\n", list.GetElem(3))
	list.InsertList(2, 19)
	list.PrintList()
	list.DeleteList(2)
	list.PrintList()
}
