package biquge

type Category struct {
	Title string `json:"title"`
	URL   string `json:"url"`
}

type Categories []Category
