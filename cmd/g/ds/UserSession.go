package ds

// UserSession Each user holds a session in each channel to save its state
type UserSession struct {
	UserChannelID   string //UserSession's unique key
	UserId          string
	ChannelID       string
	ParentChannelID string
	UserName        string
	ClearDelimiter  string
	Prompt          string
	OnConversation  bool //Whether to enable context
	OnAt            bool //Whether to reply to the robot only when AT
	ChatCount       int  //Number of messages sent by the user
}
