package main

import (
	tuls "mods/mods"
	"time"

	"github.com/gbin/goncurses"
)

func main() {
	goncurses.Init()
	goncurses.Cursor(0)
	goncurses.Echo(false)

	win, _ := goncurses.NewWindow(24, 60, 0, 0)
	win.Timeout(0)
	score, _ := goncurses.NewWindow(24, 20, 0, 61)

	var key goncurses.Key
	var stop int
	for {
		player := tuls.NewSnake()
		apple := tuls.NewApple()
		tuls.NewMenu()

		for {
			win.Box(goncurses.ACS_VLINE, goncurses.ACS_HLINE)
			score.MovePrint(10, 5, player.Score)
			player.Render(win)
			key = win.GetChar()

			apple.Spawn(win)
			apple.Render(win)
			stop = player.Move(string(key), key)
			player.Eat(&apple)
			if stop == 1 {
				break
			}
			if key == 27 {
				break
			}
			win.Refresh()
			score.Refresh()
			time.Sleep(80 * time.Millisecond)
			win.Erase()
			score.Erase()
		}
	}
}
