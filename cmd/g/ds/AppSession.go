package ds

// todo  这个要做成config匹配的yaml,AppSession和Assistant是不是可以合为一个东西呢？？？
type AppSession struct {
	DiscordBotID    string `yaml:"discordBodId"`
	Home            string `yaml:"home"`
	MaxFetchRecord  int    `yaml:"maxFetchRecord"`
	MaxUserRecord   int    `yaml:"maxUserRecord"`
	ChatModel       string `yaml:"chatModel"`
	ChatTemperature string `yaml:"chatTemperature"`
	Assistant       Assistant
}
