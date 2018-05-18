package bootstrap

import (
	"context"
	"html"
	"io"
)

type Container struct {
	ID    string
	Class string
	Fluid bool
	Yield func()
}

func (c *Container) Render(ctx context.Context, w io.Writer) {
	io.WriteString(w, `<div`)
	writeAttr(w, "id", c.ID)
	c.renderClass(ctx, w)
	io.WriteString(w, `>`)

	if c.Yield != nil {
		c.Yield()
	}

	io.WriteString(w, `</div>`)
}

func (c *Container) renderClass(ctx context.Context, w io.Writer) {
	io.WriteString(w, ` class="`)
	if c.Fluid {
		io.WriteString(w, "container-fluid")
	} else {
		io.WriteString(w, "container")
	}
	if s := c.Class; s != "" {
		io.WriteString(w, " ")
		io.WriteString(w, html.EscapeString(s))
	}

	io.WriteString(w, `"`)
}
