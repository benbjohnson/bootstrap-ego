package bootstrap

import (
	"context"
	"io"
)

type Row struct {
	ID    string
	Class string
	Yield func()
}

func (c *Row) Render(ctx context.Context, w io.Writer) {
	io.WriteString(w, `<div`)
	appendAttr(w, "id", c.ID)
	c.renderClass(ctx, w)
	io.WriteString(w, `>`)

	if c.Yield != nil {
		c.Yield()
	}

	io.WriteString(w, `</div>`)
}

func (c *Row) renderClass(ctx context.Context, w io.Writer) {
	io.WriteString(w, ` class="row`)
	appendHTML(w, c.Class)
	io.WriteString(w, `"`)
}
