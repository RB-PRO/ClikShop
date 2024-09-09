package config

type Config struct {
	Telegram struct {
		Name   string `json:"name"`
		ChatID int64  `json:"chat_id"`
		Token  string `json:"token"`
	} `json:"telegram"`
	Translator struct {
		FolderId   string `json:"folder_id"`
		OauthToken string `json:"oauth_token"`
	} `json:"translator"`
	Updater struct {
		PingLinks struct {
			MassimoDutti []string `json:"massimo_dutti"`
			HM           []string `json:"hm"`
			Zara         []string `json:"zara"`
			SneakSup     []string `json:"sneaksup"`
			Trendyol     []string `json:"trendyol"`
		} `json:"ping_links"`
	} `json:"updater"`
	Proxy []string `json:"proxy"`
}
