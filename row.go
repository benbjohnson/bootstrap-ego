package bootstrap

import (
	"context"
	"html"
	"io"
)

type Row struct {
	ID    string
	Class string
	Yield func()
}

func (c *Row) Render(ctx context.Context, w io.Writer) {
	io.WriteString(w, `<div`)
	writeAttr(w, "id", c.ID)
	c.renderClass(ctx, w)
	io.WriteString(w, `>`)

	if c.Yield != nil {
		c.Yield()
	}

	io.WriteString(w, `</div>`)
}

func (c *Row) renderClass(ctx context.Context, w io.Writer) {
	io.WriteString(w, ` class="row`)
	if s := c.Class; s != "" {
		io.WriteString(w, " ")
		io.WriteString(w, html.EscapeString(s))
	}
	io.WriteString(w, `"`)
}
