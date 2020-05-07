package lock4go

type Session interface {
	NewMutex(key string) Mutex
}
