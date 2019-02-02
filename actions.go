package main

import (
	"github.com/nlopes/slack"
)

type action interface {
	run() error
}

type deploySelectBranchAction struct {
	client *slack.Client
	channel string
}

func (a deploySelectBranchAction) run() error {
	branches := []string{"master", "deploy", "feature"}

	attachment := selectionAttachmentBuilder{
		title:      "Select a branch to deploy",
		texts:      branches,
		values:     branches,
		callbackID: deploySelectBranch,
	}

	if _, _, err := a.client.PostMessage(a.channel, slack.MsgOptionAttachments(attachment.build())); err != nil {
		return err
	}

	return nil
}

type deploySelectBuildKindAction struct {
	client *slack.Client
	channel string
}

func (a deploySelectBuildKindAction) run() error {
	buildKinds := []string{"release", "pr", "beta"}

	attachment := selectionAttachmentBuilder{
		title:      "Select a build kind",
		texts:      buildKinds,
		values:     buildKinds,
		callbackID: deploySelectBuildKind,
	}

	if _, _, err := a.client.PostMessage(a.channel, slack.MsgOptionAttachments(attachment.build())); err != nil {
		return err
	}

	return nil
}
