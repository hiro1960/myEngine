// quartenionパッケージ
package core

import "math"

type Quartenion struct {
	e0 float64
	e1 float64
	e2 float64
	e3 float64

	e0DotPre float64
	e1DotPre float64
	e2DotPre float64
	e3DotPre float64

	// Mat3型
	// ( l1 m1 n1 )
	// ( l2 m2 n2 )
	// ( l3 m3 n3 )
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

// 各値のゲッター
func (p *Quartenion) GetE0() float64 {
	return p.e0
}

func (p *Quartenion) GetE1() float64 {
	return p.e1
}

func (p *Quartenion) GetE2() float64 {
	return p.e2
}

func (p *Quartenion) GetE3() float64 {
	return p.e3
}

func (p *Quartenion) GetL1() float64 {
	return p.l1
}

func (p *Quartenion) GetL2() float64 {
	return p.l2
}

func (p *Quartenion) GetL3() float64 {
	return p.l3
}

func (p *Quartenion) GetM1() float64 {
	return p.m1
}

func (p *Quartenion) GetM2() float64 {
	return p.m2
}
func (p *Quartenion) GetM3() float64 {
	return p.m3
}

func (p *Quartenion) GetN1() float64 {
	return p.n1
}

func (p *Quartenion) GetN2() float64 {
	return p.n2
}

func (p *Quartenion) GetN3() float64 {
	return p.n3
}

// 初期化
func (p *Quartenion) Initialize(phi, theta, psi float64) {
	cosPsi2 := math.Cos(psi / 2.0)
	sinPsi2 := math.Sin(psi / 2.0)
	cosTheta2 := math.Cos(theta / 2.0)
	sinTheta2 := math.Sin(theta / 2.0)
	cosPhi2 := math.Cos(phi / 2.0)
	sinPhi2 := math.Sin(phi / 2.0)

	p.e0 = cosPsi2*cosTheta2*cosPhi2 + sinPsi2*sinTheta2*sinPhi2
	p.e1 = cosPsi2*cosTheta2*sinPhi2 - sinPsi2*sinTheta2*cosPhi2
	p.e2 = cosPsi2*sinTheta2*cosPhi2 + sinPsi2*cosTheta2*sinPhi2
	p.e3 = -cosPsi2*sinTheta2*sinPhi2 + sinPsi2*cosTheta2*cosPhi2

	p.e3DotPre = 0.0
	p.e2DotPre = 0.0
	p.e1DotPre = 0.0
	p.e0DotPre = 0.0

	p.UpdateEulerMatrix()
}

// 行列の更新
func (p *Quartenion) UpdateEulerMatrix() {
	e0e0 := p.e0 * p.e0
	e1e1 := p.e1 * p.e1
	e2e2 := p.e2 * p.e2
	e3e3 := p.e3 * p.e3

	e0e1 := p.e0 * p.e1
	e0e2 := p.e0 * p.e2
	e0e3 := p.e0 * p.e3
	e1e2 := p.e1 * p.e3
	e1e3 := p.e1 * p.e3
	e2e3 := p.e2 * p.e3

	p.l1 = e0e0 + e1e1 - e2e2 - e3e3
	p.l2 = 2.0 * (e1e2 - e2e3)
	p.l3 = 2.0 * (e0e2 + e1e3)
	p.m1 = 2.0 * (e0e3 + e1e2)
	p.m2 = e0e0 - e1e1 + e2e2 - e3e3
	p.m3 = 2.0 * (e2e3 - e0e1)
	p.n1 = 2.0 * (e1e3 - e0e2)
	p.n2 = 2.0 * (e0e1 + e2e3)
	p.n3 = e0e0 - e1e1 - e2e2 + e3e3
}

// quartenionの更新
func (v *Quartenion) UpdateQuartenion(p, q, r, dtime float64) {
	var Kq float64
	Kq = 0.0

	Keps := Kq * (1.0 - (v.e0*v.e0 + v.e1*v.e1 + v.e2*v.e2 + v.e3*v.e3))

	e0Dot := -(v.e1*p+v.e2*q+v.e3*r)/2.0 + Keps*v.e0
	e1Dot := (v.e0*p-v.e3*q+v.e2*r)/2.0 + Keps*v.e1
	e2Dot := (v.e3*p+v.e0*q-v.e1*r)/2.0 + Keps*v.e2
	e3Dot := (-v.e2*p+v.e1*q+v.e0*r)/2.0 + Keps*v.e3

	IntegImp(e0Dot, dtime, &v.e0DotPre, &v.e0)
	IntegImp(e1Dot, dtime, &v.e1DotPre, &v.e1)
	IntegImp(e2Dot, dtime, &v.e2DotPre, &v.e2)
	IntegImp(e3Dot, dtime, &v.e3DotPre, &v.e3)

	aa := math.Sqrt(v.e0*v.e0 + v.e1*v.e1 + v.e2*v.e2 + v.e3*v.e3)
	v.e0 = v.e0 / aa
	v.e1 = v.e1 / aa
	v.e2 = v.e2 / aa
	v.e3 = v.e3 / aa

	v.UpdateEulerMatrix()
}

// オイラー変換
func (p *Quartenion) EulerTrans(org, dest *Point) {
	dest.Value[0] = p.l1*org.Value[0] + p.m1*org.Value[1] + p.n1*org.Value[2]
	dest.Value[1] = p.l2*org.Value[0] + p.m2*org.Value[1] + p.n2*org.Value[2]
	dest.Value[2] = p.l3*org.Value[0] + p.m3*org.Value[1] + p.n3*org.Value[2]
}

// 逆オイラー変換
func (p *Quartenion) EulerTransInverse(org, dest *Point) {
	dest.Value[0] = p.l1*org.Value[0] + p.l2*org.Value[1] + p.l3*org.Value[2]
	dest.Value[1] = p.m1*org.Value[0] + p.m2*org.Value[1] + p.m3*org.Value[2]
	dest.Value[2] = p.n1*org.Value[0] + p.n2*org.Value[1] + p.n3*org.Value[2]
}

// オイラー変換XYZ
func (p *Quartenion) EulerTransXYZ(ox, oy, oz, dx, dy, dz float64) {
	dx = p.l1*ox + p.m1*oy + p.n1*oz
	dy = p.l2*ox + p.m2*oy + p.n2*oz
	dz = p.l3*ox + p.m3*oy + p.n3*oz
}

// 逆オイラー変換XYZ
func (p *Quartenion) EulerTransInverseXYZ(ox, oy, oz, dx, dy, dz float64) {
	dx = p.l1*ox + p.l2*oy + p.l3*oz
	dy = p.m1*ox + p.m2*oy + p.m3*oz
	dz = p.n1*ox + p.n2*oy + p.n3*oz
}
