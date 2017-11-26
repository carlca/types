package statistics

import (
	"testing"
)

func TestRegression(t *testing.T) {
	var r Regression
	r = PopulateRegression(r)
}

func PopulateRegression(r Regression) Regression {
	r.AddXY(1, 6)
	r.AddXY(1, 7)
	r.AddXY(2, 4)
	r.AddXY(3, 9)
	r.AddXY(5, 14)
	r.AddXY(4, 12)
	r.AddXY(7, 20)
	r.AddXY(5, 17)
	r.AddXY(8, 20)
	r.AddXY(2, 5)
	r.AddXY(5, 15)
	r.AddXY(5, 26)
	return r
}
