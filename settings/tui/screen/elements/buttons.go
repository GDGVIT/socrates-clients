package elements

import (
	"log"
	"github.com/GDGVIT/socrates/schema"
	"github.com/gdamore/tcell"
	"errors"
	"github.com/rivo/tview"
)

type ButtonsBox struct {
	Grid *tview.Grid
	buttons []*tview.Button
	totalButtons int
	currentFocus int
	config		*schema.Config
}

func NewButtonsBox(c *schema.Config) *ButtonsBox {
	b := ButtonsBox {
		tview.NewGrid(),
		make([]*tview.Button, 0, BtnBoxTotalCols),
		0,
		0,
		c,
	}
	
	l := make([]int, BtnBoxTotalCols)
	for i := range l {
		l[i] = -1
	}

	b.Grid = b.Grid.
		SetColumns(l...).
		SetRows(4)

	// Initialize buttonbox with topic buttons acc to topics from API
	for _, topic := range c.Topics {
		b.RenderButton(topic)
	}

	return &b
}

// RenderButton renders a button in UI but does not update the global schema.Config instance
func (b *ButtonsBox) RenderButton(text string) error {
	if b.totalButtons == BtnBoxTotalCols {
		return errors.New("Maximum domains of interest reached")
	}

	btn := tview.NewButton(text)
	btn.SetBackgroundColorActivated(tcell.ColorOrangeRed)
	b.buttons = append(b.buttons, btn)
	idx := b.totalButtons
	b.Grid.AddItem(btn, 0, idx, 1, 1, 0, 0, true)
	b.totalButtons ++

	return nil
}

// AddButton adds button to UI and updates the global schema.Config instance
func (b *ButtonsBox) AddButton(topic string) {
	err := b.RenderButton(topic)
	if err != nil {
		log.Println(err)
	}

	b.config.Topics = append(b.config.Topics, topic)
}

// RemoveButton removes button in focus from UI and updates global schema.Config instance
func (b *ButtonsBox) RemoveButton() error {
	if b.totalButtons == 0 {
		return errors.New("No domains left to remove")
	}

	idx := b.currentFocus
	btn := b.buttons[idx]

	// Remove button at idx: buttons = buttons[0..idx-1, idx+1..]
	b.buttons = append(b.buttons[:idx], b.buttons[idx+1:]...)
	// Remove topic from config at idx, same logic
	b.config.Topics = append(b.config.Topics[:idx], b.config.Topics[idx+1:]...)

	b.Grid.RemoveItem(btn)
	b.totalButtons --
	if b.currentFocus != 0 {
		b.currentFocus --
	}
	return nil
}

func (b *ButtonsBox) ScrollLeft() {
	if b.totalButtons < 1 {
		return
	}

	if b.currentFocus == 0 {
		b.currentFocus = b.totalButtons - 1
	} else {
		b.currentFocus --
	}
}

func (b *ButtonsBox) ScrollRight() {
	if b.totalButtons < 1 {
		return
	}

	if b.currentFocus == b.totalButtons - 1 {
		b.currentFocus = 0
	} else {
		b.currentFocus ++
	}
}

func (b *ButtonsBox) GetFocus() tview.Primitive {
	if b.totalButtons < 1 {
		return b.Grid
	}

	return b.buttons[b.currentFocus]
}

func (b *ButtonsBox) HasFocus() bool {
	return b.Grid.HasFocus()
}
