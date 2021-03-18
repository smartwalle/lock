package etcd

import (
	"context"
	"github.com/smartwalle/lock"
	"go.etcd.io/etcd/client/v3"
	"go.etcd.io/etcd/client/v3/concurrency"
	"path"
)

const (
	kPrefix = "lock"
)

type session struct {
	key    string
	client *clientv3.Client
	opts   []concurrency.SessionOption
}

func NewSession(key string, client *clientv3.Client, opts ...concurrency.SessionOption) lock.Session {
	var s = &session{}
	s.key = key
	s.client = client
	s.opts = opts
	return s
}

func (this *session) NewMutex(key string) lock.Mutex {
	var nPath = path.Join(kPrefix, this.key, key)
	var session, err = concurrency.NewSession(this.client, this.opts...)
	var mu = &mutex{}
	mu.err = err
	mu.session = session
	mu.mu = concurrency.NewMutex(session, nPath)
	return mu
}

type mutex struct {
	err     error
	session *concurrency.Session
	mu      *concurrency.Mutex
}

func (this *mutex) Lock() error {
	if this.err != nil {
		return this.err
	}
	var err = this.mu.Lock(context.TODO())

	if err != nil && this.session != nil {
		this.session.Close()
	}
	return err
}

func (this *mutex) Unlock() error {
	var err error
	if this.mu != nil {
		err = this.mu.Unlock(context.TODO())
	}
	if this.session != nil {
		this.session.Close()
	}
	return err
}
