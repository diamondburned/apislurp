package apislurp_test

import (
	"github.com/diamondburned/apislurp"
	"github.com/diamondburned/apislurp/internal/tests/testutil"
	"path/filepath"
	"testing"
)

func TestVoice(t *testing.T) {
	path := filepath.Join("resources", "Voice.md")
	f := testutil.ReadFile(t, path)
	d := apislurp.Parse(f)
	testutil.Equal(t, d, voiceData, "Voice")
}

var voiceData = apislurp.Document{
	Structures: []apislurp.Structure{
		apislurp.Structure{
			Name:    "Voice State",
			RawName: "Voice State Structure",
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
					Description: "the guild id this voice state is for",
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
					Description: "the channel id this user is connected to",
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
					Description: "the user id this voice state is for",
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
						RawText:   "[guild member](#DOCS_RESOURCES_GUILD/guild-member-object) object",
					},
					Optional:    true,
					Description: "the guild member this voice state is for",
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
					Description: "the session id for this voice state",
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
					Description: "whether this user is deafened by the server",
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
					Description: "whether this user is muted by the server",
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
					Description: "whether this user is locally deafened",
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
					Description: "whether this user is locally muted",
				},
				apislurp.Field{
					Name:    "self_stream",
					RawName: "self_stream?",
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
					Description: "whether this user is streaming using \"Go Live\"",
				},
				apislurp.Field{
					Name:    "self_video",
					RawName: "self_video",
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
					Description: "whether this user's camera is enabled",
				},
				apislurp.Field{
					Name:    "suppress",
					RawName: "suppress",
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
					Description: "whether this user is muted by the current user",
				},
			},
		},
	},
}
