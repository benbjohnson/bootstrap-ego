package bootstrap

import (
	"context"
	"io"
)

type FormText struct {
	ID    string
	Class string
	Yield func()
}

func (c *FormText) Render(ctx context.Context, w io.Writer) {
	io.WriteString(w, `<div`)
	appendAttr(w, "id", c.ID)
	c.renderClass(ctx, w)
	io.WriteString(w, `>`)

	if c.Yield != nil {
		c.Yield()
	}

	io.WriteString(w, `</div>`)
}

func (c *FormText) renderClass(ctx context.Context, w io.Writer) {
	io.WriteString(w, ` class="form-text`)
	appendHTML(w, c.Class)
	io.WriteString(w, `"`)
}
