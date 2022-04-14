package main

import (
	"fmt"
	"time"
)

func foo() (int, string) {
	return 10, "Q1mi"
}

const (
	_  = iota             //0
	KB = 1 << (10 * iota) // 1
	MB = 1 << (10 * iota)
	GB = 1 << (10 * iota)
	TB = 1 << (10 * iota)
	PB = 1 << (10 * iota)
)
const (
	a, b = iota + 1, iota + 2 //1,2
	c, d                      //2,3
	e, f                      //3,4
)

func main() {
	pTime()
}

/**
获取年 月 日 时 分 秒 纳秒
*/
func pTime() {
	dateTime := time.Now()
	fmt.Println("获取当前时间=", dateTime)
	fmt.Println("====================")
	year := time.Now().Year() //年
	fmt.Println(year)
	month := time.Now().Month() //月
	fmt.Println(month)
	day := time.Now().Day() //日
	fmt.Println(day)
	hour := time.Now().Hour() //小时
	fmt.Println(hour)
	minute := time.Now().Minute() //分钟
	fmt.Println(minute)
	second := time.Now().Second() //秒
	fmt.Println(second)
	nanosecond := time.Now().Nanosecond() //纳秒
	fmt.Println(nanosecond)
}

/**
获取当前时间戳
*/
func getTimeUnix() {

	fmt.Println("=====获取当前时间戳=======")
	timeUnix := time.Now().Unix()         //单位秒
	timeUnixNano := time.Now().UnixNano() //单位纳秒
	fmt.Println(timeUnix)
	fmt.Println(timeUnixNano)

}
func formatTimeToUnix() {
	//时间格式转时间戳
	fmt.Println(time.Now().Format("2006-01-02 15:04:05"))
	//时间戳转为go格式的时间

	var timeUnix int64 = 1562555859
	fmt.Println(time.Unix(timeUnix, 0)) // 之后可以用Format 比如
	fmt.Println(time.Unix(timeUnix, 0).Format("2006-01-02 15:04:05"))

}
func formatTimes() {
	//str格式化时间转时间戳

	t := time.Date(2014, 1, 7, 5, 50, 4, 0, time.Local).Unix()
	fmt.Println(t)

}
func getCurTime() {
	//获取今天0点0时0分的时间戳
	currentTime := time.Now()
	startTime := time.Date(currentTime.Year(), currentTime.Month(), currentTime.Day(), 0, 0, 0, 0, currentTime.Location())
	fmt.Println(startTime)
	fmt.Println(startTime.Format("2006/01/02 15:04:05"))

	//02: 获取今天23:59:59秒的时间戳
	//currentTime := time.Now()
	//endTime := time.Date(currentTime.Year(), currentTime.Month(), currentTime.Day(), 23, 59, 59, 0, currentTime.Location())
	//fmt.Println(endTime)
	//fmt.Println(endTime.Format("2006/01/02 15:04:05"))

	//03: 获取1分钟之前的时间
	m, _ := time.ParseDuration("-1m")
	result := currentTime.Add(m)
	fmt.Println(result)
	fmt.Println(result.Format("2006/01/02 15:04:05"))

	//04: 获取1小时之前的时间

	//m, _ := time.ParseDuration("-1h")
	//result := currentTime.Add(m)
	//fmt.Println(result)
	//fmt.Println(result.Format("2006/01/02 15:04:05"))
	//05: 获取1分钟之后的时间

	//m, _ := time.ParseDuration("1m")
	//result := currentTime.Add(m)
	//fmt.Println(result)
	//fmt.Println(result.Format("2006/01/02 15:04:05"))

}

func subTime() {
	currentTime := time.Now()
	//07 :计算两个时间戳
	afterTime, _ := time.ParseDuration("1h")
	result := currentTime.Add(afterTime)
	beforeTime, _ := time.ParseDuration("-1h")
	result2 := currentTime.Add(beforeTime)
	m := result.Sub(result2)
	fmt.Printf("%v 分钟 \n", m.Minutes())
	h := result.Sub(result2)
	fmt.Printf("%v小时 \n", h.Hours())
	d := result.Sub(result2)
	fmt.Printf("%v 天\n", d.Hours()/24)

}
func tfTime() {
	//08: 判断一个时间是否在一个时间之后
	stringTime, _ := time.Parse("2006-01-02 15:04:05", "2019-12-12 12:00:00")
	beforeOrAfter := stringTime.After(time.Now())

	if true == beforeOrAfter {
		fmt.Println("2019-12-12 12:00:00在当前时间之后!")
	} else {
		fmt.Println("2019-12-12 12:00:00在当前时间之前!")
	}

}

func howlongTime() {
	//09: 判断一个时间相比另外一个时间过去了多久
	startTime := time.Now()
	time.Sleep(time.Second * 5)
	fmt.Println("离现在过去了：", time.Since(startTime))

}
