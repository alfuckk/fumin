package biquge

type BookRuleSearch struct {
	Author       string `json:"author"`
	BookList     string `json:"bookList"`
	BookURL      string `json:"bookUrl"`
	CheckKeyWord string `json:"checkKeyWord"`
	CoverURL     string `json:"coverUrl"`
	Intro        string `json:"intro"`
	Kind         string `json:"kind"`
	Name         string `json:"name"`
}
