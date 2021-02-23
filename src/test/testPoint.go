// Pointテスト
package test

import (
	"fmt"
	"core"
	"math"
)

func TestPoint()  {
	pos := core.NewPoint([3]float64{1.0, 2.0, 3.0})
	fmt.Printf("\n")
	fmt.Printf("in TestPoint: Point's Radius = %f\n", pos.GetRadius())
	// うっかり地球半径を変更して、試験を続行してしまっていたのでコメントアウトする
	// pos.SetRadius(7000000.0)
	// pos2 := core.NewPoint([3]float64{1.0, 2.0, 3.0})

	// // 別のPointでも同じ地球半径を持つことを確認
	// fmt.Printf("in TestPoint: Point's Radius = %f\n", pos2.GetRadius())

	fmt.Printf("\n")

	// 緯度、経度の確認
	pos3 := core.NewPoint( [3]float64{7000000, 7000000, 700000})
	fmt.Printf("in TestPoint: GetLat = %f\n", pos3.GetLat()*core.RadToDeg)
	fmt.Printf("in TestPoint: GetLon = %f\n", pos3.GetLon()*core.RadToDeg)
	fmt.Printf("in TestPoint: GetLength = %f\n", pos3.GetLength())
	fmt.Printf("in TestPoint: GetAlt = %f\n", pos3.GetAlt())
	fmt.Printf("in TestPoint: earthRadius = %f\n", pos3.GetRadius())
	fmt.Printf("\n")

	pos4 := core.NewPoint( [3]float64{7000000, 7000000, 7000000})
	fmt.Printf("in TestPoint: DistanceTo = %f\n", pos4.DistanceTo(pos3))
	fmt.Printf("in TestPoint: DownRange = %f\n", pos4.DownRangeTo(pos3))
	fmt.Printf("\n")

	pos5 := pos4.Localize(pos3)
	fmt.Printf("in TestPoint: localize = %f\n", pos5)
	fmt.Printf("\n")

	pos6 := pos5.Globalize(pos3)
	fmt.Printf("in testPoint: globalize = %f\n", pos6)
	fmt.Printf("\n")

}

func TestMath() {
	fmt.Printf("\n")
	fmt.Printf("int TestMath: Pi = %f\n", math.Pi)

	fmt.Printf("\n")

}