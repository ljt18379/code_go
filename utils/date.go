package main

import (
	"fmt"
	"strconv"
	"time"
)

func main() {
	//timeFmtToTimeStr()
	//timeStrToTimeFmt()
	timeFmtToTimestamp()
	//timestampToTimeFmt()
}

const (
	TIMEFORMATSECOND = "2006-01-02 15:04:05"
	TIMEFORMATDAY    = "2006-01-02"
	INFLUXFMT        = "2006-01-02T15:04:05Z"
	layout           = "Jan 2, 2006 at 3:04pm (MST)"
	TIMEFORMAT       = "2023-04-13T15:30:00.000Z"
)

/**
标准时间格式
标准时间格式主要指的是国际标准ISO 8601和RFC 3339，它们被广泛用于表示日期和时间。12

ISO 8601的格式为YYYY-MM-DDTHH:mm:ss.sssZ，其中YYYY代表四位数年份，MM代表月份，DD代表天数，T作为日期和时间的分隔符，HH代表小时，mm代表分钟，ss.sss代表秒和毫秒，而Z代表时区。例如：2023-04-13T15:30:00.000Z表示的是2023年4月13日15时30分0秒的UTC时间。

RFC 3339是基于ISO 8601格式的扩展，它允许使用非零开头的小时、分钟和秒。例如：2023-04-13T15:30:00Z表示的是同样的时间，但省略了毫秒部分和时区偏移量。

此外，还有其他几种日期和时间表示格式，如RFC 2822（用于电子邮件）、Unix时间戳（表示自1970年1月1日00:00:00 UTC以来的秒数）、自然语言表示、中文日期表示、美国习惯用法、欧洲习惯用法等。每种格式都有其特定的用途和优势，选择哪种格式取决于具体的应用场景和需求。

在编程语言中，日期和时间库通常会提供工具来处理和编码这些不同的日期时间格式，以便于在不同的系统和应用程序之间进行数据交换和处理。
*/

// 时间格式->时间字符串
func timeFmtToTimeStr() {
	timeNow := time.Now()
	//timeFmt->timeStr
	timeStr := timeNow.Format(TIMEFORMATSECOND)
	timeStr2 := timeNow.Format(TIMEFORMATDAY)
	fmt.Println("时间格式: 2006-01-02 15:04:05\t当前时间:", timeStr)
	fmt.Println("时间格式: 2006-01-02\t当前时间:", timeStr2)

	timeStrAddDate := timeNow.AddDate(-1, 0, -2).Format("2006-01-02T00:00:00Z")
	fmt.Println("timeStrAddDate:", timeStrAddDate)

	dayStr := strconv.FormatInt(int64(timeNow.Day()), 10)
	fmt.Println("dayStr:", dayStr)

}

// 时间字符串->时间格式
func timeStrToTimeFmt() {
	timeStr := time.Now().Format(TIMEFORMATSECOND)
	//timeStr->timeFmt
	timeFmt, _ := time.Parse(TIMEFORMATSECOND, timeStr)
	fmt.Println("timeFmt", timeFmt)

	//UTC+8小时
	timeFmtCST, _ := time.ParseInLocation(TIMEFORMATSECOND, timeStr, time.Local)
	fmt.Println("timeFmtCST", timeFmtCST)
}

// 时间格式->时间戳格式
func timeFmtToTimestamp() {
	timeStr := time.Now().Format(TIMEFORMATSECOND)
	timeFmt, _ := time.Parse(TIMEFORMATSECOND, timeStr)
	fmt.Println("timeStr", timeStr)
	fmt.Println("timeFmt", timeFmt)
	//timeFmt->timestamp
	secTs := timeFmt.Unix()
	nanoTs := timeFmt.UnixNano()
	fmt.Println("secTs", secTs)
	fmt.Println("nanoTs", nanoTs)
}

// 时间戳格式->时间格式
func timestampToTimeFmt() {
	//**********second
	secTs32 := uint32(time.Now().Unix())
	secTs64 := time.Now().Unix()
	secTsInt64AddDate := time.Now().AddDate(0, 0, -2).Unix()
	fmt.Println("secondTs32:", secTs32)
	fmt.Println("secondTs64:", secTs64)
	//second->timeFmt
	secTimeFmt := time.Unix(secTs64, 0)
	secTimeFmtAddDate := time.Unix(secTsInt64AddDate, 0)
	secTimeFmtStr := time.Unix(secTs64, 0).Format(TIMEFORMATSECOND)
	fmt.Println("secTimeFmt:", secTimeFmt)
	fmt.Println("secTimeFmtAddDate:", secTimeFmtAddDate)
	fmt.Println("secTimeFmtStr:", secTimeFmtStr)

	//**********nanosecond
	nanoTsUnit64 := uint64(time.Now().UnixNano())
	nanoTsInt64 := time.Now().UnixNano()
	nanoTsInt64AddDate := time.Now().AddDate(0, 0, -2).UnixNano()
	fmt.Println("nanoTsUnit64:", nanoTsUnit64)
	fmt.Println("nanoTsInt64:", nanoTsInt64)
	//nanosecond->timeFmt
	nanoTimeFmtCST := time.Unix(0, nanoTsInt64)
	nanoTimeFmtCSTAddDate := time.Unix(0, nanoTsInt64AddDate)
	nanoTimeFmtStr2 := time.Unix(0, nanoTsInt64).Format(TIMEFORMATSECOND)
	fmt.Println("nanoTimeFmtCST:", nanoTimeFmtCST)
	fmt.Println("nanoTimeFmtCSTAddDate:", nanoTimeFmtCSTAddDate)
	fmt.Println("nanoTimeFmtStr2:", nanoTimeFmtStr2)
}

// time.Sleep
func sleepTest() {
	/*每隔30秒，传入的为一个Duration，
	如果想睡眠5s钟，不能直接写 time.Sleep(5) ，
	而应该写time.Sleep(5 * time.Second)
	其中time.Second就是一个Duration类型，表示1s的时间间隔，
	乘系数5就得到5s的时间间隔。*/
	time.Sleep(time.Duration(30) * time.Second)

}
