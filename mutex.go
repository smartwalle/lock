package lock4go

type Mutex interface {
	Lock() error

	Unlock() error
}
