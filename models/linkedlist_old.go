package models

import "fmt"

/**
链表是一种物理存储单元上非连续、非顺序的存储结构，数据元素的逻辑顺序是通过链表中的指针链接次序实现的。
链表由一系列结点（链表中每一个元素称为结点）组成，结点可以在运行时动态生成。
每个结点包括两个部分：一个是存储数据元素的数据域，另一个是存储下一个结点地址的指针域。

链表和数组的区别：
两者的区别：
	数组静态分配内存，链表动态分配内存。
	数组在内存中是连续的，链表是不连续的。
	数组利用下标定位，查找的时间复杂度是O(1)，链表通过遍历定位元素，查找的时间复杂度是O(N)。
	数组插入和删除需要移动其他元素，时间复杂度是O(N)，链表的插入或删除不需要移动其他元素，时间复杂度是O(1)。
数组的优点:
	随机访问性比较强，可以通过下标进行快速定位。
	查找速度快
数组的缺点:
	插入和删除的效率低，需要移动其他元素。
	会造成内存的浪费，因为内存是连续的，所以在申请数组的时候就必须规定七内存的大小，如果不合适，就会造成内存的浪费。
	内存空间要求高，创建一个数组，必须要有足够的连续内存空间。
	数组的大小是固定的，在创建数组的时候就已经规定好，不能动态拓展。
链表的优点:
	插入和删除的效率高，只需要改变指针的指向就可以进行插入和删除。
	内存利用率高，不会浪费内存，可以使用内存中细小的不连续的空间，只有在需要的时候才去创建空间。大小不固定，拓展很灵活。
链表的缺点:
	查找的效率低，因为链表是从第一个节点向后遍历查找。
原文链接：https://blog.csdn.net/Shuffle_Ts/article/details/95055467

list与slice的区别：
如果频繁的插入和删除建议用list,频繁的遍历查询选slice。
由于container/list不是并发安全的，所以需要自己手动添加一层并发的包装
*/

/**
一.单链表基本概念

单链表是一种顺序存储的结构。
有一个头结点，没有值域，只有链域，专门存放第一个结点的地址。
有一个尾结点，有值域，也有链域，链域值始终为NULL。

所以，在单链表中为找第i个结点或数据元素，必须先找到第i - 1 结点或数据元素，而且必须知道头结点，否者整个链表无法访问。

本文主要通过Golang实现链表的几种常见操作:
1、判断是否为空的单链表
2、单链表的长度
3、从头部添加元素
4、从尾部添加元素
5、在指定位置添加元素
6、删除指定元素
7、删除指定位置的元素
8、查看是否包含某个元素
9、遍历所有元素
*/

type Object1 interface{}

type Node struct {
	Data Object1 //定义数据域
	Next *Node   //定义地址域(指向下一个表的地址)
}

type List struct {
	headNode *Node //头节点
}

/**
1、判断链表是否为空:判断单链表是否为空，只需要判断头节点是否为空即可
*/
func (l *List) IsEmpty() bool {
	if l.headNode == nil {
		return true
	} else {
		return false
	}
}

/**
2、获取列表长度
*/
func (l *List) Length() int {
	cur := l.headNode //获取链表的头节点
	count := 0
	for cur != nil { //for循环
		count++        //如果头节点不为空，则++
		cur = cur.Next //逐渐位移
	}
	return count
}

/**
3、从链表头部添加元素
*/
func (l *List) Add(data Object) *Node {

	node := &Node{Data: data}
	node.Next = l.headNode
	l.headNode = node
	return node
}

/**
4、从链表尾部添加元素
*/
func (l *List) Append(data Object) {
	node := &Node{Data: data}
	if l.IsEmpty() { //如果头节点为空，则直接插入
		l.headNode = node
	} else {
		cur := l.headNode
		for cur.Next != nil { //判断是否为尾节点，如果为null就是尾节点
			cur = cur.Next
		}
		cur.Next = node //此时是尾节点，将地址指针指向新创建的node
	}
	return
}

/**
5、在链表指定位置添加元素
*/
func (l *List) Insert(index int, data Object) {
	node := &Node{Data: data}
	if index < 0 {
		l.Add(data)
	} else if index > l.Length() {
		l.Append(data)
	} else {
		cur := l.headNode
		count := 0
		for count < index-1 {
			cur = cur.Next
			count++
		}
		node.Next = cur.Next
		cur.Next = node
	}
}

