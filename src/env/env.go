// env型　envパッケージ
// singletonとして実装する
package env

//
type Singleton struct {
	Radius float64
}

func (s *Singleton) GetInstance() *Singleton {
	// s.radius = 6371000.0
	return s
}

// Env型
type Env struct {
	Singleton
}

func (s *Singleton) SetUp() {
	s.Radius = 6371000.0
}

var MyEnv = &Env{}
