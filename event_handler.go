package main

import (
	"fmt"
	"log"
	"strings"

	"github.com/nlopes/slack"
)

const (
	actionSelect = "select"
	actionStart  = "start"
	actionCancel = "cancel"
)

type eventHandler struct {
	client *slack.Client
	botUserID  string
}

func (h *eventHandler) listen() {
	rtm := h.client.NewRTM()

	log.Printf("[INFO] Establishing a RTM connection with Slack")
	go rtm.ManageConnection()

	for msg := range rtm.IncomingEvents {
		switch ev := msg.Data.(type) {
		case *slack.MessageEvent:
			if err := h.handleMessageEvent(ev); err != nil {
				log.Printf("[ERROR] Failed to handle message: %s", err)
			}
		}
	}
}

func (h *eventHandler) handleMessageEvent(ev *slack.MessageEvent) error {

	txt := ev.Msg.Text
	log.Printf("[INFO] Incoming message: %s", txt)

	if !strings.HasPrefix(txt, fmt.Sprintf("<@%s> ", h.botUserID)) {
		return nil
	}

	c := strings.Split(strings.TrimSpace(ev.Msg.Text), " ")[1:]
	if len(c) != 1 {
		return nil
	}

	switch c[0] {
	case "deploy":
		// TODO: Get a list of branches from GitHub
		attachment := slack.Attachment{
			Text:       "Select a branch to deploy",
			Color:      "#f9a41b",
			CallbackID: "deploy_branch",
			Actions: []slack.AttachmentAction{
				{
					Name: actionSelect,
					Type: "select",
					Options: []slack.AttachmentActionOption{
						{
							Text:  "master",
							Value: "master",
						},
						{
							Text:  "branch-1",
							Value: "branch-1",
						},
						{
							Text:  "branch-2",
							Value: "branch-2",
						},
					},
				},

				{
					Name:  actionCancel,
					Text:  "Cancel",
					Type:  "button",
					Style: "danger",
				},
			},
		}

		if _, _, err := h.client.PostMessage(ev.Channel, slack.MsgOptionAttachments(attachment)); err != nil {
			return fmt.Errorf("Something went wront. Error: %s", err)
		}

		return nil
	case "code_freeze":
		return fmt.Errorf("Sorry, `%s` is not supported yet", c)
	case "submit":
		return fmt.Errorf("Sorry, `%s` is not supported yet", c)
	case "release":
		return fmt.Errorf("Sorry, `%s` is not supported yet", c)
	default:
		return fmt.Errorf("Sorry, I don't understand %s", c)
	}
}
