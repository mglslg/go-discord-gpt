package ds

type ChatChannel struct {
	ChannelId           string  `yaml:"channelId"`
	Model               string  `yaml:"model,omitempty"`
	Temperature         float32 `yaml:"temperature,omitempty"`
	Prompt              string  `yaml:"prompt,omitempty"`
	FreeChatLimit       int     `yaml:"freeChatLimit,omitempty"`
	ConversationSupport bool    `yaml:"conversationSupport,omitempty"`
}
