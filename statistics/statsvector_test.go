package statistics

import (
	"fmt"
	"strconv"
	"testing"

	"github.com/franela/goblin"
)

func TestStatsVector(t *testing.T) {
	var v StatsVector
	v = populateVector(v)
	var expected string
	var actual string
	g := goblin.Goblin(t)

	actual = fmt.Sprintf("%10.5f", v.Avg())
	expected = "2677.04503"
	g.Describe("Avg", func() {
		g.It("Should return the average of the consistuent elements", func() {
			ReportAndAssert(g, expected, actual)
		})
	})

	actual = fmt.Sprint(v.DegFree())
	expected = "8"
	g.Describe("DegFree", func() {
		g.It("Should return the degrees of freedom of the vector (N - 2)", func() {
			ReportAndAssert(g, expected, actual)
		})
	})

	actual = strconv.FormatBool(v.HasValues())
	expected = "true"
	g.Describe("HasValues", func() {
		g.It("Should return true or false depending on whether the vector contains any values - inverse of IsEmpty", func() {
			ReportAndAssert(g, expected, actual)
		})
	})

	actual = strconv.FormatBool(v.IsEmpty())
	expected = "false"
	g.Describe("IsEmpty", func() {
		g.It("Should return true or false depending on whether the vector contains any values - inverse of HasValues", func() {
			ReportAndAssert(g, expected, actual)
		})
	})

	actual = fmt.Sprintf("%9.4f", v.Max())
	expected = "7952.3664"
	g.Describe("Max", func() {
		g.It("Should return the largest numerical value in the vector", func() {
			ReportAndAssert(g, expected, actual)
		})
	})

	actual = fmt.Sprintf("%7.3f", v.Min())
	expected = "123.522"
	g.Describe("Min", func() {
		g.It("Should return the smallest numerical value in the vector", func() {
			ReportAndAssert(g, expected, actual)
		})
	})

	actual = fmt.Sprint(v.N())
	expected = "10"
	g.Describe("N", func() {
		g.It("Should return the number of elements in the vector", func() {
			ReportAndAssert(g, expected, actual)
		})
	})

	actual = fmt.Sprintf("%10.4f", v.Sum())
	expected = "26770.4503"
	g.Describe("Sum", func() {
		g.It("Should provide the sum of the constituent elements", func() {
			ReportAndAssert(g, expected, actual)
		})
	})

	actual = fmt.Sprintf("%11.6f", v.StdDev())
	expected = "2662.969424"
	g.Describe("StdDev", func() {
		g.It("Should provide the standard deviation of the vector", func() {
			ReportAndAssert(g, expected, actual)
		})
	})

	actual = fmt.Sprintf("%11.1f", v.SumOfSquares())
	expected = "135488356.3"
	g.Describe("SumOfSquares", func() {
		g.It("Should provide the sum of squares of the vector", func() {
			ReportAndAssert(g, expected, actual)
		})
	})

	actual = fmt.Sprintf("%11.2f", v.SumSqrDev())
	expected = "63822655.39"
	g.Describe("SumSqrDev", func() {
		g.It("Should provide the sum of squares of deviations the vector", func() {
			ReportAndAssert(g, expected, actual)
		})
	})

	actual = fmt.Sprintf("%11.3f", v.Variance())
	expected = "7091406.155"
	g.Describe("Variance", func() {
		g.It("Should provide the Variance of the vector", func() {
			ReportAndAssert(g, expected, actual)
		})
	})

	actual = fmt.Sprintf("%7.3f", v.X(2))
	expected = "894.735"
	g.Describe("X", func() {
		g.It("Should provide the third element of the vector", func() {
			ReportAndAssert(g, expected, actual)
		})
	})

}

func ReportAndAssert(g *goblin.G, expected, actual string) {
	if expected == actual {
		PrintCheck()
	} else {
		PrintFail()
	}
	fmt.Print(MakeGray("Expected: " + expected))
	fmt.Println(MakeGray("   Actual: " + actual))
	g.Assert(actual).Equal(expected)
}

func PrintCheck() {
	fmt.Print("    \033[32m\u2713\033[0m ")
}

func PrintFail() {
	fmt.Print("    \033[31m" + "x" + "\033[0m ")
}

func MakeGray(s string) string {
	return "\033[90m" + s + "\033[0m"
}

func populateVector(v StatsVector) StatsVector {
	v.Add(123.522)
	v.Add(5634.93)
	v.Add(894.735)
	v.Add(4562.378)
	v.Add(3567.456)
	v.Add(125.698)
	v.Add(798.3644)
	v.Add(7952.3664)
	v.Add(2123.654)
	v.Add(987.3465)
	return v
}
