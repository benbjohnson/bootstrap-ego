package bootstrap

import (
	"context"
	"html"
	"io"
)

type Dropdown struct {
	ID    string
	Class string
	Yield func()
}

func (c *Dropdown) Render(ctx context.Context, w io.Writer) {
	io.WriteString(w, `<div`)
	writeAttr(w, "id", c.ID)
	c.renderClass(ctx, w)
	io.WriteString(w, `>`)

	if c.Yield != nil {
		c.Yield()
	}

	io.WriteString(w, `</div>`)
}

func (c *Dropdown) renderClass(ctx context.Context, w io.Writer) {
	io.WriteString(w, ` class="dropdown`)
	if s := c.Class; s != "" {
		io.WriteString(w, " ")
		io.WriteString(w, html.EscapeString(s))
	}

	io.WriteString(w, `"`)
}

type DropdownMenu struct {
	ID    string
	Class string
	Yield func()
}

func (c *DropdownMenu) Render(ctx context.Context, w io.Writer) {
	io.WriteString(w, `<div`)
	writeAttr(w, "id", c.ID)
	c.renderClass(ctx, w)
	io.WriteString(w, `>`)

	if c.Yield != nil {
		c.Yield()
	}

	io.WriteString(w, `</div>`)
}

func (c *DropdownMenu) renderClass(ctx context.Context, w io.Writer) {
	io.WriteString(w, ` class="dropdown-menu`)
	if s := c.Class; s != "" {
		io.WriteString(w, " ")
		io.WriteString(w, html.EscapeString(s))
	}

	io.WriteString(w, `"`)
}

type DropdownDivider struct {
	ID    string
	Class string
}

func (c *DropdownDivider) Render(ctx context.Context, w io.Writer) {
	io.WriteString(w, `<div`)
	writeAttr(w, "id", c.ID)
	c.renderClass(ctx, w)
	io.WriteString(w, `>`)
	io.WriteString(w, `</div>`)
}

func (c *DropdownDivider) renderClass(ctx context.Context, w io.Writer) {
	io.WriteString(w, ` class="dropdown-divider`)
	if s := c.Class; s != "" {
		io.WriteString(w, " ")
		io.WriteString(w, html.EscapeString(s))
	}

	io.WriteString(w, `"`)
}
