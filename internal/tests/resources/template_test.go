// +build !discordtests

package apislurp_test

import (
	"github.com/diamondburned/apislurp"
	"github.com/diamondburned/apislurp/internal/tests/testutil"
	"path/filepath"
	"testing"
)

func TestTemplate(t *testing.T) {
	path := filepath.Join("resources", "Template.md")
	f := testutil.ReadFile(t, path)
	d := apislurp.Parse(f)
	testutil.Equal(t, d, templateData, "Template")
}

var templateData = apislurp.Document{
	Structures: []apislurp.Structure{
		apislurp.Structure{
			Name:    "Template",
			RawName: "Template Structure",
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
					Description: "the template code (unique ID)",
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
					Description: "template name",
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
					Description: "the description for the template",
				},
				apislurp.Field{
					Name:    "usage_count",
					RawName: "usage_count",
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
					Description: "number of times this template has been used",
				},
				apislurp.Field{
					Name:    "creator_id",
					RawName: "creator_id",
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
					Description: "the ID of the user who created the template",
				},
				apislurp.Field{
					Name:    "creator",
					RawName: "creator",
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
					Description: "the user who created the template",
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
					Description: "when this template was created",
				},
				apislurp.Field{
					Name:    "updated_at",
					RawName: "updated_at",
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
					Description: "when this template was last synced to the source guild",
				},
				apislurp.Field{
					Name:    "source_guild_id",
					RawName: "source_guild_id",
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
					Description: "the ID of the guild this template is based on",
				},
				apislurp.Field{
					Name:    "serialized_source_guild",
					RawName: "serialized_source_guild",
					Type: apislurp.Type{
						Names: []apislurp.TypeName{
							"guild",
						},
						Array:     false,
						ArraySize: 0,
						Nullable:  false,
						RawText:   "partial [guild](#DOCS_RESOURCES_GUILD/guild-object) object",
					},
					Optional:    false,
					Description: "the guild snapshot this template contains",
				},
				apislurp.Field{
					Name:    "is_dirty",
					RawName: "is_dirty",
					Type: apislurp.Type{
						Names: []apislurp.TypeName{
							"boolean",
						},
						Array:     false,
						ArraySize: 0,
						Nullable:  true,
						RawText:   "?boolean",
					},
					Optional:    false,
					Description: "whether the template has unsynced changes",
				},
			},
		},
	},
}
