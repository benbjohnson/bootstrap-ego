package bootstrap

import (
	"context"
	"io"
)

type CardColumns struct {
	Class string
	Yield func()
}

func (c *CardColumns) Render(ctx context.Context, w io.Writer) {
	io.WriteString(w, `<div`)
	c.renderClass(ctx, w)
	io.WriteString(w, `>`)

	if c.Yield != nil {
		c.Yield()
	}

	io.WriteString(w, `</div>`)
}

func (c *CardColumns) renderClass(ctx context.Context, w io.Writer) {
	io.WriteString(w, ` class="card-columns`)
	if c.Class != "" {
		io.WriteString(w, ` `)
		io.WriteString(w, c.Class)
	}
	io.WriteString(w, `"`)
}
