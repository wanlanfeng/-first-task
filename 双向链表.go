package main

import "fmt"
//定义链表结构体
type Linknode struct{
	num int
	id int
	pre *Linknode
	last *Linknode
}
//创建一个链表
func createNode() *Linknode{
	var head *Linknode		//声明头指针，但此时head = nil,并未分配地址
	p := &Linknode{1, 2009, nil, nil}		//声明结构体指针p的同时显式赋值
															//结构体临时指针的初始化只能使用显式赋值
	if head == nil{											//检验头指针为空，将临时指针赋给头指针
		head = p
	}
	return head
}
//查找节点
func (head *Linknode)findNode(i int) (a bool, b int){
	temhead := head											//声明临时指针
	for {													//无限循环
		if temhead.id == i {								//如果头指针的内容是需要查找的值
			return true, temhead.num
			break											//使用break退出循环
		}
		if temhead.last == nil{									//如果找到链表末尾仍未找到该内容，则退出循环
			break
		}
		temhead = temhead.last								//如果当前指针的内容不符合i,则
															// 通过循环不断将下一个节点的地址更新到temhead

	}
	return false, 0
}
//添加新节点
func (head *Linknode)addNode(i, j int) *Linknode{
	temhead := head
	p := &Linknode{j, i,nil, nil}				//显式初始化临时结构体指针变量
	//检验头指针是否为空
	if head == nil{												//头指针为空
		head = p
	} else{														//头指针不为空
		for {													//无限循环找到当前节点的下一个节点为空的节点
			if temhead.last == nil {
				break
			}
			temhead = temhead.last								//通过循环不断将下一个节点的地址更新到temhead
			}
		}
		temhead.last = p										//将p存到下一个节点
		p.pre = temhead											//将temhead赋给p的pre
	return head
}
//删除节点
func (head *Linknode)deleteNode(i int) *Linknode{
	//如果需要删除的数据在头指针
	if head.id == i{
		head.pre = nil
		head = head.last
		//数据不在头指针
	}else{
		temhead := head
		for{
			temhead = temhead.last
			if temhead.id == i{
				temhead.pre.last = temhead.last
				temhead.last.pre = temhead.pre
				break
			}
		}
	}
	return head
}
//修改某个结点的内容
func (head *Linknode)reviseNode(i, j int) *Linknode{
	temhead := head
	for {
		if temhead.id == i {		//找到节点后修改值并退出循环
			temhead.id = j
			break
		}
		if temhead.last == nil{		//找完所有节点并未找到相应的节点
			break
		}
		temhead = temhead.last
	}
	return head
}
//打印链表中所有节点
func (head *Linknode)printf() {
	temhead := head
	for {
		fmt.Println(temhead.num ,temhead.id)
		if temhead.last == nil{
			break
		}
		temhead = temhead.last
	}
}
//实现数组迭代器
func Arr(head *Linknode) []int{
	temhead := head
	var arr []int
	for {
		if temhead.last != nil{
			arr = append(arr, temhead.id)

		}else{
			break
		}
		temhead = temhead.last
	}
	return arr
}
func main(){
	head := createNode()							//创建链表
	for i, j := 2011, 2; i <= 2022; i++{			//通过循环添加节点
		head.addNode(i,j)
		j++
	}
	fmt.Println(head.findNode(2020))
	head.deleteNode(2020)						//删除节点
	fmt.Println(head.findNode(2020))
	head.reviseNode(2009, 2010)				//修改某个节点内的内容
	head.printf()									//打印节点内容
	arr := Arr(head)								//实现数组迭代器
	fmt.Println(arr)
	return
}
