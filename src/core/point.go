// Point型　coreパッケージ
package core

import (
	_ "fmt"
	"math"
)

// Point型
type Point struct {
	Value [3]float64
}

// 地球半径[m]
var earthRadius float64 = 6371000.0

// Pointのコンストラクタ　（ECF座標系で直にセット）
func NewPoint(v [3]float64) *Point {
	// p := new(Point)
	var p *Point = new(Point) // こちらの方が型が明確に分かるので、いい気がする
	for i, _ := range v {
		p.Value[i] = v[i]
	}
	return p
}

// Pointの値セット
func (p *Point) Set(pos *Point) {
	for i, _ := range p.Value {
		p.Value[i] = pos.Value[i]
	}
}

// Pointの加算（自分自身に加算する）
// (p自体もポインタにしておかないと変更結果を外に繁栄できない)
func (p *Point) Add(v *Point) {
	for i, _ := range p.Value {
		p.Value[i] = p.Value[i] + v.Value[i]
	}
}

// Pointの減算（自分自身に減算する）
func (p *Point) Sub(v *Point) {
	for i, _ := range p.Value {
		p.Value[i] = p.Value[i] - v.Value[i]
	}
}

// Pointとして地球半径を取り出す
func (p *Point) GetRadius() float64 {
	return earthRadius
}

// Pointとして地球半径を変更する
func (p *Point) SetRadius(r float64) {
	earthRadius = r
}

// 緯度の取得 GetLat()
func (p *Point) GetLat() float64 {
	r := math.Atan2(p.Value[2], math.Sqrt(p.Value[0]*p.Value[0]+p.Value[1]*p.Value[1]))
	return r
}

// 経度の取得 GetLon()
func (p *Point) GetLon() float64 {
	r := math.Atan2(p.Value[1], p.Value[0])
	return r
}

// 高度の取得 GetAlt()
func (p *Point) GetAlt() float64 {
	r := p.GetLength() - earthRadius
	return r
}

// 長さの取得 GetLength()
func (p *Point) GetLength() float64 {
	r := math.Sqrt(p.Value[0]*p.Value[0] + p.Value[1]*p.Value[1] + p.Value[2]*p.Value[2])
	return r
}

// 長さの２乗の取得 GetDoubleLength()
func (p *Point) GetDoubleLength() float64 {
	r := p.Value[0]*p.Value[0] + p.Value[1]*p.Value[1] + p.Value[2]*p.Value[2]
	return r
}

// ２点の距離の取得 DistanceTo()
func (p *Point) DistanceTo(v *Point) float64 {
	xx := p.Value[0] - v.Value[0]
	yy := p.Value[1] - v.Value[1]
	zz := p.Value[2] - v.Value[2]
	r := math.Sqrt(xx*xx + yy*yy + zz*zz)
	return r
}

// 地表面上の距離の取得 DownRangeTo()
func (p *Point) DownRangeTo(v *Point) float64 {
	org := p.GetDoubleLength()
	dest := v.GetDoubleLength()

	rangeD := p.DistanceTo(v)
	rangeD = rangeD * rangeD

	cosValue := (org + dest - rangeD) / 2.0 / math.Sqrt(org*dest)
	if cosValue > 1.0 {
		cosValue = 1.0
	} else if cosValue < -1.0 {
		cosValue = -1.0
	}

	r := earthRadius * math.Acos(cosValue)
	return r
}

// local座標系に変換 Localize()
func (p *Point) Localize(origin *Point) *Point {
	dif := NewPoint([3]float64{})

	dif.Value[0] = p.Value[0] - origin.Value[0]
	dif.Value[1] = p.Value[1] - origin.Value[1]
	dif.Value[2] = p.Value[2] - origin.Value[2]

	localLon := origin.GetLon()
	localLat := origin.GetLat()

	sinLat := math.Sin(localLat)
	cosLat := math.Cos(localLat)
	sinLon := math.Sin(localLon)
	cosLon := math.Cos(localLon)

	xdot := -dif.Value[0]*sinLat*cosLon - dif.Value[1]*sinLat*sinLon + dif.Value[2]*cosLat
	ydot := -dif.Value[0]*sinLon + dif.Value[1]*cosLon
	zdot := -dif.Value[0]*cosLat*cosLon - dif.Value[1]*cosLat*sinLon - dif.Value[2]*sinLat

	return NewPoint([3]float64{xdot, ydot, zdot})
}

// Global座標系に変換 Globalize()
func (p *Point) Globalize(origin *Point) *Point {
	localLon := origin.GetLon()
	localLat := origin.GetLat()

	sinLat := math.Sin(localLat)
	sinLon := math.Sin(localLon)
	cosLat := math.Cos(localLat)
	cosLon := math.Cos(localLon)

	xdot := -p.Value[1]*sinLon - p.Value[0]*cosLon*sinLat - p.Value[2]*cosLat*cosLon
	ydot := p.Value[1]*cosLon - p.Value[0]*sinLat*sinLon - p.Value[2]*cosLat*sinLon
	zdot := p.Value[0]*cosLat - p.Value[2]*sinLat

	pos := NewPoint([3]float64{xdot, ydot, zdot})
	pos.Value[0] += origin.Value[0]
	pos.Value[1] += origin.Value[1]
	pos.Value[2] += origin.Value[2]

	return pos
}
