package bootstrap

import (
	"context"
	"io"
)

type ButtonToolbar struct {
	Class    string
	Size     string
	Vertical bool
	Yield    func()
}

func (c *ButtonToolbar) Render(ctx context.Context, w io.Writer) {
	io.WriteString(w, `<div role="toolbar"`)
	c.renderClass(ctx, w)
	io.WriteString(w, `>`)

	if c.Yield != nil {
		c.Yield()
	}

	io.WriteString(w, `</div>`)
}

func (c *ButtonToolbar) renderClass(ctx context.Context, w io.Writer) {
	io.WriteString(w, ` class="btn-toolbar`)
	if c.Class != "" {
		io.WriteString(w, ` `)
		io.WriteString(w, c.Class)
	}
	io.WriteString(w, `"`)
}
