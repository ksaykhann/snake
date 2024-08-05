package tuls

import (
	"os"

	"github.com/gbin/goncurses"
)

type Menus struct {
}

func NewMenu() {
	Startmenu, _ := goncurses.NewWindow(24, 80, 0, 0)
	for {
		Startmenu.Box(goncurses.ACS_VLINE, goncurses.ACS_HLINE)
		Startmenu.MovePrint(12, 29, "Press Enter to start game")
		t := Startmenu.GetChar()
		Startmenu.Erase()
		Startmenu.Refresh()
		if t == 10 {
			break
		}
		if t == 27 {
			os.Exit(0)
		}

	}
}
