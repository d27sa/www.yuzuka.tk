package translator

type result struct {
	Engine string
	Text   string
}

// newResult creates a new instance of result
func newResult(engine, text string) *result {
	var e string
	switch engine {
	case Google:
		e = "Google Translate"

	case Baidu:
		e = "百度翻译"
	}
	return &result{e, text}
}
