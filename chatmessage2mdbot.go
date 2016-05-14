package main

import (
	"fmt"
	"github.com/nlopes/slack"
	"os"
	"strconv"
	"strings"
	"time"
)

func unescapeCharacters(escaped_message string) string {
	r := strings.NewReplacer("&amp;", "&",
		"&lt;", "<",
		"&gt;", ">")
	message := r.Replace(escaped_message)
	return message
}

func convertIdToName(include_id_message string) string {
	r := strings.NewReplacer("U0BRJDKKN", "t_fujiwara")
	message := r.Replace(include_id_message)
	return message
}

func main() {
	api := slack.New(os.Getenv("SLACK_API_TOKEN"))
	params := slack.HistoryParameters{Count: 100}
	history, _ := api.GetChannelHistory("C066ZLMJS", params)
	for _, msg := range history.Messages {
		text := unescapeCharacters(msg.Text)
		text = convertIdToName(text)
		id := convertIdToName(msg.User)
		time_stamp := strings.Split(msg.Timestamp, ".")
		utime, _ := strconv.ParseInt(time_stamp[0], 10, 64)
		fmt.Printf("ID:%v, Message:%v, Timestamp:%v\n", id, text, time.Unix(utime, 0))
	}
}
