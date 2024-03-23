package g

import (
	"encoding/json"
	"fmt"
	"github.com/mglslg/go-discord-gpt/cmd/g/ds"
	"gopkg.in/yaml.v3"
	"io"
	"log"
	"os"
	"time"
)

var (
	Logger         *log.Logger
	AppSession     ds.AppSession
	SecToken       ds.Token
	Assistant      ds.Assistant
	UserSessionMap map[string]*ds.UserSession
)

// InitAppSession Reads the config file and unmarshal it into the AppSession variables
func InitAppSession(configPath string) {
	fmt.Println("Reading config file...")

	file, err := os.ReadFile(configPath)
	if err != nil {
		fmt.Println("Read config failed...", err)
		return
	}

	err = yaml.Unmarshal(file, &AppSession)

	if err != nil {
		fmt.Println("Resolve config file failed!", err)
		return
	}

	fmt.Println("Config file read successfully!")
}

// InitConfig readConfig reads the config file and unmarshals it into the config variable
func InitConfig(configPath string) {
	fmt.Println("Reading config file...")

	file, err := os.ReadFile(configPath)
	if err != nil {
		fmt.Println("Read config failed...", err)
		return
	}

	err = yaml.Unmarshal(file, &AppSession)

	if err != nil {
		fmt.Println("Resolve config file failed!", err)
		return
	}

	fmt.Println("Config file read successfully!")
}

// todo 这里都需要做改造
func InitAssistant(roleName string) {
	roleConfFile := fmt.Sprintf("role/%s.json", roleName)

	file, err := os.ReadFile(roleConfFile)
	if err != nil {
		fmt.Println("Read role config failed:", err)
	}

	Assistant.Name = roleName
	err = json.Unmarshal(file, &Assistant)

	if err != nil {
		fmt.Println("Resolve role config file failed:", err)
	}
	fmt.Println("This is " + Assistant.Name)
}

func InitLogger() *os.File {
	currentDate := time.Now().Format("2006-01-02")
	logPath := fmt.Sprintf("%s/logs", AppSession.Home)

	// Check if the logs directory exists, create it if it does not exist
	if _, err := os.Stat(logPath); os.IsNotExist(err) {
		if mkErr := os.MkdirAll(logPath, 0755); mkErr != nil {
			log.Fatalf("Unable to create log directory: %v", mkErr)
		}
	}

	logFileName := fmt.Sprintf("%s/%s-%s.log", logPath, currentDate, Assistant.Name)

	// Create a log file
	f, err := os.OpenFile(logFileName, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		log.Fatalf("Unable to open log file: %v", err)
	}

	// Create a logger
	Logger = log.New(io.MultiWriter(os.Stderr, f), "", log.LstdFlags)

	return f
}

func InitSecretConfig() {
	fmt.Println("Reading secret config file...")

	discordToken := os.Getenv(Assistant.Name + "_DISCORD_TOKEN")
	if discordToken == "" {
		log.Fatal(Assistant.Name + "_DISCORD_TOKEN is not set")
	}
	openaiToken := os.Getenv("OPENAI_API_KEY")
	if openaiToken == "" {
		log.Fatal("OPENAI_TOKEN is not set")
	}

	SecToken.Discord = discordToken
	SecToken.OpenAi = openaiToken

	Logger.Println("Secret Config file read successfully!Token:", SecToken.Discord)
}

func InitUserSession() {
	UserSessionMap = make(map[string]*ds.UserSession)
}

// GetUserSession Get the current user session, create it if it does not exist
func GetUserSession(authorId string, channelId string, authorName string) *ds.UserSession {
	key := getUserChannelId(authorId, channelId)
	_, exists := UserSessionMap[key]
	if !exists {
		UserSessionMap[key] = newUserSession(authorId, channelId, authorName)
	}
	return UserSessionMap[key]
}

func newUserSession(authorId string, channelId string, authorName string) *ds.UserSession {
	userChannelId := getUserChannelId(authorId, channelId)
	return &ds.UserSession{
		UserId:          authorId,
		UserName:        authorName,
		UserChannelID:   userChannelId,
		ChannelID:       channelId,
		ClearDelimiter:  Assistant.ClearDelimiter,
		Model:           "gpt-3.5-turbo", //todo 这里要改成直接从AppSession获取的，不写死
		Temperature:     0.7,
		Prompt:          Assistant.Characters[0].Desc,
		AllowChannelIds: Assistant.ChannelIds,
		OnConversation:  true, //Default to have context
		OnAt:            true, //Default to reply only when AT
	}
}

func getUserChannelId(authorId string, channelId string) string {
	return authorId + "_" + channelId
}
