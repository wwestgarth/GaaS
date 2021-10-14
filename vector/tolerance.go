package vector

import (
	"math"
)

const (
	spatialTolerance = 0.00000001
)

func DistEq(dist1, dist2 float64) bool {
	return math.Abs(dist1-dist2) < spatialTolerance

}

func DistZero(dist float64) bool {
	return DistEq(dist, 0.0)
}
