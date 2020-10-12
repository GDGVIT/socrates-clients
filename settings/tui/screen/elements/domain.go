package elements

import (
	// "log"
	"strings"

	"github.com/gdamore/tcell"
	"github.com/rivo/tview"
)

const (
	DomainFieldText = "Add a domain:"
)

type DomainBox struct {
	Field *tview.InputField
}

func (d DomainBox) GetText() string {
	return d.Field.GetText()
}

func NewDomainBox() DomainBox {
	d := DomainBox {
		tview.NewInputField(),
	}

	padding := strings.Repeat(" ", len(FreqFieldText) - len(DomainFieldText))

	d.Field = d.Field.
		SetLabel(DomainFieldText + padding).
		SetFieldWidth(30).
		SetDoneFunc(func(key tcell.Key) {
			return
		})

	return d
}