/**
6、删除链表指定值的元素
删除链表中的节点 (leetcode 237)
*/
func (l *List) Del(data Object) {
	cur := l.headNode
	if cur.Data == data {
		cur = cur.Next //如果删除的是头节点，那么第二个节点就是头节点
	} else {
		for cur.Next != nil { //判断是否是最后一个节点，不是继续循环
			if cur.Next.Data == data { // 如果cur.Next的data等于data,则删除节点，cur.Next指向之后节点的地址
				cur.Next = cur.Next.Next
			} else { //进行节点位移
				cur = cur.Next
			}
		}
	}
	return
}

/**
7、删除指定位置的元素
*/
func (l *List) DelAtIndex(index int) {
	cur := l.headNode
	if index <= 0 {
		cur = cur.Next //删除头节点，第二个接地那作为头节点
	} else if index > l.Length() {
		fmt.Println("超出")
		return
	} else {
		count := 0
		for count != (index-1) && cur.Next != nil {
			cur = cur.Next
			count++
		}
		cur.Next = cur.Next.Next
	}
	return
}

/**
8、查看链表是否包含某个元素
*/
func (l *List) Contain(data Object) bool {
	cur := l.headNode
	for cur != nil {
		if cur.Data == data {
			return true
		}
		cur = cur.Next
	}
	return false
}

/**
9、遍历链表所有节点
*/
func (l *List) ShowList() {
	if l.IsEmpty() {
		fmt.Println("没有任何数据")
		return
	} else {
		cur := l.headNode
		for {
			fmt.Println("showList :", cur.Data)
			if cur.Next != nil {
				cur = cur.Next
			} else {
				break
			}

		}
	}
	return
}

/**
10、从尾到头打印链表
*/
func (l *List) Reverse() []Object {
	count := 0
	res := make([]Object, count)
	if l.IsEmpty() {
		fmt.Println("链表为空")
		return res
	} else {
		cur := l.headNode
		newHead := cur
		i := 0
		for cur != nil {
			cur = cur.Next
			count++
		}
		res := make([]Object, count)
		for newHead != nil {
			res[count-i-1] = newHead.Data
			i++
			newHead = newHead.Next
		}

		return res
	}

}

/**
11、单链表反转 (leetcode 206)，返回的是节点
输入: 1->2->3->4->5->NULL
输出: 5->4->3->2->1->NULL
*/
func (l *List) Reverse1() *Node {
	cur := l.headNode
	var pre *Node = nil
	for cur != nil {
		pre, cur, cur.Next = cur, cur.Next, pre //这句话最重要
	}

	for pre != nil {
		fmt.Println("val:", pre.Data)
		pre = pre.Next
	}

	return pre

}

/**
12、截取出单链表中的后K个结点（k>0） 时间复杂度O（N）
*/
func (l *List) Trunc(num int) *Node {
	cur := l.headNode
	newCur := cur
	count := 0
	len := l.Length()
	for cur != nil {
		count++

		if count > len-num {
			//fmt.Println("count:",count,",len-num:",len-num,"cur.data:",cur.Data)
			newCur = cur
			fmt.Println("data111:", newCur.Data)
		}
		cur = cur.Next
	}
	//12345  45  len-num = 3
	for newCur != nil {
		fmt.Println("newCur :", newCur.Data)
		newCur = newCur.Next
	}
	return newCur
}

/**
12、截取出单链表中的后K个结点（k>0） --- 推荐此做法

2个指针一起，前个指针比这个指针快k,然后2个一起走
*/
func (l *List) Trunc1(k int) *Node {
	cur := l.headNode
	trunc := l.headNode

	for k >= 1 && cur != nil {
		cur = cur.Next
		k--
	}
	for cur != nil {
		cur = cur.Next
		trunc = trunc.Next
	}

	for trunc != nil {
		fmt.Println("trunc :", trunc.Data)
		trunc = trunc.Next
	}
	return trunc
}

/**
13、判断一个单链表中是否有环
思路：
如果一个链表中有环，也就是说用一个指针去遍历，是永远走不到头的。因此，我们可以用两个指针去遍历，
一个指针一次走两步，一个指针一次走一步，如果有环，两个指针肯定会在环中相遇。时间复杂度为O(n)。
*/
func (l *List) IsRound() bool {
	fast := l.headNode
	slow := l.headNode

	for fast != nil && fast.Next != nil {
		fast = fast.Next.Next
		slow = slow.Next
		if fast == slow {
			return true
		}
	}
	return false
}

/**
14、去除有序链表中的重复元素(leetcode 83)
1->1->2   ----- 1->2
1->2->2->3 -----1->2->3
https://blog.csdn.net/qq_35621006/article/details/104740484
*/
func (l *List) RemoveSame() *Node {
	cur := l.headNode
	if cur == nil {
		return nil
	}
	for cur.Next != nil {
		if cur.Data == cur.Next.Data {
			cur.Next = cur.Next.Next
		} else {
			cur = cur.Next
		}
	}
	return cur

}

