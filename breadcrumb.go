package bootstrap

import (
	"context"
	"io"
)

type Breadcrumb struct {
	ID    string
	Class string
	Yield func()
}

func (c *Breadcrumb) Render(ctx context.Context, w io.Writer) {
	io.WriteString(w, `<ol`)
	appendAttr(w, "id", c.ID)
	c.renderClass(ctx, w)
	io.WriteString(w, `>`)

	if c.Yield != nil {
		c.Yield()
	}

	io.WriteString(w, `</ol>`)
}

func (c *Breadcrumb) renderClass(ctx context.Context, w io.Writer) {
	io.WriteString(w, ` class="breadcrumb`)
	appendHTML(w, c.Class)
	io.WriteString(w, `"`)
}
