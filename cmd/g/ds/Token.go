package ds

// Token is the token for the discord bot and chatgpt
type Token struct {
	Discord string `yaml:"discord"`
	OpenAi  string `yaml:"chatgpt"`
}
