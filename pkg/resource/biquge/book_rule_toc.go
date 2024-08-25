package biquge

type BookRuleToc struct {
	ChapterList string `json:"chapterList"`
	ChapterName string `json:"chapterName"`
	ChapterURL  string `json:"chapterUrl"`
	NextTocURL  string `json:"nextTocUrl"`
}
