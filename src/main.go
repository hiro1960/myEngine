// simexecのプロトタイプ　まず1ファイルで作り、その後packageに分ける
package main

// GOPATH環境変数の設定は必要（パッケージを読めなくなる）
import (
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

	for _,s := range objList {
		fmt.Printf(" ID = %d\n", s.GetId())
		fmt.Printf(" Name = %s\n", s.GetName())
	}

}
