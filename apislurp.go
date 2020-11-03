package apislurp

import (
	"regexp"
	"strings"

	"github.com/diamondburned/apislurp/internal/mdparse/shortmd"
	"github.com/yuin/goldmark/ast"
	"github.com/yuin/goldmark/text"
)

// Document represents a single Markdown page.
type Document struct {
	Structures []Structure
}

// Parse parses the given raw Markdown text into a Document.
func Parse(md []byte) Document {
	node := defaultParser.Parse(text.NewReader(md))
	walk := walkState{
		source: md,
	}

	ast.Walk(node, func(n ast.Node, enter bool) (ast.WalkStatus, error) {
		return walk.walker(n, enter)
	})

	// Try not to keep walk alive.
	doc := walk.document
	return doc
}

// Structure is a single structure, typically represented by a Markdown table.
type Structure struct {
	Name    string
	RawName string
	Type    StructureType
	Fields  []Field
}

type StructureType string

const (
	StructureObject     StructureType = "Structure"
	StructureEvent      StructureType = "Event Fields"
	StructureProperties StructureType = "Properties"
)

var StructureTypes = []StructureType{
	StructureObject,
	StructureEvent,
	StructureProperties,
}

func isValidStructure(name string) StructureType {
	for _, suffix := range StructureTypes {
		if strings.HasSuffix(name, " "+string(suffix)) {
			return suffix
		}
	}
	return ""
}

// NewStructure creates a new structure from the raw name taken from a table
// cell. If the structure type cannot be derived from the given name, then a
// zero-value Structure is returned.
func NewStructure(rawName string) Structure {
	structType := isValidStructure(rawName)
	if structType == "" {
		return Structure{}
	}

	return Structure{
		Name:    rawName[:len(rawName)-(len(structType)+1)], // cheap suffix trim
		RawName: rawName,
	}
}

// Field is a single field inside a structure, typically represented by a table
// row.
type Field struct {
	Name        string
	RawName     string
	Type        Type
	Optional    bool
	Description string
}

// NewField creates a new field from the given parameters taken from a table
// row.
func NewField(name, desc string, t Type) Field {
	fname := name
	optional := trimReplace(&fname, "?")

	return Field{
		Name:        fname,
		RawName:     name,
		Type:        t,
		Optional:    optional,
		Description: desc,
	}
}

// TypeName represents a short, normalized type name. These names are normalized
// on a best effort basis; they may not cover every single edge case in the
// documentation. As such, one shouldn't rely on the uniqueness of those names
// too much and instead handle unknown type names (edge cases) appropriately.
type TypeName string

// Primitives.
const (
	TypeSnowflake TypeName = "snowflake"
	TypeString    TypeName = "string"
	TypeInteger   TypeName = "integer" // also "int"
	TypeBoolean   TypeName = "boolean"
	TypeArray     TypeName = "array" // use ParseTypeName to handle this.
	TypeMixed     TypeName = "mixed"
	TypeObject    TypeName = "object" // this is a thing and I don't know why.
	TypeImageData TypeName = "image data"
	TypeTimestamp TypeName = "ISO8601 timestamp"
)

// Object Types.
const (
	TypeUser               TypeName = "user"
	TypeRole               TypeName = "role"
	TypeEmoji              TypeName = "emoji"
	TypeParty              TypeName = "party"
	TypeAssets             TypeName = "assets"
	TypeAccount            TypeName = "account"
	TypeWebhook            TypeName = "webhook"
	TypeSecrets            TypeName = "secrets"
	TypeActivity           TypeName = "activity"
	TypeReaction           TypeName = "reaction"
	TypeOverwrite          TypeName = "overwrite"
	TypeAttachment         TypeName = "attachment"
	TypeTimestamps         TypeName = "timestamps"
	TypeApplication        TypeName = "application"
	TypeVoiceState         TypeName = "voice state"
	TypeClientStatus       TypeName = "client_status"
	TypeUpdateStatus       TypeName = "update status"
	TypeAllowedMentionType TypeName = "allowed mention type"

	TypePresence       TypeName = "presence"
	TypePresenceUpdate TypeName = "presence update"

	TypeIntegration               TypeName = "integration"
	TypeIntegrationExpireBehavior TypeName = "integration expire behavior"

	TypeChannel        TypeName = "channel"
	TypeChannelMention TypeName = "channel mention"

	TypeGuild            TypeName = "guild"
	TypeGuildMember      TypeName = "guild member" // also "member"
	TypeGuildFeature     TypeName = "guild feature"
	TypeUnavailableGuild TypeName = "unavailable guild"

	TypeMessage            TypeName = "message"
	TypeMessageActivity    TypeName = "message activity"
	TypeMessageReference   TypeName = "message reference"
	TypeMessageApplication TypeName = "message application"

	TypeEmbed          TypeName = "embed"
	TypeEmbedFooter    TypeName = "embed footer"
	TypeEmbedImage     TypeName = "embed image"
	TypeEmbedThumbnail TypeName = "embed thumbnail"
	TypeEmbedVideo     TypeName = "embed video"
	TypeEmbedProvider  TypeName = "embed provider"
	TypeEmbedAuthor    TypeName = "embed author"
	TypeEmbedField     TypeName = "embed field"

	TypeAuditLog       TypeName = "audit log"
	TypeAuditLogEvent  TypeName = "audit log event"
	TypeAuditLogEntry  TypeName = "audit log entry"
	TypeAuditLogChange TypeName = "audit log change"
	TypeAuditEntryInfo TypeName = "audit entry info"
)

