package apislurp_test

import (
	"github.com/diamondburned/apislurp"
	"github.com/diamondburned/apislurp/internal/tests/testutil"
	"path/filepath"
	"testing"
)

func TestInvite(t *testing.T) {
	path := filepath.Join("resources", "Invite.md")
	f := testutil.ReadFile(t, path)
	d := apislurp.Parse(f)
	testutil.Equal(t, d, inviteData, "Invite")
}

var inviteData = apislurp.Document{
	Structures: []apislurp.Structure{
		apislurp.Structure{
			Name:    "Invite",
			RawName: "Invite Structure",
			Type:    "",
			Fields: []apislurp.Field{
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
					Description: "the invite code (unique ID)",
				},
				apislurp.Field{
					Name:    "guild",
					RawName: "guild?",
					Type: apislurp.Type{
						Names: []apislurp.TypeName{
							"guild",
						},
						Array:     false,
						ArraySize: 0,
						Nullable:  false,
						RawText:   "partial [guild](#DOCS_RESOURCES_GUILD/guild-object) object",
					},
					Optional:    true,
					Description: "the guild this invite is for",
				},
				apislurp.Field{
					Name:    "channel",
					RawName: "channel",
					Type: apislurp.Type{
						Names: []apislurp.TypeName{
							"channel",
						},
						Array:     false,
						ArraySize: 0,
						Nullable:  false,
						RawText:   "partial [channel](#DOCS_RESOURCES_CHANNEL/channel-object) object",
					},
					Optional:    false,
					Description: "the channel this invite is for",
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
					Description: "the user who created the invite",
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
					Description: "approximate count of online members (only present when target_user is set)",
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
					Description: "approximate count of total members",
				},
			},
		},
		apislurp.Structure{
			Name:    "Invite Metadata",
			RawName: "Invite Metadata Structure",
			Type:    "",
			Fields: []apislurp.Field{
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
					Description: "number of times this invite has been used",
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
					Description: "max number of times this invite can be used",
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
					Description: "duration (in seconds) after which the invite expires",
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
					Description: "whether this invite only grants temporary membership",
				},
				apislurp.Field{
					Name:    "created_at",
					RawName: "created_at",
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
					Description: "when this invite was created",
				},
			},
		},
	},
}
