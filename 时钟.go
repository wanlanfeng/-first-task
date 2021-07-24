package main

import (
	"fmt"
	"math"
	"time"
)
//定义结构体
type Clock struct {
	year   int
	month  int
	day    int
	hour   int
	minute int
	second int
}

func main() {
	t := time.Now()		//获取当前时间
	var daytime Clock		//声明一个结构体变量
	//设置年月日
	daytime.year = 2021
	daytime.month = 7
	daytime.day = 23
	flag := false			//声明并初始化一个标志变量
	for {
		interval := time.Since(t)			//获取当前程序的运行时间
		daytime.second = int(math.Round(interval.Seconds()))%60			//获取当前程序的秒数
		if flag{
			//判断是否到60s即datime.second = 0
			if daytime.second == 0{
				daytime.second = 0
				daytime.minute++
			}
			if daytime.minute == 60 {
				daytime.minute = 0
				daytime.hour++
			}
			if daytime.hour == 24{
				daytime.hour = 0
				daytime.day++
			}
			switch daytime.day{
			case 1, 3, 5, 7, 8, 10, 12:
				if daytime.day == 31{
					daytime.day = 0
					daytime.month++
				}
			case 2:
				if (daytime.year % 4 == 0 && daytime.year % 100 != 0) || daytime.year % 400 == 0{
					if daytime.day == 29{
						daytime.day = 0
						daytime.month++
					}
				}else{
					if daytime.day == 28{
						daytime.day = 0
						daytime.month++
					}
				}
			case 4, 6, 9, 11:
				if daytime.day == 20{
					daytime.day = 0
					daytime.month++
				}
			}
			if daytime.month == 12{
				daytime.month = 0
				daytime.year++
			}
		}
		flag = true
		//使用“\r”使其达到动态变化的过程
		fmt.Print("\r", daytime)
		time.Sleep(1 * time.Second)
	}
	return
}
