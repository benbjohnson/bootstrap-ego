package bootstrap

import (
	"context"
	"io"
)

type DropdownDivider struct {
	ID    string
	Class string
}

func (c *DropdownDivider) Render(ctx context.Context, w io.Writer) {
	io.WriteString(w, `<div`)
	appendAttr(w, "id", c.ID)
	c.renderClass(ctx, w)
	io.WriteString(w, `>`)
	io.WriteString(w, `</div>`)
}

func (c *DropdownDivider) renderClass(ctx context.Context, w io.Writer) {
	io.WriteString(w, ` class="dropdown-divider`)
	appendHTML(w, c.Class)
	io.WriteString(w, `"`)
}
