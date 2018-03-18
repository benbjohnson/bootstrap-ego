package bootstrap

import (
	"context"
	"io"
)

type InputGroupText struct {
	ID    string
	Class string
	Yield func()
}

func (c *InputGroupText) Render(ctx context.Context, w io.Writer) {
	io.WriteString(w, `<div`)
	appendAttr(w, "id", c.ID)
	c.renderClass(ctx, w)
	io.WriteString(w, `>`)

	if c.Yield != nil {
		c.Yield()
	}
	io.WriteString(w, `</div>`)
}

func (c *InputGroupText) renderClass(ctx context.Context, w io.Writer) {
	io.WriteString(w, ` class="input-group-text`)
	appendHTML(w, c.Class)
	io.WriteString(w, `"`)
}
