package mallagent

import "github.com/cloudwego/eino/schema"

type UserMessage struct {
	ChatId  string            `json:"chat_id"`
	UserId  string            `json:"user_id"`
	Query   string            `json:"query"`
	History []*schema.Message `json:"history"`
}
