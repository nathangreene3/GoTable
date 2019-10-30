package table

import (
	"encoding/csv"
	"fmt"
	gomath "math"
	"os"
	"sort"
	"strings"
	"testing"

	"github.com/nathangreene3/math"
)

func TestImportExportCSV(t *testing.T) {
	inFile, err := os.Open("test0.csv")
	if err != nil {
		t.Fatalf("%v", err)
	}
	defer inFile.Close()

	table, err := Import(*csv.NewReader(inFile), "star wars", FltFmtNoExp, 3)
	if err != nil {
		t.Fatalf("%v", err)
	}

	outFile, err := os.OpenFile("test1.csv", os.O_WRONLY, os.ModeAppend)
	if err != nil {
		t.Fatalf("%v", err)
	}
	defer outFile.Close()

	if err = table.Export(*csv.NewWriter(outFile)); err != nil {
		t.Fatalf("%v", err)
	}

	// t.Fatalf("\n%s", table.String())
}

func TestTable1(t *testing.T) {
	type factPow struct {
		factor, power int
	}

	var (
		n     = 1 << 12
		tbl   = New("Squared-Triangle Numbers", FltFmtNoExp, 0)
		facts = func(n int) string {
			if n < 1 {
				return ""
			}

			var (
				fs      = math.Factor(n)
				factors = make([]string, 0, len(fs))
			)

			for fact, pow := range fs {
				factors = append(factors, fmt.Sprintf("%d^%d", fact, pow))
			}

			sort.Strings(factors)
			return strings.Join(factors, " * ")
		}
	)

	tbl.SetHeader(Header{"x", "y", "(x^2+x)/2", "y^2", "x+y", "x-y", "x^2-y^2", "y^2-x", "facts(x)", "facts(y)"})
	for x := 0; x < n; x++ {
		var (
			x2 = x * x
			T  = (x2 + x) >> 1
		)

		for y := 0; y < n; y++ {
			if S := y * y; T == S {
				tbl.AppendRow(Row{x, y, T, S, x + y, x - y, x2 - S, S - x, facts(x), facts(y)})
			}
		}
	}

	t.Fatalf("\n%s\n", tbl.String())
}

func TestTable2(t *testing.T) {
	var (
		x0, y0, x0to2, y0to2, x1to2 float64
		numRows                     = 10
		tbl                         = New("Solutions to Pell's Equation for n = 2", 0, 3)
	)

	tbl.SetHeader(Header{"k", "x", "y", "x^2 - 2y^2"})
	x0 = 1
	for k := 0; k < numRows; k++ {
		x0to2, y0to2 = x0*x0, y0*y0
		tbl.AppendRow(NewRow(k, x0, y0, x0to2-2*y0to2))

		x1to2 = 3.0*x0to2 + 4.0*x0*y0
		x0, y0 = gomath.Sqrt(x1to2), gomath.Sqrt((x1to2-x0to2)/2.0+y0to2)
	}

	t.Fatalf("\n%s\n", tbl.String())
}

func TestApproxSqrt2(t *testing.T) {
	var (
		x0, y0, x0to2, y0to2, x1to2 float64
		numRows                     = 10
		tbl                         = New("Approximations of sqrt(2)", 0, 9)
	)

	tbl.SetHeader(Header{"k", "x", "y", "~sqrt(2)"})
	x0 = 1
	for k := 0; k < numRows; k++ {
		x0to2, y0to2 = x0*x0, y0*y0
		tbl.AppendRow(NewRow(k, x0, y0, gomath.Sqrt(x0to2-1)/y0))

		x1to2 = 3.0*x0to2 + 4.0*x0*y0
		x0, y0 = gomath.Sqrt(x1to2), gomath.Sqrt((x1to2-x0to2)/2.0+y0to2)
	}

	t.Fatalf("\n%s\n", tbl.String())
}