/**
15、合并两个排好序的链表(leetcode 21)
https://www.cnblogs.com/TimLiuDream/p/10140132.html

将两个有序链表合并为一个新的有序链表并返回, 新链表是通过拼接给定的两个链表的所有节点组成的。

使用递归
*/
func Merge(n1 *Node, n2 *Node) *Node {
	if n1 == nil {
		return n2
	}
	if n2 == nil {
		return n1
	}
	var res *Node
	if n1.Data.(int) <= n2.Data.(int) {
		res = n1
		res.Next = Merge(n1.Next, n2)
	} else {
		res = n2
		res.Next = Merge(n1, n2.Next)
	}

	for {
		fmt.Println("showList_merge :", res.Data)
		if res.Next != nil {
			res = res.Next
		} else {
			break
		}

	}
	return res
}

/**
16、链表的中间节点(leetcode 876)
请编写一个函数，使其可以删除某个链表中给定的（非末尾）节点，你将只被给定要求被删除的节点。

输入: head = [4,5,1,9], node = 5
输出: [4,1,9]
解释: 给定你链表中值为 5 的第二个节点，那么在调用了你的函数之后，该链表应变为 4 -> 1 -> 9
*/

func (l *List) DelValue(data Object) *Node {
	cur := l.headNode
	for cur.Next != nil {
		if cur.Next.Data == data {
			cur.Next = cur.Next.Next
		} else {
			cur = cur.Next
		}
	}
	return cur
}

/**
17、删除链表的倒数第N个节点(leetcode 19)
给定一个链表: 1->2->3->4->5, 和 n = 2.
当删除了倒数第二个节点后，链表变为 1->2->3->5.
*/
func (l *List) DelByNum(index int) *Node {
	cur := l.headNode
	len := l.Length()
	delIndex := len - index
	count := 0
	for cur.Next != nil {
		count++
		fmt.Println("count:", count, ",delIndex:", delIndex, ",data:", cur.Data)
		if count == delIndex {
			fmt.Println("count111:", count, ",delIndex:", delIndex, ",data:", cur.Data)
			cur.Next = cur.Next.Next
		} else {
			cur = cur.Next
		}

	}
	return cur

}

/**
17、删除链表中的节点(leetcode 203)
题目
删除链表中等于给定值 val 的所有节点。

示例
输入: 1->2->6->3->4->5->6, val = 6
输出: 1->2->3->4->5

链表里面的元素删除，其实就是指针指向下一个元素就ok了
*/
func (l *List) DelAllValue(data Object) *Node {
	cur := l.headNode
	if cur == nil {
		return nil
	}
	for cur.Next != nil {
		if cur.Next.Data == data {
			cur.Next = cur.Next.Next
		} else {
			cur = cur.Next
		}
	}
	return cur
}

func ListDemo() {
	list := List{}
	list.Add("4")
	list.Append("1")
	list.Append("5")
	list.Append("9")

	/*	n1 := &Node{
			Data: 4,
		}
		n1.Next = &Node{
			Data: 8,
		}
		n1.Next = &Node{
			Data: 12,
		}
		n1.Next = &Node{
			Data: 16,
		}
		n2 := &Node{
			Data: 3,
		}
		n2.Next = &Node{
			Data: 9,
		}
		n2.Next = &Node{
			Data: 12,
		}

		n2.Next = &Node{
			Data: 15,
		}

		Merge(n1, n2)*/

	//list.ShowList()

	fmt.Println("截取")
	list.Trunc(2)
	fmt.Println("截取2")
	list.Trunc1(2)

	fmt.Println("reverse1:", list.Reverse1())
	list.ShowList()

	/*Reverse2(list.headNode)
	fmt.Println("11")
	list.ShowList()
	*/
	/*

		fmt.Println("flag:", list.IsRound())

		list.RemoveSame()
		list.ShowList()*/

	//list.DelAllValue("3")
	/*list.DelByNum(2)
	list.ShowList()*/

	/*list.DelValue("5")
	list.ShowList()*/
	/*list.Append("4")



	list.ShowList()
	fmt.Println("len_0:",list.Length())
	list.Insert(2,"nihao")

	list.ShowList()
	fmt.Println("len_1:",list.Length())

	list.Del("3")
	list.ShowList()
	fmt.Println("len_2:",list.Length())

	fmt.Println("contain:",list.Contain("4"))

	list.DelAtIndex(1)
	list.ShowList()
	fmt.Println("len_3:",list.Length())

	list.Insert(1,"China")
	list.ShowList()
	fmt.Println("len_4:",list.Length())*/

	//res := list.Reverse()
	//fmt.Println("结果:", res)

}
