package square


import (
	"math"
)


type (
	intCustomType int
)

const (
	SidesCircle   intCustomType = 0
	SidesTriangle intCustomType = 3
	SidesSquare   intCustomType = 4
)

func CalcSquare(sideLen float64, sidesNum intCustomType) float64 {
	if (sidesNum == SidesSquare){
		return (sideLen * sideLen)
	} else if (sidesNum == SidesTriangle){
		return((sideLen * sideLen * math.Sqrt(3)) /4 )
	}else if (sidesNum == SidesCircle){
		return ((sideLen * sideLen)* math.Pi )
	}
	return(0)
}
