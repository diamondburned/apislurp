// +build !discordtests

package apislurp_test

import (
	"github.com/diamondburned/apislurp"
	"github.com/diamondburned/apislurp/internal/tests/testutil"
	"path/filepath"
	"testing"
)

func TestEmoji(t *testing.T) {
	path := filepath.Join("resources", "Emoji.md")
	f := testutil.ReadFile(t, path)
	d := apislurp.Parse(f)
	testutil.Equal(t, d, emojiData, "Emoji")
}

var emojiData = apislurp.Document{
	Structures: []apislurp.Structure{
		apislurp.Structure{
			Name:    "Emoji",
			RawName: "Emoji Structure",
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
						Nullable:  true,
						RawText:   "?snowflake",
					},
					Optional:    false,
					Description: "[emoji id](#DOCS_REFERENCE/image-formatting)",
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
						Nullable:  true,
						RawText:   "?string (can be null only in reaction emoji objects)",
					},
					Optional:    false,
					Description: "emoji name",
				},
				apislurp.Field{
					Name:    "roles",
					RawName: "roles?",
					Type: apislurp.Type{
						Names: []apislurp.TypeName{
							"role",
						},
						Array:     false,
						ArraySize: 0,
						Nullable:  false,
						RawText:   "array of [role](#DOCS_TOPICS_PERMISSIONS/role-object) object ids",
					},
					Optional:    true,
					Description: "roles this emoji is whitelisted to",
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
					Description: "user that created this emoji",
				},
				apislurp.Field{
					Name:    "require_colons",
					RawName: "require_colons?",
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
					Description: "whether this emoji must be wrapped in colons",
				},
				apislurp.Field{
					Name:    "managed",
					RawName: "managed?",
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
					Description: "whether this emoji is managed",
				},
				apislurp.Field{
					Name:    "animated",
					RawName: "animated?",
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
					Description: "whether this emoji is animated",
				},
				apislurp.Field{
					Name:    "available",
					RawName: "available?",
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
					Description: "whether this emoji can be used, may be false due to loss of Server Boosts",
				},
			},
		},
	},
}
