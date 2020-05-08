package lock

const (
	kPrefix = "/lock/"
)

type Session interface {
	NewMutex(key string) Mutex
}
