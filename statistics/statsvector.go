package statistics

import (
	"math"
)

type StatsVector []float64

func (v *StatsVector) Add(n float64) int {
	*v = append(*v, n)
	return len(*v) - 1
}

func (v *StatsVector) Avg() float64 {
	return v.Sum() / float64(v.N())
}

func (v *StatsVector) DegFree() int {
	return v.N() - 2
}

func (v *StatsVector) HasValues() bool {
	return len(*v) > 0
}

func (v *StatsVector) IsEmpty() bool {
	return !v.HasValues()
}

func (v *StatsVector) Max() float64 {
	result := -math.MaxFloat64
	for _, n := range *v {
		result = math.Max(result, n)
	}
	return result
}

func (v *StatsVector) Min() float64 {
	result := math.MaxFloat64
	for _, n := range *v {
		result = math.Min(result, n)
	}
	return result
}

func (v *StatsVector) N() int {
	return len(*v)
}

func (v *StatsVector) Sum() float64 {
	var result float64
	for _, n := range *v {
		result = result + n
	}
	return result
}

func (v *StatsVector) StdDev() float64 {
	return math.Sqrt(v.Variance())
}

func (v *StatsVector) SumOfSquares() float64 {
	var result float64
	for _, n := range *v {
		result = result + (n * n)
	}
	return result
}

func (v *StatsVector) SumSqrDev() float64 {
	var result float64
	average := v.Avg()
	for _, n := range *v {
		result = result + (n-average)*(n-average)
	}
	return result
}

func (v *StatsVector) Variance() float64 {
	var result float64
	if v.N() > 2 {
		result = v.SumSqrDev() / float64(v.N()-1)
	}
	return result
}

func (v *StatsVector) X(i int) float64 {
	return (*v)[i]
}
