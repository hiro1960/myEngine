// 試験用の実装をおく
package test

import (
	"core"
	"fmt"
	"model"
)

// 試験用にこちらに定義しておく
var ObjList []model.SimObject

// TestMainだけは外部から参照可能とする
func TestMain() {

	testCore()

	testTank()

}

// coreパッケージのテスト
func testCore() {
	var p1, p2 *core.Point
	p1 = core.NewPoint([3]float64{1.0, 2.0, 3.0})
	p2 = core.NewPoint([3]float64{1.0, 1.0, 1.0})

	p1.Add(p2)

	fmt.Printf("--- p1 result ---\n")
	for i, v := range p1.Value {
		fmt.Printf(" i = %d, v = %f\n", i, v)
	}
}

// Tankに位置p1をセットし、値を確認する
func testTank() {
	var obj model.SimObject

	obj = ObjList[0]

	var p1 *core.Point
	p1 = core.NewPoint([3]float64{1.0, 2.0, 3.0})

	obj.SetPos(*p1)
	var p3 core.Point
	p3 = obj.GetPos()

	fmt.Printf("--- Tank 位置 ---\n")
	for i, v := range p3.Value {
		fmt.Printf(" i = %d, v = %f\n", i, v)
	}
}
