// This file is generated by gorazor 2.0.1
// DON'T modified manually
// Should edit source file and re-generate: tests/tpl/index.gohtml

package tpl

import (
	"io"
	"strings"

	"github.com/sipin/gorazor/gorazor"
	"github.com/sipin/gorazor/tests/data"
)

func Index(rows []data.BenchRow) string {
	var _b strings.Builder
	RenderIndex(&_b, rows)
	return _b.String()
}

func RenderIndex(_buffer io.StringWriter, rows []data.BenchRow) {
	_buffer.WriteString("<html>\n\t<head><title>test</title></head>\n\t<body>\n\t\t<ul>\n\t\t\n\t\t\t")
	for _, row := range rows {
		if row.Print {

			_buffer.WriteString("\n\t\t\t\t<li>ID=")
			_buffer.WriteString(gorazor.HTMLEscInt(row.ID))
			_buffer.WriteString(", Message=")
			_buffer.WriteString(gorazor.HTMLEscStr(row.Message))
			_buffer.WriteString("</li>\n\t\t\t\n\t\t\n\t\t")

		} else {
			_buffer.WriteString("\t\n\t\t\n\t\t\t")
		}
	}
	_buffer.WriteString("</ul>\n\t</body>\n</html>\n")

}
