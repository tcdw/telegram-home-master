package functions

type Config struct {
	Token string         `json:"token"`
	ChatID float64       `json:"chatID"`
	Computers []Computer `json:"computers"`
}

type Computer struct {
	Name string `json:"name"`
	Mac string `json:"mac"`
	IP *string `json:"broadcast"`
	Password *string `json:"password"`
}
