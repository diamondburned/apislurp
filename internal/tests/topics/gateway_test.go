package apislurp_test

import (
	"github.com/diamondburned/apislurp"
	"github.com/diamondburned/apislurp/internal/tests/testutil"
	"path/filepath"
	"testing"
)

func TestGateway(t *testing.T) {
	path := filepath.Join("topics", "Gateway.md")
	f := testutil.ReadFile(t, path)
	d := apislurp.Parse(f)
	testutil.Equal(t, d, gatewayData, "Gateway")
}

var gatewayData = apislurp.Document{
	Structures: []apislurp.Structure{
		apislurp.Structure{
			Name:    "Gateway Payload",
			RawName: "Gateway Payload Structure",
			Type:    "",
			Fields: []apislurp.Field{
				apislurp.Field{
					Name:    "op",
					RawName: "op",
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
					Description: "[opcode](#DOCS_TOPICS_OPCODES_AND_STATUS_CODES/gateway-opcodes) for the payload",
				},
				apislurp.Field{
					Name:    "d",
					RawName: "d",
					Type: apislurp.Type{
						Names: []apislurp.TypeName{
							"mixed",
						},
						Array:     false,
						ArraySize: 0,
						Nullable:  true,
						RawText:   "?mixed (any JSON value)",
					},
					Optional:    false,
					Description: "event data",
				},
				apislurp.Field{
					Name:    "s",
					RawName: "s",
					Type: apislurp.Type{
						Names: []apislurp.TypeName{
							"integer \\*",
						},
						Array:     false,
						ArraySize: 0,
						Nullable:  true,
						RawText:   "?integer \\*",
					},
					Optional:    false,
					Description: "sequence number, used for resuming sessions and heartbeats",
				},
				apislurp.Field{
					Name:    "t",
					RawName: "t",
					Type: apislurp.Type{
						Names: []apislurp.TypeName{
							"string \\*",
						},
						Array:     false,
						ArraySize: 0,
						Nullable:  true,
						RawText:   "?string \\*",
					},
					Optional:    false,
					Description: "the event name for this payload",
				},
			},
		},
		apislurp.Structure{
			Name:    "Identify",
			RawName: "Identify Structure",
			Type:    "",
			Fields: []apislurp.Field{
				apislurp.Field{
					Name:    "token",
					RawName: "token",
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
					Description: "authentication token",
				},
				apislurp.Field{
					Name:    "properties",
					RawName: "properties",
					Type: apislurp.Type{
						Names: []apislurp.TypeName{
							"object",
						},
						Array:     false,
						ArraySize: 0,
						Nullable:  false,
						RawText:   "object",
					},
					Optional:    false,
					Description: "[connection properties](#DOCS_TOPICS_GATEWAY/identify-identify-connection-properties)",
				},
				apislurp.Field{
					Name:    "compress",
					RawName: "compress?",
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
					Description: "whether this connection supports compression of packets",
				},
				apislurp.Field{
					Name:    "large_threshold",
					RawName: "large_threshold?",
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
					Description: "value between 50 and 250, total number of members where the gateway will stop sending offline members in the guild member list",
				},
				apislurp.Field{
					Name:    "shard",
					RawName: "shard?",
					Type: apislurp.Type{
						Names: []apislurp.TypeName{
							" two integer",
						},
						Array:     true,
						ArraySize: 0,
						Nullable:  false,
						RawText:   "array of two integers (shard_id, num_shards)",
					},
					Optional:    true,
					Description: "used for [Guild Sharding](#DOCS_TOPICS_GATEWAY/sharding)",
				},
				apislurp.Field{
					Name:    "presence",
					RawName: "presence?",
					Type: apislurp.Type{
						Names: []apislurp.TypeName{
							"update status",
						},
						Array:     false,
						ArraySize: 0,
						Nullable:  false,
						RawText:   "[update status](#DOCS_TOPICS_GATEWAY/update-status) object",
					},
					Optional:    true,
					Description: "presence structure for initial presence information",
				},
				apislurp.Field{
					Name:    "guild_subscriptions",
					RawName: "guild_subscriptions?",
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
					Description: "enables dispatching of guild subscription events (presence and typing events)",
				},
				apislurp.Field{
					Name:    "intents",
					RawName: "intents",
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
					Description: "the [Gateway Intents](#DOCS_TOPICS_GATEWAY/gateway-intents) you wish to receive",
				},
			},
		},
		apislurp.Structure{
			Name:    "Identify Connection",
			RawName: "Identify Connection Properties",
			Type:    "",
			Fields: []apislurp.Field{
				apislurp.Field{
					Name:    "\\$os",
					RawName: "\\$os",
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
					Description: "your operating system",
				},
				apislurp.Field{
					Name:    "\\$browser",
					RawName: "\\$browser",
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
					Description: "your library name",
				},
				apislurp.Field{
					Name:    "\\$device",
					RawName: "\\$device",
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
					Description: "your library name",
				},
			},
		},
		apislurp.Structure{
			Name:    "Resume",
			RawName: "Resume Structure",
			Type:    "",
			Fields: []apislurp.Field{
				apislurp.Field{
					Name:    "token",
					RawName: "token",
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
					Description: "session token",
				},
				apislurp.Field{
					Name:    "session_id",
					RawName: "session_id",
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
					Description: "session id",
				},
				apislurp.Field{
					Name:    "seq",
					RawName: "seq",
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
					Description: "last sequence number received",
				},
			},
		},
		apislurp.Structure{
			Name:    "Guild Request Members",
			RawName: "Guild Request Members Structure",
			Type:    "",
			Fields: []apislurp.Field{
				apislurp.Field{
					Name:    "guild_id",
					RawName: "guild_id",
					Type: apislurp.Type{
						Names: []apislurp.TypeName{
							"snowflake",
							" snowflake",
						},
						Array:     true,
						ArraySize: 0,
						Nullable:  false,
						RawText:   "snowflake or array of snowflakes",
					},
					Optional:    false,
					Description: "id of the guild(s) to get members for",
				},
				apislurp.Field{
					Name:    "query",
					RawName: "query?",
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
					Description: "string that username starts with, or an empty string to return all members",
				},
				apislurp.Field{
					Name:    "limit",
					RawName: "limit",
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
					Description: "maximum number of members to send matching the `query`; a limit of `0` can be used with an empty string `query` to return all members",
				},
				apislurp.Field{
					Name:    "presences",
					RawName: "presences?",
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
					Description: "used to specify if we want the presences of the matched members",
				},
				apislurp.Field{
					Name:    "user_ids",
					RawName: "user_ids?",
					Type: apislurp.Type{
						Names: []apislurp.TypeName{
							"snowflake",
							" snowflake",
						},
						Array:     true,
						ArraySize: 0,
						Nullable:  false,
						RawText:   "snowflake or array of snowflakes",
					},
					Optional:    true,
					Description: "used to specify which users you wish to fetch",
				},
				apislurp.Field{
					Name:    "nonce",
					RawName: "nonce?",
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
					Description: "nonce to identify the [Guild Members Chunk](#DOCS_TOPICS_GATEWAY/guild-members-chunk) response",
				},
			},
		},
		apislurp.Structure{
			Name:    "Gateway Voice State Update",
			RawName: "Gateway Voice State Update Structure",
			Type:    "",
			Fields: []apislurp.Field{
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
					Description: "id of the guild",
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
						Nullable:  true,
						RawText:   "?snowflake",
					},
					Optional:    false,
					Description: "id of the voice channel client wants to join (null if disconnecting)",
				},
				apislurp.Field{
					Name:    "self_mute",
					RawName: "self_mute",
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
					Description: "is the client muted",
				},
				apislurp.Field{
					Name:    "self_deaf",
					RawName: "self_deaf",
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
					Description: "is the client deafened",
				},
			},
		},
		apislurp.Structure{
			Name:    "Gateway Status Update",
			RawName: "Gateway Status Update Structure",
			Type:    "",
			Fields: []apislurp.Field{
				apislurp.Field{
					Name:    "since",
					RawName: "since",
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
					Description: "unix time (in milliseconds) of when the client went idle, or null if the client is not idle",
				},
				apislurp.Field{
					Name:    "activities",
					RawName: "activities",
					Type: apislurp.Type{
						Names: []apislurp.TypeName{
							"activity",
						},
						Array:     false,
						ArraySize: 0,
						Nullable:  true,
						RawText:   "?array of [activity](#DOCS_TOPICS_GATEWAY/activity-object) objects",
					},
					Optional:    false,
					Description: "null, or the user's activities",
				},
				apislurp.Field{
					Name:    "status",
					RawName: "status",
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
					Description: "the user's new [status](#DOCS_TOPICS_GATEWAY/update-status-status-types)",
				},
				apislurp.Field{
					Name:    "afk",
					RawName: "afk",
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
					Description: "whether or not the client is afk",
				},
			},
		},
		apislurp.Structure{
			Name:    "Hello",
			RawName: "Hello Structure",
			Type:    "",
			Fields: []apislurp.Field{
				apislurp.Field{
					Name:    "heartbeat_interval",
					RawName: "heartbeat_interval",
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
					Description: "the interval (in milliseconds) the client should heartbeat with",
				},
			},
		},
		apislurp.Structure{
			Name:    "Ready",
			RawName: "Ready Event Fields",
			Type:    "",
			Fields: []apislurp.Field{
				apislurp.Field{
					Name:    "v",
					RawName: "v",
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
					Description: "[gateway version](#DOCS_TOPICS_GATEWAY/gateways-gateway-versions)",
				},
				apislurp.Field{
					Name:    "user",
					RawName: "user",
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
					Description: "information about the user including email",
				},
				apislurp.Field{
					Name:    "private_channels",
					RawName: "private_channels",
					Type: apislurp.Type{
						Names: []apislurp.TypeName{
							"empty array",
						},
						Array:     false,
						ArraySize: 0,
						Nullable:  false,
						RawText:   "array",
					},
					Optional:    false,
					Description: "empty array",
				},
				apislurp.Field{
					Name:    "guilds",
					RawName: "guilds",
					Type: apislurp.Type{
						Names: []apislurp.TypeName{
							"Unavailable Guild",
						},
						Array:     false,
						ArraySize: 0,
						Nullable:  false,
						RawText:   "array of [Unavailable Guild](#DOCS_RESOURCES_GUILD/unavailable-guild-object) objects",
					},
					Optional:    false,
					Description: "the guilds the user is in",
				},
				apislurp.Field{
					Name:    "session_id",
					RawName: "session_id",
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
					Description: "used for resuming connections",
				},
				apislurp.Field{
					Name:    "shard",
					RawName: "shard?",
					Type: apislurp.Type{
						Names: []apislurp.TypeName{
							" two integer",
						},
						Array:     true,
						ArraySize: 0,
						Nullable:  false,
						RawText:   "array of two integers (shard_id, num_shards)",
					},
					Optional:    true,
					Description: "the [shard information](#DOCS_TOPICS_GATEWAY/sharding) associated with this session, if sent when identifying",
				},
			},
		},
		apislurp.Structure{
			Name:    "Channel Pins Update",
			RawName: "Channel Pins Update Event Fields",
			Type:    "",
			Fields: []apislurp.Field{
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
					Description: "the id of the channel",
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
						Nullable:  false,
						RawText:   "ISO8601 timestamp",
					},
					Optional:    true,
					Description: "the time at which the most recent pinned message was pinned",
				},
			},
		},
		apislurp.Structure{
			Name:    "Guild Ban Add",
			RawName: "Guild Ban Add Event Fields",
			Type:    "",
			Fields: []apislurp.Field{
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
					Description: "id of the guild",
				},
				apislurp.Field{
					Name:    "user",
					RawName: "user",
					Type: apislurp.Type{
						Names: []apislurp.TypeName{
							"user",
						},
						Array:     false,
						ArraySize: 0,
						Nullable:  false,
						RawText:   "a [user](#DOCS_RESOURCES_USER/user-object) object",
					},
					Optional:    false,
					Description: "the banned user",
				},
			},
		},
		apislurp.Structure{
			Name:    "Guild Ban Remove",
			RawName: "Guild Ban Remove Event Fields",
			Type:    "",
			Fields: []apislurp.Field{
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
					Description: "id of the guild",
				},
				apislurp.Field{
					Name:    "user",
					RawName: "user",
					Type: apislurp.Type{
						Names: []apislurp.TypeName{
							"user",
						},
						Array:     false,
						ArraySize: 0,
						Nullable:  false,
						RawText:   "a [user](#DOCS_RESOURCES_USER/user-object) object",
					},
					Optional:    false,
					Description: "the unbanned user",
				},
			},
		},
		apislurp.Structure{
			Name:    "Guild Emojis Update",
			RawName: "Guild Emojis Update Event Fields",
			Type:    "",
			Fields: []apislurp.Field{
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
					Description: "id of the guild",
				},
				apislurp.Field{
					Name:    "emojis",
					RawName: "emojis",
					Type: apislurp.Type{
						Names: []apislurp.TypeName{
							"emojis",
						},
						Array:     false,
						ArraySize: 0,
						Nullable:  false,
						RawText:   "array",
					},
					Optional:    false,
					Description: "array of [emojis](#DOCS_RESOURCES_EMOJI/emoji-object)",
				},
			},
		},
		apislurp.Structure{
			Name:    "Guild Integrations Update",
			RawName: "Guild Integrations Update Event Fields",
			Type:    "",
			Fields: []apislurp.Field{
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
					Description: "id of the guild whose integrations were updated",
				},
			},
		},
		apislurp.Structure{
			Name:    "Guild Member Remove",
			RawName: "Guild Member Remove Event Fields",
			Type:    "",
			Fields: []apislurp.Field{
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
					Description: "the id of the guild",
				},
				apislurp.Field{
					Name:    "user",
					RawName: "user",
					Type: apislurp.Type{
						Names: []apislurp.TypeName{
							"user",
						},
						Array:     false,
						ArraySize: 0,
						Nullable:  false,
						RawText:   "a [user](#DOCS_RESOURCES_USER/user-object) object",
					},
					Optional:    false,
					Description: "the user who was removed",
				},
			},
		},
		apislurp.Structure{
			Name:    "Guild Member Update",
			RawName: "Guild Member Update Event Fields",
			Type:    "",
			Fields: []apislurp.Field{
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
					Description: "the id of the guild",
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
						RawText:   "array of snowflakes",
					},
					Optional:    false,
					Description: "user role ids",
				},
				apislurp.Field{
					Name:    "user",
					RawName: "user",
					Type: apislurp.Type{
						Names: []apislurp.TypeName{
							"user",
						},
						Array:     false,
						ArraySize: 0,
						Nullable:  false,
						RawText:   "a [user](#DOCS_RESOURCES_USER/user-object) object",
					},
					Optional:    false,
					Description: "the user",
				},
				apislurp.Field{
					Name:    "nick",
					RawName: "nick?",
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
					Description: "nickname of the user in the guild",
				},
				apislurp.Field{
					Name:    "joined_at",
					RawName: "joined_at",
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
					Description: "when the user joined the guild",
				},
				apislurp.Field{
					Name:    "premium_since",
					RawName: "premium_since?",
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
					Description: "when the user starting [boosting](https://support.discord.com/hc/en-us/articles/360028038352-Server-Boosting-) the guild",
				},
			},
		},
		apislurp.Structure{
			Name:    "Guild Members Chunk",
			RawName: "Guild Members Chunk Event Fields",
			Type:    "",
			Fields: []apislurp.Field{
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
					Description: "the id of the guild",
				},
				apislurp.Field{
					Name:    "members",
					RawName: "members",
					Type: apislurp.Type{
						Names: []apislurp.TypeName{
							"guild member",
						},
						Array:     false,
						ArraySize: 0,
						Nullable:  false,
						RawText:   "array of [guild member](#DOCS_RESOURCES_GUILD/guild-member-object) objects",
					},
					Optional:    false,
					Description: "set of guild members",
				},
				apislurp.Field{
					Name:    "chunk_index",
					RawName: "chunk_index",
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
					Description: "the chunk index in the expected chunks for this response (0 <= chunk\\_index < chunk\\_count)",
				},
				apislurp.Field{
					Name:    "chunk_count",
					RawName: "chunk_count",
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
					Description: "the total number of expected chunks for this response",
				},
				apislurp.Field{
					Name:    "not_found",
					RawName: "not_found?",
					Type: apislurp.Type{
						Names: []apislurp.TypeName{
							"if passing an invalid id to `REQUEST_GUILD_MEMBERS`, it will be returned here",
						},
						Array:     false,
						ArraySize: 0,
						Nullable:  false,
						RawText:   "array",
					},
					Optional:    true,
					Description: "if passing an invalid id to `REQUEST_GUILD_MEMBERS`, it will be returned here",
				},
				apislurp.Field{
					Name:    "presences",
					RawName: "presences?",
					Type: apislurp.Type{
						Names: []apislurp.TypeName{
							"presence",
						},
						Array:     false,
						ArraySize: 0,
						Nullable:  false,
						RawText:   "array of [presence](#DOCS_TOPICS_GATEWAY/presence) objects",
					},
					Optional:    true,
					Description: "if passing true to `REQUEST_GUILD_MEMBERS`, presences of the returned members will be here",
				},
				apislurp.Field{
					Name:    "nonce",
					RawName: "nonce?",
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
					Description: "the nonce used in the [Guild Members Request](#DOCS_TOPICS_GATEWAY/request-guild-members)",
				},
			},
		},
		apislurp.Structure{
			Name:    "Guild Role Create",
			RawName: "Guild Role Create Event Fields",
			Type:    "",
			Fields: []apislurp.Field{
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
					Description: "the id of the guild",
				},
				apislurp.Field{
					Name:    "role",
					RawName: "role",
					Type: apislurp.Type{
						Names: []apislurp.TypeName{
							"role",
						},
						Array:     false,
						ArraySize: 0,
						Nullable:  false,
						RawText:   "a [role](#DOCS_TOPICS_PERMISSIONS/role-object) object",
					},
					Optional:    false,
					Description: "the role created",
				},
			},
		},
		apislurp.Structure{
			Name:    "Guild Role Update",
			RawName: "Guild Role Update Event Fields",
			Type:    "",
			Fields: []apislurp.Field{
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
					Description: "the id of the guild",
				},
				apislurp.Field{
					Name:    "role",
					RawName: "role",
					Type: apislurp.Type{
						Names: []apislurp.TypeName{
							"role",
						},
						Array:     false,
						ArraySize: 0,
						Nullable:  false,
						RawText:   "a [role](#DOCS_TOPICS_PERMISSIONS/role-object) object",
					},
					Optional:    false,
					Description: "the role updated",
				},
			},
		},
		apislurp.Structure{
			Name:    "Guild Role Delete",
			RawName: "Guild Role Delete Event Fields",
			Type:    "",
			Fields: []apislurp.Field{
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
					Description: "id of the guild",
				},
				apislurp.Field{
					Name:    "role_id",
					RawName: "role_id",
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
					Description: "id of the role",
				},
			},
		},
		apislurp.Structure{
			Name:    "Invite Create",
			RawName: "Invite Create Event Fields",
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
					Description: "the channel the invite is for",
				},
				apislurp.Field{
					Name:    "code",
					RawName: "code",
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
					Description: "the unique invite [code](#DOCS_RESOURCES_INVITE/invite-object)",
				},
				apislurp.Field{
					Name:    "created_at",
					RawName: "created_at",
					Type: apislurp.Type{
						Names: []apislurp.TypeName{
							"timestamp",
						},
						Array:     false,
						ArraySize: 0,
						Nullable:  false,
						RawText:   "timestamp",
					},
					Optional:    false,
					Description: "the time at which the invite was created",
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
					Description: "the guild of the invite",
				},
				apislurp.Field{
					Name:    "inviter",
					RawName: "inviter?",
					Type: apislurp.Type{
						Names: []apislurp.TypeName{
							"user",
						},
						Array:     false,
						ArraySize: 0,
						Nullable:  false,
						RawText:   "[user](#DOCS_RESOURCES_USER/user-object) object",
					},
					Optional:    true,
					Description: "the user that created the invite",
				},
				apislurp.Field{
					Name:    "max_age",
					RawName: "max_age",
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
					Description: "how long the invite is valid for (in seconds)",
				},
				apislurp.Field{
					Name:    "max_uses",
					RawName: "max_uses",
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
					Description: "the maximum number of times the invite can be used",
				},
				apislurp.Field{
					Name:    "target_user",
					RawName: "target_user?",
					Type: apislurp.Type{
						Names: []apislurp.TypeName{
							"user",
						},
						Array:     false,
						ArraySize: 0,
						Nullable:  false,
						RawText:   "partial [user](#DOCS_RESOURCES_USER/user-object) object",
					},
					Optional:    true,
					Description: "the target user for this invite",
				},
				apislurp.Field{
					Name:    "target_user_type",
					RawName: "target_user_type?",
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
					Description: "the [type of user target](#DOCS_RESOURCES_INVITE/invite-object-target-user-types) for this invite",
				},
				apislurp.Field{
					Name:    "temporary",
					RawName: "temporary",
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
					Description: "whether or not the invite is temporary (invited users will be kicked on disconnect unless they're assigned a role)",
				},
				apislurp.Field{
					Name:    "uses",
					RawName: "uses",
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
					Description: "how many times the invite has been used (always will be 0)",
				},
			},
		},
		apislurp.Structure{
			Name:    "Invite Delete",
			RawName: "Invite Delete Event Fields",
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
					Description: "the channel of the invite",
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
					Description: "the guild of the invite",
				},
				apislurp.Field{
					Name:    "code",
					RawName: "code",
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
					Description: "the unique invite [code](#DOCS_RESOURCES_INVITE/invite-object)",
				},
			},
		},
		apislurp.Structure{
			Name:    "Message Delete",
			RawName: "Message Delete Event Fields",
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
					Description: "the id of the message",
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
					Description: "the id of the channel",
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
			},
		},
		apislurp.Structure{
			Name:    "Message Delete Bulk",
			RawName: "Message Delete Bulk Event Fields",
			Type:    "",
			Fields: []apislurp.Field{
				apislurp.Field{
					Name:    "ids",
					RawName: "ids",
					Type: apislurp.Type{
						Names: []apislurp.TypeName{
							" snowflake",
						},
						Array:     true,
						ArraySize: 0,
						Nullable:  false,
						RawText:   "array of snowflakes",
					},
					Optional:    false,
					Description: "the ids of the messages",
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
					Description: "the id of the channel",
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
			},
		},
		apislurp.Structure{
			Name:    "Message Reaction Add",
			RawName: "Message Reaction Add Event Fields",
			Type:    "",
			Fields: []apislurp.Field{
				apislurp.Field{
					Name:    "user_id",
					RawName: "user_id",
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
					Description: "the id of the user",
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
					Description: "the id of the channel",
				},
				apislurp.Field{
					Name:    "message_id",
					RawName: "message_id",
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
					Description: "the id of the message",
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
					Name:    "member",
					RawName: "member?",
					Type: apislurp.Type{
						Names: []apislurp.TypeName{
							"guild member",
						},
						Array:     false,
						ArraySize: 0,
						Nullable:  false,
						RawText:   "[member](#DOCS_RESOURCES_GUILD/guild-member-object) object",
					},
					Optional:    true,
					Description: "the member who reacted if this happened in a guild",
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
						RawText:   "a partial [emoji](#DOCS_RESOURCES_EMOJI/emoji-object) object",
					},
					Optional:    false,
					Description: "the emoji used to react - [example](#DOCS_RESOURCES_EMOJI/emoji-object-gateway-reaction-standard-emoji-example)",
				},
			},
		},
		apislurp.Structure{
			Name:    "Message Reaction Remove",
			RawName: "Message Reaction Remove Event Fields",
			Type:    "",
			Fields: []apislurp.Field{
				apislurp.Field{
					Name:    "user_id",
					RawName: "user_id",
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
					Description: "the id of the user",
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
					Description: "the id of the channel",
				},
				apislurp.Field{
					Name:    "message_id",
					RawName: "message_id",
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
					Description: "the id of the message",
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
					Name:    "emoji",
					RawName: "emoji",
					Type: apislurp.Type{
						Names: []apislurp.TypeName{
							"emoji",
						},
						Array:     false,
						ArraySize: 0,
						Nullable:  false,
						RawText:   "a partial [emoji](#DOCS_RESOURCES_EMOJI/emoji-object) object",
					},
					Optional:    false,
					Description: "the emoji used to react - [example](#DOCS_RESOURCES_EMOJI/emoji-object-gateway-reaction-standard-emoji-example)",
				},
			},
		},
		apislurp.Structure{
			Name:    "Message Reaction Remove All",
			RawName: "Message Reaction Remove All Event Fields",
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
					Description: "the id of the channel",
				},
				apislurp.Field{
					Name:    "message_id",
					RawName: "message_id",
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
					Description: "the id of the message",
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
			},
		},
		apislurp.Structure{
			Name:    "Presence Update",
			RawName: "Presence Update Event Fields",
			Type:    "",
			Fields: []apislurp.Field{
				apislurp.Field{
					Name:    "user",
					RawName: "user",
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
					Description: "the user presence is being updated for",
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
					Description: "id of the guild",
				},
				apislurp.Field{
					Name:    "status",
					RawName: "status",
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
					Description: "either \"idle\", \"dnd\", \"online\", or \"offline\"",
				},
				apislurp.Field{
					Name:    "activities",
					RawName: "activities",
					Type: apislurp.Type{
						Names: []apislurp.TypeName{
							"activity",
						},
						Array:     false,
						ArraySize: 0,
						Nullable:  false,
						RawText:   "array of [activity](#DOCS_TOPICS_GATEWAY/activity-object) objects",
					},
					Optional:    false,
					Description: "user's current activities",
				},
				apislurp.Field{
					Name:    "client_status",
					RawName: "client_status",
					Type: apislurp.Type{
						Names: []apislurp.TypeName{
							"client_status",
						},
						Array:     false,
						ArraySize: 0,
						Nullable:  false,
						RawText:   "[client_status](#DOCS_TOPICS_GATEWAY/client-status-object) object",
					},
					Optional:    false,
					Description: "user's platform-dependent status",
				},
				apislurp.Field{
					Name:    "desktop",
					RawName: "desktop?",
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
					Description: "the user's status set for an active desktop (Windows, Linux, Mac) application session",
				},
				apislurp.Field{
					Name:    "mobile",
					RawName: "mobile?",
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
					Description: "the user's status set for an active mobile (iOS, Android) application session",
				},
				apislurp.Field{
					Name:    "web",
					RawName: "web?",
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
					Description: "the user's status set for an active web (browser, bot account) application session",
				},
			},
		},
		apislurp.Structure{
			Name:    "Activity",
			RawName: "Activity Structure",
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
					Description: "the activity's name",
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
					Description: "[activity type](#DOCS_TOPICS_GATEWAY/activity-object-activity-types)",
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
						Nullable:  true,
						RawText:   "?string",
					},
					Optional:    true,
					Description: "stream url, is validated when type is 1",
				},
				apislurp.Field{
					Name:    "created_at",
					RawName: "created_at",
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
					Description: "unix timestamp of when the activity was added to the user's session",
				},
				apislurp.Field{
					Name:    "timestamps",
					RawName: "timestamps?",
					Type: apislurp.Type{
						Names: []apislurp.TypeName{
							"timestamps",
						},
						Array:     false,
						ArraySize: 0,
						Nullable:  false,
						RawText:   "[timestamps](#DOCS_TOPICS_GATEWAY/activity-object-activity-timestamps) object",
					},
					Optional:    true,
					Description: "unix timestamps for start and/or end of the game",
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
					Description: "application id for the game",
				},
				apislurp.Field{
					Name:    "details",
					RawName: "details?",
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
					Description: "what the player is currently doing",
				},
				apislurp.Field{
					Name:    "state",
					RawName: "state?",
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
					Description: "the user's current party status",
				},
				apislurp.Field{
					Name:    "emoji",
					RawName: "emoji?",
					Type: apislurp.Type{
						Names: []apislurp.TypeName{
							"emoji",
						},
						Array:     false,
						ArraySize: 0,
						Nullable:  true,
						RawText:   "?[emoji](#DOCS_TOPICS_GATEWAY/activity-object-activity-emoji) object",
					},
					Optional:    true,
					Description: "the emoji used for a custom status",
				},
				apislurp.Field{
					Name:    "party",
					RawName: "party?",
					Type: apislurp.Type{
						Names: []apislurp.TypeName{
							"party",
						},
						Array:     false,
						ArraySize: 0,
						Nullable:  false,
						RawText:   "[party](#DOCS_TOPICS_GATEWAY/activity-object-activity-party) object",
					},
					Optional:    true,
					Description: "information for the current party of the player",
				},
				apislurp.Field{
					Name:    "assets",
					RawName: "assets?",
					Type: apislurp.Type{
						Names: []apislurp.TypeName{
							"assets",
						},
						Array:     false,
						ArraySize: 0,
						Nullable:  false,
						RawText:   "[assets](#DOCS_TOPICS_GATEWAY/activity-object-activity-assets) object",
					},
					Optional:    true,
					Description: "images for the presence and their hover texts",
				},
				apislurp.Field{
					Name:    "secrets",
					RawName: "secrets?",
					Type: apislurp.Type{
						Names: []apislurp.TypeName{
							"secrets",
						},
						Array:     false,
						ArraySize: 0,
						Nullable:  false,
						RawText:   "[secrets](#DOCS_TOPICS_GATEWAY/activity-object-activity-secrets) object",
					},
					Optional:    true,
					Description: "secrets for Rich Presence joining and spectating",
				},
				apislurp.Field{
					Name:    "instance",
					RawName: "instance?",
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
					Description: "whether or not the activity is an instanced game session",
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
					Description: "[activity flags](#DOCS_TOPICS_GATEWAY/activity-object-activity-flags) `OR`d together, describes what the payload includes",
				},
			},
		},
		apislurp.Structure{
			Name:    "Typing Start",
			RawName: "Typing Start Event Fields",
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
					Description: "id of the channel",
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
					Description: "id of the guild",
				},
				apislurp.Field{
					Name:    "user_id",
					RawName: "user_id",
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
					Description: "id of the user",
				},
				apislurp.Field{
					Name:    "timestamp",
					RawName: "timestamp",
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
					Description: "unix time (in seconds) of when the user started typing",
				},
				apislurp.Field{
					Name:    "member",
					RawName: "member?",
					Type: apislurp.Type{
						Names: []apislurp.TypeName{
							"guild member",
						},
						Array:     false,
						ArraySize: 0,
						Nullable:  false,
						RawText:   "[member](#DOCS_RESOURCES_GUILD/guild-member-object) object",
					},
					Optional:    true,
					Description: "the member who started typing if this happened in a guild",
				},
			},
		},
		apislurp.Structure{
			Name:    "Voice Server Update",
			RawName: "Voice Server Update Event Fields",
			Type:    "",
			Fields: []apislurp.Field{
				apislurp.Field{
					Name:    "token",
					RawName: "token",
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
					Description: "voice connection token",
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
					Description: "the guild this voice server update is for",
				},
				apislurp.Field{
					Name:    "endpoint",
					RawName: "endpoint",
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
					Description: "the voice server host",
				},
			},
		},
		apislurp.Structure{
			Name:    "Webhook Update",
			RawName: "Webhook Update Event Fields",
			Type:    "",
			Fields: []apislurp.Field{
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
					Description: "id of the guild",
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
					Description: "id of the channel",
				},
			},
		},
	},
}
