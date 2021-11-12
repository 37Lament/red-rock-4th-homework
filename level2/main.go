package main

import (
	"fmt"
	"time"
)

// 复读机
func repeater(content string, d time.Duration) {
	ticker := time.NewTicker(d)
	for {
		select {
		case <-ticker.C://利用ticker用来作为时间周期执行命令
			fmt.Println(content)
		}
	}
}
// 定时器
func timing(content string, h, m int) {
	// 先算出第一个输出的时间点
	now := time.Now()
	t := time.Date(now.Year(), now.Month(), now.Day(), h, m, 0, 0, time.Local)
	if t.Before(now) {
		t = t.Add(24 * time.Hour)//判断是否在今天并且调整到对应天数
	}

	timer := time.NewTimer(t.Sub(now))
	for {
		<-timer.C
		fmt.Println(content)
		timer.Reset(24 * time.Hour) // 每天重置时间
	}
}

func main() {
	go repeater("芜湖！起飞！", 30*time.Second)
	go timing("谁能比我卷！", 2, 0)
	go timing("早八算什么，早六才是吾辈应起之时", 6, 0)

	// 保持定时器运行
	control := make(chan bool)//定义并且初始化并运行
	<-control
}
