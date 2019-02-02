package main

import (
	"log"
	"net/http"
	"os"

	"github.com/joho/godotenv"
	"github.com/nlopes/slack"
)

func main() {
	os.Exit(_main(os.Args[1:]))
}

func _main(args []string) int {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("[ERROR] Failed to load .env file")
	}

	client := slack.New(os.Getenv("BOT_USER_OAUTH_ACCESS_TOKEN"))
	eventHandler := &eventHandler{
		client: client,
		botUserID: os.Getenv("BOT_USER_ID"),
	}
	go eventHandler.listen()

	//http.Handle("/", urlVerification{})

	http.Handle("/interaction", interactionHandler{
		client: client,
		verificationToken: os.Getenv("VERIFICATION_TOKEN"),
	})

	log.Printf("[INFO] Listening port :3000")
	if err := http.ListenAndServe(":3000", nil); err != nil {
		log.Printf("[ERROR] %s", err)
		return 1
	}

	return 0
}
