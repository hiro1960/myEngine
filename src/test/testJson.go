// json読み込みテスト
package test

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"os"
)

// jsonを読み込む構造体
type Scenario struct {
	ID  int32
	Pos []float64
}

func testJson() {
	fmt.Println(" in testJson ")

	// jsonファイル読み込み（ファイルはmain.goと同じ階層に置くこと。そこがカレントディレクトリとなる。）
	f, err := os.Open("sample.json")

	// 読み取り時の例外処理
	if err != nil {
		log.Fatal(err)
	}
	// 関数が終了した際に確実にクローズするようにする
	defer f.Close()

	// 一気に全部読み取り
	b, err := ioutil.ReadAll(f)

	// 出力
	fmt.Println(string(b))

	var sce Scenario
	// jsonを構造体にエンコード(既に使用しているerr変数は何故か代入できず、エラーになる)
	er := json.Unmarshal(b, &sce)

	if er != nil {
		log.Fatal(er)
	}

	// エンコード結果を出力
	fmt.Println(sce)

}
