package statistics

import (
	"math"
)

var CorrelationMode int

const (
	ForcedOrigin = 1 + iota
	Intercept
)

func init() {
	SetCorrelationMode(ForcedOrigin)
}

func SetCorrelationMode(mode int) {
	CorrelationMode = mode
}

type Regression struct {
	xVector StatsVector
	yVector StatsVector
}

func (r *Regression) AddXY(x, y float64) {
	r.xVector.Add(x)
	r.yVector.Add(y)
}

func (r *Regression) Correlation() float64 {
	return r.CoVariance() / math.Sqrt(r.xVector.Variance()*r.yVector.Variance())
}

func (r *Regression) CoVariance() float64 {
	return r.Sxy()
}

func (r *Regression) N() int {
	return r.xVector.N()
}

func (r *Regression) ResidualSumSqr() float64 {
	return (1 - r.RSquared()) * r.yVector.SumSqrDev()
}

func (r *Regression) RSquared() float64 {
	return r.Correlation() * r.Correlation()
}

func (r *Regression) SumOfProducts() float64 {
	if r.xVector.N() != r.yVector.N() {
		panic("x and y vectors are of different size")
	}
	i := 0
	var result float64
	for {
		result = result + r.xVector[i]*r.yVector[i]
		if i == r.xVector.N()-1 {
			break
		}
		i++
	}
	return result
}

func Sqr(n float64) float64 {
	return n * n
}

func (r *Regression) Sxx() float64 {
	return r.xVector.SumOfSquares() - (Sqr(r.xVector.Sum()) / float64(r.N()))
}

func (r *Regression) Sxy() float64 {
	return r.SumOfProducts() - r.xVector.Sum()*r.yVector.Sum()/float64(r.N())
}

func (r *Regression) Syy() float64 {
	return r.yVector.SumOfSquares() - (Sqr(r.yVector.Sum()) / float64(r.N()))
}

func (r *Regression) S2() float64 {
	switch {
	case CorrelationMode == ForcedOrigin:
		return (r.yVector.SumOfSquares() - Sqr(r.SumOfProducts())/r.xVector.SumOfSquares()) / float64(r.N()-1)
	case CorrelationMode == Intercept:
		return (r.Syy() - Sqr(r.Sxy())/r.Sxx()) / (float64(r.N()) - 2)
	default:
		return 0
	}
}

func (r *Regression) Slope() float64 {
	switch {
	case CorrelationMode == ForcedOrigin:
		return r.SumOfProducts() / r.xVector.SumOfSquares()
	case CorrelationMode == Intercept:
		return r.Sxy() / r.Sxx()
	default:
		return 0
	}
}

func (r *Regression) StdErrSlope() float64 {
	switch {
	case CorrelationMode == ForcedOrigin:
		return math.Sqrt(r.S2() / r.xVector.SumOfSquares())
	case CorrelationMode == Intercept:
		return math.Sqrt(r.S2() / r.Sxx())
	default:
		return 0
	}
}

func (r *Regression) YIntercept() float64 {
	switch {
	case CorrelationMode == ForcedOrigin:
		return 0
	case CorrelationMode == Intercept:
		return r.yVector.Avg() - r.Slope()/r.xVector.Avg()
	default:
		return 0
	}
}
