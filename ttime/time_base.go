package ttime

import (
	"strconv"
	"time"
)

//输出文本时间格式 2020-11-22 20:54:05 ，24 小时制
func ToStr(t time.Time) string {
	// 格式化的模板为 Go 的出生时间 2006 年 1 月 2 号 15 点 04 分 Mon Jan
	return t.Format("2006-01-02 15:04:05") // 2020-11-22 20:54:05 24 小时制
}

//取文本时间差， now-t   输出格式 : type=0, 3天2小时5分 ; type=1  122h5m
func Nowsub(t time.Time, Type int) string {
	m := int(time.Since(t).Minutes()) //间隔时间 分
	var d, h = 0, 0
	if m > 1440 {
		d = int(m / 1440)
		h = int((m % 1440) / 60)
		m = int((m % 1440) % 60)
	} else {
		h = int(m / 60)
		m = m % 60
	}

	if Type == 0 {
		return strconv.Itoa(d) + "天" + strconv.Itoa(h) + "小时" + strconv.Itoa(m) + "分"
	} else {

		if d > 0 {
			h += d * 24
		}
		return strconv.Itoa(h) + "h" + strconv.Itoa(m) + "m"
	}
}

//取耗时 now-t   返回时间间隔 毫秒
func NowsubInt13(t time.Time) int {
	m := int(time.Since(t).Milliseconds())
	return m
}

//13位时间戳
func Int13(t time.Time) int64 {
	return t.UnixNano() / 1e6
}

//10位时间戳
func Int10(t time.Time) int64 {
	return t.Unix()
}

//今日10位时间戳
func Int10ToDay() int64 {

	//var cstSh, _ = time.LoadLocation("UTC") //上海
	t := time.Now()
	t2 := time.Date(t.Year(), t.Month(), t.Day(), 0, 0, 0, 0, time.Local).Unix()
	//a := t.Unix()
	return t2
	/* a := t2
	b := a % (24 * 60 * 60)
	return (a - b) */

}
