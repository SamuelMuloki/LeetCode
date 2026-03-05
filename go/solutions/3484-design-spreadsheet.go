package solutions

import (
	"strconv"
	"strings"
)

type Spreadsheet struct {
	rows  int
	cells map[string]int
}

func SpreadsheetConstructor(rows int) Spreadsheet {
	return Spreadsheet{
		rows:  rows,
		cells: make(map[string]int),
	}
}

func (this *Spreadsheet) SetCell(cell string, value int) {
	this.cells[cell] = value
}

func (this *Spreadsheet) ResetCell(cell string) {
	this.cells[cell] = 0
}

func (this *Spreadsheet) GetValue(formula string) int {
	eq := string(formula[1:])
	arr := strings.Split(eq, "+")

	res := 0
	if val1, ok := this.cells[arr[0]]; ok {
		res += val1
	} else if num1, err := strconv.Atoi(arr[0]); err == nil {
		res += num1
	}

	if val2, ok := this.cells[arr[1]]; ok {
		res += val2
	} else if num2, err := strconv.Atoi(arr[1]); err == nil {
		res += num2
	}

	return res
}

/**
 * Your Spreadsheet object will be instantiated and called as such:
 * obj := Constructor(rows);
 * obj.SetCell(cell,value);
 * obj.ResetCell(cell);
 * param_3 := obj.GetValue(formula);
 */
