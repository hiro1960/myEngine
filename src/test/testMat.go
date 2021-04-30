// Mat3テスト
package test

import (
	"core"
	"fmt"
)

func TestMat() {
	v := [][]float64{{1, 2, 3}, {4, 5, 6}, {7, 8, 9}}

	fmt.Printf("%v\n", v)
	mat1 := core.NewMat3(v)
	fmt.Printf("mat1 = \n")
	mat1.Output()

	fmt.Printf("%f\n", mat1.GetL2())

	v2 := [][]float64{{1, 1, 1}, {1, 1, 1}, {1, 1, 1}}
	mat2 := core.NewMat3(v2)
	fmt.Printf("mat2 = \n")
	mat2.Output()
	v3 := [][]float64{{1, 1, 1}, {1, 1, 1}, {1, 1, 1}}
	mat3 := core.NewMat3(v3)

	mat3.MultiMatrix(mat1, mat2)
	fmt.Printf("mat3 = \n")
	mat3.Output()
	fmt.Printf("\n")

}
