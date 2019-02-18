package document

import (
	"fmt"
	"html"
	"strings"
)

type htmlFormatVisitor interface {
	visitParagraph(paragraph) string
	visitImage(image) string
	visitDocument(Document) string
}

type htmlFormatAcceptor interface {
	acceptVisitor(htmlFormatVisitor) string
}

type htmlFormatVisitorImpl struct{}

func (htmlFormatVisitorImpl) visitParagraph(p paragraph) string {
	return fmt.Sprintf("<p>%s</p>", html.EscapeString(p.getText()))
}

func (htmlFormatVisitorImpl) visitImage(i image) string {
	w, h := i.getSize()
	return fmt.Sprintf(`<img src="%s" width="%d" height="%d" />`, html.EscapeString(i.getPath()), w, h)
}

func (v htmlFormatVisitorImpl) visitDocument(d Document) string {
	const htmlTemplate = `<!DOCTYPE html>
<html>
<head>
<meta charset="utf-8">
<title>%s</title>
</head>
<body>%s
</body>
</html>
`

	var builder strings.Builder
	for _, elem := range d.getElements() {
		builder.WriteString(fmt.Sprintf("\n%s", elem.(element).acceptVisitor(&v)))
	}

	return fmt.Sprintf(htmlTemplate, html.EscapeString(d.GetTitle()), builder.String())
}
