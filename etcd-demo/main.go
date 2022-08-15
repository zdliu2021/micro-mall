package main

import (
	"fmt"
	"time"
)

//主协程退出了，其它子协程也要跟着退出
func main() {

	//ch := make(chan os.Signal, 1)
	//signal.Notify(ch, syscall.SIGTERM, syscall.SIGINT, syscall.SIGKILL, syscall.SIGHUP, syscall.SIGQUIT)
	defer func() {
		//_ = <-ch
		i := 0
		for {
			i++
			fmt.Println("子协程 i = ", i)
			time.Sleep(time.Second)
		}

	}() //别忘了()

	for true {

	}
}
