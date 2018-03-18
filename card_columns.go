package bootstrap

import (
	"context"
	"io"
)

type CardColumns struct {
	ID    string
	Class string
	Yield func()
}

func (c *CardColumns) Render(ctx context.Context, w io.Writer) {
	io.WriteString(w, `<div`)
	appendAttr(w, "id", c.ID)
	c.renderClass(ctx, w)
	io.WriteString(w, `>`)

	if c.Yield != nil {
		c.Yield()
	}

	io.WriteString(w, `</div>`)
}

func (c *CardColumns) renderClass(ctx context.Context, w io.Writer) {
	io.WriteString(w, ` class="card-columns`)
	appendHTML(w, c.Class)
	io.WriteString(w, `"`)
}
