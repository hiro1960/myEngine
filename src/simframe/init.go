// 初期化モジュール
package simframe

import (
	. "core"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"model"
	"os"
)

// simframeで扱うモデルの管理リスト
var ObjList []model.SimObject

// jsonを読み込む構造体
type Scenario struct {
	ID     int32
	Name   string
	Category   string
	Pos    [3]float64
	Vel    [3]float64
	Weight float64
	Mass   float64
}

func Initialize() {

	fmt.Println(" in initilize() ")

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
	// fmt.Println(string(b))

	var sce []Scenario

	// jsonを構造体にエンコード(既に使用しているerr変数は何故か代入できず、エラーになる)
	er := json.Unmarshal(b, &sce)

	if er != nil {
		log.Fatal(er)
	}

	// エンコード結果を出力
	// fmt.Println(sce)
	// fmt.Println(len(sce))

	for i, v := range sce {
		fmt.Printf("%d, %d, %s, %s, %f, %f\n", i, v.ID, v.Category, v.Name, v.Pos, v.Vel)
		// PosとVelは配列として、個々の値を取得可能
	}

	fmt.Printf("\n")

	// シナリオから読み込んだオブジェクトを、objListに登録する
	for _, v := range sce {
		var obj model.SimObject
		var pos *Point
		var vel *Point

		switch v.Category {
		case "vehicle":
			obj = model.NewVehicle(v.ID, v.Name, v.Weight)
			pos = NewPoint(v.Pos)
			obj.SetPos(*pos)
			vel = NewPoint(v.Vel)
			obj.SetVel(*vel)
			ObjList = append(ObjList, obj)

		case "ship":
			obj = model.NewShip(v.ID, v.Name, v.Mass)
			pos = NewPoint(v.Pos)
			obj.SetPos(*pos)
			vel = NewPoint(v.Vel)
			obj.SetVel(*vel)
			ObjList = append(ObjList, obj)

		}
	}

	for i, v := range ObjList {
		var pos Point = v.GetPos()
		var vel Point = v.GetVel()
		fmt.Printf("%d, %d, , %s, %f, %f\n", i, v.GetId(), v.GetName(), pos, vel)
	}

}