// ParseTypeName parses an individual type name. It does NOT split on "or"
// words. An empty TypeName is returned if parsing fails.
func ParseTypeName(fname string) (tname TypeName, array bool, arraySize int) {
	if linkText := shortmd.LinkText(fname); linkText != "" {
		fname = linkText
	}

	fname = strings.TrimSpace(fname)
	fname = strings.TrimPrefix(fname, "optional ")

	// Derive the array for each type name instead.
	array = (array || (trimReplace(&fname, "array of") || trimReplace(&fname, "list of")))

	// Brilliant.
	arraySize, fname = englishNumber(fname)
	array = (array || arraySize > 1)

	// Sometimes, I question myself if the abstraction I'm writing is a
	// mistake.
	if array {
		fname = strings.TrimSuffix(fname, "s")
	}

	// Handle footnotes not making sense.
	trimReplace(&fname, ` \\*`)

	tname = TypeName(fname)

	// Handle additional special cases.
	switch fname {
	case "int":
		tname = TypeInteger
	case "member":
		tname = TypeGuildMember
	}

	return
}

// Type represents a type for each field. Types may be duplicated or slightly
// different from other types elsewhere in other fields and structures.
type Type struct {
	Names     []TypeName // or types
	Array     bool
	ArraySize int // 0 == slice/not an array
	Nullable  bool
	RawText   string
}

// parenthesesRegex is the regex to trim all parentheses texts and the prefixed
// space.
var parenthesesRegex = regexp.MustCompile(` \(.*\) ?`)

// NewType creates a new normalized Type instance from the given raw type name.
// If description is not empty, then the function will try and parse the new
// type from it if the type is "array".
func NewType(typeField, description string) Type {
	fname := typeField

	nullable := trimReplace(&fname, "?")
	arraySize := 0
	array := false

	fname = parenthesesRegex.ReplaceAllLiteralString(fname, "")
	fname = strings.TrimSpace(fname)

	// Handle "or" types.
	fnames := strings.Split(fname, " or ")
	tNames := make([]TypeName, len(fnames))

	// Normalize and handle more special cases.
	for i, fname := range fnames {
		typeName, isArray, size := ParseTypeName(fname)
		array = (array || isArray)

		// Human-written documentations are always error-prone. Unfortunately.
		if typeName == TypeArray && description != "" {
			typeName, _, size = ParseTypeName(description)
		}

		arraySize += size

		// Ensure no duplicate types are added.
		if !typeNameContains(tNames, typeName) {
			tNames[i] = typeName
		}
	}

	return Type{
		Names:     tNames,
		Array:     array,
		ArraySize: arraySize,
		Nullable:  nullable,
		RawText:   typeField,
	}
}

func typeNameContains(typeNames []TypeName, contains TypeName) bool {
	for _, prev := range typeNames {
		if prev == contains {
			return true
		}
	}
	return false
}

// I kid you not.
func englishNumber(words string) (number int, word string) {
	splitWords := strings.SplitN(words, " ", 2)
	if len(splitWords) != 2 {
		return 0, words
	}

	word = splitWords[1]

	switch splitWords[0] {
	case "one":
		number = 1
	case "two":
		number = 2
	case "three":
		number = 3
	default:
		word = words
	}

	return
}
