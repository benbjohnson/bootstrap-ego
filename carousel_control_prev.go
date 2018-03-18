package bootstrap

import (
	"context"
	"io"
)

type CarouselControlPrev struct {
	ID     string
	Class  string
	Target string
}

func (c *CarouselControlPrev) Render(ctx context.Context, w io.Writer) {
	io.WriteString(w, `<a`)
	appendAttr(w, "id", c.ID)
	c.renderClass(ctx, w)
	c.renderHref(ctx, w)
	io.WriteString(w, ` role="button" data-slide="prev"`)
	io.WriteString(w, `>`)
	io.WriteString(w, `<span class="carousel-control-prev-icon" aria-hidden="true"></span>`)
	io.WriteString(w, `<span class="sr-only">Previous</span>`)
	io.WriteString(w, `</a>`)
}

func (c *CarouselControlPrev) renderClass(ctx context.Context, w io.Writer) {
	io.WriteString(w, ` class="carousel-control-prev`)
	appendHTML(w, c.Class)
	io.WriteString(w, `"`)
}

func (c *CarouselControlPrev) renderHref(ctx context.Context, w io.Writer) {
	io.WriteString(w, ` href="#`)
	appendHTML(w, c.Target)
	io.WriteString(w, `"`)
}
