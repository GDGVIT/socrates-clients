package screen

import (
	"github.com/GDGVIT/socrates/schema"
	"tui/screen/elements"
	"github.com/rivo/tview"
	"log"
)

type Screen struct {
	ButtonGrid 	*elements.ButtonsBox
	DomainField *elements.DomainBox
	FreqField 	*elements.FreqBox

	appGrid 	*tview.Grid
	inFocus		uint

	config		*schema.Config
	port		string

	app			*tview.Application
}

func New(app *tview.Application, config *schema.Config, port string) *Screen {
	dmnBox := elements.NewDomainBox()
	btnBox := elements.NewButtonsBox(config)
	freqBox := elements.NewFreqBox(config)

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
		config,
		port,
		app,
	}
}

// Start the application
func (s *Screen) Start() {
	s.app.SetRoot(s.appGrid, true).SetFocus(s.FreqField.GetFocus())

	if err := s.app.Run(); err != nil {
		log.Fatal(err)
	}
}