package biquge

type Biquge struct {
	BookSourceComment string          `json:"bookSourceComment"`
	BookSourceGroup   string          `json:"bookSourceGroup"`
	BookSourceName    string          `json:"bookSourceName"`
	BookSourceType    int             `json:"bookSourceType"`
	BookSourceURL     string          `json:"bookSourceUrl"`
	BookURLPattern    string          `json:"bookUrlPattern"`
	CustomOrder       int             `json:"customOrder"`
	Enabled           bool            `json:"enabled"`
	EnabledCookieJar  bool            `json:"enabledCookieJar"`
	EnabledExplore    bool            `json:"enabledExplore"`
	ExploreURL        string          `json:"exploreUrl"`
	LastUpdateTime    string          `json:"lastUpdateTime"`
	RespondTime       int             `json:"respondTime"`
	RuleBookInfo      BookInfo        `json:"ruleBookInfo"`
	RuleContent       BookRuleContent `json:"ruleContent"`
	RuleExplore       struct {
	} `json:"ruleExplore"`
	RuleReview struct {
	} `json:"ruleReview"`
	RuleSearch BookRuleSearch `json:"ruleSearch"`
	RuleToc    BookRuleToc    `json:"ruleToc"`
	SearchURL  string         `json:"searchUrl"`
	Weight     int            `json:"weight"`
}
