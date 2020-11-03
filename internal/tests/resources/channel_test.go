// +build !nodiscordtests

package apislurp_test

import (
	"github.com/diamondburned/apislurp"
	"github.com/diamondburned/apislurp/internal/tests/testutil"
	"path/filepath"
	"testing"
)

func TestChannel(t *testing.T) {
	path := filepath.Join("resources", "Channel.md")
	f := testutil.ReadFile(t, path)
	d := apislurp.Parse(f)
	testutil.Equal(t, d, channelData, "Channel")
}

var channelData = apislurp.Document{
	Structures: []apislurp.Structure{
		apislurp.Structure{
			Name:    "Channel",
			RawName: "Channel Structure",
			Type:    "",
			Fields: []apislurp.Field{
				apislurp.Field{
					Name:    "id",
					RawName: "id",
					Type: apislurp.Type{
						Names: []apislurp.TypeName{
							"snowflake",
						},
						Array:     false,
						ArraySize: 0,
						Nullable:  false,
						RawText:   "snowflake",
					},
					Optional:    false,
					Description: "the id of this channel",
				},
				apislurp.Field{
					Name:    "type",
					RawName: "type",
					Type: apislurp.Type{
						Names: []apislurp.TypeName{
							"integer",
						},
						Array:     false,
						ArraySize: 0,
						Nullable:  false,
						RawText:   "integer",
					},
					Optional:    false,
					Description: "the [type of channel](#DOCS_RESOURCES_CHANNEL/channel-object-channel-types)",
				},
				apislurp.Field{
					Name:    "guild_id",
					RawName: "guild_id?",
					Type: apislurp.Type{
						Names: []apislurp.TypeName{
							"snowflake",
						},
						Array:     false,
						ArraySize: 0,
						Nullable:  false,
						RawText:   "snowflake",
					},
					Optional:    true,
					Description: "the id of the guild",
				},
				apislurp.Field{
					Name:    "position",
					RawName: "position?",
					Type: apislurp.Type{
						Names: []apislurp.TypeName{
							"integer",
						},
						Array:     false,
						ArraySize: 0,
						Nullable:  false,
						RawText:   "integer",
					},
					Optional:    true,
					Description: "sorting position of the channel",
				},
				apislurp.Field{
					Name:    "permission_overwrites",
					RawName: "permission_overwrites?",
					Type: apislurp.Type{
						Names: []apislurp.TypeName{
							"overwrite",
						},
						Array:     false,
						ArraySize: 0,
						Nullable:  false,
						RawText:   "array of [overwrite](#DOCS_RESOURCES_CHANNEL/overwrite-object) objects",
					},
					Optional:    true,
					Description: "explicit permission overwrites for members and roles",
				},
				apislurp.Field{
					Name:    "name",
					RawName: "name?",
					Type: apislurp.Type{
						Names: []apislurp.TypeName{
							"string",
						},
						Array:     false,
						ArraySize: 0,
						Nullable:  false,
						RawText:   "string",
					},
					Optional:    true,
					Description: "the name of the channel (2-100 characters)",
				},
				apislurp.Field{
					Name:    "topic",
					RawName: "topic?",
					Type: apislurp.Type{
						Names: []apislurp.TypeName{
							"string",
						},
						Array:     false,
						ArraySize: 0,
						Nullable:  true,
						RawText:   "?string",
					},
					Optional:    true,
					Description: "the channel topic (0-1024 characters)",
				},
				apislurp.Field{
					Name:    "nsfw",
					RawName: "nsfw?",
					Type: apislurp.Type{
						Names: []apislurp.TypeName{
							"boolean",
						},
						Array:     false,
						ArraySize: 0,
						Nullable:  false,
						RawText:   "boolean",
					},
					Optional:    true,
					Description: "whether the channel is nsfw",
				},
				apislurp.Field{
					Name:    "last_message_id",
					RawName: "last_message_id?",
					Type: apislurp.Type{
						Names: []apislurp.TypeName{
							"snowflake",
						},
						Array:     false,
						ArraySize: 0,
						Nullable:  true,
						RawText:   "?snowflake",
					},
					Optional:    true,
					Description: "the id of the last message sent in this channel (may not point to an existing or valid message)",
				},
				apislurp.Field{
					Name:    "bitrate",
					RawName: "bitrate?",
					Type: apislurp.Type{
						Names: []apislurp.TypeName{
							"integer",
						},
						Array:     false,
						ArraySize: 0,
						Nullable:  false,
						RawText:   "integer",
					},
					Optional:    true,
					Description: "the bitrate (in bits) of the voice channel",
				},
				apislurp.Field{
					Name:    "user_limit",
					RawName: "user_limit?",
					Type: apislurp.Type{
						Names: []apislurp.TypeName{
							"integer",
						},
						Array:     false,
						ArraySize: 0,
						Nullable:  false,
						RawText:   "integer",
					},
					Optional:    true,
					Description: "the user limit of the voice channel",
				},
				apislurp.Field{
					Name:    "rate_limit_per_user",
					RawName: "rate_limit_per_user?",
					Type: apislurp.Type{
						Names: []apislurp.TypeName{
							"integer",
						},
						Array:     false,
						ArraySize: 0,
						Nullable:  false,
						RawText:   "integer",
					},
					Optional:    true,
					Description: "amount of seconds a user has to wait before sending another message (0-21600); bots, as well as users with the permission `manage_messages` or `manage_channel`, are unaffected",
				},
				apislurp.Field{
					Name:    "recipients",
					RawName: "recipients?",
					Type: apislurp.Type{
						Names: []apislurp.TypeName{
							"user",
						},
						Array:     false,
						ArraySize: 0,
						Nullable:  false,
						RawText:   "array of [user](#DOCS_RESOURCES_USER/user-object) objects",
					},
					Optional:    true,
					Description: "the recipients of the DM",
				},
				apislurp.Field{
					Name:    "icon",
					RawName: "icon?",
					Type: apislurp.Type{
						Names: []apislurp.TypeName{
							"string",
						},
						Array:     false,
						ArraySize: 0,
						Nullable:  true,
						RawText:   "?string",
					},
					Optional:    true,
					Description: "icon hash",
				},
				apislurp.Field{
					Name:    "owner_id",
					RawName: "owner_id?",
					Type: apislurp.Type{
						Names: []apislurp.TypeName{
							"snowflake",
						},
						Array:     false,
						ArraySize: 0,
						Nullable:  false,
						RawText:   "snowflake",
					},
					Optional:    true,
					Description: "id of the DM creator",
				},
				apislurp.Field{
					Name:    "application_id",
					RawName: "application_id?",
					Type: apislurp.Type{
						Names: []apislurp.TypeName{
							"snowflake",
						},
						Array:     false,
						ArraySize: 0,
						Nullable:  false,
						RawText:   "snowflake",
					},
					Optional:    true,
					Description: "application id of the group DM creator if it is bot-created",
				},
				apislurp.Field{
					Name:    "parent_id",
					RawName: "parent_id?",
					Type: apislurp.Type{
						Names: []apislurp.TypeName{
							"snowflake",
						},
						Array:     false,
						ArraySize: 0,
						Nullable:  true,
						RawText:   "?snowflake",
					},
					Optional:    true,
					Description: "id of the parent category for a channel (each parent category can contain up to 50 channels)",
				},
				apislurp.Field{
					Name:    "last_pin_timestamp",
					RawName: "last_pin_timestamp?",
					Type: apislurp.Type{
						Names: []apislurp.TypeName{
							"ISO8601 timestamp",
						},
						Array:     false,
						ArraySize: 0,
						Nullable:  true,
						RawText:   "?ISO8601 timestamp",
					},
					Optional:    true,
					Description: "when the last pinned message was pinned. This may be `null` in events such as `GUILD_CREATE` when a message is not pinned.",
				},
			},
		},
		apislurp.Structure{
			Name:    "Message",
			RawName: "Message Structure",
			Type:    "",
			Fields: []apislurp.Field{
				apislurp.Field{
					Name:    "id",
					RawName: "id",
					Type: apislurp.Type{
						Names: []apislurp.TypeName{
							"snowflake",
						},
						Array:     false,
						ArraySize: 0,
						Nullable:  false,
						RawText:   "snowflake",
					},
					Optional:    false,
					Description: "id of the message",
				},
				apislurp.Field{
					Name:    "channel_id",
					RawName: "channel_id",
					Type: apislurp.Type{
						Names: []apislurp.TypeName{
							"snowflake",
						},
						Array:     false,
						ArraySize: 0,
						Nullable:  false,
						RawText:   "snowflake",
					},
					Optional:    false,
					Description: "id of the channel the message was sent in",
				},
				apislurp.Field{
					Name:    "guild_id",
					RawName: "guild_id?",
					Type: apislurp.Type{
						Names: []apislurp.TypeName{
							"snowflake",
						},
						Array:     false,
						ArraySize: 0,
						Nullable:  false,
						RawText:   "snowflake",
					},
					Optional:    true,
					Description: "id of the guild the message was sent in",
				},
				apislurp.Field{
					Name:    "author\\*",
					RawName: "author\\*",
					Type: apislurp.Type{
						Names: []apislurp.TypeName{
							"user",
						},
						Array:     false,
						ArraySize: 0,
						Nullable:  false,
						RawText:   "[user](#DOCS_RESOURCES_USER/user-object) object",
					},
					Optional:    false,
					Description: "the author of this message (not guaranteed to be a valid user, see below)",
				},
				apislurp.Field{
					Name:    "member?\\*\\*",
					RawName: "member?\\*\\*",
					Type: apislurp.Type{
						Names: []apislurp.TypeName{
							"guild member",
						},
						Array:     false,
						ArraySize: 0,
						Nullable:  false,
						RawText:   "partial [guild member](#DOCS_RESOURCES_GUILD/guild-member-object) object",
					},
					Optional:    false,
					Description: "member properties for this message's author",
				},
				apislurp.Field{
					Name:    "content",
					RawName: "content",
					Type: apislurp.Type{
						Names: []apislurp.TypeName{
							"string",
						},
						Array:     false,
						ArraySize: 0,
						Nullable:  false,
						RawText:   "string",
					},
					Optional:    false,
					Description: "contents of the message",
				},
				apislurp.Field{
					Name:    "timestamp",
					RawName: "timestamp",
					Type: apislurp.Type{
						Names: []apislurp.TypeName{
							"ISO8601 timestamp",
						},
						Array:     false,
						ArraySize: 0,
						Nullable:  false,
						RawText:   "ISO8601 timestamp",
					},
					Optional:    false,
					Description: "when this message was sent",
				},
				apislurp.Field{
					Name:    "edited_timestamp",
					RawName: "edited_timestamp",
					Type: apislurp.Type{
						Names: []apislurp.TypeName{
							"ISO8601 timestamp",
						},
						Array:     false,
						ArraySize: 0,
						Nullable:  true,
						RawText:   "?ISO8601 timestamp",
					},
					Optional:    false,
					Description: "when this message was edited (or null if never)",
				},
				apislurp.Field{
					Name:    "tts",
					RawName: "tts",
					Type: apislurp.Type{
						Names: []apislurp.TypeName{
							"boolean",
						},
						Array:     false,
						ArraySize: 0,
						Nullable:  false,
						RawText:   "boolean",
					},
					Optional:    false,
					Description: "whether this was a TTS message",
				},
				apislurp.Field{
					Name:    "mention_everyone",
					RawName: "mention_everyone",
					Type: apislurp.Type{
						Names: []apislurp.TypeName{
							"boolean",
						},
						Array:     false,
						ArraySize: 0,
						Nullable:  false,
						RawText:   "boolean",
					},
					Optional:    false,
					Description: "whether this message mentions everyone",
				},
				apislurp.Field{
					Name:    "mentions\\*\\*\\*",
					RawName: "mentions\\*\\*\\*",
					Type: apislurp.Type{
						Names: []apislurp.TypeName{
							"user",
						},
						Array:     false,
						ArraySize: 0,
						Nullable:  false,
						RawText:   "array of [user](#DOCS_RESOURCES_USER/user-object) objects, with an additional partial [member](#DOCS_RESOURCES_GUILD/guild-member-object) field",
					},
					Optional:    false,
					Description: "users specifically mentioned in the message",
				},
				apislurp.Field{
					Name:    "mention_roles",
					RawName: "mention_roles",
					Type: apislurp.Type{
						Names: []apislurp.TypeName{
							"role",
						},
						Array:     false,
						ArraySize: 0,
						Nullable:  false,
						RawText:   "array of [role](#DOCS_TOPICS_PERMISSIONS/role-object) object ids",
					},
					Optional:    false,
					Description: "roles specifically mentioned in this message",
				},
				apislurp.Field{
					Name:    "mention_channels?\\*\\*\\*\\*",
					RawName: "mention_channels?\\*\\*\\*\\*",
					Type: apislurp.Type{
						Names: []apislurp.TypeName{
							"channel mention",
						},
						Array:     false,
						ArraySize: 0,
						Nullable:  false,
						RawText:   "array of [channel mention](#DOCS_RESOURCES_CHANNEL/channel-mention-object) objects",
					},
					Optional:    false,
					Description: "channels specifically mentioned in this message",
				},
				apislurp.Field{
					Name:    "attachments",
					RawName: "attachments",
					Type: apislurp.Type{
						Names: []apislurp.TypeName{
							"attachment",
						},
						Array:     false,
						ArraySize: 0,
						Nullable:  false,
						RawText:   "array of [attachment](#DOCS_RESOURCES_CHANNEL/attachment-object) objects",
					},
					Optional:    false,
					Description: "any attached files",
				},
				apislurp.Field{
					Name:    "embeds",
					RawName: "embeds",
					Type: apislurp.Type{
						Names: []apislurp.TypeName{
							"embed",
						},
						Array:     false,
						ArraySize: 0,
						Nullable:  false,
						RawText:   "array of [embed](#DOCS_RESOURCES_CHANNEL/embed-object) objects",
					},
					Optional:    false,
					Description: "any embedded content",
				},
				apislurp.Field{
					Name:    "reactions",
					RawName: "reactions?",
					Type: apislurp.Type{
						Names: []apislurp.TypeName{
							"reaction",
						},
						Array:     false,
						ArraySize: 0,
						Nullable:  false,
						RawText:   "array of [reaction](#DOCS_RESOURCES_CHANNEL/reaction-object) objects",
					},
					Optional:    true,
					Description: "reactions to the message",
				},
				apislurp.Field{
					Name:    "nonce",
					RawName: "nonce?",
					Type: apislurp.Type{
						Names: []apislurp.TypeName{
							"integer",
							"string",
						},
						Array:     false,
						ArraySize: 0,
						Nullable:  false,
						RawText:   "integer or string",
					},
					Optional:    true,
					Description: "used for validating a message was sent",
				},
				apislurp.Field{
					Name:    "pinned",
					RawName: "pinned",
					Type: apislurp.Type{
						Names: []apislurp.TypeName{
							"boolean",
						},
						Array:     false,
						ArraySize: 0,
						Nullable:  false,
						RawText:   "boolean",
					},
					Optional:    false,
					Description: "whether this message is pinned",
				},
				apislurp.Field{
					Name:    "webhook_id",
					RawName: "webhook_id?",
					Type: apislurp.Type{
						Names: []apislurp.TypeName{
							"snowflake",
						},
						Array:     false,
						ArraySize: 0,
						Nullable:  false,
						RawText:   "snowflake",
					},
					Optional:    true,
					Description: "if the message is generated by a webhook, this is the webhook's id",
				},
				apislurp.Field{
					Name:    "type",
					RawName: "type",
					Type: apislurp.Type{
						Names: []apislurp.TypeName{
							"integer",
						},
						Array:     false,
						ArraySize: 0,
						Nullable:  false,
						RawText:   "integer",
					},
					Optional:    false,
					Description: "[type of message](#DOCS_RESOURCES_CHANNEL/message-object-message-types)",
				},
				apislurp.Field{
					Name:    "activity",
					RawName: "activity?",
					Type: apislurp.Type{
						Names: []apislurp.TypeName{
							"message activity",
						},
						Array:     false,
						ArraySize: 0,
						Nullable:  false,
						RawText:   "[message activity](#DOCS_RESOURCES_CHANNEL/message-object-message-activity-structure) object",
					},
					Optional:    true,
					Description: "sent with Rich Presence-related chat embeds",
				},
				apislurp.Field{
					Name:    "application",
					RawName: "application?",
					Type: apislurp.Type{
						Names: []apislurp.TypeName{
							"message application",
						},
						Array:     false,
						ArraySize: 0,
						Nullable:  false,
						RawText:   "[message application](#DOCS_RESOURCES_CHANNEL/message-object-message-application-structure) object",
					},
					Optional:    true,
					Description: "sent with Rich Presence-related chat embeds",
				},
				apislurp.Field{
					Name:    "message_reference",
					RawName: "message_reference?",
					Type: apislurp.Type{
						Names: []apislurp.TypeName{
							"message reference",
						},
						Array:     false,
						ArraySize: 0,
						Nullable:  false,
						RawText:   "[message reference](#DOCS_RESOURCES_CHANNEL/message-object-message-reference-structure) object",
					},
					Optional:    true,
					Description: "reference data sent with crossposted messages",
				},
				apislurp.Field{
					Name:    "flags",
					RawName: "flags?",
					Type: apislurp.Type{
						Names: []apislurp.TypeName{
							"integer",
						},
						Array:     false,
						ArraySize: 0,
						Nullable:  false,
						RawText:   "integer",
					},
					Optional:    true,
					Description: "[message flags](#DOCS_RESOURCES_CHANNEL/message-object-message-flags) combined as a [bitfield](https://en.wikipedia.org/wiki/Bit_field)",
				},
			},
		},
		apislurp.Structure{
			Name:    "Message Activity",
			RawName: "Message Activity Structure",
			Type:    "",
			Fields: []apislurp.Field{
				apislurp.Field{
					Name:    "type",
					RawName: "type",
					Type: apislurp.Type{
						Names: []apislurp.TypeName{
							"integer",
						},
						Array:     false,
						ArraySize: 0,
						Nullable:  false,
						RawText:   "integer",
					},
					Optional:    false,
					Description: "[type of message activity](#DOCS_RESOURCES_CHANNEL/message-object-message-activity-types)",
				},
				apislurp.Field{
					Name:    "party_id",
					RawName: "party_id?",
					Type: apislurp.Type{
						Names: []apislurp.TypeName{
							"string",
						},
						Array:     false,
						ArraySize: 0,
						Nullable:  false,
						RawText:   "string",
					},
					Optional:    true,
					Description: "party_id from a [Rich Presence event](#DOCS_RICH_PRESENCE_HOW_TO/updating-presence-update-presence-payload-fields)",
				},
			},
		},
		apislurp.Structure{
			Name:    "Message Application",
			RawName: "Message Application Structure",
			Type:    "",
			Fields: []apislurp.Field{
				apislurp.Field{
					Name:    "id",
					RawName: "id",
					Type: apislurp.Type{
						Names: []apislurp.TypeName{
							"snowflake",
						},
						Array:     false,
						ArraySize: 0,
						Nullable:  false,
						RawText:   "snowflake",
					},
					Optional:    false,
					Description: "id of the application",
				},
				apislurp.Field{
					Name:    "cover_image",
					RawName: "cover_image?",
					Type: apislurp.Type{
						Names: []apislurp.TypeName{
							"string",
						},
						Array:     false,
						ArraySize: 0,
						Nullable:  false,
						RawText:   "string",
					},
					Optional:    true,
					Description: "id of the embed's image asset",
				},
				apislurp.Field{
					Name:    "description",
					RawName: "description",
					Type: apislurp.Type{
						Names: []apislurp.TypeName{
							"string",
						},
						Array:     false,
						ArraySize: 0,
						Nullable:  false,
						RawText:   "string",
					},
					Optional:    false,
					Description: "application's description",
				},
				apislurp.Field{
					Name:    "icon",
					RawName: "icon",
					Type: apislurp.Type{
						Names: []apislurp.TypeName{
							"string",
						},
						Array:     false,
						ArraySize: 0,
						Nullable:  true,
						RawText:   "?string",
					},
					Optional:    false,
					Description: "id of the application's icon",
				},
				apislurp.Field{
					Name:    "name",
					RawName: "name",
					Type: apislurp.Type{
						Names: []apislurp.TypeName{
							"string",
						},
						Array:     false,
						ArraySize: 0,
						Nullable:  false,
						RawText:   "string",
					},
					Optional:    false,
					Description: "name of the application",
				},
			},
		},
		apislurp.Structure{
			Name:    "Message Reference",
			RawName: "Message Reference Structure",
			Type:    "",
			Fields: []apislurp.Field{
				apislurp.Field{
					Name:    "message_id",
					RawName: "message_id?",
					Type: apislurp.Type{
						Names: []apislurp.TypeName{
							"snowflake",
						},
						Array:     false,
						ArraySize: 0,
						Nullable:  false,
						RawText:   "snowflake",
					},
					Optional:    true,
					Description: "id of the originating message",
				},
				apislurp.Field{
					Name:    "channel_id",
					RawName: "channel_id",
					Type: apislurp.Type{
						Names: []apislurp.TypeName{
							"snowflake",
						},
						Array:     false,
						ArraySize: 0,
						Nullable:  false,
						RawText:   "snowflake",
					},
					Optional:    false,
					Description: "id of the originating message's channel",
				},
				apislurp.Field{
					Name:    "guild_id",
					RawName: "guild_id?",
					Type: apislurp.Type{
						Names: []apislurp.TypeName{
							"snowflake",
						},
						Array:     false,
						ArraySize: 0,
						Nullable:  false,
						RawText:   "snowflake",
					},
					Optional:    true,
					Description: "id of the originating message's guild",
				},
			},
		},
		apislurp.Structure{
			Name:    "Followed Channel",
			RawName: "Followed Channel Structure",
			Type:    "",
			Fields: []apislurp.Field{
				apislurp.Field{
					Name:    "channel_id",
					RawName: "channel_id",
					Type: apislurp.Type{
						Names: []apislurp.TypeName{
							"snowflake",
						},
						Array:     false,
						ArraySize: 0,
						Nullable:  false,
						RawText:   "snowflake",
					},
					Optional:    false,
					Description: "source channel id",
				},
				apislurp.Field{
					Name:    "webhook_id",
					RawName: "webhook_id",
					Type: apislurp.Type{
						Names: []apislurp.TypeName{
							"snowflake",
						},
						Array:     false,
						ArraySize: 0,
						Nullable:  false,
						RawText:   "snowflake",
					},
					Optional:    false,
					Description: "created target webhook id",
				},
			},
		},
		apislurp.Structure{
			Name:    "Reaction",
			RawName: "Reaction Structure",
			Type:    "",
			Fields: []apislurp.Field{
				apislurp.Field{
					Name:    "count",
					RawName: "count",
					Type: apislurp.Type{
						Names: []apislurp.TypeName{
							"integer",
						},
						Array:     false,
						ArraySize: 0,
						Nullable:  false,
						RawText:   "integer",
					},
					Optional:    false,
					Description: "times this emoji has been used to react",
				},
				apislurp.Field{
					Name:    "me",
					RawName: "me",
					Type: apislurp.Type{
						Names: []apislurp.TypeName{
							"boolean",
						},
						Array:     false,
						ArraySize: 0,
						Nullable:  false,
						RawText:   "boolean",
					},
					Optional:    false,
					Description: "whether the current user reacted using this emoji",
				},
				apislurp.Field{
					Name:    "emoji",
					RawName: "emoji",
					Type: apislurp.Type{
						Names: []apislurp.TypeName{
							"emoji",
						},
						Array:     false,
						ArraySize: 0,
						Nullable:  false,
						RawText:   "partial [emoji](#DOCS_RESOURCES_EMOJI/emoji-object) object",
					},
					Optional:    false,
					Description: "emoji information",
				},
			},
		},
		apislurp.Structure{
			Name:    "Overwrite",
			RawName: "Overwrite Structure",
			Type:    "",
			Fields: []apislurp.Field{
				apislurp.Field{
					Name:    "id",
					RawName: "id",
					Type: apislurp.Type{
						Names: []apislurp.TypeName{
							"snowflake",
						},
						Array:     false,
						ArraySize: 0,
						Nullable:  false,
						RawText:   "snowflake",
					},
					Optional:    false,
					Description: "role or user id",
				},
				apislurp.Field{
					Name:    "type",
					RawName: "type",
					Type: apislurp.Type{
						Names: []apislurp.TypeName{
							"integer",
						},
						Array:     false,
						ArraySize: 0,
						Nullable:  false,
						RawText:   "int",
					},
					Optional:    false,
					Description: "either 0 (role) or 1 (member)",
				},
				apislurp.Field{
					Name:    "allow",
					RawName: "allow",
					Type: apislurp.Type{
						Names: []apislurp.TypeName{
							"string",
						},
						Array:     false,
						ArraySize: 0,
						Nullable:  false,
						RawText:   "string",
					},
					Optional:    false,
					Description: "permission bit set",
				},
				apislurp.Field{
					Name:    "deny",
					RawName: "deny",
					Type: apislurp.Type{
						Names: []apislurp.TypeName{
							"string",
						},
						Array:     false,
						ArraySize: 0,
						Nullable:  false,
						RawText:   "string",
					},
					Optional:    false,
					Description: "permission bit set",
				},
			},
		},
		apislurp.Structure{
			Name:    "Embed",
			RawName: "Embed Structure",
			Type:    "",
			Fields: []apislurp.Field{
				apislurp.Field{
					Name:    "title",
					RawName: "title?",
					Type: apislurp.Type{
						Names: []apislurp.TypeName{
							"string",
						},
						Array:     false,
						ArraySize: 0,
						Nullable:  false,
						RawText:   "string",
					},
					Optional:    true,
					Description: "title of embed",
				},
				apislurp.Field{
					Name:    "type",
					RawName: "type?",
					Type: apislurp.Type{
						Names: []apislurp.TypeName{
							"string",
						},
						Array:     false,
						ArraySize: 0,
						Nullable:  false,
						RawText:   "string",
					},
					Optional:    true,
					Description: "[type of embed](#DOCS_RESOURCES_CHANNEL/embed-object-embed-types) (always \"rich\" for webhook embeds)",
				},
				apislurp.Field{
					Name:    "description",
					RawName: "description?",
					Type: apislurp.Type{
						Names: []apislurp.TypeName{
							"string",
						},
						Array:     false,
						ArraySize: 0,
						Nullable:  false,
						RawText:   "string",
					},
					Optional:    true,
					Description: "description of embed",
				},
				apislurp.Field{
					Name:    "url",
					RawName: "url?",
					Type: apislurp.Type{
						Names: []apislurp.TypeName{
							"string",
						},
						Array:     false,
						ArraySize: 0,
						Nullable:  false,
						RawText:   "string",
					},
					Optional:    true,
					Description: "url of embed",
				},
				apislurp.Field{
					Name:    "timestamp",
					RawName: "timestamp?",
					Type: apislurp.Type{
						Names: []apislurp.TypeName{
							"ISO8601 timestamp",
						},
						Array:     false,
						ArraySize: 0,
						Nullable:  false,
						RawText:   "ISO8601 timestamp",
					},
					Optional:    true,
					Description: "timestamp of embed content",
				},
				apislurp.Field{
					Name:    "color",
					RawName: "color?",
					Type: apislurp.Type{
						Names: []apislurp.TypeName{
							"integer",
						},
						Array:     false,
						ArraySize: 0,
						Nullable:  false,
						RawText:   "integer",
					},
					Optional:    true,
					Description: "color code of the embed",
				},
				apislurp.Field{
					Name:    "footer",
					RawName: "footer?",
					Type: apislurp.Type{
						Names: []apislurp.TypeName{
							"embed footer",
						},
						Array:     false,
						ArraySize: 0,
						Nullable:  false,
						RawText:   "[embed footer](#DOCS_RESOURCES_CHANNEL/embed-object-embed-footer-structure) object",
					},
					Optional:    true,
					Description: "footer information",
				},
				apislurp.Field{
					Name:    "image",
					RawName: "image?",
					Type: apislurp.Type{
						Names: []apislurp.TypeName{
							"embed image",
						},
						Array:     false,
						ArraySize: 0,
						Nullable:  false,
						RawText:   "[embed image](#DOCS_RESOURCES_CHANNEL/embed-object-embed-image-structure) object",
					},
					Optional:    true,
					Description: "image information",
				},
				apislurp.Field{
					Name:    "thumbnail",
					RawName: "thumbnail?",
					Type: apislurp.Type{
						Names: []apislurp.TypeName{
							"embed thumbnail",
						},
						Array:     false,
						ArraySize: 0,
						Nullable:  false,
						RawText:   "[embed thumbnail](#DOCS_RESOURCES_CHANNEL/embed-object-embed-thumbnail-structure) object",
					},
					Optional:    true,
					Description: "thumbnail information",
				},
				apislurp.Field{
					Name:    "video",
					RawName: "video?",
					Type: apislurp.Type{
						Names: []apislurp.TypeName{
							"embed video",
						},
						Array:     false,
						ArraySize: 0,
						Nullable:  false,
						RawText:   "[embed video](#DOCS_RESOURCES_CHANNEL/embed-object-embed-video-structure) object",
					},
					Optional:    true,
					Description: "video information",
				},
				apislurp.Field{
					Name:    "provider",
					RawName: "provider?",
					Type: apislurp.Type{
						Names: []apislurp.TypeName{
							"embed provider",
						},
						Array:     false,
						ArraySize: 0,
						Nullable:  false,
						RawText:   "[embed provider](#DOCS_RESOURCES_CHANNEL/embed-object-embed-provider-structure) object",
					},
					Optional:    true,
					Description: "provider information",
				},
				apislurp.Field{
					Name:    "author",
					RawName: "author?",
					Type: apislurp.Type{
						Names: []apislurp.TypeName{
							"embed author",
						},
						Array:     false,
						ArraySize: 0,
						Nullable:  false,
						RawText:   "[embed author](#DOCS_RESOURCES_CHANNEL/embed-object-embed-author-structure) object",
					},
					Optional:    true,
					Description: "author information",
				},
				apislurp.Field{
					Name:    "fields",
					RawName: "fields?",
					Type: apislurp.Type{
						Names: []apislurp.TypeName{
							"embed field",
						},
						Array:     false,
						ArraySize: 0,
						Nullable:  false,
						RawText:   "array of [embed field](#DOCS_RESOURCES_CHANNEL/embed-object-embed-field-structure) objects",
					},
					Optional:    true,
					Description: "fields information",
				},
			},
		},
		apislurp.Structure{
			Name:    "Embed Thumbnail",
			RawName: "Embed Thumbnail Structure",
			Type:    "",
			Fields: []apislurp.Field{
				apislurp.Field{
					Name:    "url",
					RawName: "url?",
					Type: apislurp.Type{
						Names: []apislurp.TypeName{
							"string",
						},
						Array:     false,
						ArraySize: 0,
						Nullable:  false,
						RawText:   "string",
					},
					Optional:    true,
					Description: "source url of thumbnail (only supports http(s) and attachments)",
				},
				apislurp.Field{
					Name:    "proxy_url",
					RawName: "proxy_url?",
					Type: apislurp.Type{
						Names: []apislurp.TypeName{
							"string",
						},
						Array:     false,
						ArraySize: 0,
						Nullable:  false,
						RawText:   "string",
					},
					Optional:    true,
					Description: "a proxied url of the thumbnail",
				},
				apislurp.Field{
					Name:    "height",
					RawName: "height?",
					Type: apislurp.Type{
						Names: []apislurp.TypeName{
							"integer",
						},
						Array:     false,
						ArraySize: 0,
						Nullable:  false,
						RawText:   "integer",
					},
					Optional:    true,
					Description: "height of thumbnail",
				},
				apislurp.Field{
					Name:    "width",
					RawName: "width?",
					Type: apislurp.Type{
						Names: []apislurp.TypeName{
							"integer",
						},
						Array:     false,
						ArraySize: 0,
						Nullable:  false,
						RawText:   "integer",
					},
					Optional:    true,
					Description: "width of thumbnail",
				},
			},
		},
		apislurp.Structure{
			Name:    "Embed Video",
			RawName: "Embed Video Structure",
			Type:    "",
			Fields: []apislurp.Field{
				apislurp.Field{
					Name:    "url",
					RawName: "url?",
					Type: apislurp.Type{
						Names: []apislurp.TypeName{
							"string",
						},
						Array:     false,
						ArraySize: 0,
						Nullable:  false,
						RawText:   "string",
					},
					Optional:    true,
					Description: "source url of video",
				},
				apislurp.Field{
					Name:    "height",
					RawName: "height?",
					Type: apislurp.Type{
						Names: []apislurp.TypeName{
							"integer",
						},
						Array:     false,
						ArraySize: 0,
						Nullable:  false,
						RawText:   "integer",
					},
					Optional:    true,
					Description: "height of video",
				},
				apislurp.Field{
					Name:    "width",
					RawName: "width?",
					Type: apislurp.Type{
						Names: []apislurp.TypeName{
							"integer",
						},
						Array:     false,
						ArraySize: 0,
						Nullable:  false,
						RawText:   "integer",
					},
					Optional:    true,
					Description: "width of video",
				},
			},
		},
		apislurp.Structure{
			Name:    "Embed Image",
			RawName: "Embed Image Structure",
			Type:    "",
			Fields: []apislurp.Field{
				apislurp.Field{
					Name:    "url",
					RawName: "url?",
					Type: apislurp.Type{
						Names: []apislurp.TypeName{
							"string",
						},
						Array:     false,
						ArraySize: 0,
						Nullable:  false,
						RawText:   "string",
					},
					Optional:    true,
					Description: "source url of image (only supports http(s) and attachments)",
				},
				apislurp.Field{
					Name:    "proxy_url",
					RawName: "proxy_url?",
					Type: apislurp.Type{
						Names: []apislurp.TypeName{
							"string",
						},
						Array:     false,
						ArraySize: 0,
						Nullable:  false,
						RawText:   "string",
					},
					Optional:    true,
					Description: "a proxied url of the image",
				},
				apislurp.Field{
					Name:    "height",
					RawName: "height?",
					Type: apislurp.Type{
						Names: []apislurp.TypeName{
							"integer",
						},
						Array:     false,
						ArraySize: 0,
						Nullable:  false,
						RawText:   "integer",
					},
					Optional:    true,
					Description: "height of image",
				},
				apislurp.Field{
					Name:    "width",
					RawName: "width?",
					Type: apislurp.Type{
						Names: []apislurp.TypeName{
							"integer",
						},
						Array:     false,
						ArraySize: 0,
						Nullable:  false,
						RawText:   "integer",
					},
					Optional:    true,
					Description: "width of image",
				},
			},
		},
		apislurp.Structure{
			Name:    "Embed Provider",
			RawName: "Embed Provider Structure",
			Type:    "",
			Fields: []apislurp.Field{
				apislurp.Field{
					Name:    "name",
					RawName: "name?",
					Type: apislurp.Type{
						Names: []apislurp.TypeName{
							"string",
						},
						Array:     false,
						ArraySize: 0,
						Nullable:  false,
						RawText:   "string",
					},
					Optional:    true,
					Description: "name of provider",
				},
				apislurp.Field{
					Name:    "url",
					RawName: "url?",
					Type: apislurp.Type{
						Names: []apislurp.TypeName{
							"string",
						},
						Array:     false,
						ArraySize: 0,
						Nullable:  false,
						RawText:   "string",
					},
					Optional:    true,
					Description: "url of provider",
				},
			},
		},
		apislurp.Structure{
			Name:    "Embed Author",
			RawName: "Embed Author Structure",
			Type:    "",
			Fields: []apislurp.Field{
				apislurp.Field{
					Name:    "name",
					RawName: "name?",
					Type: apislurp.Type{
						Names: []apislurp.TypeName{
							"string",
						},
						Array:     false,
						ArraySize: 0,
						Nullable:  false,
						RawText:   "string",
					},
					Optional:    true,
					Description: "name of author",
				},
				apislurp.Field{
					Name:    "url",
					RawName: "url?",
					Type: apislurp.Type{
						Names: []apislurp.TypeName{
							"string",
						},
						Array:     false,
						ArraySize: 0,
						Nullable:  false,
						RawText:   "string",
					},
					Optional:    true,
					Description: "url of author",
				},
				apislurp.Field{
					Name:    "icon_url",
					RawName: "icon_url?",
					Type: apislurp.Type{
						Names: []apislurp.TypeName{
							"string",
						},
						Array:     false,
						ArraySize: 0,
						Nullable:  false,
						RawText:   "string",
					},
					Optional:    true,
					Description: "url of author icon (only supports http(s) and attachments)",
				},
				apislurp.Field{
					Name:    "proxy_icon_url",
					RawName: "proxy_icon_url?",
					Type: apislurp.Type{
						Names: []apislurp.TypeName{
							"string",
						},
						Array:     false,
						ArraySize: 0,
						Nullable:  false,
						RawText:   "string",
					},
					Optional:    true,
					Description: "a proxied url of author icon",
				},
			},
		},
		apislurp.Structure{
			Name:    "Embed Footer",
			RawName: "Embed Footer Structure",
			Type:    "",
			Fields: []apislurp.Field{
				apislurp.Field{
					Name:    "text",
					RawName: "text",
					Type: apislurp.Type{
						Names: []apislurp.TypeName{
							"string",
						},
						Array:     false,
						ArraySize: 0,
						Nullable:  false,
						RawText:   "string",
					},
					Optional:    false,
					Description: "footer text",
				},
				apislurp.Field{
					Name:    "icon_url",
					RawName: "icon_url?",
					Type: apislurp.Type{
						Names: []apislurp.TypeName{
							"string",
						},
						Array:     false,
						ArraySize: 0,
						Nullable:  false,
						RawText:   "string",
					},
					Optional:    true,
					Description: "url of footer icon (only supports http(s) and attachments)",
				},
				apislurp.Field{
					Name:    "proxy_icon_url",
					RawName: "proxy_icon_url?",
					Type: apislurp.Type{
						Names: []apislurp.TypeName{
							"string",
						},
						Array:     false,
						ArraySize: 0,
						Nullable:  false,
						RawText:   "string",
					},
					Optional:    true,
					Description: "a proxied url of footer icon",
				},
			},
		},
		apislurp.Structure{
			Name:    "Embed Field",
			RawName: "Embed Field Structure",
			Type:    "",
			Fields: []apislurp.Field{
				apislurp.Field{
					Name:    "name",
					RawName: "name",
					Type: apislurp.Type{
						Names: []apislurp.TypeName{
							"string",
						},
						Array:     false,
						ArraySize: 0,
						Nullable:  false,
						RawText:   "string",
					},
					Optional:    false,
					Description: "name of the field",
				},
				apislurp.Field{
					Name:    "value",
					RawName: "value",
					Type: apislurp.Type{
						Names: []apislurp.TypeName{
							"string",
						},
						Array:     false,
						ArraySize: 0,
						Nullable:  false,
						RawText:   "string",
					},
					Optional:    false,
					Description: "value of the field",
				},
				apislurp.Field{
					Name:    "inline",
					RawName: "inline?",
					Type: apislurp.Type{
						Names: []apislurp.TypeName{
							"boolean",
						},
						Array:     false,
						ArraySize: 0,
						Nullable:  false,
						RawText:   "boolean",
					},
					Optional:    true,
					Description: "whether or not this field should display inline",
				},
			},
		},
		apislurp.Structure{
			Name:    "Attachment",
			RawName: "Attachment Structure",
			Type:    "",
			Fields: []apislurp.Field{
				apislurp.Field{
					Name:    "id",
					RawName: "id",
					Type: apislurp.Type{
						Names: []apislurp.TypeName{
							"snowflake",
						},
						Array:     false,
						ArraySize: 0,
						Nullable:  false,
						RawText:   "snowflake",
					},
					Optional:    false,
					Description: "attachment id",
				},
				apislurp.Field{
					Name:    "filename",
					RawName: "filename",
					Type: apislurp.Type{
						Names: []apislurp.TypeName{
							"string",
						},
						Array:     false,
						ArraySize: 0,
						Nullable:  false,
						RawText:   "string",
					},
					Optional:    false,
					Description: "name of file attached",
				},
				apislurp.Field{
					Name:    "size",
					RawName: "size",
					Type: apislurp.Type{
						Names: []apislurp.TypeName{
							"integer",
						},
						Array:     false,
						ArraySize: 0,
						Nullable:  false,
						RawText:   "integer",
					},
					Optional:    false,
					Description: "size of file in bytes",
				},
				apislurp.Field{
					Name:    "url",
					RawName: "url",
					Type: apislurp.Type{
						Names: []apislurp.TypeName{
							"string",
						},
						Array:     false,
						ArraySize: 0,
						Nullable:  false,
						RawText:   "string",
					},
					Optional:    false,
					Description: "source url of file",
				},
				apislurp.Field{
					Name:    "proxy_url",
					RawName: "proxy_url",
					Type: apislurp.Type{
						Names: []apislurp.TypeName{
							"string",
						},
						Array:     false,
						ArraySize: 0,
						Nullable:  false,
						RawText:   "string",
					},
					Optional:    false,
					Description: "a proxied url of file",
				},
				apislurp.Field{
					Name:    "height",
					RawName: "height",
					Type: apislurp.Type{
						Names: []apislurp.TypeName{
							"integer",
						},
						Array:     false,
						ArraySize: 0,
						Nullable:  true,
						RawText:   "?integer",
					},
					Optional:    false,
					Description: "height of file (if image)",
				},
				apislurp.Field{
					Name:    "width",
					RawName: "width",
					Type: apislurp.Type{
						Names: []apislurp.TypeName{
							"integer",
						},
						Array:     false,
						ArraySize: 0,
						Nullable:  true,
						RawText:   "?integer",
					},
					Optional:    false,
					Description: "width of file (if image)",
				},
			},
		},
		apislurp.Structure{
			Name:    "Channel Mention",
			RawName: "Channel Mention Structure",
			Type:    "",
			Fields: []apislurp.Field{
				apislurp.Field{
					Name:    "id",
					RawName: "id",
					Type: apislurp.Type{
						Names: []apislurp.TypeName{
							"snowflake",
						},
						Array:     false,
						ArraySize: 0,
						Nullable:  false,
						RawText:   "snowflake",
					},
					Optional:    false,
					Description: "id of the channel",
				},
				apislurp.Field{
					Name:    "guild_id",
					RawName: "guild_id",
					Type: apislurp.Type{
						Names: []apislurp.TypeName{
							"snowflake",
						},
						Array:     false,
						ArraySize: 0,
						Nullable:  false,
						RawText:   "snowflake",
					},
					Optional:    false,
					Description: "id of the guild containing the channel",
				},
				apislurp.Field{
					Name:    "type",
					RawName: "type",
					Type: apislurp.Type{
						Names: []apislurp.TypeName{
							"integer",
						},
						Array:     false,
						ArraySize: 0,
						Nullable:  false,
						RawText:   "integer",
					},
					Optional:    false,
					Description: "the [type of channel](#DOCS_RESOURCES_CHANNEL/channel-object-channel-types)",
				},
				apislurp.Field{
					Name:    "name",
					RawName: "name",
					Type: apislurp.Type{
						Names: []apislurp.TypeName{
							"string",
						},
						Array:     false,
						ArraySize: 0,
						Nullable:  false,
						RawText:   "string",
					},
					Optional:    false,
					Description: "the name of the channel",
				},
			},
		},
		apislurp.Structure{
			Name:    "Allowed Mentions",
			RawName: "Allowed Mentions Structure",
			Type:    "",
			Fields: []apislurp.Field{
				apislurp.Field{
					Name:    "parse",
					RawName: "parse",
					Type: apislurp.Type{
						Names: []apislurp.TypeName{
							" allowed mention type",
						},
						Array:     true,
						ArraySize: 0,
						Nullable:  false,
						RawText:   "array of allowed mention types",
					},
					Optional:    false,
					Description: "An array of [allowed mention types](#DOCS_RESOURCES_CHANNEL/allowed-mentions-object-allowed-mention-types) to parse from the content.",
				},
				apislurp.Field{
					Name:    "roles",
					RawName: "roles",
					Type: apislurp.Type{
						Names: []apislurp.TypeName{
							" snowflake",
						},
						Array:     true,
						ArraySize: 0,
						Nullable:  false,
						RawText:   "list of snowflakes",
					},
					Optional:    false,
					Description: "Array of role_ids to mention (Max size of 100)",
				},
				apislurp.Field{
					Name:    "users",
					RawName: "users",
					Type: apislurp.Type{
						Names: []apislurp.TypeName{
							" snowflake",
						},
						Array:     true,
						ArraySize: 0,
						Nullable:  false,
						RawText:   "list of snowflakes",
					},
					Optional:    false,
					Description: "Array of user_ids to mention (Max size of 100)",
				},
			},
		},
	},
}
