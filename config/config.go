package config

type Config struct {
	Password string `json:"password"`
	BotToken string `json:"botToken"`
	ChatID float64 `json:"chatID"`
	Computers []Computer `json:"computers"`
}

type Computer struct {
	Name string `json:"name"`
	Mac string `json:"mac"`
	IP *string `json:"broadcast"`
	Password *string `json:"password"`
}
