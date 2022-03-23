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

type Object interface{}

type LinkedNode struct {
	Data     Object      //定义数据域
	NextNode *LinkedNode //定义地址域(指向下一个表的地址)
}

type HeadNode struct {
	headNode *LinkedNode //头节点
}

/**
1、判断链表是否为空:判断单链表是否为空，只需要判断头节点是否为空即可
*/
func (headNode *HeadNode) IsEmpty() bool {
	if headNode.headNode == nil {
		return true
	}
	return false
}

/**
2、获取列表长度
*/
func (headNode *HeadNode) Length() int {
	cur := headNode.headNode //获取链表的头节点
	count := 0
	for cur != nil {
		cur = cur.NextNode //逐渐位移
		count++            //如果头节点不为空，则++
	}
	return count
}

/**
3、从链表头部添加元素
*/
func (headNode *HeadNode) Add(data Object) *LinkedNode {
	node := &LinkedNode{Data: data}
	node.NextNode = headNode.headNode
	headNode.headNode = node
	return node
}

/**
4、从链表尾部添加元素
*/
func (headNode *HeadNode) Append(data Object) {
	cur := headNode.headNode
	for cur.NextNode != nil { //判断是否为尾节点，如果为null就是尾节点
		cur = cur.NextNode
	}
	node := &LinkedNode{Data: data}
	cur.NextNode = node //此时是尾节点，将地址指针指向新创建的node
	return
}

/**
5、在链表指定位置添加元素
*/
func (headNode *HeadNode) Insert(index int, data Object) {
	node := &LinkedNode{Data: data}
	if index < 0 {
		headNode.Add(data)
	} else if index > headNode.Length() {
		headNode.Append(data)
	} else {
		cur := headNode.headNode
		count := 0
		for count < index-1 {
			count++
			cur = cur.NextNode
		}
		node.NextNode = cur.NextNode
		cur.NextNode = node
	}
}

/**
6、删除指定位置的元素
*/
func (headNode *HeadNode) DelAtIndex(index int) {
	if index < 0 || index > headNode.Length() {
		fmt.Println("index有误")
		return
	}
	cur := headNode.headNode
	count := 0
	for count < index-1 {
		count++
		cur = cur.NextNode
	}
	cur.NextNode = cur.NextNode.NextNode
}

/**
7、删除链表指定值的元素
删除链表中的节点 (leetcode 237)
*/
func (headNode *HeadNode) Del(data Object) {
	cur := headNode.headNode
	if cur.Data == data { //如果删除的是头节点，那么第二个节点就是头节点
		cur = cur.NextNode
	} else {
		for cur.NextNode != nil { //判断是否是最后一个节点，不是继续循环
			if cur.Data == data {
				cur.NextNode = cur.NextNode.NextNode // 如果cur.Next的data等于data,则删除节点，cur.Next指向之后节点的地址
			} else {
				cur = cur.NextNode //进行节点位移
			}
		}
	}
	return
}

/**
8、查看链表是否包含某个元素
*/
func (headNode *HeadNode) Contain(data Object) bool {
	cur := headNode.headNode
	for cur.NextNode != nil {
		if cur.Data == data {
			return true
		} else {
			cur = cur.NextNode
		}
	}
	return false
}

/**
9、遍历链表所有节点
*/
