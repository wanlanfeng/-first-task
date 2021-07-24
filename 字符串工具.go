package main

import (
	"fmt"
	"strconv"
)

type String struct{
	num int
	s [8]string
}
func Part(s1 string) String{
	i, j := 0, 0
	var str String
	str.num = 0
	for {
		i = j
		for {
			if j == len(s1) || string(s1[j]) == "," {
				break
			}
			j++
		}
		if j < len(s1){
			str.s[str.num] = s1[i: j]
		}else {
			str.s[str.num] = s1[i:]
			break
		}
		str.num++
		j++
	}
	return str
}
//判断是否全是数字
func Isalnum(s string) bool {
	for _, v := range s{
		if v > 57 || v < 48{
			return false
		}
	}
	return true
}
//判断是否全是大写字母
func Isupper(s string) bool{
	for _, v := range s{
		if v > 90 || v < 65{
			return false
		}
	}
	return true
}
//判断是否全是小写字母
func Islower(s string) bool{
	for _, v := range s{
		if v > 122 || v < 97{
			return false
		}
	}
	return true
}
//判断是否全是字母
func Isalpha(s string) bool{
	for _, v := range s{
		if v > 122 || (v < 97 && v > 90) || v < 65{
			return false
		}
	}
	return true
}
//将字母转为大写
func Upper(s string) string{
	var a string
	for _, v := range s{
		if v <= 122 && v >= 97{
			a += string(v - 32)
		}else{
			a += string(v)
		}
	}
	return a
}
//字符串转为int数组
func Change(s string) []int{
	var a []int
	for _, v := range s{
		k, _ :=  strconv.Atoi(string(v))
		a = append(a, k)
	}
	return a
}
//冒泡排序
func Bubsort(a []int) []int{
	for i := 1; i < len(a); i++{
		for j := 0; j < len(a) - i; j++{
			if a[j] > a[j + 1]{
				a[j + 1], a[j] = a[j], a[j + 1]
			}
		}
	}
	return a
}
//快速排序
func Quicksort(a []int, left, right int){
	if left < right{
		i, j := left, right
		x := a[left]
		for {
			for ; a[i] < x && i < j; i++{}
			for ; a[j] > x         ; j--{}
			if i >= j{
				break
			}
			a[i], a[j] = a[j], a[i]
		}
		a[left], a[j] = a[j], x
		Quicksort(a, left, j - 1)
		Quicksort(a, j + 1, right)
	}
	return
}
func main(){
	var str String
	s1 := "acbdw,1269547,AASIDX,AIUydjs,12sjaA,3819247,ausSHSzio,IUFISsi"
	str = Part(s1)
	fmt.Println("判断是否全是数字：")
	for i := 0; i <= str.num; i++{
		fmt.Print(" ", Isalnum(str.s[i]))
	}
	fmt.Println("\n判断是否全是大写字母：")
	for i := 0; i <= str.num; i++{
		fmt.Print(" ", Isupper(str.s[i]))
	}
	fmt.Println("\n判断是否全是小写字母：")
	for i := 0; i <= str.num; i++{
		fmt.Print(" ", Islower(str.s[i]))
	}
	fmt.Println("\n判断是否全是字母并将其大写化：")
	for i := 0; i <= str.num; i++{
		if Isalpha(str.s[i]){
			fmt.Println(Upper(str.s[i]))
		}
	}
	fmt.Println("\n判断是否全是数字并将其从小到大排列：")
	for i := 0; i <= str.num; i++{
		if Isalnum(str.s[i]){
			a := Change(str.s[i])
			fmt.Println(Bubsort(a))
			Quicksort(a, 0, 6)
			fmt.Println(a)
		}
	}
	return
}
