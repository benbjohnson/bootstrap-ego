package bootstrap

import (
	"context"
	"io"
)

type FormRow struct {
	ID    string
	Class string
	Yield func()
}

func (c *FormRow) Render(ctx context.Context, w io.Writer) {
	io.WriteString(w, `<div`)
	appendAttr(w, "id", c.ID)
	c.renderClass(ctx, w)
	io.WriteString(w, `>`)

	if c.Yield != nil {
		c.Yield()
	}

	io.WriteString(w, `</div>`)
}

func (c *FormRow) renderClass(ctx context.Context, w io.Writer) {
	io.WriteString(w, ` class="form-row`)
	appendHTML(w, c.Class)
	io.WriteString(w, `"`)
}