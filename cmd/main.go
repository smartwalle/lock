package main

import (
	"fmt"
	"github.com/coreos/etcd/clientv3"
	"github.com/samuel/go-zookeeper/zk"
	"github.com/smartwalle/lock"
	"time"
)

func main() {
	run(getETCDSession())
}

func getZKSession() lock.Session {
	zkConn, _, err := zk.Connect([]string{"127.0.0.1:2181"}, time.Second*5)
	if err != nil {
		fmt.Println(err)
		return nil
	}

	//defer zkConn.Close()
	return lock.NewZookeeperSession("test", zkConn, zk.WorldACL(zk.PermAll))
}

func getETCDSession() lock.Session {
	etcdCli, err := clientv3.New(clientv3.Config{Endpoints: []string{"127.0.0.1:2379"}})
	if err != nil {
		return nil
	}
	//defer etcdCli.Close()

	return lock.NewETCDSession("test", etcdCli)
}

func run(s lock.Session) {
	for i := 0; i < 100; i++ {
		var mu = s.NewMutex("/hihi/sss/s/")
		if err := mu.Lock(); err != nil {
			fmt.Println("加锁失败:", err)
			continue
		}
		fmt.Println("加锁成功:", i, time.Now())
		time.Sleep(time.Minute * 3)
		fmt.Println("释放锁:", i, time.Now())
		if err := mu.Unlock(); err != nil {
			fmt.Println("解锁失败:", err)
			continue
		}
	}
}
