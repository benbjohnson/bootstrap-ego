package bootstrap

import (
	"context"
	"io"
)

type FormCheck struct {
	ID     string
	Class  string
	Inline bool
	Yield  func()
}

func (c *FormCheck) Render(ctx context.Context, w io.Writer) {
	io.WriteString(w, `<div`)
	appendAttr(w, "id", c.ID)
	c.renderClass(ctx, w)
	io.WriteString(w, `>`)

	if c.Yield != nil {
		c.Yield()
	}

	io.WriteString(w, `</div>`)
}

func (c *FormCheck) renderClass(ctx context.Context, w io.Writer) {
	io.WriteString(w, ` class="form-check`)
	if c.Inline {
		io.WriteString(w, ` form-check-inline`)
	}
	appendHTML(w, c.Class)
	io.WriteString(w, `"`)
}
