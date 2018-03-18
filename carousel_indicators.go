package bootstrap

import (
	"context"
	"io"
)

type CarouselIndicators struct {
	ID    string
	Class string
	Yield func()
}

func (c *CarouselIndicators) Render(ctx context.Context, w io.Writer) {
	io.WriteString(w, `<ol`)
	appendAttr(w, "id", c.ID)
	c.renderClass(ctx, w)
	io.WriteString(w, `>`)

	if c.Yield != nil {
		c.Yield()
	}

	io.WriteString(w, `</ol>`)
}

func (c *CarouselIndicators) renderClass(ctx context.Context, w io.Writer) {
	io.WriteString(w, ` class="carousel-indicators`)
	appendHTML(w, c.Class)
	io.WriteString(w, `"`)
}
