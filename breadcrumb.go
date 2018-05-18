package bootstrap

import (
	"context"
	"html"
	"io"
)

type Breadcrumb struct {
	ID    string
	Class string
	Yield func()
}

func (c *Breadcrumb) Render(ctx context.Context, w io.Writer) {
	io.WriteString(w, `<ol`)
	writeAttr(w, "id", c.ID)
	c.renderClass(ctx, w)
	io.WriteString(w, `>`)

	if c.Yield != nil {
		c.Yield()
	}

	io.WriteString(w, `</ol>`)
}

func (c *Breadcrumb) renderClass(ctx context.Context, w io.Writer) {
	io.WriteString(w, ` class="breadcrumb`)
	if s := c.Class; s != "" {
		io.WriteString(w, " ")
		io.WriteString(w, html.EscapeString(s))
	}
	io.WriteString(w, `"`)
}

type BreadcrumbItem struct {
	ID     string
	Class  string
	Active bool
	Yield  func()
}

func (c *BreadcrumbItem) Render(ctx context.Context, w io.Writer) {
	io.WriteString(w, `<li`)
	writeAttr(w, "id", c.ID)
	c.renderClass(ctx, w)
	io.WriteString(w, `>`)

	if c.Yield != nil {
		c.Yield()
	}

	io.WriteString(w, `</li>`)
}

func (c *BreadcrumbItem) renderClass(ctx context.Context, w io.Writer) {
	io.WriteString(w, ` class="breadcrumb-item`)
	if c.Active {
		io.WriteString(w, ` active`)
	}
	if s := c.Class; s != "" {
		io.WriteString(w, " ")
		io.WriteString(w, html.EscapeString(s))
	}

	io.WriteString(w, `"`)
}
