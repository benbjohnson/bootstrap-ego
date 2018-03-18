package bootstrap

import (
	"context"
	"io"
)

type BreadcrumbItem struct {
	Class  string
	Active bool
	Yield  func()
}

func (c *BreadcrumbItem) Render(ctx context.Context, w io.Writer) {
	io.WriteString(w, `<li`)
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
	if c.Class != "" {
		io.WriteString(w, ` `)
		io.WriteString(w, c.Class)
	}
	io.WriteString(w, `"`)
}
