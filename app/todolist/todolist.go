package todolist

// List stores things to do - temp
var List []string

// temp -
func init() {
	for i := 0; i < 30; i++ {
		List = append(List, "I must finish this app.")
	}
}
