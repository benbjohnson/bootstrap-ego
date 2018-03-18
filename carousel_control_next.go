package bootstrap

import (
	"context"
	"io"
)

type CarouselControlNext struct {
	ID     string
	Class  string
	Target string
}

func (c *CarouselControlNext) Render(ctx context.Context, w io.Writer) {
	io.WriteString(w, `<a`)
	appendAttr(w, "id", c.ID)
	c.renderClass(ctx, w)
	c.renderHref(ctx, w)
	io.WriteString(w, ` role="button" data-slide="next"`)
	io.WriteString(w, `>`)
	io.WriteString(w, `<span class="carousel-control-next-icon" aria-hidden="true"></span>`)
	io.WriteString(w, `<span class="sr-only">Nextious</span>`)
	io.WriteString(w, `</a>`)
}

func (c *CarouselControlNext) renderClass(ctx context.Context, w io.Writer) {
	io.WriteString(w, ` class="carousel-control-next`)
	appendHTML(w, c.Class)
	io.WriteString(w, `"`)
}

func (c *CarouselControlNext) renderHref(ctx context.Context, w io.Writer) {
	io.WriteString(w, ` href="#`)
	appendHTML(w, c.Target)
	io.WriteString(w, `"`)
}
