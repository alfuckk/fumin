package endpoints

type NovelListReq struct {
	Keyword string `json:"keyword"`
	Page    int64  `json:"page"`
}
