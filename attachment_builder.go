package main

import (
	"github.com/nlopes/slack"
)

type attachmentBuilder interface {
	build() slack.Attachment
}

type selectionAttachmentBuilder struct {
	title         string
	texts, values []string
	callbackID    string
}

func (b selectionAttachmentBuilder) build() slack.Attachment {
	options := make([]slack.AttachmentActionOption, len(b.texts))
	for i, text := range b.texts {
		options[i] = slack.AttachmentActionOption{Text: text, Value: b.values[i]}
	}

	return slack.Attachment{
		Text:       b.title,
		CallbackID: b.callbackID,
		Actions: []slack.AttachmentAction{
			{
				Name:    actionSelect,
				Type:    "select",
				Options: options,
			},

			{
				Name:  actionCancel,
				Text:  "Cancel",
				Type:  "button",
				Style: "danger",
			},
		},
	}
}
