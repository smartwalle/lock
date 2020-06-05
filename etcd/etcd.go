package etcd

import (
	"context"
	"github.com/coreos/etcd/clientv3"
	"github.com/coreos/etcd/clientv3/concurrency"
	"github.com/smartwalle/lock"
	"path"
)

const (
	kPrefix = "lock"
)

type etcdSession struct {
	key    string
	client *clientv3.Client
	opts   []concurrency.SessionOption
}

func (this *etcdSession) NewMutex(key string) lock.Mutex {
	var nPath = path.Join(kPrefix, this.key, key)
	var session, err = concurrency.NewSession(this.client, this.opts...)
	var mu = &etcdMutex{}
	mu.err = err
	mu.session = session
	mu.mu = concurrency.NewMutex(session, nPath)
	return mu
}

func NewSession(key string, client *clientv3.Client, opts ...concurrency.SessionOption) lock.Session {
	var s = &etcdSession{}
	s.key = key
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
