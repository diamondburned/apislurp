package apislurp_test

import (
	"github.com/diamondburned/apislurp"
	"github.com/diamondburned/apislurp/internal/tests/testutil"
	"path/filepath"
	"testing"
)

func TestGuild(t *testing.T) {
	path := filepath.Join("resources", "Guild.md")
	f := testutil.ReadFile(t, path)
	d := apislurp.Parse(f)
	testutil.Equal(t, d, guildData, "Guild")
}

var guildData = apislurp.Document{
	Structures: []apislurp.Structure{
		apislurp.Structure{
			Name:    "Guild",
			RawName: "Guild Structure",
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
					Description: "guild id",
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
					Description: "guild name (2-100 characters, excluding trailing and leading whitespace)",
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
					Description: "[icon hash](#DOCS_REFERENCE/image-formatting)",
				},
				apislurp.Field{
					Name:    "icon_hash",
					RawName: "icon_hash?",
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
					Description: "[icon hash](#DOCS_REFERENCE/image-formatting), returned when in the template object",
				},
				apislurp.Field{
					Name:    "splash",
					RawName: "splash",
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
					Description: "[splash hash](#DOCS_REFERENCE/image-formatting)",
				},
				apislurp.Field{
					Name:    "discovery_splash",
					RawName: "discovery_splash",
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
					Description: "[discovery splash hash](#DOCS_REFERENCE/image-formatting); only present for guilds with the \"DISCOVERABLE\" feature",
				},
				apislurp.Field{
					Name:    "owner? \\*\\*",
					RawName: "owner? \\*\\*",
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
					Description: "true if [the user](#DOCS_RESOURCES_USER/get-current-user-guilds) is the owner of the guild",
				},
				apislurp.Field{
					Name:    "owner_id",
					RawName: "owner_id",
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
					Description: "id of owner",
				},
				apislurp.Field{
					Name:    "permissions? \\*\\*",
					RawName: "permissions? \\*\\*",
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
					Description: "total permissions for [the user](#DOCS_RESOURCES_USER/get-current-user-guilds) in the guild (excludes overrides)",
				},
				apislurp.Field{
					Name:    "region",
					RawName: "region",
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
					Description: "[voice region](#DOCS_RESOURCES_VOICE/voice-region-object) id for the guild",
				},
				apislurp.Field{
					Name:    "afk_channel_id",
					RawName: "afk_channel_id",
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
					Description: "id of afk channel",
				},
				apislurp.Field{
					Name:    "afk_timeout",
					RawName: "afk_timeout",
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
					Description: "afk timeout in seconds",
				},
				apislurp.Field{
					Name:    "widget_enabled",
					RawName: "widget_enabled?",
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
					Description: "true if the server widget is enabled",
				},
				apislurp.Field{
					Name:    "widget_channel_id",
					RawName: "widget_channel_id?",
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
					Description: "the channel id that the widget will generate an invite to, or `null` if set to no invite",
				},
				apislurp.Field{
					Name:    "verification_level",
					RawName: "verification_level",
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
					Description: "[verification level](#DOCS_RESOURCES_GUILD/guild-object-verification-level) required for the guild",
				},
				apislurp.Field{
					Name:    "default_message_notifications",
					RawName: "default_message_notifications",
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
					Description: "default [message notifications level](#DOCS_RESOURCES_GUILD/guild-object-default-message-notification-level)",
				},
				apislurp.Field{
					Name:    "explicit_content_filter",
					RawName: "explicit_content_filter",
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
					Description: "[explicit content filter level](#DOCS_RESOURCES_GUILD/guild-object-explicit-content-filter-level)",
				},
				apislurp.Field{
					Name:    "roles",
					RawName: "roles",
					Type: apislurp.Type{
						Names: []apislurp.TypeName{
							"role",
						},
						Array:     false,
						ArraySize: 0,
						Nullable:  false,
						RawText:   "array of [role](#DOCS_TOPICS_PERMISSIONS/role-object) objects",
					},
					Optional:    false,
					Description: "roles in the guild",
				},
				apislurp.Field{
					Name:    "emojis",
					RawName: "emojis",
					Type: apislurp.Type{
						Names: []apislurp.TypeName{
							"emoji",
						},
						Array:     false,
						ArraySize: 0,
						Nullable:  false,
						RawText:   "array of [emoji](#DOCS_RESOURCES_EMOJI/emoji-object) objects",
					},
					Optional:    false,
					Description: "custom guild emojis",
				},
				apislurp.Field{
					Name:    "features",
					RawName: "features",
					Type: apislurp.Type{
						Names: []apislurp.TypeName{
							"guild feature",
						},
						Array:     false,
						ArraySize: 0,
						Nullable:  false,
						RawText:   "array of [guild feature](#DOCS_RESOURCES_GUILD/guild-object-guild-features) strings",
					},
					Optional:    false,
					Description: "enabled guild features",
				},
				apislurp.Field{
					Name:    "mfa_level",
					RawName: "mfa_level",
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
					Description: "required [MFA level](#DOCS_RESOURCES_GUILD/guild-object-mfa-level) for the guild",
				},
				apislurp.Field{
					Name:    "application_id",
					RawName: "application_id",
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
					Description: "application id of the guild creator if it is bot-created",
				},
				apislurp.Field{
					Name:    "system_channel_id",
					RawName: "system_channel_id",
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
					Description: "the id of the channel where guild notices such as welcome messages and boost events are posted",
				},
				apislurp.Field{
					Name:    "system_channel_flags",
					RawName: "system_channel_flags",
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
					Description: "[system channel flags](#DOCS_RESOURCES_GUILD/guild-object-system-channel-flags)",
				},
				apislurp.Field{
					Name:    "rules_channel_id",
					RawName: "rules_channel_id",
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
					Description: "the id of the channel where Community guilds can display rules and/or guidelines",
				},
				apislurp.Field{
					Name:    "joined_at? \\*",
					RawName: "joined_at? \\*",
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
					Description: "when this guild was joined at",
				},
				apislurp.Field{
					Name:    "large? \\*",
					RawName: "large? \\*",
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
					Description: "true if this is considered a large guild",
				},
				apislurp.Field{
					Name:    "unavailable? \\*",
					RawName: "unavailable? \\*",
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
					Description: "true if this guild is unavailable due to an outage",
				},
				apislurp.Field{
					Name:    "member_count? \\*",
					RawName: "member_count? \\*",
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
					Description: "total number of members in this guild",
				},
				apislurp.Field{
					Name:    "voice_states? \\*",
					RawName: "voice_states? \\*",
					Type: apislurp.Type{
						Names: []apislurp.TypeName{
							"voice state",
						},
						Array:     false,
						ArraySize: 0,
						Nullable:  false,
						RawText:   "array of partial [voice state](#DOCS_RESOURCES_VOICE/voice-state-object) objects",
					},
					Optional:    false,
					Description: "states of members currently in voice channels; lacks the `guild_id` key",
				},
				apislurp.Field{
					Name:    "members? \\*",
					RawName: "members? \\*",
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
					Description: "users in the guild",
				},
				apislurp.Field{
					Name:    "channels? \\*",
					RawName: "channels? \\*",
					Type: apislurp.Type{
						Names: []apislurp.TypeName{
							"channel",
						},
						Array:     false,
						ArraySize: 0,
						Nullable:  false,
						RawText:   "array of [channel](#DOCS_RESOURCES_CHANNEL/channel-object) objects",
					},
					Optional:    false,
					Description: "channels in the guild",
				},
				apislurp.Field{
					Name:    "presences? \\*",
					RawName: "presences? \\*",
					Type: apislurp.Type{
						Names: []apislurp.TypeName{
							"presence update",
						},
						Array:     false,
						ArraySize: 0,
						Nullable:  false,
						RawText:   "array of partial [presence update](#DOCS_TOPICS_GATEWAY/presence-update) objects",
					},
					Optional:    false,
					Description: "presences of the members in the guild, will only include non-offline members if the size is greater than `large threshold`",
				},
				apislurp.Field{
					Name:    "max_presences",
					RawName: "max_presences?",
					Type: apislurp.Type{
						Names: []apislurp.TypeName{
							"integer",
						},
						Array:     false,
						ArraySize: 0,
						Nullable:  true,
						RawText:   "?integer",
					},
					Optional:    true,
					Description: "the maximum number of presences for the guild (the default value, currently 25000, is in effect when `null` is returned)",
				},
				apislurp.Field{
					Name:    "max_members",
					RawName: "max_members?",
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
					Description: "the maximum number of members for the guild",
				},
				apislurp.Field{
					Name:    "vanity_url_code",
					RawName: "vanity_url_code",
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
					Description: "the vanity url code for the guild",
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
						Nullable:  true,
						RawText:   "?string",
					},
					Optional:    false,
					Description: "the description for the guild, if the guild is discoverable",
				},
				apislurp.Field{
					Name:    "banner",
					RawName: "banner",
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
					Description: "[banner hash](#DOCS_REFERENCE/image-formatting)",
				},
				apislurp.Field{
					Name:    "premium_tier",
					RawName: "premium_tier",
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
					Description: "[premium tier](#DOCS_RESOURCES_GUILD/guild-object-premium-tier) (Server Boost level)",
				},
				apislurp.Field{
					Name:    "premium_subscription_count",
					RawName: "premium_subscription_count?",
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
					Description: "the number of boosts this guild currently has",
				},
				apislurp.Field{
					Name:    "preferred_locale",
					RawName: "preferred_locale",
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
					Description: "the preferred locale of a Community guild; used in server discovery and notices from Discord; defaults to \"en-US\"",
				},
				apislurp.Field{
					Name:    "public_updates_channel_id",
					RawName: "public_updates_channel_id",
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
					Description: "the id of the channel where admins and moderators of Community guilds receive notices from Discord",
				},
				apislurp.Field{
					Name:    "max_video_channel_users",
					RawName: "max_video_channel_users?",
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
					Description: "the maximum amount of users in a video channel",
				},
				apislurp.Field{
					Name:    "approximate_member_count",
					RawName: "approximate_member_count?",
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
					Description: "approximate number of members in this guild, returned from the `GET /guild/<id>` endpoint when `with_counts` is `true`",
				},
				apislurp.Field{
					Name:    "approximate_presence_count",
					RawName: "approximate_presence_count?",
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
					Description: "approximate number of non-offline members in this guild, returned from the `GET /guild/<id>` endpoint when `with_counts` is `true`",
				},
			},
		},
		apislurp.Structure{
			Name:    "Guild Preview",
			RawName: "Guild Preview Structure",
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
					Description: "guild id",
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
					Description: "guild name (2-100 characters)",
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
					Description: "[icon hash](#DOCS_REFERENCE/image-formatting)",
				},
				apislurp.Field{
					Name:    "splash",
					RawName: "splash",
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
					Description: "[splash hash](#DOCS_REFERENCE/image-formatting)",
				},
				apislurp.Field{
					Name:    "discovery_splash",
					RawName: "discovery_splash",
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
					Description: "[discovery splash hash](#DOCS_REFERENCE/image-formatting)",
				},
				apislurp.Field{
					Name:    "emojis",
					RawName: "emojis",
					Type: apislurp.Type{
						Names: []apislurp.TypeName{
							"emoji",
						},
						Array:     false,
						ArraySize: 0,
						Nullable:  false,
						RawText:   "array of [emoji](#DOCS_RESOURCES_EMOJI/emoji-object) objects",
					},
					Optional:    false,
					Description: "custom guild emojis",
				},
				apislurp.Field{
					Name:    "features",
					RawName: "features",
					Type: apislurp.Type{
						Names: []apislurp.TypeName{
							"guild feature",
						},
						Array:     false,
						ArraySize: 0,
						Nullable:  false,
						RawText:   "array of [guild feature](#DOCS_RESOURCES_GUILD/guild-object-guild-features) strings",
					},
					Optional:    false,
					Description: "enabled guild features",
				},
				apislurp.Field{
					Name:    "approximate_member_count",
					RawName: "approximate_member_count",
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
					Description: "approximate number of members in this guild",
				},
				apislurp.Field{
					Name:    "approximate_presence_count",
					RawName: "approximate_presence_count",
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
					Description: "approximate number of online members in this guild",
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
						Nullable:  true,
						RawText:   "?string",
					},
					Optional:    false,
					Description: "the description for the guild",
				},
			},
		},
		apislurp.Structure{
			Name:    "Guild Widget",
			RawName: "Guild Widget Structure",
			Type:    "",
			Fields: []apislurp.Field{
				apislurp.Field{
					Name:    "enabled",
					RawName: "enabled",
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
					Description: "whether the widget is enabled",
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
					Description: "the widget channel id",
				},
			},
		},
		apislurp.Structure{
			Name:    "Guild Member",
			RawName: "Guild Member Structure",
			Type:    "",
			Fields: []apislurp.Field{
				apislurp.Field{
					Name:    "user",
					RawName: "user?",
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
					Description: "the user this guild member represents",
				},
				apislurp.Field{
					Name:    "nick",
					RawName: "nick",
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
					Description: "this users guild nickname",
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
					Description: "array of [role](#DOCS_TOPICS_PERMISSIONS/role-object) object ids",
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
					Description: "when the user started [boosting](https://support.discord.com/hc/en-us/articles/360028038352-Server-Boosting-) the guild",
				},
				apislurp.Field{
					Name:    "deaf",
					RawName: "deaf",
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
					Description: "whether the user is deafened in voice channels",
				},
				apislurp.Field{
					Name:    "mute",
					RawName: "mute",
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
					Description: "whether the user is muted in voice channels",
				},
			},
		},
		apislurp.Structure{
			Name:    "Integration",
			RawName: "Integration Structure",
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
					Description: "integration id",
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
					Description: "integration name",
				},
				apislurp.Field{
					Name:    "type",
					RawName: "type",
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
					Description: "integration type (twitch, youtube, or discord)",
				},
				apislurp.Field{
					Name:    "enabled",
					RawName: "enabled",
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
					Description: "is this integration enabled",
				},
				apislurp.Field{
					Name:    "syncing",
					RawName: "syncing",
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
					Description: "is this integration syncing",
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
					Description: "id that this integration uses for \"subscribers\", or the guild id for discord integrations",
				},
				apislurp.Field{
					Name:    "enable_emoticons",
					RawName: "enable_emoticons?",
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
					Description: "whether emoticons should be synced for this integration (twitch only currently)",
				},
				apislurp.Field{
					Name:    "expire_behavior",
					RawName: "expire_behavior",
					Type: apislurp.Type{
						Names: []apislurp.TypeName{
							"integration expire behavior",
						},
						Array:     false,
						ArraySize: 0,
						Nullable:  false,
						RawText:   "[integration expire behavior](#DOCS_RESOURCES_GUILD/integration-object-integration-expire-behaviors)",
					},
					Optional:    false,
					Description: "the behavior of expiring subscribers",
				},
				apislurp.Field{
					Name:    "expire_grace_period",
					RawName: "expire_grace_period",
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
					Description: "the grace period (in days) before expiring subscribers",
				},
				apislurp.Field{
					Name:    "user",
					RawName: "user?",
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
					Description: "user for this integration",
				},
				apislurp.Field{
					Name:    "account",
					RawName: "account",
					Type: apislurp.Type{
						Names: []apislurp.TypeName{
							"account",
						},
						Array:     false,
						ArraySize: 0,
						Nullable:  false,
						RawText:   "[account](#DOCS_RESOURCES_GUILD/integration-account-object) object",
					},
					Optional:    false,
					Description: "integration account information",
				},
				apislurp.Field{
					Name:    "synced_at",
					RawName: "synced_at",
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
					Description: "when this integration was last synced",
				},
				apislurp.Field{
					Name:    "subscriber_count",
					RawName: "subscriber_count",
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
					Description: "how many subscribers this integration has (0 for discord integrations)",
				},
				apislurp.Field{
					Name:    "revoked",
					RawName: "revoked",
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
					Description: "has this integration been revoked",
				},
				apislurp.Field{
					Name:    "application",
					RawName: "application?",
					Type: apislurp.Type{
						Names: []apislurp.TypeName{
							"application",
						},
						Array:     false,
						ArraySize: 0,
						Nullable:  false,
						RawText:   "[application](#DOCS_RESOURCES_GUILD/integration-account-object) object",
					},
					Optional:    true,
					Description: "The bot/OAuth2 application for discord integrations",
				},
			},
		},
		apislurp.Structure{
			Name:    "Integration Account",
			RawName: "Integration Account Structure",
			Type:    "",
			Fields: []apislurp.Field{
				apislurp.Field{
					Name:    "id",
					RawName: "id",
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
					Description: "id of the account",
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
					Description: "name of the account",
				},
			},
		},
		apislurp.Structure{
			Name:    "Integration Application",
			RawName: "Integration Application Structure",
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
					Description: "the id of the app",
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
					Description: "the name of the app",
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
					Description: "the [icon hash](#DOCS_REFERENCE/image-formatting) of the app",
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
					Description: "the description of the app",
				},
				apislurp.Field{
					Name:    "summary",
					RawName: "summary",
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
					Description: "the description of the app",
				},
				apislurp.Field{
					Name:    "bot",
					RawName: "bot?",
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
					Description: "the bot associated with this application",
				},
			},
		},
		apislurp.Structure{
			Name:    "Ban",
			RawName: "Ban Structure",
			Type:    "",
			Fields: []apislurp.Field{
				apislurp.Field{
					Name:    "reason",
					RawName: "reason",
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
					Description: "the reason for the ban",
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
					Description: "the banned user",
				},
			},
		},
	},
}
