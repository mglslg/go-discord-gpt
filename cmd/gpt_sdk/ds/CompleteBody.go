package ds

import (
	"github.com/mglslg/go-discord-gpt/cmd/g/ds"
)

type CompleteBody struct {
	Model       string           `json:"model"`
	Messages    []ds.ChatMessage `json:"messaged"`
	Temperature int              `json:"temperature"`
	Prompt      string           `json:"prompt"`
}

var DefaultCompleteBody = CompleteBody{
	Model:       "gpt-3.5-turbo",
	Temperature: 1,
}
