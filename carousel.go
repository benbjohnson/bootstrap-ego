package bootstrap

import (
	"context"
	"io"
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
	appendAttr(w, "id", c.ID)
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
	appendHTML(w, c.Class)
	io.WriteString(w, `"`)
}
