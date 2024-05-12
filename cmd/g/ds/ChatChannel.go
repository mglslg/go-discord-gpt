package ds

type ChatChannel struct {
	ChannelId           string  `yaml:"channelId"`
	ChatModel           string  `yaml:"chatModel,omitempty"`
	ChatTemperature     float32 `yaml:"chatTemperature,omitempty"`
	Character           string  `yaml:"character,omitempty"`
	FreeChatLimit       int     `yaml:"freeChatLimit,omitempty"`
	ConversationSupport bool    `yaml:"conversationSupport,omitempty"`
}
