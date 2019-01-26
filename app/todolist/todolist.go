package todolist

import (
	"container/list"
	"strconv"
)

const (
	Up   = uint8(0)
	Down = uint8(1)
)

// List stores things to do - temp
var List *list.List

// temp -
func init() {
	List = list.New()
	for i := 0; i < 30; i++ {
		List.PushBack("I must finish this app." + strconv.Itoa(i))
	}
}

// Del deletes a string from a list
func Del(l *list.List, s string) {
	e := l.Front()
	for e != nil {
		if e.Value.(string) == s {
			l.Remove(e)
			return
		}
		e = e.Next()
	}
}

func Move(l *list.List, s string, dir uint8) {
	e := l.Front()
	for e != nil {
		if e.Value.(string) == s {
			if dir == Up {
				if e == l.Front() {
					return
				}
				p := e.Prev()
				l.Remove(e)
				l.InsertBefore(s, p)
			} else {
				if e.Next() == nil {
					return
				}
				n := e.Next()
				l.Remove(e)
				l.InsertAfter(s, n)
			}
		}
		e = e.Next()
	}
}
