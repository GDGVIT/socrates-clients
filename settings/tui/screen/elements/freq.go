package elements

import (
	"fmt"
	"github.com/GDGVIT/socrates/schema"
	"github.com/gdamore/tcell"
	"github.com/rivo/tview"
)

type FreqBox struct {
	Field *tview.InputField
}

func (f *FreqBox) GetText() string {
	return f.Field.GetText()
}

func NewFreqBox(c *schema.Config) *FreqBox {
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

	if c.Freq != 0 {
		f.Field.SetText(
			fmt.Sprint(c.Freq),
		)
	}

	return &f
}

func (f *FreqBox) HasFocus() bool {
	return f.Field.HasFocus()
}

func (f *FreqBox) GetFocus() *tview.InputField {
	return f.Field
}