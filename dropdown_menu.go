package bootstrap

import (
	"context"
	"io"
)

type DropdownMenu struct {
	ID    string
	Class string
	Yield func()
}

func (c *DropdownMenu) Render(ctx context.Context, w io.Writer) {
	io.WriteString(w, `<div`)
	appendAttr(w, "id", c.ID)
	c.renderClass(ctx, w)
	io.WriteString(w, `>`)

	if c.Yield != nil {
		c.Yield()
	}

	io.WriteString(w, `</div>`)
}

func (c *DropdownMenu) renderClass(ctx context.Context, w io.Writer) {
	io.WriteString(w, ` class="dropdown-menu`)
	appendHTML(w, c.Class)
	io.WriteString(w, `"`)
}
