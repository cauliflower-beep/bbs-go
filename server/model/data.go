package model

// 站点导航
type ActionLink struct {
	Title string `json:"title"`
	Url   string `json:"url"`
}

// 积分配置
type ScoreConfig struct {
	PostTopicScore   int `json:"postTopicScore"`   // 发帖获得积分
	PostCommentScore int `json:"postCommentScore"` // 跟帖获得积分
	CheckInScore     int `json:"checkInScore"`     // 签到积分
}

// 登录方式
type LoginMethod struct {
	Password bool `json:"password"` // 账户密码登录
	QQ       bool `json:"qq"`       // qq登录
	Github   bool `json:"github"`   // GitHub登录
	Osc      bool `json:"osc"`      //
}

// SysConfigResponse 配置返回结构体
type SysConfigResponse struct {
	SiteTitle                  string       `json:"siteTitle"`
	SiteDescription            string       `json:"siteDescription"`
	SiteKeywords               []string     `json:"siteKeywords"`
	SiteNavs                   []ActionLink `json:"siteNavs"`
	SiteNotification           string       `json:"siteNotification"`
	RecommendTags              []string     `json:"recommendTags"`
	UrlRedirect                bool         `json:"urlRedirect"`
	ScoreConfig                ScoreConfig  `json:"scoreConfig"`
	DefaultNodeId              int64        `json:"defaultNodeId"`
	ArticlePending             bool         `json:"articlePending"`
	TopicCaptcha               bool         `json:"topicCaptcha"`
	UserObserveSeconds         int          `json:"userObserveSeconds"`
	TokenExpireDays            int          `json:"tokenExpireDays"`
	LoginMethod                LoginMethod  `json:"loginMethod"`
	CreateTopicEmailVerified   bool         `json:"createTopicEmailVerified"`
	CreateArticleEmailVerified bool         `json:"createArticleEmailVerified"`
	CreateCommentEmailVerified bool         `json:"createCommentEmailVerified"`
	EnableHideContent          bool         `json:"enableHideContent"`
}
