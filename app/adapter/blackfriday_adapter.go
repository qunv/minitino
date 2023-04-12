package adapter

import (
	"bytes"
	"github.com/russross/blackfriday"
	"strings"
)

type blackFridayAdapter struct {
	renderer blackfriday.Renderer
}

func newBlackFridayAdapter(renderer blackfriday.Renderer) blackfriday.Renderer {
	return blackFridayAdapter{
		renderer: renderer,
	}
}

func (b blackFridayAdapter) BlockCode(out *bytes.Buffer, text []byte, info string) {
	doubleSpace(out)

	endOfLang := strings.IndexAny(info, "\t ")
	if endOfLang < 0 {
		endOfLang = len(info)
	}
	lang := info[:endOfLang]
	if len(lang) == 0 || lang == "." {
		out.WriteString("<pre><code>")
	} else if lang == "mermaid" {
		out.WriteString("<div class=\"mermaid\">")
		attrEscape(out, text)
		out.WriteString("</div>\n")
		return
	} else {
		out.WriteString("<pre><code class=\"language-")
		attrEscape(out, []byte(lang))
		out.WriteString("\">")
	}
	attrEscape(out, text)
	out.WriteString("</code></pre>\n")
}

func (b blackFridayAdapter) BlockQuote(out *bytes.Buffer, text []byte) {
	b.renderer.BlockQuote(out, text)
}

func (b blackFridayAdapter) BlockHtml(out *bytes.Buffer, text []byte) {
	b.renderer.BlockHtml(out, text)
}

func (b blackFridayAdapter) Header(out *bytes.Buffer, text func() bool, level int, id string) {
	b.renderer.Header(out, text, level, id)
}

func (b blackFridayAdapter) HRule(out *bytes.Buffer) {
	b.renderer.HRule(out)
}

func (b blackFridayAdapter) List(out *bytes.Buffer, text func() bool, flags int) {
	b.renderer.List(out, text, flags)
}

func (b blackFridayAdapter) ListItem(out *bytes.Buffer, text []byte, flags int) {
	b.renderer.ListItem(out, text, flags)
}

func (b blackFridayAdapter) Paragraph(out *bytes.Buffer, text func() bool) {
	b.renderer.Paragraph(out, text)
}

func (b blackFridayAdapter) Table(out *bytes.Buffer, header []byte, body []byte, columnData []int) {
	b.renderer.Table(out, header, body, columnData)
}

func (b blackFridayAdapter) TableRow(out *bytes.Buffer, text []byte) {
	b.renderer.TableRow(out, text)
}

func (b blackFridayAdapter) TableHeaderCell(out *bytes.Buffer, text []byte, flags int) {
	b.renderer.TableHeaderCell(out, text, flags)
}

func (b blackFridayAdapter) TableCell(out *bytes.Buffer, text []byte, flags int) {
	b.renderer.TableCell(out, text, flags)
}

func (b blackFridayAdapter) Footnotes(out *bytes.Buffer, text func() bool) {
	b.renderer.Footnotes(out, text)
}

func (b blackFridayAdapter) FootnoteItem(out *bytes.Buffer, name, text []byte, flags int) {
	b.renderer.FootnoteItem(out, name, text, flags)
}

func (b blackFridayAdapter) TitleBlock(out *bytes.Buffer, text []byte) {
	b.renderer.TitleBlock(out, text)
}

func (b blackFridayAdapter) AutoLink(out *bytes.Buffer, link []byte, kind int) {
	b.renderer.AutoLink(out, link, kind)
}

func (b blackFridayAdapter) CodeSpan(out *bytes.Buffer, text []byte) {
	b.renderer.CodeSpan(out, text)
}

func (b blackFridayAdapter) DoubleEmphasis(out *bytes.Buffer, text []byte) {
	b.renderer.DoubleEmphasis(out, text)
}

func (b blackFridayAdapter) Emphasis(out *bytes.Buffer, text []byte) {
	b.renderer.Emphasis(out, text)
}

func (b blackFridayAdapter) Image(out *bytes.Buffer, link []byte, title []byte, alt []byte) {
	b.renderer.Image(out, link, title, alt)
}

func (b blackFridayAdapter) LineBreak(out *bytes.Buffer) {
	b.renderer.LineBreak(out)
}

func (b blackFridayAdapter) Link(out *bytes.Buffer, link []byte, title []byte, content []byte) {
	b.renderer.Link(out, link, title, content)
}

func (b blackFridayAdapter) RawHtmlTag(out *bytes.Buffer, tag []byte) {
	b.renderer.RawHtmlTag(out, tag)
}

func (b blackFridayAdapter) TripleEmphasis(out *bytes.Buffer, text []byte) {
	b.renderer.TripleEmphasis(out, text)
}

func (b blackFridayAdapter) StrikeThrough(out *bytes.Buffer, text []byte) {
	b.renderer.StrikeThrough(out, text)
}

func (b blackFridayAdapter) FootnoteRef(out *bytes.Buffer, ref []byte, id int) {
	b.renderer.FootnoteRef(out, ref, id)
}

func (b blackFridayAdapter) Entity(out *bytes.Buffer, entity []byte) {
	b.renderer.Entity(out, entity)
}

func (b blackFridayAdapter) NormalText(out *bytes.Buffer, text []byte) {
	b.renderer.NormalText(out, text)
}

func (b blackFridayAdapter) DocumentHeader(out *bytes.Buffer) {
	b.renderer.DocumentHeader(out)
}

func (b blackFridayAdapter) DocumentFooter(out *bytes.Buffer) {
	b.renderer.DocumentFooter(out)
}

func (b blackFridayAdapter) GetFlags() int {
	return b.renderer.GetFlags()
}

func doubleSpace(out *bytes.Buffer) {
	if out.Len() > 0 {
		out.WriteByte('\n')
	}
}

// Using if statements is a bit faster than a switch statement. As the compiler
// improves, this should be unnecessary this is only worthwhile because
// attrEscape is the single largest CPU user in normal use.
// Also tried using map, but that gave a ~3x slowdown.
func escapeSingleChar(char byte) (string, bool) {
	if char == '"' {
		return "&quot;", true
	}
	if char == '&' {
		return "&amp;", true
	}
	if char == '<' {
		return "&lt;", true
	}
	if char == '>' {
		return "&gt;", true
	}
	return "", false
}

func attrEscape(out *bytes.Buffer, src []byte) {
	org := 0
	for i, ch := range src {
		if entity, ok := escapeSingleChar(ch); ok {
			if i > org {
				// copy all the normal characters since the last escape
				out.Write(src[org:i])
			}
			org = i + 1
			out.WriteString(entity)
		}
	}
	if org < len(src) {
		out.Write(src[org:])
	}
}
