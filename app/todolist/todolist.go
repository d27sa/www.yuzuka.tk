package todolist

import (
	"container/list"
	"fmt"
	"strconv"
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
	f := l.Front()
	for f != nil {
		if f.Value.(string) == s {
			l.Remove(f)
			fmt.Println(s + " deleted.")
			return
		}
		f = f.Next()
	}
}
