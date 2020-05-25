package lock

type Session interface {
	NewMutex(key string) Mutex
}
