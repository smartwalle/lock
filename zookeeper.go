package lock

import (
	"github.com/samuel/go-zookeeper/zk"
	"path"
)

type zookeeperSession struct {
	key  string
	conn *zk.Conn
	acl  []zk.ACL
}

func (this *zookeeperSession) NewMutex(key string) Mutex {
	var nPath = path.Join(kPrefix, this.key, key)
	return zk.NewLock(this.conn, nPath, this.acl)
}

func NewZookeeperSession(key string, conn *zk.Conn, acl []zk.ACL) Session {
	var s = &zookeeperSession{}
	s.key = key
	s.conn = conn
	s.acl = acl
	return s
}
