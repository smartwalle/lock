package main

import (
	"fmt"
	"github.com/coreos/etcd/clientv3"
	"github.com/samuel/go-zookeeper/zk"
	"github.com/smartwalle/lock4go"
	"time"
)

func main() {
	run(getETCDSession())
}

func getZKSession() lock4go.Session {
	zkConn, _, err := zk.Connect([]string{"127.0.0.1:2181"}, time.Second*5)
	if err != nil {
		fmt.Println(err)
		return nil
	}
	//defer zkConn.Close()
	return lock4go.NewSessionWithZookeeper("test", zkConn, zk.WorldACL(zk.PermAll))
}

func getETCDSession() lock4go.Session {
	etcdCli, err := clientv3.New(clientv3.Config{Endpoints: []string{"127.0.0.1:2379"}})
	if err != nil {
		return nil
	}
	//defer etcdCli.Close()

	return lock4go.NewSessionWithETCD("test", etcdCli)
}

func run(s lock4go.Session) {
	for i := 0; i < 100; i++ {
		var mu = s.NewMutex("haha")
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
