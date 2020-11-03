package apislurp_test

import (
	"github.com/diamondburned/apislurp"
	"github.com/diamondburned/apislurp/internal/tests/testutil"
	"path/filepath"
	"testing"
)

func TestUser(t *testing.T) {
	path := filepath.Join("resources", "User.md")
	f := testutil.ReadFile(t, path)
	d := apislurp.Parse(f)
	testutil.Equal(t, d, userData, "User")
}

var userData = apislurp.Document{
	Structures: []apislurp.Structure{
		apislurp.Structure{
			Name:    "User",
			RawName: "User Structure",
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
					Description: "the user's id",
				},
				apislurp.Field{
					Name:    "username",
					RawName: "username",
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
					Description: "the user's username, not unique across the platform",
				},
				apislurp.Field{
					Name:    "discriminator",
					RawName: "discriminator",
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
					Description: "the user's 4-digit discord-tag",
				},
				apislurp.Field{
					Name:    "avatar",
					RawName: "avatar",
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
					Description: "the user's [avatar hash](#DOCS_REFERENCE/image-formatting)",
				},
				apislurp.Field{
					Name:    "bot",
					RawName: "bot?",
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
					Description: "whether the user belongs to an OAuth2 application",
				},
				apislurp.Field{
					Name:    "system",
					RawName: "system?",
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
					Description: "whether the user is an Official Discord System user (part of the urgent message system)",
				},
				apislurp.Field{
					Name:    "mfa_enabled",
					RawName: "mfa_enabled?",
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
					Description: "whether the user has two factor enabled on their account",
				},
				apislurp.Field{
					Name:    "locale",
					RawName: "locale?",
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
					Description: "the user's chosen language option",
				},
				apislurp.Field{
					Name:    "verified",
					RawName: "verified?",
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
					Description: "whether the email on this account has been verified",
				},
				apislurp.Field{
					Name:    "email",
					RawName: "email?",
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
					Description: "the user's email",
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
					Description: "the [flags](#DOCS_RESOURCES_USER/user-object-user-flags) on a user's account",
				},
				apislurp.Field{
					Name:    "premium_type",
					RawName: "premium_type?",
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
					Description: "the [type of Nitro subscription](#DOCS_RESOURCES_USER/user-object-premium-types) on a user's account",
				},
				apislurp.Field{
					Name:    "public_flags",
					RawName: "public_flags?",
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
					Description: "the public [flags](#DOCS_RESOURCES_USER/user-object-user-flags) on a user's account",
				},
			},
		},
		apislurp.Structure{
			Name:    "Connection",
			RawName: "Connection Structure",
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
					Description: "id of the connection account",
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
					Description: "the username of the connection account",
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
					Description: "the service of the connection (twitch, youtube)",
				},
				apislurp.Field{
					Name:    "revoked",
					RawName: "revoked?",
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
					Description: "whether the connection is revoked",
				},
				apislurp.Field{
					Name:    "integrations",
					RawName: "integrations?",
					Type: apislurp.Type{
						Names: []apislurp.TypeName{
							"server integrations",
						},
						Array:     false,
						ArraySize: 0,
						Nullable:  false,
						RawText:   "array",
					},
					Optional:    true,
					Description: "an array of partial [server integrations](#DOCS_RESOURCES_GUILD/integration-object)",
				},
				apislurp.Field{
					Name:    "verified",
					RawName: "verified",
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
					Description: "whether the connection is verified",
				},
				apislurp.Field{
					Name:    "friend_sync",
					RawName: "friend_sync",
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
					Description: "whether friend sync is enabled for this connection",
				},
				apislurp.Field{
					Name:    "show_activity",
					RawName: "show_activity",
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
					Description: "whether activities related to this connection will be shown in presence updates",
				},
				apislurp.Field{
					Name:    "visibility",
					RawName: "visibility",
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
					Description: "[visibility](#DOCS_RESOURCES_USER/user-object-visibility-types) of this connection",
				},
			},
		},
	},
}
