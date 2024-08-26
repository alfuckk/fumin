package collyfx

import "github.com/gocolly/colly/v2"

type Colly struct {
	*colly.Collector
}

func New() *Colly {
	return &Colly{colly.NewCollector()}
}
