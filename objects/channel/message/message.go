package message

import (
	"github.com/rxdn/gdl/objects/channel"
	"github.com/rxdn/gdl/objects/channel/embed"
	"github.com/rxdn/gdl/objects/member"
	"github.com/rxdn/gdl/objects/user"
	"github.com/rxdn/gdl/utils"
	"time"
)

type Message struct {
	Id               uint64                  `json:"id,string"`
	ChannelId        uint64                  `json:"channel_id,string"`
	GuildId          uint64                  `json:"guild_id,string"`
	Author           *user.User              `json:"author"`
	Member           *member.Member          `json:"member"`
	Content          string                  `json:"content"`
	Timestamp        time.Time               `json:"timestamp"`
	EditedTimestamp  *time.Time              `json:"edited_timestamp"`
	Tts              bool                    `json:"tts"`
	MentionEveryone  bool                    `json:"mention_everyone"`
	Mentions         []*MessageMentionedUser `json:"mentions"` // The user objects in the mentions array will only have the partial member field present in MESSAGE_CREATE and MESSAGE_UPDATE events from text-based guild channels
	MentionRoles     utils.Uint64StringSlice `json:"mention_roles,string"`
	MentionChannels  []*ChannelMention       `json:"mention_channels"`
	Attachments      []channel.Attachment    `json:"attachments"`
	Embeds           []embed.Embed           `json:"embed"`
	Reactions        []Reaction              `json:"reactions"`
	Nonce            interface{}             `json:"nonce"`
	Pinned           bool                    `json:"pinned"`
	WebhookId        uint64                  `json:"webhook_id,string"` // if the message is generated by a webhook, this is the webhook's id
	Type             MessageType             `json:"message_type"`
	Activity         MessageActivity         `json:"activity"`
	Application      MessageApplication      `json:"application"`
	MessageReference MessageReference        `json:"message_reference"` // reference data sent with crossposted messages
	Flags            int                     `json:"flags"`
}
