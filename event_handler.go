package main

import (
	"fmt"
	"log"
	"strings"

	"github.com/nlopes/slack"
)

type eventHandler struct {
	client    *slack.Client
	botUserID string
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

	var action action
	switch c[0] {
	case commandDeploy:
		action = deploySelectBranchAction{h.client, ev.Channel}
	case commandCodeFreeze:
		return fmt.Errorf("Sorry, `%s` is not supported yet", c)
	case commandSubmit:
		return fmt.Errorf("Sorry, `%s` is not supported yet", c)
	case commandRelease:
		return fmt.Errorf("Sorry, `%s` is not supported yet", c)
	default:
		return fmt.Errorf("Sorry, I don't understand %s", c)
	}

	action.run()
	return nil
}
