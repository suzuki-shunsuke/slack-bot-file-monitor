package main

import (
	"fmt"

	"github.com/nlopes/slack"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/viper"

	"github.com/suzuki-shunsuke/slack-bot-file-monitor/constants/logTypes"
)

var (
	intEnvs   []string                    = []string{}
	envs      []string                    = []string{"slack_app_bot_token"}
	msgParams slack.PostMessageParameters = slack.PostMessageParameters{
		Attachments: []slack.Attachment{{
			Color: "warning",
			Text:  "Please be careful not to upload files containing personal information.",
		}},
	}
)

func init() {
	log.SetFormatter(&log.JSONFormatter{})
	// log.SetLevel(log.DebugLevel)
	bindEnvs()
}

func main() {
	log.WithFields(log.Fields{"type": logTypes.Info}).Info("Start bot")
	if err := validateFlag(); err != nil {
		log.Fatal(err)
	}
	bot := slack.New(viper.GetString("slack_app_bot_token"))
	rtm := bot.NewRTM()
	go rtm.ManageConnection()

	for {
		select {
		case msg := <-rtm.IncomingEvents:
			switch ev := msg.Data.(type) {
			case *slack.ConnectedEvent:
				log.WithFields(log.Fields{
					"type": logTypes.Info, "name": ev.Info.User.Name,
					"id": ev.Info.User.ID,
				}).Info("Connect bot to slack")
			case *slack.FileSharedEvent:
				doFileSharedEvent(bot, ev)
			}
		}
	}
}

func doFileSharedEvent(bot *slack.Client, ev *slack.FileSharedEvent) {
	// https://api.slack.com/events/file_shared
	log.WithFields(
		log.Fields{"type": logTypes.Info}).Info("File shared event")
	// Get channels
	file, _, _, err := bot.GetFileInfo(ev.FileID, 100, 1)
	if err != nil {
		log.WithFields(log.Fields{
			"error": err, "file_id": ev.FileID, "file_name": ev.File.Name,
		}).Error("Failed to get shared file info")
		return
	}
	log.WithFields(log.Fields{
		"channels": file.Channels,
	}).Debug("shared file.Channels")
	if len(file.Channels) > 1 {
		return
	}
	for _, channelId := range file.Channels {
		_, _, err = bot.PostMessage(channelId, "", msgParams)
		if err != nil {
			log.WithFields(log.Fields{
				"error": err, "channel_id": channelId,
				"text":  msgParams.Attachments[0].Text,
				"color": "warning",
			}).Error("Failed to post Message to slack")
		}
	}
}

func bindEnvs() {
	for _, e := range intEnvs {
		viper.BindEnv(e)
	}
	for _, e := range envs {
		viper.BindEnv(e)
	}
}

func validateFlag() error {
	for _, e := range intEnvs {
		if viper.GetInt(e) == 0 {
			return fmt.Errorf("[FAIRURE] %s is required", e)
		}
	}
	for _, e := range envs[1:] {
		if viper.GetString(e) == "" {
			return fmt.Errorf("[FAIRURE] %s is required", e)
		}
	}
	return nil
}
