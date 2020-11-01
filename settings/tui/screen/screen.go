package screen

import (
	"settings/tui/screen/elements"
	"github.com/rivo/tview"
	"github.com/gdamore/tcell"
	"log"
)

type Screen struct {
	ButtonGrid 	*elements.ButtonsBox
	DomainField *elements.DomainBox
	FreqField 	*elements.FreqBox

	appGrid 	*tview.Grid
	inFocus		uint

	app			*tview.Application
}

func New(app *tview.Application) *Screen {
	btnBox := elements.NewButtonsBox()
	dmnBox := elements.NewDomainBox()
	freqBox := elements.NewFreqBox()

	// totalRows in the UI screen (excluding final padding row)
	const totalRows = 7				

	// rowDims is the height of each row. 1 => unit size, -1 => take all remaining space
	rowDims := make([]int, totalRows + 1)
	for i := 0; i < totalRows; i++ {
		rowDims[i] = 1
	}
	rowDims[totalRows] = -1		// padding row 

	paddingBx := tview.NewBox()

	appGrid := tview.NewGrid().
		SetColumns(-1).
		SetRows(rowDims...).
		// AddItem(primitive, rowNo, colNo, rowSpan, colSpan, minHeight, minWidth, visible)
		AddItem(freqBox.Field, 0, 0, 1, 1, 0, 0, true).
		AddItem(dmnBox.Field, 1, 0, 1, 1, 0, 0, true).
		AddItem(paddingBx, 2, 0, 3, 1, 0, 0, false).
		AddItem(btnBox.Grid, 5, 0, 1, 1, 0, 0, true).
		AddItem(paddingBx, 6, 0, -1, -1, 0, 0, false)

	appGrid.SetBackgroundColor(tview.Styles.PrimitiveBackgroundColor)
	
	return &Screen{
		btnBox,
		dmnBox,
		freqBox,
		appGrid,
		0,
		app,
	}
}

func (s *Screen) HandleInput(event *tcell.EventKey) *tcell.EventKey {
	switch event.Key() {
	case tcell.KeyEnter:
		if s.FreqField.HasFocus() {
			// freq := freqField.GetText()
		} else if s.DomainField.HasFocus() {
		domain := s.DomainField.GetText()
			s.DomainField.SetText("")
			err := s.ButtonGrid.AddButton(domain)
			if err != nil {
				log.Fatal(err)
			}
		} else if s.ButtonGrid.HasFocus() {
			s.ButtonGrid.RemoveButton()
			s.refreshFocus()
		}

	case tcell.KeyEsc:
		s.app.Stop()
		return nil
	
	case tcell.KeyUp:
		s.scrollUp()
	case tcell.KeyDown:
		s.scrollDown()
	case tcell.KeyLeft:
		s.scrollLeft()
	case tcell.KeyRight:
		s.scrollRight()
	}
	return event
}

func (s *Screen) scrollLeft() {
	if s.inFocus == 2 {
		s.ButtonGrid.ScrollLeft()
		s.refreshFocus()
	} 
	
}

func (s *Screen) scrollRight() {
	if s.inFocus == 2 {
		s.ButtonGrid.ScrollRight()
		s.refreshFocus()
	}
	
}

func (s *Screen) scrollUp() {
	if s.inFocus == 0 {
		s.inFocus = 2
	} else {
		s.inFocus = (s.inFocus - 1) % 3
	}
	s.refreshFocus()
}

func (s *Screen) scrollDown() {
	s.inFocus = (s.inFocus + 1) % 3
	s.refreshFocus()
}

func (s *Screen) refreshFocus() {
	if s.inFocus == 0 {
		s.app.SetFocus(s.FreqField.GetFocus())
	} else if s.inFocus == 1 {
		s.app.SetFocus(s.DomainField.GetFocus())
	} else {
		s.app.SetFocus(s.ButtonGrid.GetFocus())
	}
}

// Start the application
func (s *Screen) Start() {
	s.app.SetRoot(s.appGrid, true).SetFocus(s.FreqField.GetFocus())

	if err := s.app.Run(); err != nil {
		log.Fatal(err)
	}
}