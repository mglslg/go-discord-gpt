package main

import (
	"github.com/mglslg/go-discord-gpt/cmd/g"
	"github.com/mglslg/go-discord-gpt/cmd/g/ds"
)

func setChannelStatus(us *ds.UserSession) {
	channelId := us.ChannelID

	//gpt-4 channel
	if channelId == "1127815740725153812" {
		us.Prompt = g.Assistant.Characters[1].Desc
		us.Model = "gpt-4-0125-preview"
	}
	//translate channel
	if channelId == "1095947683597914162" {
		us.OnConversation = false
		//us.OnAt = false
		us.Prompt = g.Assistant.Characters[2].Desc
	}
}
