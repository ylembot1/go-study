package singleton

type Singleton struct {
	a string
}

var singleton *Singleton

func init() {
	singleton = &Singleton{a: "test"}
}

func GetInstance() *Singleton {
	return singleton
}
