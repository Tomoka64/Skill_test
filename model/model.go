package model

type Result struct {
	Filename string `json:"filename"`
	Keyword  string `json:"keyword"`
	Line     int    `json:"line"`
	Detail   string `json:"detail"`
}

func NewResult(Filename, Keyword, Detail string, Line int) Result {
	return Result{
		Filename: Filename,
		Keyword:  Keyword,
		Detail:   Detail,
		Line:     Line,
	}
}
