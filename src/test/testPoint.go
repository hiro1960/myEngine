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
	pos.SetRadius(7000000.0)
	pos2 := core.NewPoint([3]float64{1.0, 2.0, 3.0})

	// 別のPointでも同じ地球半径を持つことを確認
	fmt.Printf("in TestPoint: Point's Radius = %f\n", pos2.GetRadius())

	fmt.Printf("\n")

}

func TestMath() {
	fmt.Printf("\n")
	fmt.Printf("int TestMath: Pi = %f\n", math.Pi)

	fmt.Printf("\n")

}