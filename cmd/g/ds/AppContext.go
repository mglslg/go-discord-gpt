package ds

type AppContext struct {
	LogFilePath         string        `yaml:"logFilePath"`
	MaxFetchRecord      int           `yaml:"maxFetchRecord"`
	MaxUserRecord       int           `yaml:"maxUserRecord"`
	OpenaiApiKey        string        `yaml:"openaiApiKey"`
	ApplicationId       string        `yaml:"applicationId"`
	BotName             string        `yaml:"botName"`
	BotToken            string        `yaml:"botToken"`
	ClearCmd            string        `yaml:"clearCmd"`
	ClearCmdDesc        string        `yaml:"clearCmdDesc"`
	ClearDelimiter      string        `yaml:"clearDelimiter"`
	Model               string        `yaml:"model"`
	Temperature         float32       `yaml:"temperature"`
	Prompt              string        `yaml:"prompt"`
	FreeChatLimit       int           `yaml:"creeChatLimit"`
	ConversationSupport bool          `yaml:"conversationSupport"`
	ChannelConfig       []ChatChannel `yaml:"channelConfig"`
	BotId               string
	ConfigFilePath      string
}
