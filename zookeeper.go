package lock4go

import (
	"github.com/samuel/go-zookeeper/zk"
	"path/filepath"
)

type zookeeperSession struct {
	prefix string
	conn   *zk.Conn
	acl    []zk.ACL
}

func (this *zookeeperSession) NewMutex(key string) Mutex {
	var path = filepath.Join("/", this.prefix, key)
	return zk.NewLock(this.conn, path, this.acl)
}

func NewSessionWithZookeeper(prefix string, conn *zk.Conn, acl []zk.ACL) Session {
	var s = &zookeeperSession{}
	s.prefix = prefix
	s.conn = conn
	s.acl = acl
	return s
}
