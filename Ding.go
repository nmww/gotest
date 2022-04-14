package main

import (
	"github.com/robfig/cron"
	"sync"
	"time"
)

//cron 定时
//https://www.cnblogs.com/shiluoliming/p/8310725.html
//every 使用
//https://blog.csdn.net/slphahaha/article/details/119458696

var mutex sync.Mutex

// 任务列表
var taskList []timeTask

// 初始化并启动定时任务
func InitProcessTimer() {
	c := cron.New()
	c.AddFunc("@every 5S", ProcessTimerTaskHandler)
	c.Start()
}

// 定时要执行的回调函数
type ProcessTimerCallback func(id string, err error) error

// 定时任务类
type timeTask struct {
	id      string
	create  time.Time
	timeOut time.Duration
	cb      ProcessTimerCallback
}

// 插入定时任务
func InserTimerTask(id string, timeout time.Duration, cb ProcessTimerCallback) {
	var task timeTask
	task.id = id
	task.create = time.Now()
	task.timeOut = timeout
	task.cb = cb

	mutex.Lock() // blocks until the mutex is available.
	defer mutex.Unlock()

	for i := 0; i < len(taskList); i++ {
		// 如果id存在，只刷新创建时间和超时时间
		if taskList[i].id == task.id {
			taskList[i].create = task.create
			taskList[i].timeOut = task.timeOut
			return
		}
	}

	// 如果id不存在。增加一个task
	taskList = append(taskList, task)
}

//删除定时任务
func RemoveTimeTask(id string) {
	mutex.Lock()
	defer mutex.Unlock()

	for i := 0; i < len(taskList); i++ {
		if taskList[i].id == id {
			taskList = append(taskList[:i], taskList[i+1:]...)
			return
		}
	}
}

// 定时处理让任务
func ProcessTimerTaskHandler() {
	var task timeTask
	mutex.Lock()
	defer mutex.Unlock()

	if len(taskList) == 0 {
		return
	}

	for i := 0; i < len(taskList); {
		task = taskList[i]
		if time.Now().Sub(task.create) > task.timeOut {
			// 执行回调函数
			task.cb(task.id, nil)
			// 删除定时任务
			taskList = append(taskList[:i], taskList[i+1:]...)
		} else {
			i++
		}
	}
}
