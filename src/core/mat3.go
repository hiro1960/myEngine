// 3x3行列　計算用
package core

import (
	"fmt"
	"math"
)

// Mat3型
// ( l1 m1 n1 )
// ( l2 m2 n2 )
// ( l3 m3 n3 )
type Mat3 struct {
	l1 float64
	l2 float64
	l3 float64
	m1 float64
	m2 float64
	m3 float64
	n1 float64
	n2 float64
	n3 float64
}

// Mat3型　コンストラクタ　（[][]float64を設定値としてもらう）
func NewMat3(val [][]float64) *Mat3 {
	p := new(Mat3)

	p.l1 = val[0][0]
	p.l2 = val[0][1]
	p.l3 = val[0][2]
	p.m1 = val[1][0]
	p.m2 = val[1][1]
	p.m3 = val[1][2]
	p.n1 = val[2][0]
	p.n2 = val[2][1]
	p.n3 = val[2][2]

	return p
}

func NewEulerMat3(phi, theta, psi float64) *Mat3 {
	p := new(Mat3)

	s1 := math.Sin(phi)
	c1 := math.Cos(phi)
	s2 := math.Sin(theta)
	c2 := math.Cos(theta)
	s3 := math.Sin(psi)
	c3 := math.Cos(psi)

	p.l1 = c2 * c3
	p.l2 = s1*s2*c3 - c1*s3
	p.l3 = c1*s2*c3 + s1*s3
	p.m1 = c2 * s3
	p.m2 = s1*s2*s3 + c1*c3
	p.m3 = c1*s2*s3 - s1*c3
	p.n1 = -s2
	p.n2 = s1 * c2
	p.n3 = c1 * c2

	return p
}

// 値確認用
func (p *Mat3) Output() {
	fmt.Printf("%f %f %f\n", p.l1, p.m1, p.n1)
	fmt.Printf("%f %f %f\n", p.l2, p.m2, p.n2)
	fmt.Printf("%f %f %f\n", p.l3, p.m3, p.n3)
}

// 各値のゲッター
func (p *Mat3) GetL1() float64 {
	return p.l1
}

func (p *Mat3) GetL2() float64 {
	return p.l2
}

func (p *Mat3) GetL3() float64 {
	return p.l3
}

func (p *Mat3) GetM1() float64 {
	return p.m1
}

func (p *Mat3) GetM2() float64 {
	return p.m2
}
func (p *Mat3) GetM3() float64 {
	return p.m3
}

func (p *Mat3) GetN1() float64 {
	return p.n1
}

func (p *Mat3) GetN2() float64 {
	return p.n2
}

func (p *Mat3) GetN3() float64 {
	return p.n3
}

func (p *Mat3) InverseEulerMarix() {
	var tempForSwap float64

	tempForSwap = p.l2
	p.l2 = p.m1
	p.m1 = tempForSwap

	tempForSwap = p.l3
	p.l3 = p.n1
	p.n1 = tempForSwap

	tempForSwap = p.m3
	p.m3 = p.n2
	p.n2 = tempForSwap
}

func (p *Mat3) MultiMatrix(mt1, mt2 *Mat3) {
	p.l1 = mt1.l1*mt2.l1 + mt1.m1*mt2.l2 + mt1.n1*mt2.l3
	p.l2 = mt1.l2*mt2.l1 + mt1.m2*mt2.l2 + mt1.n2*mt2.l3
	p.l3 = mt1.l3*mt2.l1 + mt1.m3*mt2.l2 + mt1.n3*mt2.l3

	p.m1 = mt1.l1*mt2.m1 + mt1.m1*mt2.m2 + mt1.n1*mt2.m3
	p.m2 = mt1.l2*mt2.m1 + mt1.m2*mt2.m2 + mt1.n2*mt2.m3
	p.m3 = mt1.l3*mt2.m1 + mt1.m3*mt2.m2 + mt1.n3*mt2.m3

	p.n1 = mt1.l1*mt2.n1 + mt1.m1*mt2.n2 + mt1.n1*mt2.n3
	p.n2 = mt1.l2*mt2.n1 + mt1.m2*mt2.n2 + mt1.n2*mt2.n3
	p.n3 = mt1.l3*mt2.n1 + mt1.m3*mt2.n2 + mt1.n3*mt2.n3
}

func (p *Mat3) EulerTrans(org, dest *Point) {
	dest.Value[0] = p.l1*org.Value[0] + p.m1*org.Value[1] + p.n1*org.Value[2]
	dest.Value[1] = p.l2*org.Value[0] + p.m2*org.Value[1] + p.n2*org.Value[2]
	dest.Value[2] = p.l3*org.Value[0] + p.m3*org.Value[1] + p.n3*org.Value[2]
}

func (p *Mat3) EulerTransInverse(org, dest *Point) {
	dest.Value[0] = p.l1*org.Value[0] + p.l2*org.Value[1] + p.l3*org.Value[2]
	dest.Value[1] = p.m1*org.Value[0] + p.m2*org.Value[1] + p.m3*org.Value[2]
	dest.Value[2] = p.n1*org.Value[0] + p.n2*org.Value[1] + p.n3*org.Value[2]
}

func (p *Mat3) EulerTransXYZ(ox, oy, oz, dx, dy, dz float64) {
	dx = p.l1*ox + p.m1*oy + p.n1*oz
	dy = p.l2*ox + p.m2*oy + p.n2*oz
	dz = p.l3*ox + p.m3*oy + p.n3*oz
}

func (p *Mat3) EulerTransInverseXYZ(ox, oy, oz, dx, dy, dz float64) {
	dx = p.l1*ox + p.l2*oy + p.l3*oz
	dy = p.m1*ox + p.m2*oy + p.m3*oz
	dz = p.n1*ox + p.n2*oy + p.n3*oz
}
