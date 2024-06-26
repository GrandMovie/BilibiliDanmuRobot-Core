package config

import (
	"github.com/zeromicro/go-zero/core/logx"
)

type Config struct {
	//rest.RestConf
	Log         logx.LogConf
	RoomId      int    `json:",default=4699397"`
	WsServerUrl string `json:",default=wss://broadcastlv.chat.bilibili.com:2245/sub"`
	//QrCodePath   string `json:"qr_code_path"`
	TalkRobotCmd       string            `json:",default=test"`
	FuzzyMatchCmd      bool              `json:",default=false"`
	RobotName          string            `json:",default=花花"`
	DanmuLen           int               `json:",default=20"`
	EntryEffect        bool              `json:",default=false"`
	EntryMsg           string            `json:",default=off"`
	WelcomeDanmu       []string          `json:",default='欢迎 {user} ~'"`
	ThanksGift         bool              `json:",default=false"`
	ThanksGiftTimeout  int               `json:",default=3"`
	ThanksMinCost      int               `json:",default=0"`
	CustomizeBullet    bool              `json:",default=false"`
	InteractWord       bool              `json:",default=false"`
	InteractWordByTime bool              `json:",default=false"`
	WelcomeSwitch      bool              `json:",default=false"`
	WelcomeString      map[string]string `json:",optional"`
	RobotMode          string            `json:",default=QingYunKe,options=QingYunKe|ChatGPT"`
	ChatGPT            struct {
		APIUrl   string `json:",default=https://api.openai.com/v1"`
		APIToken string `json:",optional"`
		Prompt   string `json:",default=你是一个非常幽默的机器人助理，可以使用emoji表情符号，可以使用颜文字"`
		Limit    bool   `json:",default=true"`
	}
	CronDanmu bool `json:",default=false"`
	// CronSupportSec bool `json:",default=false"`
	CronDanmuList []struct {
		Cron   string   `json:",optional"`
		Random bool     `json:",default=false"`
		Danmu  []string `json:",optional"`
	} `json:",optional"`
	WelcomeDanmuByTime []struct {
		Enabled bool     `json:",optional"`
		Key     string   `json:",optional"`
		Random  bool     `json:",default=false"`
		Danmu   []string `json:",optional"`
	} `json:",optional"`
	FocusDanmu           []string `json:",optional"`
	PKNotice             bool     `json:",default=true"`
	DrawLotsList         []string `json:",optional"`
	WelcomeBlacklistWide []string `json:",optional"`
	WelcomeBlacklist     []string `json:",optional"`
	DrawByLot            bool     `json:",default=true"`
	//TraditionalToSimplifiedConversion   bool     `json:",default=false"`
	//WelcomeTimeLimiter int    `json:",default=1"`
	SignInEnable     bool              `json:",default=true"`
	DBPath           string            `json:",default=./db"`
	DBName           string            `json:",default=sqliteDataBase.db"`
	ShowBlockMsg     bool              `json:",default=true"`  // 禁言提醒开关
	KeywordReply     bool              `json:",default=false"` //关键词回复开关
	KeywordReplyList map[string]string `json:",optional"`      // 关键词回复列表
}
