package openaisdk

import (
	"bytes"
	"encoding/json"
	"github.com/mglslg/go-discord-gpt/cmd/g"
	"github.com/mglslg/go-discord-gpt/cmd/g/ds"
	ds2 "github.com/mglslg/go-discord-gpt/cmd/openaisdk/ds"
	"io/ioutil"
	"net/http"
)

func Chat(msg []ds.ChatMessage) (string, error) {
	api := "https://api.openai.com/v1/chat/completions"
	payload := map[string]interface{}{
		"model":       g.AppContext.Model,
		"messages":    msg,
		"temperature": g.AppContext.Temperature,
	}

	body, err := json.Marshal(payload)

	req, err := http.NewRequest("POST", api, bytes.NewBuffer(body))
	if err != nil {
		g.Logger.Println("Error creating request:", err)
		return "[Error creating request:" + err.Error() + "]", err
	}

	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", "Bearer "+g.SecToken.OpenAi)

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		g.Logger.Println("Error sending request", err)
		return "[Error sending request:" + err.Error() + "]", err
	}
	defer resp.Body.Close()

	body, err = ioutil.ReadAll(resp.Body)
	if err != nil {
		g.Logger.Println("Error reading response", err)
		return "[Error reading response:" + err.Error() + "]", err
	}

	chatGptResponse := ds2.ChatGPTResponse{}
	err = json.Unmarshal(body, &chatGptResponse)
	if err != nil {
		g.Logger.Println("[Error unmarshalling response]", err)
		return "[Error unmarshalling response:" + err.Error() + "]", err
	}

	if len(chatGptResponse.Choices) == 0 {
		return "[Failed to get gpt response]", nil
	}
	g.Logger.Println(">>>>>gpt model:", g.AppContext.Model)
	g.Logger.Println(">>>>>gpt response:", chatGptResponse.Choices[0].Message.Content)
	g.Logger.Println(">>>>>finish reason:", chatGptResponse.Choices[0].FinishReason)
	g.Logger.Println(">>>>>total token:", chatGptResponse.Usage.TotalTokens)

	return chatGptResponse.Choices[0].Message.Content, nil
}
