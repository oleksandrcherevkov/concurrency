package locks

type Lock interface {
	Lock()
	Unlock()
}
