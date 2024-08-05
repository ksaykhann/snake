package tuls

import (
	"math/rand"

	"github.com/gbin/goncurses"
)

type Apple struct {
	Exist      bool
	coordinats Position
}

func (a *Apple) Spawn(win *goncurses.Window) {
	if a.Exist {
		return
	}
	for !a.Exist {
		a.coordinats.X = rand.Intn(112)/4*2 + 2
		a.coordinats.Y = rand.Intn(21) + 1
		win.Move(a.coordinats.Y, a.coordinats.X)
		if win.InChar() == 32 {
			a.Exist = true

		}
	}

}

func (a Apple) Render(win *goncurses.Window) {
	win.MoveAddChar(a.coordinats.Y, a.coordinats.X, goncurses.ACS_DIAMOND)
}

func NewApple() Apple {
	apple := Apple{}
	apple.Exist = false
	return apple
}
