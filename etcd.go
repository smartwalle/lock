package lock4go

import (
	"context"
	"github.com/coreos/etcd/clientv3"
	"github.com/coreos/etcd/clientv3/concurrency"
	"path/filepath"
)

type etcdSession struct {
	client *clientv3.Client
	opts   []concurrency.SessionOption
}

func (this *etcdSession) NewMutex(key string) Mutex {
	key = filepath.Clean("/" + key + "/")

	session, err := concurrency.NewSession(this.client, this.opts...)

	var mu = &etcdMutex{}
	mu.err = err
	mu.session = session
	mu.mu = concurrency.NewMutex(session, key)

	return mu
}

func NewSessionWithETCD(client *clientv3.Client, opts ...concurrency.SessionOption) Session {
	var s = &etcdSession{}
	s.client = client
	s.opts = opts
	return s
}

type etcdMutex struct {
	err     error
	session *concurrency.Session
	mu      *concurrency.Mutex
}

func (this *etcdMutex) Lock() error {
	if this.err != nil {
		return this.err
	}
	var err = this.mu.Lock(context.TODO())
	if err != nil {
		this.session.Close()
	}
	return err
}

func (this *etcdMutex) Unlock() error {
	if this.err != nil {
		return this.err
	}

	defer this.session.Close()
	return this.mu.Unlock(context.TODO())
}
