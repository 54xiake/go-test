package main

import (
	"fmt"
	"math"
	"strconv"
)

// 数学公式arithmetic formula

func main() {
	fmt.Println("长2米宽4米矩形周长：", RectangularCircumference(2, 4), "m")
	fmt.Println("长2米宽4米矩形面积：", RectangularArea(2, 4), "m²")
	fmt.Println("半径4米圆面积：", CircularArea(4), "m²")
}

// 周长面积体积公式
// 长方形的周长=（长+宽）×2
func RectangularCircumference(x, y float64) float64 {
	return (x + y) * 2
}

// 求矩形面积 = 长*宽
func RectangularArea(x, y float64) float64 {
	return x * y
}

// 圆周长
func Circumference() {

}

// 圆面积 S=πr²,S=π(d/2)²,S=πd²/4
func CircularArea(r float64) float64 {
	return Decimal(math.Pi * math.Sqrt(r))
}

// 先通过Sprintf保留两位小数，再转成float64
func Decimal(value float64) float64 {
	value, _ = strconv.ParseFloat(fmt.Sprintf("%.2f", value), 64)
	return value
}

// 方差公式

// 力学公式

// 电学公式
//1、电流强度：I＝Q电量/t
//2、电阻：R=ρL/S
//3、欧姆定律：I＝U/R
//4、焦耳定律：
//（1）、Q＝I2Rt普适公式)
//（2）、Q＝UIt＝Pt＝UQ电量＝U2t/R (纯电阻公式)
//5、串联电路：
//（1）、I＝I1＝I2
//（2）、U＝U1＋U2
//（3）、R＝R1＋R2
//（4）、U1/U2＝R1/R2 (分压公式)
//（5）、P1/P2＝R1/R2
//6、并联电路：
//（1）、I＝I1＋I2
//（2）、U＝U1＝U2
//（3）、1/R＝1/R1＋1/R2 [ R＝R1R2/(R1＋R2)]
//（4）、I1/I2＝R2/R1(分流公式)
//（5）、P1/P2＝R2/R1
//7定值电阻：
//（1）、I1/I2＝U1/U2
//（2）、P1/P2＝I12/I22
//（3）、P1/P2＝U12/U22
//8电功：
//（1）、W＝UIt＝Pt＝UQ (普适公式)
//（2）、W＝I2Rt＝U2t/R (纯电阻公式)
//9电功率：
//（1）、P＝W/t＝UI (普适公式)
//（2）、P＝I2R＝U2/R (纯电阻公式)
