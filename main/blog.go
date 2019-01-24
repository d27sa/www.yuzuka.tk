package main

// temp -
type postIndex struct {
	Title  string
	Author string
}

// temp -
func newPostIndex(title, author string) *postIndex {
	return &postIndex{title, author}
}

// temp -
func initPostList() {
	for i := 0; i < 50; i++ {
		postList = append(postList, newPostIndex("this is just a test post owo", "admin"))
	}
}
