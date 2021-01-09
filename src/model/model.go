// modelパッケージ
// (ディレクトリ名とパッケージ名は同じ必要があるが、ファイル名は関係ない)
package model

// 本インターフェースを利用する全てのstructの持つメソッドをリストアップすること
// 公開するのは、このインターフェース名だけにする
type SimObject interface {
	GetId() int32
	GetName() string
	GetWeight() float64 // どのstructが使うメソッドかコメントした方がいいか
	GetMass() float64
}

// 基底となるデータ構造
// class名自体は隠ぺいするため、小文字とする
type baseObject struct {
	ID   int32
	Name string
}

// IDを返す
func (b baseObject) GetId() int32 {
	return b.ID
}

// 名前を返す
func (b baseObject) GetName() string {
	return b.Name
}
