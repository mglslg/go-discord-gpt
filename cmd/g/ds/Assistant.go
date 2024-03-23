package ds

// todo 这个要改名字叫assistant
type Assistant struct {
	Name           string   `json:"name"`
	ApplicationId  string   `json:"applicationId"`
	ChannelIds     []string `json:"channelIds"`
	ClearDelimiter string   `json:"clearDelimiter"`
	Characters     []struct {
		Desc string `json:"desc"`
	} `json:"characters"`
}
