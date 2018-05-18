package bootstrap

import (
	"context"
	"html"
	"io"
	"strconv"
)

type Carousel struct {
	ID         string
	Class      string
	Indicators func()
	Inner      func()
	Controls   func()
}

func (c *Carousel) Render(ctx context.Context, w io.Writer) {
	io.WriteString(w, `<div`)
	writeAttr(w, "id", c.ID)
	c.renderClass(ctx, w)
	io.WriteString(w, ` data-ride="carousel"`)
	io.WriteString(w, `>`)

	if c.Indicators != nil {
		c.Indicators()
	}

	if c.Inner != nil {
		io.WriteString(w, `<div class="carousel-inner">`)
		c.Inner()
		io.WriteString(w, `</div>`)
	}

	if c.Controls != nil {
		c.Controls()
	}

	io.WriteString(w, `</div>`)
}

func (c *Carousel) renderClass(ctx context.Context, w io.Writer) {
	io.WriteString(w, ` class="carousel slide`)
	if s := c.Class; s != "" {
		io.WriteString(w, " ")
		io.WriteString(w, html.EscapeString(s))
	}

	io.WriteString(w, `"`)
}

type CarouselItem struct {
	ID     string
	Class  string
	Active bool
	Yield  func()
}

func (c *CarouselItem) Render(ctx context.Context, w io.Writer) {
	io.WriteString(w, `<div`)
	writeAttr(w, "id", c.ID)
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
	if s := c.Class; s != "" {
		io.WriteString(w, " ")
		io.WriteString(w, html.EscapeString(s))
	}

	io.WriteString(w, `"`)
}

type CarouselIndicators struct {
	ID    string
	Class string
	Yield func()
}

func (c *CarouselIndicators) Render(ctx context.Context, w io.Writer) {
	io.WriteString(w, `<ol`)
	writeAttr(w, "id", c.ID)
	c.renderClass(ctx, w)
	io.WriteString(w, `>`)

	if c.Yield != nil {
		c.Yield()
	}

	io.WriteString(w, `</ol>`)
}

func (c *CarouselIndicators) renderClass(ctx context.Context, w io.Writer) {
	io.WriteString(w, ` class="carousel-indicators`)
	if s := c.Class; s != "" {
		io.WriteString(w, " ")
		io.WriteString(w, html.EscapeString(s))
	}

	io.WriteString(w, `"`)
}

type CarouselIndicator struct {
	ID      string
	Class   string
	Target  string
	SlideTo int
	Active  bool
}

func (c *CarouselIndicator) Render(ctx context.Context, w io.Writer) {
	io.WriteString(w, `<li`)
	writeAttr(w, "id", c.ID)
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

type CarouselControlPrev struct {
	ID     string
	Class  string
	Target string
}

func (c *CarouselControlPrev) Render(ctx context.Context, w io.Writer) {
	io.WriteString(w, `<a`)
	writeAttr(w, "id", c.ID)
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
	if s := c.Class; s != "" {
		io.WriteString(w, " ")
		io.WriteString(w, html.EscapeString(s))
	}
	io.WriteString(w, `"`)
}

func (c *CarouselControlPrev) renderHref(ctx context.Context, w io.Writer) {
	io.WriteString(w, ` href="#`)
	if s := c.Target; s != "" {
		io.WriteString(w, " ")
		io.WriteString(w, html.EscapeString(s))
	}
	io.WriteString(w, `"`)
}

type CarouselControlNext struct {
	ID     string
	Class  string
	Target string
}

func (c *CarouselControlNext) Render(ctx context.Context, w io.Writer) {
	io.WriteString(w, `<a`)
	writeAttr(w, "id", c.ID)
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
	if s := c.Class; s != "" {
		io.WriteString(w, " ")
		io.WriteString(w, html.EscapeString(s))
	}

	io.WriteString(w, `"`)
}

func (c *CarouselControlNext) renderHref(ctx context.Context, w io.Writer) {
	io.WriteString(w, ` href="#`)
	if s := c.Target; s != "" {
		io.WriteString(w, " ")
		io.WriteString(w, html.EscapeString(s))
	}
	io.WriteString(w, `"`)
}
