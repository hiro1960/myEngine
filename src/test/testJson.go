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
	ID     int32
	Name   string
	Tyep   string
	Pos    []float64
	Vel    []float64
	Weight float64
	Mass   float64
}

type Item struct {
	OBJ []Scenario
}

type SceS struct {
	V []Item
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

	// ファイル内容を出力
	fmt.Println(string(b))

	var sce []Scenario

	// jsonを構造体にエンコード(既に使用しているerr変数は何故か代入できず、エラーになる)
	er := json.Unmarshal(b, &sce)

	if er != nil {
		log.Fatal(er)
	}

	// エンコード結果を出力
	fmt.Println(sce)
	fmt.Println(len(sce))

	for i, v := range sce {
		fmt.Printf("%d, %d, %s, %f, %f\n", i, v.ID, v.Name, v.Pos, v.Vel)
		// PosとVelは配列として、個々の値を取得可能
	}

}
