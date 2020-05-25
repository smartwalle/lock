package main

import (
	"fmt"
	"github.com/samuel/go-zookeeper/zk"
	"github.com/smartwalle/lock/zookeeper"
	"time"
)

func main() {
	zkConn, _, err := zk.Connect([]string{"127.0.0.1:2181"}, time.Second*5)
	if err != nil {
		fmt.Println(err)
		return
	}

	defer zkConn.Close()
	var ress = zookeeper.NewSession("test", zkConn, zk.WorldACL(zk.PermAll))

	for i := 0; i < 100; i++ {
		var mu = ress.NewMutex("/f1/f2/f3/")
		if err := mu.Lock(); err != nil {
			fmt.Println("加锁失败:", err)
			continue
		}
		fmt.Println("加锁成功:", i, time.Now())
		time.Sleep(time.Second * 3)
		fmt.Println("释放锁:", i, time.Now())
		if err := mu.Unlock(); err != nil {
			fmt.Println("解锁失败:", err)
			continue
		}
	}
}
