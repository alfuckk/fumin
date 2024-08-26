package novel

type RuleBookInfo struct {
	Author      string `json:"author"`
	CanReName   string `json:"canReName"`
	CoverURL    string `json:"coverUrl"`
	Init        string `json:"init"`
	Intro       string `json:"intro"`
	Kind        string `json:"kind"`
	LastChapter string `json:"lastChapter"`
	Name        string `json:"name"`
	TocURL      string `json:"tocUrl"`
	WordCount   string `json:"wordCount"`
}

type RuleContent struct {
	Content        string `json:"content"`
	ImageStyle     string `json:"imageStyle"`
	NextContentURL string `json:"nextContentUrl"`
	ReplaceRegex   string `json:"replaceRegex"`
	SourceRegex    string `json:"sourceRegex"`
}

type RuleExplore struct {
	Author      string `json:"author"`
	BookList    string `json:"bookList"`
	BookURL     string `json:"bookUrl"`
	CoverURL    string `json:"coverUrl"`
	Intro       string `json:"intro"`
	Kind        string `json:"kind"`
	LastChapter string `json:"lastChapter"`
	Name        string `json:"name"`
	WordCount   string `json:"wordCount"`
}

type RuleSearch struct {
	Author       string `json:"author"`
	BookList     string `json:"bookList"`
	BookURL      string `json:"bookUrl"`
	CheckKeyWord string `json:"checkKeyWord"`
	CoverURL     string `json:"coverUrl"`
	Intro        string `json:"intro"`
	Kind         string `json:"kind"`
	LastChapter  string `json:"lastChapter"`
	Name         string `json:"name"`
	WordCount    string `json:"wordCount"`
}

type RuleToc struct {
	ChapterList string `json:"chapterList"`
	ChapterName string `json:"chapterName"`
	ChapterURL  string `json:"chapterUrl"`
	FormatJs    string `json:"formatJs"`
	IsVip       string `json:"isVip"`
	IsVolume    string `json:"isVolume"`
	PreUpdateJs string `json:"preUpdateJs"`
	UpdateTime  string `json:"updateTime"`
}

type Novel struct {
	BookSourceComment string       `json:"bookSourceComment,omitempty"`
	BookSourceGroup   string       `json:"bookSourceGroup"`
	BookSourceName    string       `json:"bookSourceName"`
	BookSourceType    int          `json:"bookSourceType"`
	BookSourceURL     string       `json:"bookSourceUrl"`
	CustomOrder       int          `json:"customOrder"`
	Enabled           bool         `json:"enabled"`
	EnabledCookieJar  bool         `json:"enabledCookieJar"`
	EnabledExplore    bool         `json:"enabledExplore"`
	ExploreURL        string       `json:"exploreUrl,omitempty"`
	Header            string       `json:"header,omitempty"`
	LastUpdateTime    int64        `json:"lastUpdateTime"`
	LoginUI           string       `json:"loginUi,omitempty"`
	LoginURL          string       `json:"loginUrl,omitempty"`
	RespondTime       int          `json:"respondTime"`
	RuleBookInfo      RuleBookInfo `json:"ruleBookInfo"`
	RuleContent       RuleContent  `json:"ruleContent"`
	RuleExplore       RuleExplore  `json:"ruleExplore"`
	RuleSearch        RuleSearch   `json:"ruleSearch"`
	RuleToc           RuleToc      `json:"ruleToc"`
	SearchURL         string       `json:"searchUrl,omitempty"`
	VariableComment   string       `json:"variableComment,omitempty"`
	Weight            int          `json:"weight"`
	BookURLPattern    string       `json:"bookUrlPattern,omitempty"`
	LoginCheckJs      string       `json:"loginCheckJs,omitempty"`
	ConcurrentRate    string       `json:"concurrentRate,omitempty"`
	CoverDecodeJs     string       `json:"coverDecodeJs,omitempty"`
}
