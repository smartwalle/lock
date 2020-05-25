package zookeeper

import (
	"github.com/samuel/go-zookeeper/zk"
	"github.com/smartwalle/lock"
	"path"
)

const (
	kPrefix = "/lock/"
)

type zookeeperSession struct {
	key  string
	conn *zk.Conn
	acl  []zk.ACL
}

func (this *zookeeperSession) NewMutex(key string) lock.Mutex {
	var nPath = path.Join(kPrefix, this.key, key)
	return zk.NewLock(this.conn, nPath, this.acl)
}

func NewSession(key string, conn *zk.Conn, acl []zk.ACL) lock.Session {
	var s = &zookeeperSession{}
	s.key = key
	s.conn = conn
	s.acl = acl
	return s
}
