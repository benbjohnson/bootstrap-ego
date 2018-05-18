package bootstrap

import (
	"context"
	"html"
	"io"
)

type Jumbotron struct {
	ID    string
	Class string
	Yield func()
}

func (c *Jumbotron) Render(ctx context.Context, w io.Writer) {
	io.WriteString(w, `<div`)
	writeAttr(w, "id", c.ID)
	c.renderClass(ctx, w)
	io.WriteString(w, `>`)

	if c.Yield != nil {
		c.Yield()
	}
	io.WriteString(w, `</div>`)
}

func (c *Jumbotron) renderClass(ctx context.Context, w io.Writer) {
	io.WriteString(w, ` class="jumbotron`)
	if s := c.Class; s != "" {
		io.WriteString(w, " ")
		io.WriteString(w, html.EscapeString(s))
	}
	io.WriteString(w, `"`)
}
