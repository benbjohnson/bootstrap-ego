package bootstrap

import (
	"context"
	"io"
)

type CarouselItem struct {
	ID     string
	Class  string
	Active bool
	Yield  func()
}

func (c *CarouselItem) Render(ctx context.Context, w io.Writer) {
	io.WriteString(w, `<div`)
	appendAttr(w, "id", c.ID)
	c.renderClass(ctx, w)
	io.WriteString(w, `>`)

	if c.Yield != nil {
		c.Yield()
	}

	io.WriteString(w, `</div>`)
}

func (c *CarouselItem) renderClass(ctx context.Context, w io.Writer) {
	io.WriteString(w, ` class="carousel-item`)
	if c.Active {
		io.WriteString(w, ` active`)
	}
	appendHTML(w, c.Class)
	io.WriteString(w, `"`)
}
