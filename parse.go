package apislurp

import (
	"strings"

	"github.com/yuin/goldmark/ast"
	"github.com/yuin/goldmark/extension"
	extast "github.com/yuin/goldmark/extension/ast"
	"github.com/yuin/goldmark/parser"
	"github.com/yuin/goldmark/text"
	"github.com/yuin/goldmark/util"
)

var defaultParser = parser.NewParser(
	parser.WithInlineParsers(),
	parser.WithBlockParsers(
		util.Prioritized(parser.NewATXHeadingParser(), 100),
		util.Prioritized(parser.NewParagraphParser(), 200),
	),
	parser.WithParagraphTransformers(
		util.Prioritized(extension.NewTableParagraphTransformer(), 100),
	),
)

type walkState struct {
	document Document
	source   []byte

	table tableState
}

func (w *walkState) AddStructure(structure Structure) {
	if structure.Name != "" && len(structure.Fields) > 0 {
		w.document.Structures = append(w.document.Structures, structure)
	}
}

// tableState contains the current state of a table being parsed.
type tableState struct {
	structure Structure
	field     Field
	header    bool
	columns   []string // column names from table header
}

func (t *tableState) Reset(l6Header string) {
	*t = tableState{}
	t.structure = NewStructure(l6Header)
}

func (t *tableState) AddField(field Field) {
	t.structure.Fields = append(t.structure.Fields, field)
}

func (w *walkState) walker(n ast.Node, enter bool) (ast.WalkStatus, error) {
	switch n := n.(type) {
	case *ast.Heading:
		if n.Level == 6 && enter {
			w.AddStructure(w.table.structure)
			w.table.Reset(firstLine(n, w.source))
		}

	case *ast.Text:
		// fmt.Printf("Text(%q)\n", string(n.Segment.Value(w.source)))

	case *ast.String:
		// fmt.Printf("String(%q)\n", string(n.Value))

	case *ast.Paragraph:
		// fmt.Println()

	case *extast.Table:

	case *extast.TableHeader:
		// fmt.Printf("TableHeader(%t)\n", enter)
		w.table.header = enter

		if !enter && !stringsEq(w.table.columns, []string{"Field", "Type", "Description"}) {
			// log.Printf("Table with invalid colums %q found, skipping.", w.table.columns)
			return ast.WalkSkipChildren, nil
		}

	case *extast.TableRow:
		// fmt.Printf("TableRow(%t)\n", enter)
		if !enter && len(w.table.columns) >= 3 {
			w.table.AddField(NewField(
				w.table.columns[0],
				w.table.columns[2],
				NewType(w.table.columns[1], w.table.columns[2]),
			))
		} else {
			w.table.columns = w.table.columns[:0]
		}

	case *extast.TableCell:
		// fmt.Printf("TableCell(%t)(%q)\n", enter, joinLines(n, w.source))
		if enter {
			w.table.columns = append(w.table.columns, firstLine(n, w.source))
		}
	}

	return ast.WalkContinue, nil
}

type liner interface {
	Lines() *text.Segments
}

func firstLine(liner liner, src []byte) string {
	segments := liner.Lines()
	line := segments.At(0)
	return string(line.Value(src))
}

func stringsEq(strings, eq []string) bool {
	if len(strings) != len(eq) {
		return false
	}

	for i := range strings {
		if strings[i] != eq[i] {
			return false
		}
	}

	return true
}

func trimReplace(str *string, contains string) bool {
	switch {
	case strings.HasPrefix(*str, contains):
		*str = (*str)[len(contains):]
	case strings.HasSuffix(*str, contains):
		*str = (*str)[:len(*str)-len(contains)]
	default:
		return false
	}
	return true
}
