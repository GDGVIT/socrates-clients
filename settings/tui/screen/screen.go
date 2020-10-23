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
	bb := elements.NewButtonsBox()
	db := elements.NewDomainBox()
	fb := elements.NewFreqBox()

	paddingBx := tview.NewBox()

	appGrid := tview.NewGrid().
		SetColumns(-1).
		SetRows(1, 1, 1, 1, 1, 1, -1).
		AddItem(fb.Field, 0, 0, 1, 1, 0, 0, true).
		AddItem(db.Field, 1, 0, 1, 1, 0, 0, true).
		AddItem(paddingBx, 2, 0, 3, 1, 0, 0, false).
		AddItem(bb.Grid, 5, 0, 1, 1, 0, 0, true).
		AddItem(paddingBx, 6, 0, -1, -1, 0, 0, false)

	return &Screen{
		bb,
		db,
		fb,
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

func (s *Screen) Start() {
	s.app.SetRoot(s.appGrid, true).SetFocus(s.FreqField.GetFocus())

	if err := s.app.Run(); err != nil {
		log.Fatal(err)
	}
}