package events

import (
	"github.com/rxdn/gdl/objects/guild/emoji"
	"github.com/rxdn/gdl/objects/member"
)

type MessageReactionAdd struct {
	UserId    uint64         `json:"user_id,string"`
	ChannelId uint64         `json:"channel_id,string"`
	MessageId uint64         `json:"message_id,string"`
	GuildId   uint64         `json:"guild_id,string"`
	Member    *member.Member `json:"member,string"`
	Emoji     emoji.Emoji    `json:"emoji,string"` // Partial emoji object; https://discordapp.com/developers/docs/resources/emoji#emoji-object-gateway-reaction-standard-emoji-example
}
