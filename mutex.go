package lock

type Mutex interface {
	Lock() error

	Unlock() error
}
