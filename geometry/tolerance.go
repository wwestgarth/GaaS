package geometry

import (
	"math"
)

const (
	spatialTolerance = 0.00000001
)

func distEq(dist1, dist2 float64) bool {
	return math.Abs(dist1-dist2) < spatialTolerance

}

func distZero(dist float64) bool {
	return distEq(dist, 0.0)
}
