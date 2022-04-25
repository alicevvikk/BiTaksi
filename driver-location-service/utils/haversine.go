package utils

import (
	"math"
)

const r = 6371.0

func toRadian(angle float64) float64 {
	return (angle * math.Pi) / 180
}

func hav(angle float64) float64{
	return (1 - math.Cos(angle)) / 2
}

func CalculateDistance(c1 []float64, c2 []float64) (float64) {
	
	long1   := toRadian(c1[1])
	lat1	:= toRadian(c1[0])
	long2	:= toRadian(c2[1])
	lat2	:= toRadian(c2[0])

	var first = hav(lat2 - lat1)
	var sec = 1 - hav(lat1 - lat2) - hav(lat1 + lat2)
	var th = hav(long2 - long1)

	return 2 * r * math.Asin(math.Sqrt(first + (sec * th)))
}


