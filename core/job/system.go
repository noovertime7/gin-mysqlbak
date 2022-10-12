package job

type Jobber interface {
	Start()
	Stop()
	GetErr() ([]string, bool)
}
