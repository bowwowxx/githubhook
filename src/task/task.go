package task

import (
	"../config"
	"../utils"
	"fmt"
	"os/exec"
)

var running = false
var queue []*structTaskQueue

type structTaskQueue struct {
	requestBodyString string
}

// AddNewTask add new task
func AddNewTask(bodyContent string,serviceName string) {
	queue = append(queue, newStructTaskQueue(bodyContent))
	checkoutTaskStatus(serviceName)
}

func newStructTaskQueue(body string) *structTaskQueue {
	return &structTaskQueue{body}
}

func checkoutTaskStatus(serviceName string) {
	if running {
		return
	}
	if len(queue) > 0 {
		queue = queue[:0:0]
		go startTask(serviceName)
	}
}

func startTask(serviceName string) {
	running = true
	cmd := exec.Command("/bin/sh", config.GetShell(),serviceName)
	_, err := cmd.Output()
	if err == nil {
		running = false
		utils.Log2file("執行成功")
		checkoutTaskStatus(serviceName)
	} else {
		running = false
		utils.Log2file(fmt.Sprintf("執行失败:\n %s", err))
		checkoutTaskStatus(serviceName)
	}
	utils.Log2file(fmt.Sprintf("service name is %s\n", serviceName))
	utils.Log2file(fmt.Sprintf("pid is %v\n", cmd.Process.Pid))
	
}
