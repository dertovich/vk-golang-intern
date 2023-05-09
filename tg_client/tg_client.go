package tg_client

import (
	"fmt"
	"log"
	"os"
	"strconv"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	"github.com/joho/godotenv"
)

func CreateBot(config Config) (*tgbotapi.BotAPI, error) {
	bot, err := tgbotapi.NewBotAPI(config.BotToken)
	if err != nil {
		return nil, fmt.Errorf("failed to create bot: %w", err)
	}

	bot.Debug = config.DebugMode

	log.Printf("Authorized on account %s", bot.Self.UserName)

	return bot, nil
}

func LoadConfig() (Config, error) {
	err := godotenv.Load("config/.env")
	if err != nil {
		return Config{}, fmt.Errorf("failed to load env vars: %w", err)
	}

	token := os.Getenv("TELEGRAM_TOKEN")

	debugMode, err := strconv.ParseBool(os.Getenv("DEBUG"))
	if err != nil {
		return Config{}, fmt.Errorf("failed to parse DEBUG env var: %w", err)
	}

	return Config{
		BotToken:  token,
		DebugMode: debugMode,
	}, nil
}
