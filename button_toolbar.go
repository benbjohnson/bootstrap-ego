package bootstrap

import (
	"context"
	"io"
)

type ButtonToolbar struct {
	ID       string
	Class    string
	Size     string
	Vertical bool
	Yield    func()
}

func (c *ButtonToolbar) Render(ctx context.Context, w io.Writer) {
	io.WriteString(w, `<div role="toolbar"`)
	appendAttr(w, "id", c.ID)
	c.renderClass(ctx, w)
	io.WriteString(w, `>`)

	if c.Yield != nil {
		c.Yield()
	}

	io.WriteString(w, `</div>`)
}

func (c *ButtonToolbar) renderClass(ctx context.Context, w io.Writer) {
	io.WriteString(w, ` class="btn-toolbar`)
	appendHTML(w, c.Class)
	io.WriteString(w, `"`)
}
