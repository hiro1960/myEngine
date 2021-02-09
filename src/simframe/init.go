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
	ID        int32
	Name      string
	Category  string
	Pos       [3]float64
	Vel       [3]float64
	Weight    float64
	Mass      float64
	Count     int32
	DeltaTime float64
}

type SimCtrl struct {
	Count     int32
	DeltaTime float64
}

// シミュレーションの時間管理
var TimeSet SimCtrl

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
	fmt.Println(sce)
	fmt.Println(len(sce))

	// シナリオから読み込んだオブジェクトを、objListに登録する
	for _, v := range sce {
		var obj model.SimObject
		var pos *Point
		var vel *Point
		var objDB model.ObjData

		switch v.Category {
		case "SimCtrl":
			TimeSet.Count = v.Count
			TimeSet.DeltaTime = v.DeltaTime

		case "vehicle":
			obj = model.NewVehicle(v.ID, v.Name, v.Weight)
			pos = NewPoint(v.Pos) // ポインタを返してくるのに注意
			obj.SetPos(*pos)
			vel = NewPoint(v.Vel) // ポインタを返してくるのに注意
			obj.SetVel(*vel)
			ObjList = append(ObjList, obj)
			// objDataDBへの登録
			objDB.ID = v.ID
			objDB.Name = v.Name
			objDB.Pos = pos // ポインタを代入していることに注意
			objDB.Vel = vel
			objDB.UpdatedPos = new(Point)
			objDB.UpdatedVel = new(Point)
			model.ObjDataDB = append(model.ObjDataDB, objDB)

		case "ship":
			obj = model.NewShip(v.ID, v.Name, v.Mass)
			pos = NewPoint(v.Pos)
			obj.SetPos(*pos)
			vel = NewPoint(v.Vel)
			obj.SetVel(*vel)
			ObjList = append(ObjList, obj)
			// objDataDBへの登録
			objDB.ID = v.ID
			objDB.Name = v.Name
			objDB.Pos = pos
			objDB.Vel = vel
			objDB.UpdatedPos = new(Point)
			objDB.UpdatedVel = new(Point)
			model.ObjDataDB = append(model.ObjDataDB, objDB)

		}
	}

	for i, v := range ObjList {
		var pos Point = v.GetPos()
		var vel Point = v.GetVel()
		fmt.Printf("%d, %d, %s, %f, %f\n", i, v.GetId(), v.GetName(), pos, vel)
	}

}
