package bootstrap

import (
	"context"
	"io"
)

type ValidFeedback struct {
	ID    string
	Class string
	Yield func()
}

func (c *ValidFeedback) Render(ctx context.Context, w io.Writer) {
	io.WriteString(w, `<div`)
	appendAttr(w, "id", c.ID)
	c.renderClass(ctx, w)
	io.WriteString(w, `>`)

	if c.Yield != nil {
		c.Yield()
	}

	io.WriteString(w, `</div>`)
}

func (c *ValidFeedback) renderClass(ctx context.Context, w io.Writer) {
	io.WriteString(w, ` class="valid-feedback`)
	appendHTML(w, c.Class)
	io.WriteString(w, `"`)
}
