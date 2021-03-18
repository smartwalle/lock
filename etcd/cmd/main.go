package main

import (
	"fmt"
	"go.etcd.io/etcd/client/v3"

	"github.com/smartwalle/lock/etcd"
	"time"
)

func main() {
	etcdCli, err := clientv3.New(clientv3.Config{Endpoints: []string{"192.168.1.77:2379"}})
	if err != nil {
		fmt.Println(err)
		return
	}
	defer etcdCli.Close()

	var sess = etcd.NewSession("test", etcdCli)

	for i := 0; i < 100; i++ {
		var mu = sess.NewMutex("/f1/f2/f3/")
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
