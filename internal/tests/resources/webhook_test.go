// +build !nodiscordtests

package apislurp_test

import (
	"github.com/diamondburned/apislurp"
	"github.com/diamondburned/apislurp/internal/tests/testutil"
	"path/filepath"
	"testing"
)

func TestWebhook(t *testing.T) {
	path := filepath.Join("resources", "Webhook.md")
	f := testutil.ReadFile(t, path)
	d := apislurp.Parse(f)
	testutil.Equal(t, d, webhookData, "Webhook")
}

var webhookData = apislurp.Document{
	Structures: []apislurp.Structure{
		apislurp.Structure{
			Name:    "Webhook",
			RawName: "Webhook Structure",
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
					Description: "the id of the webhook",
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
					Description: "the [type](#DOCS_RESOURCES_WEBHOOK/webhook-object-webhook-types) of the webhook",
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
					Description: "the guild id this webhook is for",
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
					Description: "the channel id this webhook is for",
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
					Description: "the user this webhook was created by (not returned when getting a webhook with its token)",
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
						RawText:   "?string",
					},
					Optional:    false,
					Description: "the default name of the webhook",
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
					Description: "the default avatar of the webhook",
				},
				apislurp.Field{
					Name:    "token",
					RawName: "token?",
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
					Description: "the secure token of the webhook (returned for Incoming Webhooks)",
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
					Description: "the bot/OAuth2 application that created this webhook",
				},
			},
		},
	},
}
