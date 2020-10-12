package elements

import (
	// "log"
	// "github.com/gdamore/tcell"
	"errors"
	"github.com/rivo/tview"
)

const totalCols = 8

type ButtonsBox struct {
	Grid *tview.Grid
	buttons []*tview.Button
}

func NewButtonsBox() ButtonsBox {
	b := ButtonsBox {
		tview.NewGrid(),
		make([]*tview.Button, totalCols),
	}

	l := make([]int, totalCols)
	for i := range l {
		l[i] = -1
	}

	b.Grid = b.Grid.
		SetColumns(l...).
		SetRows(4)

	return b
}

func (b ButtonsBox) AddButton(text string) error {
	if len(b.buttons) == totalCols {
		return errors.New("Maximum domains of interest reached")
	}

	btn := tview.NewButton(text)
	b.buttons = append(b.buttons, btn)
	idx := len(b.buttons)
	b.Grid.AddItem(btn, 0, idx, 1, 1, 0, 0, true)

	return nil
}

func (b ButtonsBox) RemoveButton(btn *tview.Button) error {
	if len(b.buttons) == 0 {
		return errors.New("No domains left to remove")
	}

	idx := b.findButton(btn)
	if idx == -1 {
		return errors.New("Error: Attempt to remove domain not in list")
	}

	b.buttons = append(b.buttons[:idx], b.buttons[idx+1:]...)
	b.Grid.RemoveItem(btn)
	return nil
}

func (b ButtonsBox) findButton(btn *tview.Button) int {
	for i := range b.buttons {
		if b.buttons[i] == btn {
			return i
		}
	}

	return -1
}