package bootstrap

import (
	"context"
	"html"
	"io"
	"strconv"
)

type CarouselIndicator struct {
	ID      string
	Class   string
	Target  string
	SlideTo int
	Active  bool
}

func (c *CarouselIndicator) Render(ctx context.Context, w io.Writer) {
	io.WriteString(w, `<li`)
	appendAttr(w, "id", c.ID)
	c.renderClass(ctx, w)
	c.renderDataTarget(ctx, w)
	c.renderDataSlideTo(ctx, w)
	io.WriteString(w, `>`)
	io.WriteString(w, `</li>`)
}

func (c *CarouselIndicator) renderClass(ctx context.Context, w io.Writer) {
	if !c.Active && c.Class == "" {
		return
	}
	io.WriteString(w, ` class="`)
	if c.Active {
		io.WriteString(w, `active`)
	}
	if c.Active && c.Class != "" {
		io.WriteString(w, ` `)
	}
	if c.Class != "" {
		io.WriteString(w, html.EscapeString(c.Class))
	}
	io.WriteString(w, `"`)
}

func (c *CarouselIndicator) renderDataTarget(ctx context.Context, w io.Writer) {
	if c.Target == "" {
		return
	}
	io.WriteString(w, ` data-target="#`)
	io.WriteString(w, html.EscapeString(c.Target))
	io.WriteString(w, `"`)
}

func (c *CarouselIndicator) renderDataSlideTo(ctx context.Context, w io.Writer) {
	io.WriteString(w, ` data-slide-to="`)
	io.WriteString(w, strconv.Itoa(c.SlideTo))
	io.WriteString(w, `"`)
}
