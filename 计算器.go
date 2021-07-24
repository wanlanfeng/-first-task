package main

import (
	"fmt"
	"strconv"
)
//定义结构体栈
type stack struct{
	exp []string
}
//入栈，栈中的数据
func (s *stack) Push(i string)  {
	s.exp = append(s.exp, i)
}
//出栈
func (s *stack) Pop() string {
	if len(s.exp) != 0{
		exp := s.exp[len(s.exp) - 1]
		s.exp = s.exp[0:len(s.exp) - 1]
		return exp
	}
	return "false"
}
//获取字符串表达式
func Getexpression() string{
	var exp string
	fmt.Scan(&exp)
	return exp
}
//判断是否是数字
func Isdigit(i int) bool{
	if i <= 57 && i >= 48{
		return true
	}
	return false
}
//判断栈是否为空
func (s *stack)Isempty() bool{
	return len(s.exp) == 0
}
//获取栈顶元素
func (s *stack)Top() string{
	if len(s.exp) != 0{
		return s.exp[len(s.exp) - 1]
	}
	return "failed"
}
//将中缀表达式改为后缀表达式
func Changeform(exp string) string{
	s := stack{}				//用来存储计算符号
	var newexp string		//用来存放后缀表达式
	for i := 0; i < len(exp); i++{
		_, err := strconv.Atoi(string(exp[i]))
		if err == nil{
			newexp += string(exp[i])
		}else{
			if s.Isempty(){
				s.Push(string(exp[i]))
			}else {
				switch string(exp[i]){
				case "+", "-":
					for !s.Isempty() && s.Top() != "("{
						k := s.Pop()
						newexp += k
					}
					s.Push(string(exp[i]))
				case "*", "/":
					for {
						if ( s.Top() != "+" && s.Top() != "-" && s.Top()!= "(" ) && !s.Isempty() {
							k := s.Pop()
							newexp += k
						}else {
							break
						}
					}
					s.Push(string(exp[i]))
				case "(":
					s.Push(string(exp[i]))
				case ")":
					for {
						if s.Top() == "("{
							break
						}else {
							k := s.Pop()
							newexp += k
						}
					}
					s.Pop()
					}
				}
			}
		}
		for ; !s.Isempty(); {
			newexp += s.Pop()
		}
		//fmt.Println(s.exp)
		return newexp
	}
//计算后缀表达式
func  Calculation(exp string) int{
	s := stack{}
	for i := 0; i < len(exp); i++{
		if Isdigit(int(exp[i])) {
			s.Push(string(exp[i]))
		}else{
			n1, _ := strconv.Atoi(s.Pop())
			n2, _ := strconv.Atoi(s.Pop())
			switch string(exp[i]){
			case "+":
				s.Push(strconv.Itoa(n2+n1))
			case "-":
				s.Push(strconv.Itoa(n2-n1))
				fmt.Println()
			case "*":
				s.Push(strconv.Itoa(n2*n1))
			case "/":
				s.Push(strconv.Itoa(n2/n1))
			}
		}
	}
	res, _ := strconv.Atoi(s.Pop())
	return res
}
func main(){
	str := Getexpression()		//获取字符串形式的表达式
	exp := Changeform(str)		//将中缀表达式转为后缀
	res := Calculation(exp)		//计算后缀表达式的值
	fmt.Println(exp)
	fmt.Println(res)
	return
}
