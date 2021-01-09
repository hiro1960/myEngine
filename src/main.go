// simexecのプロトタイプ　まず1ファイルで作り、その後packageに分ける
package main

// GOPATH環境変数の設定は必要（パッケージを読めなくなる）
import (
	"core"
	"fmt"
	"model"
)

// modelのクラスはmodelパッケージ内に全て定義してある

func main() {
	var obj model.SimObject
	var objList []model.SimObject

	obj = model.NewVehicle(101, "Tank", 124.0)
	objList = append(objList, obj)
	fmt.Printf(" ID = %d\n", obj.GetId())
	fmt.Printf(" Name = %s\n", obj.GetName())
	fmt.Printf(" Weight = %f\n", obj.GetWeight())

	fmt.Printf(" objList length = %d\n", len(objList))

	obj = model.NewShip(102, "Aegis", 201.0)
	objList = append(objList, obj)
	fmt.Printf(" ID = %d\n", obj.GetId())
	fmt.Printf(" Name = %s\n", obj.GetName())
	fmt.Printf(" Mass = %f\n", obj.GetMass())

	fmt.Printf(" objList length = %d\n", len(objList))

	for _, s := range objList {
		fmt.Printf(" ID = %d\n", s.GetId())
		fmt.Printf(" Name = %s\n", s.GetName())
	}

	// coreパッケージのテスト
	var p1, p2 *core.Point
	p1 = core.NewPoint([3]float64{1.0, 2.0, 3.0})
	p2 = core.NewPoint([3]float64{1.0, 1.0, 1.0})

	p1.Add(p2)

	fmt.Printf("--- p1 result ---\n")
	for i, v := range p1.Value {
		fmt.Printf(" i = %d, v = %f\n", i, v)
	}

	// Tankに位置p1をセットし、値を確認する
	obj = objList[0]

	obj.SetPos(*p1)
	var p3 core.Point
	p3 = obj.GetPos()

	fmt.Printf("--- Tank 位置 ---\n")
	for i, v := range p3.Value {
		fmt.Printf(" i = %d, v = %f\n", i, v)
	}

}
