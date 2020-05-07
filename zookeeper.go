package lock4go

import (
	"github.com/samuel/go-zookeeper/zk"
	"path/filepath"
)

type zookeeperSession struct {
	conn *zk.Conn
	acl  []zk.ACL
}

func (this *zookeeperSession) NewMutex(key string) Mutex {
	key = filepath.Clean("/" + key)
	return zk.NewLock(this.conn, key, this.acl)
}

func NewSessionWithZookeeper(c *zk.Conn, acl []zk.ACL) Session {
	var s = &zookeeperSession{}
	s.conn = c
	s.acl = acl
	return s
}
