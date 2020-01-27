package main

import (
	"fyne.io/fyne/widget"
	"strings"
)

// field.go handles numeric entry fields (used for entry of hours/mins/secs)

type (
	NumericField struct {
		*widget.Entry
		width int
		valid func(int) bool // indicates if the values is valid
	}
)

func NewNumericFieldRange(min, max int) *NumericField {
	if min < 0 || min >= max {
		panic("Invalid range for NewNumericFieldRange")
	}
	entry := widget.NewEntry()
	digits := countDigits(max)
	entry.Text = strings.Repeat("0", digits)
	return &NumericField{
		Entry: entry,
		width: digits,
		valid: func(v int) bool {
			return v >= min && v < max
		},
	}
}

func NewNumericField(width int, valid func(int) bool) *NumericField {
	entry := widget.NewEntry()
	entry.Text = strings.Repeat("0", width)
	return &NumericField{Entry: entry, width: width, valid: valid}
}

func countDigits(v int) int {
	count := 0
	for ; v != 0; v /= 10 {
		count++
	}
	return count
}
