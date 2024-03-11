package ds

// UserSession Each user holds a session in each channel to save its state
type UserSession struct {
	UserChannelID   string //UserSession's unique key
	UserId          string
	ChannelID       string
	UserName        string
	ClearDelimiter  string
	Model           string
	Temperature     float64
	Prompt          string
	AllowChannelIds []string //Channel permission, effective for non-VIP users
	OnConversation  bool     //Whether to enable context
	OnAt            bool     //Whether to reply to the robot only when AT
}
