// logwriter型　logwriterパッケージ
// singletonとして実装する
package logwriter

import (
	"os"
)

// 単にこのoutfという変数をグローバル変数として、外部から参照できるようにするだけでよかったか？
// 複数のファイルを開きたい時、その数だけこのlog.goのコピーが必要になってしまう
type singleton struct {
	Outf *os.File
	// outf *os.File
}

func (s *singleton) GetInstance() *singleton {
	return s
}

// logwriter型
type Logw struct {
	singleton
}

func (s *singleton) Open( name string) ( error) {
	outf, err := os.Create(name)
	s.Outf = outf
	return err
}

func (s *singleton) Close() {
	s.Outf.Close()
}

func (s *singleton) WriteS( data string) {
	s.Outf.WriteString(data)
}

var LogWriter = &Logw{}
