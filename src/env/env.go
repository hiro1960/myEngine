// env型　envパッケージ
// singletonとして実装する
package env

//
type singleton struct {
	Radius float64
}

func (s *singleton) GetInstance() *singleton {
	// s.radius = 6371000.0
	return s
}

// Env型
type Env struct {
	singleton
}

func (s *singleton) SetUp() {
	s.Radius = 6371000.0
}

var MyEnv = &Env{}
