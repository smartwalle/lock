package lock4go

import (
	"context"
	"github.com/coreos/etcd/clientv3"
	"github.com/coreos/etcd/clientv3/concurrency"
	"path"
)

type etcdSession struct {
	prefix string
	client *clientv3.Client
	opts   []concurrency.SessionOption
}

func (this *etcdSession) NewMutex(key string) Mutex {
	var nPath = path.Join("/", this.prefix, key)
	var session, err = concurrency.NewSession(this.client, this.opts...)
	var mu = &etcdMutex{}
	mu.err = err
	mu.session = session
	mu.mu = concurrency.NewMutex(session, nPath)
	return mu
}

func NewSessionWithETCD(prefix string, client *clientv3.Client, opts ...concurrency.SessionOption) Session {
	var s = &etcdSession{}
	s.prefix = prefix
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

	if err != nil && this.session != nil {
		this.session.Close()
	}
	return err
}

func (this *etcdMutex) Unlock() error {
	var err error
	if this.mu != nil {
		err = this.mu.Unlock(context.TODO())
	}
	if this.session != nil {
		this.session.Close()
	}
	return err
}
