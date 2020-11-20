package screen

import (
	"github.com/gdamore/tcell"
	"log"
	"strconv"
)

func (s *Screen) HandleInput(event *tcell.EventKey) *tcell.EventKey {
	switch event.Key() {
	case tcell.KeyEnter:
		if s.FreqField.HasFocus() {
			freqText := s.FreqField.GetText()
			freq, err := strconv.Atoi(freqText)
			if err != nil {
				log.Fatal(err)
			}
			s.config.Freq = freq

		} else if s.DomainField.HasFocus() {
			domain := s.DomainField.GetText()
			s.DomainField.SetText("")
			s.ButtonGrid.AddButton(domain)
			
		} else if s.ButtonGrid.HasFocus() {
			s.ButtonGrid.RemoveButton()
			s.refreshFocus()
		}
		// Make PUT request to API
		s.putUpdate()

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
