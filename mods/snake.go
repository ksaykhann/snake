// package tuls

// import (
// 	"github.com/gammazero/deque"
// 	"github.com/gbin/goncurses"
// )

// type Snake struct {
// 	Size       int
// 	OldVect    string
// 	Vect       string
// 	coordinats deque.Deque[position]
// 	DirectX    int
// 	DirectY    int
// 	Eated      bool
// 	Score      int
// }

// func (s *Snake) Move(a string, key goncurses.Key) int {
// 	head := s.coordinats.Back()
// 	if a == "a" || a == "s" || a == "d" || a == "w" && a != "" {
// 		s.Vect = a
// 	}
// 	if a != "" {
// 		switch {
// 		case s.Vect == "d" && s.OldVect != "a":
// 			s.DirectX = 2
// 			s.DirectY = 0
// 			s.OldVect = s.Vect
// 		case s.Vect == "a" && s.OldVect != "d":
// 			s.DirectX = -2
// 			s.DirectY = 0
// 			s.OldVect = s.Vect
// 		case s.Vect == "s" && s.OldVect != "w":
// 			s.DirectX = 0
// 			s.DirectY = 1
// 			s.OldVect = s.Vect
// 		case s.Vect == "w" && s.OldVect != "s":
// 			s.DirectX = 0
// 			s.DirectY = -1
// 			s.OldVect = s.Vect
// 		}
// 	}

// 	head.X += s.DirectX
// 	head.Y += s.DirectY

// 	if head.X <= 0 || head.X >= 59 || head.Y == 0 || head.Y == 23 {
// 		return 1
// 	}
// 	for i := 0; i < s.Size; i++ {
// 		if head == s.coordinats.At(i) {
// 			return 1
// 		}
// 		// if key == 27 {
// 		// 	os.Exit(0)
// 		// }
// 	}

// 	s.coordinats.PushBack(head)
// 	if s.Eated {
// 		s.Size++
// 		s.Score++
// 		return 0
// 	}
// 	s.coordinats.PopFront()

// 	return 0
// }

// func (s Snake) Render(win *goncurses.Window) {
// 	var k position
// 	for i := 0; i < s.Size; i++ {
// 		k = s.coordinats.At(i)
// 		win.MoveAddChar(k.Y, k.X, goncurses.ACS_DIAMOND)

// 	}
// }

// func NewSnake() Snake {
// 	s := Snake{}
// 	s.Size = 2
// 	s.OldVect = "d"
// 	s.Vect = "d"
// 	s.coordinats.PushBack(position{8, 24})
// 	s.coordinats.PushBack(position{8, 26})
// 	return s
// }

// func (s *Snake) Eat(a *Apple) {
// 	s.Eated = false
// 	head := s.coordinats.Back()
// 	if head != a.coordinats {
// 		return
// 	}
// 	a.Exist = false
// 	s.Eated = true
// 	goncurses.Beep()
// }

// type position struct {
// 	Y int
// 	X int
// }

package tuls

import (
	"github.com/gbin/goncurses"
)

type Snake struct {
	Size       int
	OldVect    string
	Vect       string
	coordinats QueList
	DirectX    int
	DirectY    int
	Eated      bool
	Score      int
}

func (s *Snake) Move(a string, key goncurses.Key) int {
	head := *s.coordinats.firstElem
	if a == "a" || a == "s" || a == "d" || a == "w" && a != "" {
		s.Vect = a
	}
	if a != "" {
		switch {
		case s.Vect == "d" && s.OldVect != "a":
			s.DirectX = 2
			s.DirectY = 0
			s.OldVect = s.Vect
		case s.Vect == "a" && s.OldVect != "d":
			s.DirectX = -2
			s.DirectY = 0
			s.OldVect = s.Vect
		case s.Vect == "s" && s.OldVect != "w":
			s.DirectX = 0
			s.DirectY = 1
			s.OldVect = s.Vect
		case s.Vect == "w" && s.OldVect != "s":
			s.DirectX = 0
			s.DirectY = -1
			s.OldVect = s.Vect
		}
	}

	head.queue.X += s.DirectX
	head.queue.Y += s.DirectY

	if head.queue.X <= 0 || head.queue.X >= 59 || head.queue.Y == 0 || head.queue.Y == 23 {
		return 1
	}
	for i := 1; i < s.Size; i++ {
		if head == *s.coordinats.At(i) {
			return 1
		}
		if key == 27 {
			return 1
		}
	}

	s.coordinats.PushFront(head.queue)
	if s.Eated {
		s.Size++
		s.Score++
		return 0
	}
	s.coordinats.PopBack()

	return 0
}

func (s Snake) Render(win *goncurses.Window) {
	var k *QueElem
	for i := 0; i < s.Size; i++ {
		k = s.coordinats.At(i)
		win.MoveAddChar(k.queue.Y, k.queue.X, goncurses.ACS_DIAMOND)

	}

}

func NewSnake() Snake {
	s := Snake{}
	s.Size = 2
	s.OldVect = "d"
	s.Vect = "d"
	// s.coordinats.head.queue.X = 24
	// s.coordinats.head.queue.Y = 8

	s.coordinats.PushFront(Position{8, 24})
	s.coordinats.PushFront(Position{8, 26})
	return s
}

func (s *Snake) Eat(a *Apple) {
	s.Eated = false

	current := s.coordinats.firstElem
	for current.next != nil {
		current = current.next
	}

	head := current

	if head.queue != a.coordinats {
		return
	}
	a.Exist = false
	s.Eated = true
	goncurses.Beep()
}

type Position struct {
	Y int
	X int
}
