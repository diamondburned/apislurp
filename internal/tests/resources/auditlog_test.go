// +build !discordtests

package apislurp_test

import (
	"github.com/diamondburned/apislurp"
	"github.com/diamondburned/apislurp/internal/tests/testutil"
	"path/filepath"
	"testing"
)

func TestAuditLog(t *testing.T) {
	path := filepath.Join("resources", "Audit_Log.md")
	f := testutil.ReadFile(t, path)
	d := apislurp.Parse(f)
	testutil.Equal(t, d, auditlogData, "Audit_Log")
}

var auditlogData = apislurp.Document{
	Structures: []apislurp.Structure{
		apislurp.Structure{
			Name:    "Audit Log",
			RawName: "Audit Log Structure",
			Type:    "",
			Fields: []apislurp.Field{
				apislurp.Field{
					Name:    "webhooks",
					RawName: "webhooks",
					Type: apislurp.Type{
						Names: []apislurp.TypeName{
							"webhook",
						},
						Array:     false,
						ArraySize: 0,
						Nullable:  false,
						RawText:   "array of [webhook](#DOCS_RESOURCES_WEBHOOK/webhook-object) objects",
					},
					Optional:    false,
					Description: "list of webhooks found in the audit log",
				},
				apislurp.Field{
					Name:    "users",
					RawName: "users",
					Type: apislurp.Type{
						Names: []apislurp.TypeName{
							"user",
						},
						Array:     false,
						ArraySize: 0,
						Nullable:  false,
						RawText:   "array of [user](#DOCS_RESOURCES_USER/user-object) objects",
					},
					Optional:    false,
					Description: "list of users found in the audit log",
				},
				apislurp.Field{
					Name:    "audit_log_entries",
					RawName: "audit_log_entries",
					Type: apislurp.Type{
						Names: []apislurp.TypeName{
							"audit log entry",
						},
						Array:     false,
						ArraySize: 0,
						Nullable:  false,
						RawText:   "array of [audit log entry](#DOCS_RESOURCES_AUDIT_LOG/audit-log-entry-object) objects",
					},
					Optional:    false,
					Description: "list of audit log entries",
				},
				apislurp.Field{
					Name:    "integrations",
					RawName: "integrations",
					Type: apislurp.Type{
						Names: []apislurp.TypeName{
							"integration",
						},
						Array:     false,
						ArraySize: 0,
						Nullable:  false,
						RawText:   "array of partial [integration](#DOCS_RESOURCES_GUILD/integration-object) objects",
					},
					Optional:    false,
					Description: "list of partial integration objects",
				},
			},
		},
		apislurp.Structure{
			Name:    "Audit Log Entry",
			RawName: "Audit Log Entry Structure",
			Type:    "",
			Fields: []apislurp.Field{
				apislurp.Field{
					Name:    "target_id",
					RawName: "target_id",
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
					Description: "id of the affected entity (webhook, user, role, etc.)",
				},
				apislurp.Field{
					Name:    "changes",
					RawName: "changes?",
					Type: apislurp.Type{
						Names: []apislurp.TypeName{
							"audit log change",
						},
						Array:     false,
						ArraySize: 0,
						Nullable:  false,
						RawText:   "array of [audit log change](#DOCS_RESOURCES_AUDIT_LOG/audit-log-change-object) objects",
					},
					Optional:    true,
					Description: "changes made to the target_id",
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
					Description: "the user who made the changes",
				},
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
					Description: "id of the entry",
				},
				apislurp.Field{
					Name:    "action_type",
					RawName: "action_type",
					Type: apislurp.Type{
						Names: []apislurp.TypeName{
							"audit log event",
						},
						Array:     false,
						ArraySize: 0,
						Nullable:  false,
						RawText:   "[audit log event](#DOCS_RESOURCES_AUDIT_LOG/audit-log-entry-object-audit-log-events)",
					},
					Optional:    false,
					Description: "type of action that occurred",
				},
				apislurp.Field{
					Name:    "options",
					RawName: "options?",
					Type: apislurp.Type{
						Names: []apislurp.TypeName{
							"audit entry info",
						},
						Array:     false,
						ArraySize: 0,
						Nullable:  false,
						RawText:   "[optional audit entry info](#DOCS_RESOURCES_AUDIT_LOG/audit-log-entry-object-optional-audit-entry-info)",
					},
					Optional:    true,
					Description: "additional info for certain action types",
				},
				apislurp.Field{
					Name:    "reason",
					RawName: "reason?",
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
					Description: "the reason for the change (0-512 characters)",
				},
			},
		},
		apislurp.Structure{
			Name:    "Audit Log Change",
			RawName: "Audit Log Change Structure",
			Type:    "",
			Fields: []apislurp.Field{
				apislurp.Field{
					Name:    "new_value",
					RawName: "new_value?",
					Type: apislurp.Type{
						Names: []apislurp.TypeName{
							"mixed",
						},
						Array:     false,
						ArraySize: 0,
						Nullable:  false,
						RawText:   "[mixed](#DOCS_RESOURCES_AUDIT_LOG/audit-log-change-object-audit-log-change-key)",
					},
					Optional:    true,
					Description: "new value of the key",
				},
				apislurp.Field{
					Name:    "old_value",
					RawName: "old_value?",
					Type: apislurp.Type{
						Names: []apislurp.TypeName{
							"mixed",
						},
						Array:     false,
						ArraySize: 0,
						Nullable:  false,
						RawText:   "[mixed](#DOCS_RESOURCES_AUDIT_LOG/audit-log-change-object-audit-log-change-key)",
					},
					Optional:    true,
					Description: "old value of the key",
				},
				apislurp.Field{
					Name:    "key",
					RawName: "key",
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
					Description: "name of audit log [change key](#DOCS_RESOURCES_AUDIT_LOG/audit-log-change-object-audit-log-change-key)",
				},
			},
		},
	},
}
