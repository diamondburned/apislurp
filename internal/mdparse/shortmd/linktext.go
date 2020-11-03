package shortmd

import (
	"github.com/yuin/goldmark/ast"
	"github.com/yuin/goldmark/parser"
	"github.com/yuin/goldmark/text"
	"github.com/yuin/goldmark/util"
)

var linkParser = parser.NewParser(
	parser.WithInlineParsers(
		util.Prioritized(parser.NewLinkParser(), 100),
	),
	parser.WithBlockParsers(
		util.Prioritized(parser.NewParagraphParser(), 200),
	),
	parser.WithParagraphTransformers(),
)

// LinkText extracts the first text that is in a hyperlink and return that.
func LinkText(md string) string {
	src := []byte(md)
	node := linkParser.Parse(text.NewReader(src))
	link := ""

	// states
	insideLink := false

	ast.Walk(node, func(n ast.Node, enter bool) (ast.WalkStatus, error) {
		switch n := n.(type) {
		case *ast.Link:
			insideLink = (enter && len(n.Destination) > 0)

		case *ast.Text:
			if insideLink && enter {
				link = string(n.Segment.Value(src))

				if link != "" {
					return ast.WalkStop, nil
				}
			}
		}

		return ast.WalkContinue, nil
	})

	return link
}
