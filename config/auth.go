package config

type Auth struct {
	Time float64 `json:"time" binding:"required"`
	Token string `json:"token" binding:"required"`
}