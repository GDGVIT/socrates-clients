package elements

import (
	// "log"
	// "strings"

	"github.com/gdamore/tcell"
	"github.com/rivo/tview"
)

const (
	FreqFieldText = "Number of papers per day (max 5):"
)

type FreqBox struct {
	Field *tview.InputField
}

func (f FreqBox) GetText() string {
	return f.Field.GetText()
}

func NewFreqBox() FreqBox {
	f := FreqBox {
		tview.NewInputField(),
	}

	padding := ""

	f.Field = f.Field.
		SetLabel(FreqFieldText + padding).
		SetFieldWidth(30).
		SetDoneFunc(func(key tcell.Key) {
			return
		})

	return f
}