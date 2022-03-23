package models

import "fmt"

/**
一、链表
https://goa.lenggirl.com/algorithm/link.html
讲数据结构就离不开讲链表。因为数据结构是用来组织数据的，
如何将一个数据关联到另外一个数据呢？链表可以将数据和数据之间关联起来，从一个数据指向另外一个数据。

定义：
链表由一个个数据节点组成的，它是一个递归结构，要么它是空的，要么它存在一个指向另外一个数据节点的引用。
链表，可以说是最基础的数据结构。
*/

type LinkedNode1 struct {
	Data     int64        //存放数据
	nextNode *LinkedNode1 //指向下一个节点，这种从一个数据节点指向下一个数据节点的结构，都可以叫做链表。
}

func LinkedDome() {
	//新的节点
	node := new(LinkedNode1)
	node.Data = 1

	//新的节点
	node1 := new(LinkedNode1)
	node1.Data = 2
	node.nextNode = node1

	//新的节点
	node2 := new(LinkedNode1)
	node2.Data = 3
	node1.nextNode = node2

	// 按顺序打印数据
	newNode := node
	for {
		if newNode != nil {
			// 打印节点值
			fmt.Println(newNode.Data)
			// 获取下一个节点
			newNode = newNode.nextNode
		} else {
			// 如果下一个节点为空，表示链表结束了
			break
		}
	}
}

type Ring struct {
	Pre, next *Ring
	Data      int64
}

//初始化空的循环链表，前驱和后驱都指向自己，因为是循环的
func (r *Ring) init() *Ring {
	r.Pre = r
	r.next = r
	return r
}

//创建一个指定大小 N 的循环链表，值全为空：
func New(n int) *Ring {
	r := new(Ring)
	p := r
	for i := 1; i < n; i++ {
		p.next = &Ring{Pre: p}
		p = p.next
	}
	p.next = r
	r.Pre = p
	return r
}

//获取上一个或下一个节点
func (r *Ring) Next() *Ring {
	if r.next == nil {
		return r.init()
	}
	return r.next
}

func (r *Ring) Prev() *Ring {
	if r.Pre == nil {
		return r.init()
	}
	return r.Pre
}

//获取第 n 个节点
//因为链表是循环的，当 n 为负数，表示从前面往前遍历，否则往后面遍历：
func (r *Ring) Move(n int) *Ring {
	if r.next == nil {
		return r.init()
	}
	switch {
	case n < 0:
		for ; n < 0; n++ {
			r = r.Pre
		}
	case n > 0:
		for ; n > 0; n-- {
			r = r.next
		}
	}
	return r
}

//获取链表长度
func (r *Ring) Len() int {
	n := 0
	if r != nil {
		n = 1
		for p := r.next; p != r; p = p.next {
			n++
		}
	}

	return n
}

// 往节点A，链接一个节点，并且返回之前节点A的后驱节点
func (r *Ring) Link(s *Ring) *Ring {
	n := r.Next()
	if s != nil {
		p := s.Prev()
		r.next = s
		s.Pre = r
		n.Pre = p
		p.next = n
	}
	return n
}

// 删除节点后面的 n 个节点
func (r *Ring) Unlink(n int) *Ring {
	if n < 0 {
		return nil
	}
	return r.Link(r.Move(n + 1))
}

func RingDemo() {
	r := new(Ring)
	r.init() //因为绑定前驱和后驱节点为自己，没有循环，时间复杂度为：O(1)。

	New(2) //会连续绑定前驱和后驱节点，时间复杂度为：O(n)。

	r.Next()
	r.Prev() //获取前驱或后驱节点，时间复杂度为：O(1)。

	r.Move(2)  //因为需要遍历 n 次，所以时间复杂度为：O(n)。
	r.Move(-2) //当 n 为负数，表示从前面往前遍历，否则往后面遍历：

	r.Len() //通过循环，当引用回到自己，那么计数完毕，时间复杂度：O(n)。

}
